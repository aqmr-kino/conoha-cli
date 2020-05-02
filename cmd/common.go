package cmd

import (
	"conoha-cli/conoha/account"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

// ConohaAPIEndpoint :
// Conoha API エンドポイント設定
type ConohaAPIEndpoint struct {
	Account       string
	Compute       string
	Volume        string
	Database      string
	Image         string
	DNS           string
	ObjectStorage string
	Mail          string
	Idenity       string
	Network       string
}

// ConohaCLIConfig :
// Conoha CLI設定
type ConohaCLIConfig struct {
	Credential *account.Credentials `json:"credential"`
	Endpoint   *ConohaAPIEndpoint   `json:"endpoint"`
}

// グローバル設定変数
var (
	Configure      ConohaCLIConfig
	ConfigFilename string
)

// SaveAs :
// 設定をJSONファイルへ保存
func (conf *ConohaCLIConfig) SaveAs(fname string) error {
	data, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	perm := "0600"
	perm32, _ := strconv.ParseUint(perm, 8, 32)
	ioutil.WriteFile(fname, data, os.FileMode(perm32))

	return nil
}

// GetConfigFromFile :
// ファイルから設定を読み込む
func GetConfigFromFile(fname string) (*ConohaCLIConfig, error) {
	var ret ConohaCLIConfig

	buf, err := ioutil.ReadFile(fname)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(buf, &ret)

	return &ret, nil
}

// GetOrCreateConfigFromFile :
// ファイルから設定を読み込む（存在しない場合は新規作成）
func GetOrCreateConfigFromFile(fname string) (*ConohaCLIConfig, error) {
	var ret *ConohaCLIConfig

	_, err := os.Stat(fname)

	if os.IsNotExist(err) {
		ret := &ConohaCLIConfig{
			Credential: &account.Credentials{},
			Endpoint:   &ConohaAPIEndpoint{},
		}
		ret.SaveAs(fname)
	}

	ret, _ = GetConfigFromFile(fname)

	return ret, nil
}

// LoadGlobalConfigure :
// 設定ファイル読み込み
func LoadGlobalConfigure() error {
	// 設定ファイル名取得
	usr, _ := user.Current()
	ConfigFilename = usr.HomeDir + "/.conoha"

	// 読み込み
	tmp, err := GetOrCreateConfigFromFile(ConfigFilename)
	Configure = *tmp

	return err
}
