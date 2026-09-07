/*
Copyright © 2026 Peter Puczkowskyj
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/puczkowskyjp/seclens/internal/cli/jwt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seclens",
	Short: "Local security encoding and decoding toolkit",
	Long:  `Seclens is a local-first CLI for encoding, decoding, and inspecting security-related data.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(jwt.JwtCommand())
}
