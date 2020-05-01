package cmd

import (
	"conoha-cli/conoha/account"
	"fmt"

	"github.com/spf13/cobra"
)

type configOptions struct {
	User     string
	Pass     string
	Tenant   string
	Endpoint string
	Testmode bool
}

var (
	configOpts = &configOptions{}
)

// configCmd :
// 設定コマンド
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "account configuration",
	Long:  `account configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Printf("user=%s\n", Configure.Credential.Auth.PasswordCredentials.Username)
			fmt.Printf("pass=%s\n", Configure.Credential.Auth.PasswordCredentials.Password)
			fmt.Printf("tenant=%s\n", Configure.Credential.Auth.TenantID)
			fmt.Printf("endpoint=%s\n", Configure.Endpoint)
		} else {
			if cmd.Flag("user").Changed {
				Configure.Credential.Auth.PasswordCredentials.Username = configOpts.User
			}
			if cmd.Flag("pass").Changed {
				Configure.Credential.Auth.PasswordCredentials.Password = configOpts.Pass
			}
			if cmd.Flag("tenent").Changed {
				Configure.Credential.Auth.TenantID = configOpts.Tenant
			}
			if cmd.Flag("endpoint").Changed {
				Configure.Endpoint = configOpts.Endpoint
			}

			if configOpts.Testmode {
				_, err := account.GetToken(Configure.Endpoint, Configure.Credential)

				if err != nil {
					fmt.Printf("failed (%s)\n", err)
				} else {
					fmt.Printf("OK\n")
				}
			}

			err := Configure.SaveAs(ConfigFilename)

			if err != nil {
				fmt.Println("config file save error")
			}
		}
	},
}

func init() {
	configCmd.Flags().StringVarP(&configOpts.User, "user", "u", "", "set (or change) username")
	configCmd.Flags().StringVarP(&configOpts.Pass, "pass", "p", "", "set (or change) password")
	configCmd.Flags().StringVarP(&configOpts.Tenant, "tenent", "t", "", "set (or change) tenant-id")
	configCmd.Flags().StringVarP(&configOpts.Endpoint, "endpoint", "e", "", "set (or change) endpoint api")
	configCmd.Flags().BoolVarP(&configOpts.Testmode, "verify", "v", false, "verifing configure by connecting to token API server with using account crediential")

	configCmd.MarkPersistentFlagRequired("user")
	configCmd.MarkPersistentFlagRequired("pass")
	configCmd.MarkPersistentFlagRequired("tenant")
	configCmd.MarkPersistentFlagRequired("endpoint")
}
