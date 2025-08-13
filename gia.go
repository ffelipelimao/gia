package main

import (
	"fmt"
	"os"

	"github.com/ffelipelimao/gia/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gia",
		Short: "Gia - A CLI tool for AI-driven task execution",
		Long:  `Gia is a CLI tool that uses AI to help with git commits `,
	}

	rootCmd.AddCommand(commands.NewCommitCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
