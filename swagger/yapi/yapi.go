package yapi

import (
	"errors"
	"log"
	"net/http"

	"github.com/fzf-labs/api-import/utils"
	"github.com/go-resty/resty/v2"
)

type YaPi struct {
	token     string
	url       string
	inPutPath string
}

func NewYaPi(token, url, inPutPath string) *YaPi {
	return &YaPi{
		token:     token,
		url:       url,
		inPutPath: inPutPath,
	}
}

func (y *YaPi) Run() {
	// 查询指定目录下的swagger的json文件
	files, err := utils.ReadDirFilesWithSuffix(y.inPutPath, ".swagger.json")
	if err != nil {
		log.Print(err.Error())
		return
	}
	if len(files) == 0 {
		log.Print("no swagger json file found")
		return
	}
	for _, v := range files {
		// 读取文件内容
		toString, err2 := utils.ReadFileToString(v)
		if err2 != nil {
			log.Print(err2.Error())
			return
		}
		err2 = y.syncYaPi(y.token, toString)
		if err2 != nil {
			log.Printf("syncYaPi err:%s", err2.Error())
		}
	}
	log.Println("yapi sync successful")
}

// syncYaPi 同步api文档到yapi
func (y *YaPi) syncYaPi(token, data string) error {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	pathParams := map[string]string{
		"type":  "swagger",
		"merge": "merge",
		"token": token,
		"json":  data,
	}
	resp, err := resty.New().R().SetHeaders(headers).SetFormData(pathParams).Post(y.url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return errors.New("http request err")
	}
	return nil
}
