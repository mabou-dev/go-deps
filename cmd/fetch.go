package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	fetchOutput string
)

func FetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch dependencies",
		Long: `Fetch all dependencies from a project.
You can specify the path to scan and the output folder for fetched dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runFetch(); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(&fetchOutput, "output", "text", "output folder for fetched dependencies")

	return cmd
}

func runFetch() error {
	log.Debug("method runFetch is not implemented")
	panic("not implemented")
}
