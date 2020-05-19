package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/compute"

	"github.com/spf13/cobra"
)

type vmListOptions struct {
	DetailMode bool
}

var (
	vmListOpts = &vmListOptions{}
)

var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "list virtual machines",
	Long:  `Listing virtual machines`,
	Run: func(cmd *cobra.Command, args []string) {
		// API トークン取得
		token, err := account.GetToken(Configure.Endpoint.Idenity, Configure.Credential)

		if err != nil {
			cmd.Printf("Error: Get API token failed. (%s)\n", err)
			return
		}

		// VM取得
		mgr := &compute.VMManager{
			Token:    token,
			Endpoint: Configure.Endpoint.Compute,
		}

		vm, err2 := mgr.GetVirtualMachines()

		if err2 != nil {
			cmd.Printf("Error: Get virtual machine infomation failed. (%s)\n", err2)
			return
		}

		if vmListOpts.DetailMode {
			for _, s := range vm.Servers {
				flav, _ := mgr.FindFlavor(s.Flavor.ID)
				img, _ := mgr.FindVMImage(s.Image.ID)

				cmd.Printf("Virtual Machine: %s (%s)\n", s.Metadata.InstanceNameTag, s.ID)
				cmd.Printf("Status: %s\n", s.Status)

				cmd.Printf("IP Addresses:\n")
				for _, v := range s.Addresses {
					for _, a := range v {
						cmd.Printf("  %s\n", a.Addr)
					}
				}

				cmd.Printf("Security Groups:\n")
				for _, sg := range s.SecurityGroups {
					cmd.Printf("  %s\n", sg.Name)
				}

				cmd.Printf("Flavor: %s (%d Core(s) CPU, %d MB RAM, %d GB SSD)\n", flav.Name, flav.CPUs, flav.RAM, flav.Disk)
				cmd.Printf("Base Image: %s\n", img.Name)
				cmd.Printf("Created At: %s\n", s.Created)
				cmd.Printf("Updated At: %s\n", s.Updated)
				cmd.Println()
			}
		} else {
			cmd.Printf("%-20s %-36s %-8s\n", "Name", "ID", "Status")
			for _, s := range vm.Servers {
				cmd.Printf("%-20s %-36s %-8s\n", s.Metadata.InstanceNameTag, s.ID, s.Status)
			}
		}
	},
}

func init() {
	vmListCmd.Flags().BoolVarP(&vmListOpts.DetailMode, "detail", "d", false, "list infomation with details")
}
