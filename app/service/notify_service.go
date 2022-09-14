package service

import (
	"context"
	"github.com/jjonline/go-lib-backend/ding"
	"github.com/jjonline/go-lib-backend/feishu"
	"github.com/tvb-sz/cloud-build-notify/client"
	"net/http"
	"strings"
	"time"
)

var (
	// httpClient http client config
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 10 * time.Second,
			DisableCompression:  true,
			MaxIdleConns:        400,
			MaxIdleConnsPerHost: 20,
			MaxConnsPerHost:     50,
			IdleConnTimeout:     120 * time.Second,
		},
	}
	// NotifyService service singleton instance
	NotifyService = &notifyService{}
)

type notifyService struct {
}

// InitDing init ding client
func (n *notifyService) InitDing(token, secret string) {
	client.DingTalk = ding.New(token, secret, true, httpClient)
}

// InitFeishu init feishu client
func (n *notifyService) InitFeishu(token, secret string) {
	client.FeiShu = feishu.New(token, secret, true, httpClient)
}

// DingNotifyOk 钉钉发送`成功`状态消息
func (n *notifyService) DingNotifyOk(title, content string) error {
	return client.DingTalk.Markdown(title, strings.ReplaceAll(content, "\\n", "\n"), nil, false)
}

// FeishuNotifyOk 飞书发送`成功`状态消息
func (n *notifyService) FeishuNotifyOk(title, content string) error {
	ctx := context.Background()
	return client.FeiShu.Info(ctx, title, strings.ReplaceAll(content, "\\n", "\n"), time.Now())
}
