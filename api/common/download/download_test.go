package download_test

import (
	"anime-list/common/download"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {

	params := download.DownloadParams{
		Path: "/tmp/temp-file-name.tmp",
		Url:  "https://cdn.xxacrvxx.ydns.eu/text.txt",
	}

	if err := download.Download(params); err != nil {
		t.Fatal(err)
	}

	t.Logf("downloaded file: %s", params.Path)

	if err := download.Download(params); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(params.Path); err != nil {
		t.Fatal(err)
	}
}
