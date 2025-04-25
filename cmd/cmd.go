package cmd

import (
	"github.com/ondrejsika/hellogophercamptest2/cmd/root"
	_ "github.com/ondrejsika/hellogophercamptest2/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
