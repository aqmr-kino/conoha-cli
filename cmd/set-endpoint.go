package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type setEndpointOptions struct {
	Account       string
	Compute       string
	Volume        string
	Database      string
	Image         string
	DNS           string
	ObjectStorage string
	Mail          string
	Idenity       string
	Network       string
}

var (
	setEndpointOpts = &setEndpointOptions{}
)

var setEndpointCmd = &cobra.Command{
	Use:   "set-endpoint",
	Short: "set (or change) Conoha API endpoint servers",
	Long:  `set (or change) Conoha API endpoint servers`,
	Args: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().NFlag() == 0 {
			return errors.New("requires set any flags")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 0 {
			if cmd.Flag("network").Changed {
				Configure.Endpoint.Network = setEndpointOpts.Network
			}
			if cmd.Flag("identity").Changed {
				Configure.Endpoint.Idenity = setEndpointOpts.Idenity
			}
			if cmd.Flag("account").Changed {
				Configure.Endpoint.Account = setEndpointOpts.Account
			}
			if cmd.Flag("compute").Changed {
				Configure.Endpoint.Compute = setEndpointOpts.Compute
			}

			Configure.SaveAs(ConfigFilename)
		}
	},
}

func init() {
	setEndpointCmd.Flags().StringVarP(&setEndpointOpts.Network, "network", "n", "", "set (or change) Conoha network service API endpoint")
	setEndpointCmd.Flags().StringVarP(&setEndpointOpts.Idenity, "identity", "i", "", "set (or change) Conoha identity service API endpoint")
	setEndpointCmd.Flags().StringVarP(&setEndpointOpts.Account, "account", "a", "", "set (or change) Conoha account service API endpoint")
	setEndpointCmd.Flags().StringVarP(&setEndpointOpts.Compute, "compute", "c", "", "set (or change) Conoha compute service API endpoint")

	setEndpointCmd.MarkPersistentFlagRequired("network")
	setEndpointCmd.MarkPersistentFlagRequired("identity")
	setEndpointCmd.MarkPersistentFlagRequired("account")
	setEndpointCmd.MarkPersistentFlagRequired("compute")
}
