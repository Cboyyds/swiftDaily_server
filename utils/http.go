package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// HttpRequest 发起HTTP请求并返回响应结果。
// method: HTTP方法，如GET、POST等。
// urlStr: 请求的URL地址。
// headers: 请求头，用于设置请求的头部信息。
// params: 查询参数，将附加到URL的查询字符串中。
// data: 请求体数据，将被序列化为JSON格式。
// 返回值: *http.Response类型的响应对象和error类型的错误信息。
func HttpRequest(
	method string,
	urlStr string,
	headers map[string]string,
	params map[string]string,
	data any) (*http.Response, error) {
	// 解析URL地址，确保其有效性。
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// 构建查询字符串。
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()
	// 构建请求体。
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	// 创建新的HTTP请求。
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	// 设置请求头。
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// 设置Content-Type为application/json，仅当有请求体时设置。
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// 发起HTTP请求并获取响应。
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
