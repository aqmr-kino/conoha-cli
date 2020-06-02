package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/network"

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
			cmd.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// セキュリティグループ ルール取得
		mgr := &network.SecurityGroupManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Network,
		}

		sg, err2 := mgr.GetRules()

		if err2 != nil {
			cmd.Printf("Error: Get security group rule infomation failed. (%s)\n", err2)
			return
		}

		// 結果出力
		if !secgroupListRuleOpts.DetailMode {
			cmd.Printf("%-36s %-9s %-9s %-11s %8s %s\n", "ID", "Direction", "Ethertype", "PortRange", "Protocol", "(from)")
		}

		for _, r := range sg.SecurityGroupRules {
			if len(args) == 0 || includes(r.ID, args) {
				if secgroupListRuleOpts.DetailMode {
					cmd.Printf("ID              : %s\n", r.ID)
					cmd.Printf("Ethertype       : %s\n", r.Ethertype)
					cmd.Printf("Protocol        : %s\n", r.Protocol)
					cmd.Printf("Port Range      : %d-%d\n", r.PortRangeMin, r.PortRangeMax)
					cmd.Printf("Direction       : %s\n", r.Direction)
					cmd.Printf("Remote IP       : %s\n", r.RemoteIPPrefix)
					cmd.Printf("Remote SecGroup : %s\n", r.RemoteGroupID)
					cmd.Printf("Security Group  : %s\n", r.SecurityGroupID)
					cmd.Printf("Tenant ID       : %s\n", r.TenantID)
					cmd.Println()
				} else {
					cmd.Printf("%-36s %9s %9s %5d-%5d %8s", r.ID, r.Direction, r.Ethertype, r.PortRangeMin, r.PortRangeMax, r.Protocol)
					if len(r.RemoteIPPrefix) != 0 {
						cmd.Printf(" (ip: %s)", r.RemoteIPPrefix)
					}

					if len(r.RemoteGroupID) != 0 {
						cmd.Printf(" (secgroup: %s)", r.RemoteGroupID)
					}

					cmd.Printf("\n")
				}
			}
		}
	},
}

func init() {
	secgroupListRuleCmd.Flags().BoolVarP(&secgroupListRuleOpts.DetailMode, "detail", "d", false, "list infomation with details")
}
