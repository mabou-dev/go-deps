package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List dependencies",
		Long: `List all dependencies from a project.
You can specify the path to scan.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runList(); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func runList() error {
	err := errors.New("method runList is not implemented")
	log.Error(err.Error())
	return err
}
