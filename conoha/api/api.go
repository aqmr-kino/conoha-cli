package api

import "io"

// HTTP reuest methods
const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

// APIRequest :
// Conoha APIリクエスト定義
type APIRequest struct {
	URI        string
	ReqHeaders map[string]string
	Method     string
	GetParams  io.Reader
	PostBody   io.Reader
}
