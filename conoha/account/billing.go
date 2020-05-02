package account

import (
	"conoha-cli/conoha/api"
	"encoding/json"
)

// Item :
// アイテム
type Item struct {
	UUID             string `json:"uu_id"`
	ServiceName      string `json:"service_name"`
	ServiceStartDate string `json:"service_start_date"`
	ItemStatus       string `json:"item_status"`
}

// Items :
// アイテム一覧
type Items struct {
	OrderItems []Item `json:"order_items"`
}

// BillingInvoice :
// 請求書
type BillingInvoice struct {
	InvoiceID         int    `json:"invoice_id"`
	PaymentMethodType string `json:"payment_method_type"`
	InvoiceDate       string `json:"invoice_date"`
	Bill              int    `json:"bill_plus_tax"`
	DueDate           string `json:"due_date"`
}

// BillingInvoices :
// 請求書一覧
type BillingInvoices struct {
	Invoices []BillingInvoice `json:"billing_invoices"`
}

// BillingManager :
// 入出金管理
//   Token:    Conoha API アクセストークン
//   Endpoint: Conoha Account/Billing API エンドポイント
type BillingManager struct {
	Token    *ConohaToken
	Endpoint string
}

//
// アイテム操作
//

// GetItems :
// アイテム一覧の取得
func (mgr *BillingManager) GetItems() (*Items, error) {
	var ret Items

	uri := mgr.Endpoint + "/order-items"
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

//
// 請求書操作
//

// GetInvoices :
// 請求書一覧の取得
func (mgr *BillingManager) GetInvoices() (*BillingInvoices, error) {
	var ret BillingInvoices

	uri := mgr.Endpoint + "/billing-invoices"
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
