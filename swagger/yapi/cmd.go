package yapi

import (
	"github.com/spf13/cobra"
)

var CmdYaPi = &cobra.Command{
	Use:   "yapi",
	Short: "sync api documentation to yapi",
	Long:  "sync api documentation to yapi. Example: api-import swagger yapi",
	Run:   YaPiRun,
}

var (
	yaPiURL       string
	yaPiToken     string
	yaPiInPutPath string
)

//nolint:gochecknoinits
func init() {
	CmdYaPi.Flags().StringVarP(&yaPiURL, "url", "u", "", "url")
	CmdYaPi.Flags().StringVarP(&yaPiToken, "token", "t", "", "token")
	CmdYaPi.Flags().StringVarP(&yaPiInPutPath, "inPutPath", "i", "./api", "inPutPath")
}

func YaPiRun(_ *cobra.Command, _ []string) {
	NewYaPi(yaPiToken, yaPiURL, yaPiInPutPath).Run()
}
