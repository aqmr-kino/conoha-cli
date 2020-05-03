package cmd

import (
	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "billing management",
	Long:  `Management Conoha items and billing`,
}

func init() {
	billingCmd.AddCommand(billingListItemCmd)
	billingCmd.AddCommand(billingListInvoiceCmd)
}
