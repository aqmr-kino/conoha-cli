package compute

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/api"
	"encoding/json"
)

// Flavor :
// VMプラン
type Flavor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CPUs int    `json:"vcpus"`
	RAM  int    `json:"ram"`
	Disk int    `json:"disk"`
}

// Flavors :
// VMプラン一覧
type Flavors struct {
	Flavors []Flavor `json:"flavors"`
}

// VMImage :
// VMイメージ
type VMImage struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Size     int    `json:"OS-EXT-IMG-SIZE:size"`
	MetaData struct {
		App    string `json:"app"`
		OSType string `json:"os_type"`
	} `json:"metadata"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// VMImages :
// VMイメージ一覧
type VMImages struct {
	Images []VMImage `json:"images"`
}

// VirtualMachine :
// 仮想マシン
type VirtualMachine struct {
	ID         string `json:"id"`
	HostID     string `json:"hostId"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	AccessIPv4 string `json:"accessIPv4"`
	AccessIPv6 string `json:"accessIPv6"`
	Addresses  map[string][]struct {
		Addr    string `json:"addr"`
		Version int    `json:"version"`
	} `json:"addresses"`
	Flavor struct {
		ID    string `json:"id"`
		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
	} `json:"flavor"`
	Image struct {
		ID    string `json:"id"`
		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
	} `json:"image"`
	VolumesAttached []struct {
		ID string `json:"id"`
	} `json:"os-extended-volumes:volumes_attached"`
	SecurityGroups []struct {
		Name string `json:"name"`
	} `json:"security_groups"`
	Metadata struct {
		InstanceNameTag string `json:"instance_name_tag"`
	} `json:"metadata"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	ConfigDrive string `json:"config_drive"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	UserID      string `json:"user_id"`
	TenantID    string `json:"tenant_id"`
}

// VirtualMachines :
// 仮想マシン一覧
type VirtualMachines struct {
	Servers []VirtualMachine `json:"servers"`
}

// VMManager :
// 仮想マシン管理
//   Token:    Conoha API アクセストークン
//   Endpoint: Conoha Compute API エンドポイント
type VMManager struct {
	Endpoint string
	Token    *account.ConohaToken
}

//
// プラン情報管理
//

// GetFlavors :
// プラン情報一覧取得
func (mgr *VMManager) GetFlavors() (*Flavors, error) {
	var ret Flavors

	uri := mgr.Endpoint + "/flavors/detail"
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	err2 := json.Unmarshal(byteArray, &ret)

	if err2 != nil {
		return nil, err2
	}

	return &ret, nil
}

// FindFlavor :
// プラン情報取得
func (mgr *VMManager) FindFlavor(id string) (*Flavor, error) {
	var ret Flavor

	uri := mgr.Endpoint + "/flavors" + "/" + id
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	var tmp struct {
		Flavor Flavor `json:"flavor"`
	}

	err2 := json.Unmarshal(byteArray, &tmp)

	if err2 != nil {
		return nil, err2
	}

	ret = tmp.Flavor

	return &ret, nil
}

//
// VMイメージ管理
//

// GetVMImages :
// VMイメージ情報一覧取得
func (mgr *VMManager) GetVMImages() (*VMImages, error) {
	var ret VMImages

	uri := mgr.Endpoint + "/images/detail"
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	err2 := json.Unmarshal(byteArray, &ret)

	if err2 != nil {
		return nil, err2
	}

	return &ret, nil
}

// FindVMImage :
// VMイメージ情報取得
func (mgr *VMManager) FindVMImage(id string) (*VMImage, error) {
	var ret VMImage

	uri := mgr.Endpoint + "/images" + "/" + id
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	var tmp struct {
		Image VMImage `json:"image"`
	}

	err2 := json.Unmarshal(byteArray, &tmp)

	if err2 != nil {
		return nil, err2
	}

	ret = tmp.Image

	return &ret, nil
}

//
// 仮想サーバー管理
//

// GetVirtualMachines :
// 仮想サーバ一覧取得
func (mgr *VMManager) GetVirtualMachines() (*VirtualMachines, error) {
	var ret VirtualMachines

	uri := mgr.Endpoint + "/servers/detail"
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	err2 := json.Unmarshal(byteArray, &ret)

	if err2 != nil {
		return nil, err2
	}

	return &ret, nil
}
