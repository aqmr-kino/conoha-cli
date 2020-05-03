package cmd

import (
	"conoha-cli/conoha/account"
	"fmt"

	"github.com/spf13/cobra"
)

var billingListInvoiceCmd = &cobra.Command{
	Use:   "list-invoice",
	Short: "Listing billing invoices",
	Long:  `Listing billing invoices`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
		}

		// 請求書取得
		mgr := &account.BillingManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Account,
		}

		inv, err2 := mgr.GetInvoices()

		if err2 != nil {
			fmt.Printf("Error: Get billing invoices infomation failed. (%s)\n", err2)
		}

		// 結果出力
		fmt.Printf("%-10s %-12s %-20s %-8s %-20s\n", "ID", "Method", "InvoiceDate", "Amount", "Due")
		for _, i := range inv.Invoices {
			fmt.Printf("%10d %-12s %-20s %8d %-20s\n", i.InvoiceID, i.PaymentMethodType, i.InvoiceDate, i.Bill, i.DueDate)
		}
	},
}

func init() {
}
