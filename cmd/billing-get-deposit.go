package cmd

import (
	"conoha-cli/conoha/account"
	"fmt"

	"github.com/spf13/cobra"
)

var billingGetDepositCmd = &cobra.Command{
	Use:   "get-deposit",
	Short: "Get current deposit",
	Long:  `Get current deposit`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// 入金残高取得
		mgr := &account.BillingManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Account,
		}

		deposit, err2 := mgr.GetPaymentSummary()

		if err2 != nil {
			fmt.Printf("Error: Get payment deposit infomation failed. (%s)\n", err2)
			return
		}

		// 結果出力
		fmt.Printf("%d\n", deposit.TotalDepositAmount)
	},
}

func init() {
}
