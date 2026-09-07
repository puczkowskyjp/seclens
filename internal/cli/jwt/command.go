/*
Copyright © 2026 Peter Puczkowskyj
*/
package jwt

import (
	"github.com/spf13/cobra"
)

func JwtCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jwt",
		Short: "Commands for working with JWTs",
		Long:  `The jwt command allows you to encode, decode, and inspect JSON Web Tokens (JWTs).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return cmd
}
