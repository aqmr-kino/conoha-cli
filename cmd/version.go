package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version :
// バージョン情報
var Version string = "0.0.0"

// versionCmd :
// バージョン情報出力コマンド
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version infomation",
	Long:  `Show conoha-cli tool version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Conoha CLI Tool v%s", Version)
	},
}
