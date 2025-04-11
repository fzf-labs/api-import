package apifox

import (
	"github.com/spf13/cobra"
)

var CmdAPIFox = &cobra.Command{
	Use:   "apifox",
	Short: "sync api documentation to apifox",
	Long:  "sync api documentation to apifox. Example: api-import swagger apifox",
	Run:   apiFoxRun,
}

var (
	apiFoxProjectId string
	apiFoxToken     string
	apiFoxInPutPath string
)

//nolint:gochecknoinits
func init() {
	CmdAPIFox.Flags().StringVarP(&apiFoxProjectId, "projectId", "p", "", "projectId")
	CmdAPIFox.Flags().StringVarP(&apiFoxToken, "token", "t", "", "token")
	CmdAPIFox.Flags().StringVarP(&apiFoxInPutPath, "inPutPath", "i", "./api", "inPutPath")
}

func apiFoxRun(_ *cobra.Command, _ []string) {
	NewAPIFox(apiFoxToken, apiFoxProjectId, apiFoxInPutPath).Run()
}
