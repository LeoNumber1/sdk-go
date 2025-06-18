package client

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	BaseURL    string
	HttpClient *http.Client

	AccessKey string
	SecretKey string
}

func NewClient(baseURL, accessKey, secretKey string) *Client {
	// 初始化客户端逻辑
	return &Client{
		BaseURL: baseURL,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 忽略证书验证
				},
			},
		},

		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}
