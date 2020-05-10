package cmd

import (
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "virtual machine management",
	Long:  `Management Conoha virtual machines`,
}

func init() {
	vmCmd.AddCommand(vmListCmd)
}
