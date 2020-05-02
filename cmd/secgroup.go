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
	// secgroup list --detail
	// secgroup create (--name) aaa (--description) bbb
	// secgroup delete [id] [-f]
	// secgroup modify [id] --opt value...

	// secgroup addrule [id] --rule
	// secgroup deleterule [id] --ruleid
}
