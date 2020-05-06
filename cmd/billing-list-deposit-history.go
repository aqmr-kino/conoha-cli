package cmd

import (
	"conoha-cli/conoha/account"
	"fmt"

	"github.com/spf13/cobra"
)

var billingListDepositHistoryCmd = &cobra.Command{
	Use:   "list-deposit-history",
	Short: "Listing deposit history",
	Long:  `Listing deposit history`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// 入金履歴取得
		mgr := &account.BillingManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Account,
		}

		history, err2 := mgr.GetPaymentHistories()

		if err2 != nil {
			fmt.Printf("Error: Get payment history infomation failed. (%s)\n", err2)
			return
		}

		// 結果出力
		fmt.Printf("%-12s %-8s %-20s\n", "MoneyType", "Amount", "ReceivedDate")
		for _, h := range history.Histories {
			fmt.Printf("%-12s %8d %-20s\n", h.MoneyType, h.DepositAmount, h.ReceivedDate)
		}
	},
}

func init() {
}
