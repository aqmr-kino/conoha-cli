package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/network"

	"github.com/spf13/cobra"
)

type secgroupListOptions struct {
	DetailMode bool
}

var (
	secgroupListOpts = &secgroupListOptions{}
)

var secgroupListCmd = &cobra.Command{
	Use:   "list",
	Short: "list security groups",
	Long:  `Listing security groups`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			cmd.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// セキュリティグループ取得
		mgr := &network.SecurityGroupManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Network,
		}

		sg, err2 := mgr.GetGroups()

		if err2 != nil {
			cmd.Printf("Error: Get security group infomation failed. (%s)\n", err2)
			return
		}

		// 結果出力
		for _, s := range sg.SecurityGroups {
			if len(args) == 0 || includes(s.Name, args) || includes(s.ID, args) {
				if secgroupListOpts.DetailMode {
					cmd.Printf("SecurityGroup: %s\n", s.Name)
					cmd.Printf("ID           : %s\n", s.ID)
					cmd.Printf("Description  : %s\n", s.Description)
					cmd.Printf("Rules        :\n")

					for _, r := range s.Rules {
						cmd.Printf("  %s %-7s %s %d-%d %s", r.ID, r.Direction, r.Ethertype, r.PortRangeMin, r.PortRangeMax, r.Protocol)

						if len(r.RemoteIPPrefix) != 0 {
							cmd.Printf(" (from: %s)", r.RemoteIPPrefix)
						}

						if len(r.RemoteGroupID) != 0 {
							cmd.Printf(" (from-secgroup: %s)", r.RemoteGroupID)
						}

						cmd.Printf("\n")
					}

					cmd.Println()
				} else {
					cmd.Printf("%s %s\n", s.Name, s.ID)
				}
			}
		}

	},
}

func init() {
	secgroupListCmd.Flags().BoolVarP(&secgroupListOpts.DetailMode, "detail", "d", false, "list infomation with details")
}
