package main

import (
	"github.com/fzf-labs/api-import/swagger"
	"github.com/spf13/cobra"
)

var CmdAPIImport = &cobra.Command{
	Use:   "api-import",
	Short: "api import",
	Long:  "api import. Example: api-import",
}

//nolint:gochecknoinits
func init() {
	CmdAPIImport.AddCommand(swagger.CmdSwagger)
}
