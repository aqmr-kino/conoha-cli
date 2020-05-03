package network

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/api"
	"encoding/json"
)

// SecurityGroupRule :
// セキュリティグループ ルール定義
type SecurityGroupRule struct {
	ID              string `json:"id"`
	Ethertype       string `json:"ethertype"`
	Protocol        string `json:"protocol"`
	PortRangeMin    int    `json:"port_range_min"`
	PortRangeMax    int    `json:"port_range_max"`
	Direction       string `json:"direction"`
	TenantID        string `json:"tenant_id"`
	SecurityGroupID string `json:"security_group_id"`
	RemoteIPPrefix  string `json:"remote_ip_prefix"`
	RemoteGroupID   string `json:"remote_group_id"`
}

// SecurityGroup :
// セキュリティグループ定義
type SecurityGroup struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	TenantID    string              `json:"tenant_id"`
	Rules       []SecurityGroupRule `json:"security_group_rules"`
}

// SecurityGroupRules :
// セキュリティグループ ルール定義一覧
type SecurityGroupRules struct {
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
}

// SecurityGroups :
// セキュリティグループ定義一覧
type SecurityGroups struct {
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

// SecurityGroupManager :
// セキュリティグループ管理
//   Token:    Conoha API アクセストークン
//   Endpoint: Conoha Network API エンドポイント
type SecurityGroupManager struct {
	Token    *account.ConohaToken
	Endpoint string
}

//
// セキュリティグループ操作
//

// GetGroups :
// セキュリティグループ一覧の取得
func (mgr *SecurityGroupManager) GetGroups() (*SecurityGroups, error) {
	var ret SecurityGroups

	uri := mgr.Endpoint + "/v2.0/security-groups"
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

// GetGroupDetail :
// セキュリティグループ詳細情報の取得
func (mgr *SecurityGroupManager) GetGroupDetail(sgid string) (*SecurityGroup, error) {
	var ret SecurityGroup

	uri := mgr.Endpoint + "/v2.0/security-groups" + "/" + sgid
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	var tmp struct {
		SG SecurityGroup `json:"security_group"`
	}

	err2 := json.Unmarshal(byteArray, &tmp)

	if err2 != nil {
		return nil, err2
	}

	ret = tmp.SG

	return &ret, nil
}

//
// セキュリティグループ ルール操作
//

// GetRules :
// ルール一覧の取得
func (mgr *SecurityGroupManager) GetRules() (*SecurityGroupRules, error) {
	var ret SecurityGroupRules

	uri := mgr.Endpoint + "/v2.0/security-group-rules"
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

// GetRuleDetail :
// ルール詳細の取得
func (mgr *SecurityGroupManager) GetRuleDetail(rid string) (*SecurityGroupRule, error) {
	var ret SecurityGroupRule

	uri := mgr.Endpoint + "/v2.0/security-group-rules" + "/" + rid
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	var tmp struct {
		SGRule SecurityGroupRule `json:"security_group_rule"`
	}

	err2 := json.Unmarshal(byteArray, &tmp)

	if err2 != nil {
		return nil, err2
	}

	ret = tmp.SGRule

	return &ret, nil
}
