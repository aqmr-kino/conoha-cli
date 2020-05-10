package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/compute"
	"fmt"

	"github.com/spf13/cobra"
)

var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "list virtual machines",
	Long:  `Listing virtual machines`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			fmt.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// VM取得
		mgr := &compute.VMManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Compute,
		}

		vm, err2 := mgr.GetVirtualMachines()

		if err2 != nil {
			fmt.Printf("Error: Get virtual machine infomation failed. (%s)\n", err2)
			return
		}

		fmt.Printf("%-20s %-36s %-8s\n", "Name", "ID", "Status")
		for _, s := range vm.Servers {
			fmt.Printf("%-20s %-36s %-8s\n", s.Metadata.InstanceNameTag, s.ID, s.Status)
		}

	},
}

func init() {

}
