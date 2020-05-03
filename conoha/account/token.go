package account

import (
	"bytes"
	"conoha-cli/conoha/api"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Credentials :
// Conoha API アクセストークン取得のための認証情報
type Credentials struct {
	Auth struct {
		PasswordCredentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"passwordCredentials"`
		TenantID string `json:"tenantId"`
	} `json:"auth"`
}

// ConohaToken :
// Conoha API アクセストークン
type ConohaToken struct {
	Access struct {
		Token struct {
			Issuedat string `json:"issued_at"`
			Expires  string `json:"expires"`
			ID       string `json:"id"`
		} `json:"token"`
	} `json:"access"`
}

// Request :
// APIリクエスト発行
// returns (response []byte, HTTPstatus int, error error)
func (token *ConohaToken) Request(req *api.APIRequest) ([]byte, int, error) {
	// HTTP クライアント生成
	var httpreq *http.Request

	if req.Method == api.POST || req.Method == api.PUT {
		httpreq, _ = http.NewRequest(req.Method, req.URI, req.PostBody)
		httpreq.Header.Set("Content-Type", "application/json")
	} else {
		httpreq, _ = http.NewRequest(req.Method, req.URI, nil)
	}

	httpreq.Header.Set("X-Auth-Token", token.Access.Token.ID)
	httpreq.Header.Set("Accept", "application/json, */*")

	// HTTP リクエスト実行
	client := new(http.Client)
	response, err := client.Do(httpreq)

	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, response.StatusCode, errors.New(response.Status)
	}

	byteArray, _ := ioutil.ReadAll(response.Body)

	return byteArray, response.StatusCode, nil
}

// GetToken :
// APIトークンの取得
func GetToken(endpoint string, cred *Credentials) (*ConohaToken, error) {
	var ret ConohaToken

	uri := endpoint + "/v2.0/tokens"

	// 認証情報をjson化
	data, _ := json.Marshal(cred)

	// HTTP クライアント生成
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, */*")

	// HTTP リクエスト実行
	client := new(http.Client)
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, errors.New(response.Status)
	}

	byteArray, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(byteArray, &ret)

	return &ret, nil
}
