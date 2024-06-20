package images_test

import (
	"anime-list/common/config"
	"anime-list/services/images"
	"testing"
)

func TestDownload(t *testing.T) {

	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	imageParams := images.ImageParams{
		Id:      20,
		CoverLg: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx20792-Q53sZsUAh5FF.jpg",
		CoverMd: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/medium/bx20792-Q53sZsUAh5FF.jpg",
		CoverSm: "https://s4.anilist.co/file/anilistcdn/media/anime/cover/small/bx20792-Q53sZsUAh5FF.jpg",
		Banner:  "https://s4.anilist.co/file/anilistcdn/media/anime/banner/103572-qRtBguYOOR2j.jpg",
	}

	if err := images.Download(imageParams); err != nil {
		t.Fatal(err)
	}
}
