package cmd

import (
	"conoha-cli/conoha/account"
	"fmt"

	"github.com/spf13/cobra"
)

var billingListItemCmd = &cobra.Command{
	Use:   "list-item",
	Short: "Listing Conoha items",
	Long:  `Listing Conoha items`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
		}

		// アイテム取得
		mgr := &account.BillingManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Account,
		}

		item, err2 := mgr.GetItems()

		if err2 != nil {
			fmt.Printf("Error: Get Conoha items infomation failed. (%s)\n", err2)
		}

		// 結果出力
		fmt.Printf("%-36s %-15s %-10s %s\n", "UUID", "ServiceName", "Status", "StartAt")

		for _, i := range item.OrderItems {
			fmt.Printf("%-36s %-15s %-10s %s\n", i.UUID, i.ServiceName, i.ItemStatus, i.ServiceStartDate)
		}
	},
}

func init() {
}
