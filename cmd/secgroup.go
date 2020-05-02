package cmd

import (
	"github.com/spf13/cobra"
)

var secgroupCmd = &cobra.Command{
	Use:   "secgroup",
	Short: "security group management",
	Long:  `Management security group and rule`,
}

func init() {
	secgroupCmd.AddCommand(secgroupListCmd)
	secgroupCmd.AddCommand(secgroupListRuleCmd)
}
