package dingding

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type DingDing interface {
	SendMessage(token string, msg *DingMessage) error
}

type dingding struct {
	httpClient *http.Client
}

func NewDingDing(
	httpClient *http.Client,
) DingDing {
	return &dingding{
		httpClient: httpClient,
	}
}

func (s *dingding) send(token, msg string) error {

	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
	resp, err := s.httpClient.Post(url, "application/json", strings.NewReader(msg))
	if err != nil {
		return errors.Wrap(err, "send dingding message error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response error")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status is not 200[code=%d,body=%s]", resp.StatusCode, body)
	}

	return nil
}

func (s *dingding) SendMessage(token string, msg *DingMessage) error {
	bytes, _ := json.Marshal(msg)
	return s.send(token, string(bytes))
}

// Markdown markdown消息
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// ActionCard  独立跳转action card类型
type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []Btn  `json:"btns"`
}

// Btn 按钮
type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// At 用来@别人
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// DingMessage 文本格式
type DingMessage struct {
	MessageType string   `json:"msgtype"`
	Markdown    Markdown `json:"markdown"`
	At          At       `json:"at"`
}

// DingActionMessage .
type DingActionMessage struct {
	MessageType string     `json:"msgtype"`
	ActionCard  ActionCard `json:"actionCard"`
}

var ProviderSet = wire.NewSet(NewDingDing)
