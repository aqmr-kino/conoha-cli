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

// PaymentHistory :
// 入金履歴
type PaymentHistory struct {
	MoneyType     string `json:"money_type"`
	DepositAmount int    `json:"deposit_amount"`
	ReceivedDate  string `json:"received_date"`
}

// PaymentHistories :
// 入金履歴一覧
type PaymentHistories struct {
	Histories []PaymentHistory `json:"payment_history"`
}

// PaymentSummary :
// 入金状況
type PaymentSummary struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
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

//
// 入金履歴操作
//

// GetPaymentHistories :
// 入金履歴一覧の取得
func (mgr *BillingManager) GetPaymentHistories() (*PaymentHistories, error) {
	var ret PaymentHistories

	uri := mgr.Endpoint + "/payment-history"
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

// GetPaymentSummary :
// 入金状況取得
func (mgr *BillingManager) GetPaymentSummary() (*PaymentSummary, error) {
	var ret PaymentSummary

	uri := mgr.Endpoint + "/payment-summary"
	byteArray, _, err := mgr.Token.Request(&api.APIRequest{
		URI:       uri,
		Method:    api.GET,
		GetParams: nil,
	})

	if err != nil {
		return nil, err
	}

	var tmp struct {
		Summary PaymentSummary `json:"payment_summary"`
	}

	err2 := json.Unmarshal(byteArray, &ret)

	if err2 != nil {
		return nil, err2
	}

	ret = tmp.Summary

	return &ret, nil
}
