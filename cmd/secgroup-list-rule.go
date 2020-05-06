package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/network"
	"fmt"

	"github.com/spf13/cobra"
)

type secgroupListRuleOptions struct {
	DetailMode bool
}

var (
	secgroupListRuleOpts = &secgroupListRuleOptions{}
)

var secgroupListRuleCmd = &cobra.Command{
	Use:   "list-rule",
	Short: "list security group rules",
	Long:  `Listing security group rules`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// セキュリティグループ ルール取得
		mgr := &network.SecurityGroupManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Network,
		}

		sg, err2 := mgr.GetRules()

		if err2 != nil {
			fmt.Printf("Error: Get security group rule infomation failed. (%s)\n", err2)
			return
		}

		// 結果出力
		for _, r := range sg.SecurityGroupRules {
			if len(args) == 0 || includes(r.ID, args) {
				if secgroupListRuleOpts.DetailMode {
					fmt.Printf("ID              : %s\n", r.ID)
					fmt.Printf("Ethertype       : %s\n", r.Ethertype)
					fmt.Printf("Protocol        : %s\n", r.Protocol)
					fmt.Printf("Port Range      : %d-%d\n", r.PortRangeMin, r.PortRangeMax)
					fmt.Printf("Direction       : %s\n", r.Direction)
					fmt.Printf("Remote IP       : %s\n", r.RemoteIPPrefix)
					fmt.Printf("Remote SecGroup : %s\n", r.RemoteGroupID)
					fmt.Printf("Security Group  : %s\n", r.SecurityGroupID)
					fmt.Printf("Tenant ID       : %s\n", r.TenantID)
					fmt.Println()
				} else {
					fmt.Printf("%s %-7s %s %d-%d %s", r.ID, r.Direction, r.Ethertype, r.PortRangeMin, r.PortRangeMax, r.Protocol)
					if len(r.RemoteIPPrefix) != 0 {
						fmt.Printf(" (from: %s)", r.RemoteIPPrefix)
					}

					if len(r.RemoteGroupID) != 0 {
						fmt.Printf(" (from-secgroup: %s)", r.RemoteGroupID)
					}

					fmt.Printf("\n")
				}
			}
		}
	},
}

func init() {
	secgroupListRuleCmd.Flags().BoolVarP(&secgroupListRuleOpts.DetailMode, "detail", "d", false, "list infomation with details")
}
