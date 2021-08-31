package command

import (
	"github.com/pyroscope-io/pyroscope/benchmark/config"
	"github.com/pyroscope-io/pyroscope/benchmark/loadgen"
	"github.com/pyroscope-io/pyroscope/pkg/cli"
	"github.com/spf13/cobra"
)

func newLoadGen(cfg *config.LoadGen) *cobra.Command {
	vpr := newViper()
	loadgenCmd := &cobra.Command{
		Use:   "loadgen [flags]",
		Short: "Generates load",
		RunE: createCmdRunFn(cfg, vpr, false, func(cmd *cobra.Command, args []string, logger config.LoggerFunc) error {
			return loadgen.Cli(cfg)
		}),
	}

	cli.PopulateFlagSet(cfg, loadgenCmd.Flags(), vpr)
	return loadgenCmd
}
