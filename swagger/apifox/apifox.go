package apifox

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fzf-labs/api-import/utils"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type APIFox struct {
	token     string
	projectId string
	inPutPath string
}

func NewAPIFox(token, projectId, inPutPath string) *APIFox {
	return &APIFox{
		token:     token,
		projectId: projectId,
		inPutPath: inPutPath,
	}
}

func (a *APIFox) Run() {
	// 查询指定目录下的swagger的json文件
	files, err := utils.ReadDirFilesWithSuffix(a.inPutPath, ".swagger.json")
	if err != nil {
		log.Print(err.Error())
		return
	}
	if len(files) == 0 {
		log.Print("no swagger json file found")
		return
	}
	// 读取文件夹
	foxFolders, err := a.getAPIFoxFolders(a.token, a.projectId)
	if err != nil {
		return
	}
	for _, v := range files {
		// 读取文件内容
		toString, err2 := utils.ReadFileToString(v)
		if err2 != nil {
			log.Print(err2.Error())
			return
		}
		title := gjson.Get(toString, "info.title").String()
		apiFoxFolder := strings.Split(filepath.Dir(filepath.FromSlash(title)), "/")[0]
		apiFolderID, ok := foxFolders[apiFoxFolder]
		if !ok {
			apiFolderID, err = a.createAPIFoxFolders(a.token, a.projectId, apiFoxFolder)
			if err != nil {
				log.Print(err.Error())
				return
			}
			foxFolders[apiFoxFolder] = apiFolderID
		}
		err2 = a.syncAPIFox(a.token, a.projectId, apiFolderID, toString)
		if err2 != nil {
			log.Printf("syncAPIFox err2:%s", err2.Error())
			return
		}
	}
	log.Print("apifox sync successful")
}

// syncAPIFox 同步api文档到apifox
func (a *APIFox) syncAPIFox(token, projectId, apiFolderID, data string) error {
	url := fmt.Sprintf("https://api.apifox.com/api/v1/projects/%s/import-data", projectId)
	type APIFoxHTTPParam struct {
		// 导入数据格式，目前只支持`openapi`，表示 Swagger 或 OpenAPI 格式
		ImportFormat string `json:"importFormat"`
		// 要导入的数据，Swagger（OpenAPI） 格式 json 字符串，支持 OpenAPI 3、Swagger 1、2、3 数据格式
		Data string `json:"data"`
		// 导入到目标目录的ID，不传表示导入到根目录
		APIFolderID int `json:"apiFolderId,omitempty"`
		// 覆盖模式，匹配到相同接口时的覆盖模式，不传表示忽略
		APIOverwriteMode string `json:"apiOverwriteMode,omitempty"`
		// 是否在接口路径加上basePath，建议不传，即为 false，推荐将 BasePath 放到环境里的”前置 URL“里
		ImportBasePath bool `json:"importBasePath,omitempty"`
		// 覆盖模式，匹配到相同数据模型时的覆盖模式，不传表示忽略
		SchemaOverwriteMode string `json:"schemaOverwriteMode,omitempty"`
		// 是否同步更新接口所在目录
		SyncAPIFolder bool `json:"syncApiFolder,omitempty"`
	}
	headers := map[string]string{
		"X-Apifox-Version": "2022-11-16",
		"Authorization":    fmt.Sprintf("Bearer %s", token),
		"Content-Type":     "application/json",
		"User-Agent":       "resty/1.0.0",
		"Host":             "api.apifox.com",
	}
	folderID, _ := strconv.Atoi(apiFolderID)
	body := APIFoxHTTPParam{
		ImportFormat:        "openapi",
		Data:                data,
		APIFolderID:         folderID,
		APIOverwriteMode:    "methodAndPath",
		SchemaOverwriteMode: "name",
	}
	resp, err := resty.New().R().SetHeaders(headers).SetBody(body).Post(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Println(resp.String())
		return errors.New("http request err")
	}
	return nil
}

// getAPIFoxFolders 获取apifox的文件夹
func (a *APIFox) getAPIFoxFolders(token, projectId string) (map[string]string, error) {
	reply := make(map[string]string)
	url := fmt.Sprintf("https://api.apifox.com/api/v1/projects/%s/api-folders", projectId)
	headers := map[string]string{
		"X-Apifox-Version": "2022-11-16",
		"Authorization":    fmt.Sprintf("Bearer %s", token),
		"Content-Type":     "application/json",
		"User-Agent":       "resty/1.0.0",
		"Host":             "api.apifox.com",
	}
	resp, err := resty.New().R().SetHeaders(headers).Get(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, errors.New("http request err")
	}
	gjson.Get(resp.String(), "data").ForEach(func(key, value gjson.Result) bool {
		reply[value.Get("name").String()] = value.Get("id").String()
		return true
	})
	return reply, nil
}

// createAPIFoxFolders 创建apifox的文件夹
func (a *APIFox) createAPIFoxFolders(token, projectId, path string) (string, error) {
	url := fmt.Sprintf("https://api.apifox.com/api/v1/projects/%s/api-folders", projectId)
	headers := map[string]string{
		"X-Apifox-Version": "2022-11-16",
		"Authorization":    fmt.Sprintf("Bearer %s", token),
		"Content-Type":     "application/x-www-form-urlencoded",
		"User-Agent":       "resty/1.0.0",
		"Host":             "api.apifox.com",
	}
	resp, err := resty.New().R().SetHeaders(headers).SetFormData(map[string]string{
		"name":     path,
		"parentId": "0",
	}).Post(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", errors.New("http request err")
	}
	return gjson.Get(resp.String(), "data.id").String(), nil
}
