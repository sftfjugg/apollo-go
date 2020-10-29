package services

import (
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/single_queue"
	"time"
)

type ReleaseMessageService interface {
	Poll()
}

type releaseMessageService struct {
	repository repositories.ReleaseMessageRepository
}

func NewReleaseMessageService(repository repositories.ReleaseMessageRepository) ReleaseMessageService {
	return &releaseMessageService{
		repository: repository,
	}
}

//一直运行,需要单独启动
func (s releaseMessageService) Poll() {
	releases, err := s.repository.FindAll()
	if err != nil {
		errors.Wrap(err, "mysql poll failed，Please restart")
	}
	m := single_queue.GetV()
	for _, r := range releases {
		m[r.Message] = r.Id
	}
	max := 0
	if len(releases) > 1 {
		max = int(releases[len(releases)-1].Id)
	}
	ticker := time.NewTicker(1 * time.Second)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for range ticker.C {
			releases, err := s.repository.FindOffsetByMax(max)
			if err != nil {
				errors.Wrap(err, "mysql poll failed")
			}
			for _, r := range releases {
				m[r.Message] = r.Id
			}
			if len(releases) > 0 {
				max = int(releases[len(releases)-1].Id)
			}
		}
	}(ticker)

}
