package http

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http/middlewares/ginprom"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/utils/netutil"
	junogin "go.didapinche.com/juno-gin/v2"
	"go.didapinche.com/juno-go/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Options struct {
	Port int
	Mode string
}

// Server define address server
type Server struct {
	o          *Options
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

type InitControllers func(r *gin.Engine)

// NewRouter initialize gin router
func NewRouter(o *Options, logger *zap.Logger, init InitControllers) (*gin.Engine, error) {

	// 配置gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery()) // panic之后自动恢复
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(ginprom.New(r).Middleware()) // 添加prometheus 监控
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"address://*", "https://*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origins", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/juno", gin.WrapH(juno.MetaHandler()))

	pprof.Register(r)
	junogin.Apply(r)

	init(r)

	return r, nil
}

// New is constructor of address server
func New(o *Options, logger *zap.Logger, router *gin.Engine) (*Server, error) {
	var s = &Server{
		logger: logger.With(zap.String("type", "address.Server")),
		router: router,
		o:      o,
	}

	return s, nil
}

// Host returns address host
func (s *Server) Host() string {
	return s.host
}

// Port returns address port
func (s *Server) Port() int {
	return s.port
}

// Start address server
func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}

	s.host = netutil.GetLocalIP4()

	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("address server starting ...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start address server err", zap.Error(err))
			return
		}
	}()

	return nil
}

// Stop address server
func (s *Server) Stop() error {
	s.logger.Info("stopping address server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown address server error")
	}

	return nil
}

// ProviderSet define provider set of address package
var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)
