package api

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"sdk/client"
	"sdk/models"
	"strconv"
	"time"
)

func generateHeader(client *client.Client, method, path string,
	extraHeaders map[string]interface{}, body []byte) http.Header {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	// nonce 是12位随机字符串
	nonce := generateNonce(12)
	// 生成签名
	signature := generateSignature(client.SecretKey, method, path, timestamp, nonce, body)
	headers := http.Header{}
	headers.Set("Authorization", client.AccessKey+":"+signature)
	headers.Set("X-Timestamp", timestamp)
	headers.Set("X-Nonce", nonce)
	for key, value := range extraHeaders {
		headers.Set(key, fmt.Sprintf("%v", value))
	}
	return headers
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func generateNonce(length int) string {
	result := make([]byte, length)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func generateSignature(secretKey, method, path, timestamp, nonce string, body []byte) string {
	// 1. 构建签名字符串
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s",
		method,
		url.QueryEscape(path),
		timestamp,
		nonce,
		hex.EncodeToString(sha256.New().Sum(body)))

	// 2. 计算HMAC-SHA256
	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write([]byte(stringToSign))

	// 3. Base64编码
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}

func httpRequest(client *http.Client, url, method string, headers http.Header,
	payload io.Reader, res interface{}) (resp models.Response, err error) {
	defer func() {
		if err := recover(); err != nil {
			// 处理 panic
			fmt.Printf("panic in send request(%s): %v", url, err)
			return
		}
	}()
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	req.Header = headers
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if response.Body != nil {
			err := response.Body.Close()
			if err != nil {
				fmt.Printf("close response body failed: %v", err)
			}
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New("call http response not 200 OK")
		if response.StatusCode == http.StatusForbidden {
			if err = json.Unmarshal(body, &resp); err != nil {
				return
			}
			err = errors.New(resp.Message)
			return
		}
		return
	}
	resp.Data = res
	if err = json.Unmarshal(body, &resp); err != nil {
		return
	}
	return resp, nil
}
