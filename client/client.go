package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL    string
	HttpClient *http.Client

	AccessKey string
	SecretKey string
}

type Config struct {
	BaseURL             string
	AccessKey           string
	SecretKey           string
	Timeout             time.Duration
	DialTimeout         time.Duration
	HeaderTimeout       time.Duration
	TLSHandshakeTimeout time.Duration
}

func NewClient(cfg Config) *Client {
	// 检查参数是否为空
	if cfg.BaseURL == "" || cfg.AccessKey == "" || cfg.SecretKey == "" {
		panic("baseURL, accessKey, and secretKey cannot be empty")
	}
	// 设置默认值
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}
	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = 15 * time.Second
	}
	if cfg.HeaderTimeout == 0 {
		cfg.HeaderTimeout = 10 * time.Second
	}
	if cfg.TLSHandshakeTimeout == 0 {
		cfg.TLSHandshakeTimeout = 5 * time.Second
	}
	// 初始化客户端逻辑
	return &Client{
		BaseURL: strings.TrimSuffix(cfg.BaseURL, "/"),
		HttpClient: &http.Client{
			Timeout: cfg.Timeout, // 设置总超时时间（包括连接、请求和响应）
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 忽略证书验证
				},
				DialContext: (&net.Dialer{
					Timeout:   cfg.DialTimeout, // 连接超时
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ResponseHeaderTimeout: cfg.HeaderTimeout,       // 响应头超时
				TLSHandshakeTimeout:   cfg.TLSHandshakeTimeout, // TLS握手超时
			},
		},

		AccessKey: cfg.AccessKey,
		SecretKey: cfg.SecretKey,
	}
}
