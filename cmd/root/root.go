package root

import (
	"github.com/ondrejsika/hellogophercamptest2/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "hellogophercamptest2",
	Short: "hellogophercamptest2, " + version.Version,
}
