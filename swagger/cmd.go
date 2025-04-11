package swagger

import (
	"github.com/fzf-labs/api-import/swagger/apifox"
	"github.com/fzf-labs/api-import/swagger/yapi"
	"github.com/spf13/cobra"
)

var CmdSwagger = &cobra.Command{
	Use:   "swagger",
	Short: "swagger api import",
	Long:  "swagger api import. Example: api-import swagger",
}

//nolint:gochecknoinits
func init() {
	CmdSwagger.AddCommand(apifox.CmdAPIFox)
	CmdSwagger.AddCommand(yapi.CmdYaPi)
}
