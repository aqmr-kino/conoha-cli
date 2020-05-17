package cmd

import (
	"conoha-cli/conoha/account"
	"conoha-cli/conoha/compute"
	"fmt"

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

		if vmListOpts.DetailMode {
			for _, s := range vm.Servers {
				flav, _ := mgr.FindFlavor(s.Flavor.ID)
				img, _ := mgr.FindVMImage(s.Image.ID)

				fmt.Printf("Virtual Machine: %s (%s)\n", s.Metadata.InstanceNameTag, s.ID)
				fmt.Printf("Status: %s\n", s.Status)

				fmt.Printf("IP Addresses:\n")
				for _, v := range s.Addresses {
					for _, a := range v {
						fmt.Printf("  %s\n", a.Addr)
					}
				}

				fmt.Printf("Security Groups:\n")
				for _, sg := range s.SecurityGroups {
					fmt.Printf("  %s\n", sg.Name)
				}

				fmt.Printf("Flavor: %s (%d Core(s) CPU, %d MB RAM, %d GB SSD)\n", flav.Name, flav.CPUs, flav.RAM, flav.Disk)
				fmt.Printf("Base Image: %s\n", img.Name)
				fmt.Printf("Created At: %s\n", s.Created)
				fmt.Printf("Updated At: %s\n", s.Updated)
				fmt.Println()
			}
		} else {
			fmt.Printf("%-20s %-36s %-8s\n", "Name", "ID", "Status")
			for _, s := range vm.Servers {
				fmt.Printf("%-20s %-36s %-8s\n", s.Metadata.InstanceNameTag, s.ID, s.Status)
			}
		}
	},
}

func init() {
	vmListCmd.Flags().BoolVarP(&vmListOpts.DetailMode, "detail", "d", false, "list infomation with details")
}
