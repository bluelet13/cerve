package verify

import (
	"fmt"
	"ngen.co.jp/cerve/pkg/cerve"
	"os"

	"github.com/spf13/cobra"
)

var (
	verifyExample = `  # From the outside of the server, verify TLS certificate using define file
  cerve verify <define-file>`
)

func NewVerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:                   "verify <define-file>",
		DisableFlagsInUseLine: true,
		Short:                 "Verify TLS certificate",
		Long:                  `Verify TLS certificate from the outside of the server`,
		Example:               verifyExample,
		Run:                   runCommand,
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	option := cerve.Option{}
	c, err := cerve.NewCerve(args, option)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	v, err := c.Verify()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(v.JSON())
}
