package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conoha-cli",
	Short: "Conoha API Operation CLI Tool",
	Long:  `Conoha API Operation CLI Tool`,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(getEndpointCmd)
	rootCmd.AddCommand(setEndpointCmd)
	rootCmd.AddCommand(secgroupCmd)
	rootCmd.AddCommand(billingCmd)
	rootCmd.AddCommand(vmCmd)
}

// Execute :
// CLIコマンド関数エントリーポイント
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
