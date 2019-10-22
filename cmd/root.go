package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
  "ngen.co.jp/cerve/cmd/verify"
  "os"
)

var rootCmd = &cobra.Command{
	Use:   "cerve",
	Long: `The Cerve is CLI tool which verify a TLS certificate from the outside of the server.`,
}

func init() {
  rootCmd.AddCommand(verify.NewVerifyCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
