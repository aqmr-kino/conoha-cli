package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/compute"

	"github.com/spf13/cobra"
)

var vmListFlavorCmd = &cobra.Command{
	Use:   "list-flavor",
	Short: "list virtual machine flavors",
	Long:  `Listing virtual machine flavors`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			cmd.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// Flavor一覧情報取得
		mgr := &compute.VMManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Compute,
		}

		flav, err2 := mgr.GetFlavors()

		if err2 != nil {
			cmd.Printf("Error: Get virtual machine flavor infomation failed. (%s)\n", err2)
			return
		}

		cmd.Printf("%-16s %-36s %-9s %-7s %-8s\n", "Name", "ID", "CPU(Core)", "RAM(MB)", "Disk(GB)")
		for _, f := range flav.Flavors {
			cmd.Printf("%-16s %-36s %9d %7d %8d\n", f.Name, f.ID, f.CPUs, f.RAM, f.Disk)
		}
	},
}

func init() {

}
