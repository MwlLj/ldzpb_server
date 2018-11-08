package httpfile

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func eixsts(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createDir(root string) {
	if eixsts(root) == false {
		os.MkdirAll(root, 0777)
	}
}

func genUuidFile(file string) string {
	base := path.Base(file)
	ext := path.Ext(base)
	u, _ := uuid.NewV4()
	return strings.Join([]string{u.String(), ext}, "")
}

func DownloadFile(r *http.Request, formName string, dstRoot string, isUseUuid bool) ([]string, error) {
	createDir(dstRoot)
	var urls []string
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(1, err)
		return urls, err
	}
	m := r.MultipartForm
	files := m.File[formName]
	for _, f := range files {
		file, err := f.Open()
		defer file.Close()
		if err != nil {
			return urls, err
		}
		filename := f.Filename
		if isUseUuid == true {
			filename = genUuidFile(filename)
		}
		url := strings.Join([]string{dstRoot, filename}, "/")
		urls = append(urls, url)
		dst, err := os.Create(url)
		defer dst.Close()
		if err != nil {
			return urls, err
		}
		if _, err = io.Copy(dst, file); err != nil {
			return urls, err
		}
	}
	return urls, nil
}

func UploadFile(formname string, filename string, targetUrl string, response *string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(formname, filename)
	if err != nil {
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	request, err := http.NewRequest("POST", targetUrl, bodyBuf)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("cookieid", "70be01e6-cf71-11e8-93bc-3c970efa4c76")
	request.Header.Set("pictype", "test")
	// resp, err := http.Post(targetUrl, contentType, bodyBuf)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	*response = string(resp_body)
	return nil
}
