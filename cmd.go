package main

import (
	"log"

	"github.com/fzf-labs/api-import/swagger"
	"github.com/spf13/cobra"
)

var CmdAPIImport = &cobra.Command{
	Use:   "api-import",
	Short: "api import",
	Long:  "api import. Example: api-import",
}

func init() {
	CmdAPIImport.AddCommand(swagger.CmdSwagger)
}

func main() {
	if err := CmdAPIImport.Execute(); err != nil {
		log.Fatal(err)
	}
}
