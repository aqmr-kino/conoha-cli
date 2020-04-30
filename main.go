package main

import "conoha-cli/cmd"

func main() {
	cmd.LoadGlobalConfigure()

	cmd.Execute()
}
