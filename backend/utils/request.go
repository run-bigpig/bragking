package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// GetUrlParam 解析url 获取指定参数的值
func GetUrlParam(url, param string) string {
	nowparam := "&" + param + "="
	index := strings.LastIndex(url, nowparam)
	if index == -1 {
		return ""
	}
	index += len(param) + 2
	end := strings.Index(url[index:], "&")
	if end == -1 {
		return url[index:]
	}
	return url[index : index+end]
}

// SendRequest 请求
func SendRequest(method string, url string, headers map[string]string, data interface{}) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	client := &http.Client{
		Timeout: time.Second * 50,
	}
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}
	// 设置请求头（如果有）
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParseCookie 解析cookie
func ParseCookie(cookie string) map[string]string {
	cookies := strings.Split(cookie, ";")
	cookieMap := make(map[string]string)
	for _, v := range cookies {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			cookieMap[kv[0]] = kv[1]
		}
	}
	return cookieMap
}

// ExtractUid 提取uid
func ExtractUid(cookie string) int64 {
	cookieMap := ParseCookie(cookie)
	for key, _ := range cookieMap {
		if strings.Index(key, "GET") != -1 {
			uidStr := strings.ReplaceAll(key, "GET", "")
			return StringToInt64(uidStr)
		}
	}
	return 0
}
