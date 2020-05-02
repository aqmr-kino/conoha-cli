package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getEndpointCmd = &cobra.Command{
	Use:   "get-endpoint",
	Short: "Listing current Conoha API endpoint servers",
	Long:  `Listing current Conoha API endpoint servers`,
	Run: func(cmd *cobra.Command, args []string) {
		var showEndpoints *[]string

		// 表示内容設定
		if len(args) > 0 {
			showEndpoints = &args
		} else {
			showEndpoints = &[]string{
				"identity",
				"network",
			}
		}

		// APIエンドポイント情報表示
		for _, e := range *showEndpoints {
			switch e {
			case "identity":
				fmt.Printf("identity=%s\n", Configure.Endpoint.Idenity)
			case "network":
				fmt.Printf("netork=%s\n", Configure.Endpoint.Network)
			default:
				fmt.Printf("Error: unknown endpoint type \"" + e + "\"\n")
			}
		}
	},
}

func init() {
}