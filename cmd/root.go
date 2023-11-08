package cmd

import (
	"fmt"
	"os"
	"tree-test/tree"
)
import "github.com/spf13/cobra"

func rootCmd() *cobra.Command {
	r := &cobra.Command{
		Use:   "tree-test",
		Short: "merkle tree cli tree",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			fileName := args[0]

			size, err := cmd.Flags().GetInt64("size")
			if err != nil {
				panic(err)
			}

			tree.BuildTree(fileName, size)

		},
	}

	r.Flags().Int64P("size", "s", 10240, "the size of the chunks to use for generating the tree")

	return r
}

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
