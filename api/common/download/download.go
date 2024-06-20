package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/charmbracelet/log"
)

type DownloadParams struct {
	Path             string
	Url              string
	MaxRetries       int
	DownloadIfExists bool
}

func Download(params DownloadParams) error {

	dir := path.Dir(params.Path)

	// create file directory if not exists
	_, err := os.Stat(dir)
	if err != nil {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	var head *http.Response
	var headRetries int

	for true {
		head, err = http.Head(params.Url)
		if err == nil {
			break
		}

		if headRetries >= params.MaxRetries {
			return fmt.Errorf("max retries reached (%d): %v", params.MaxRetries, err)
		}

		headRetries++
	}

	fileInfo, err := os.Stat(params.Path)
	if err == nil && fileInfo.Size() == head.ContentLength && !params.DownloadIfExists {
		log.Warn("omiting download, file already exists.", "path", params.Path)
		return nil
	}

	file, err := os.Create(params.Path)
	if err != nil {
		return fmt.Errorf("creating file: %v", err)
	}

	defer file.Close()

	var resp *http.Response
	var respRetries int

	for true {
		resp, err = http.Get(params.Url)
		if err == nil {
			break
		}

		if respRetries >= params.MaxRetries {
			return fmt.Errorf("max retries reached (%d): %v", params.MaxRetries, err)
		}

		respRetries++
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("copying the file: %v", err)
	}

	return nil
}
