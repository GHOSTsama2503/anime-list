package images

import (
	"anime-list/common/config"
	"anime-list/common/download"
	"fmt"
	"path"
)

type ImageParams struct {
	Id      int64
	CoverLg string
	CoverMd string
	CoverSm string
	Banner  string
}

func Download(params ImageParams) error {

	stParams := download.DownloadParams{
		Path: path.Join(config.Env.ImagesPath, "st", fmt.Sprintf("%d.png", params.Id)),
		Url:  fmt.Sprintf("https://img.anili.st/media/%d", params.Id),
	}

	if err := download.Download(stParams); err != nil {
		return err
	}

	bannerParams := download.DownloadParams{
		Path: path.Join(config.Env.ImagesPath, "banners", fmt.Sprintf("%d.jpeg", params.Id)),
		Url:  params.Banner,
	}

	if err := download.Download(bannerParams); err != nil {
		return err
	}

	lgParams := download.DownloadParams{
		Path: path.Join(config.Env.ImagesPath, "covers", "large", fmt.Sprintf("%d.jpeg", params.Id)),
		Url:  params.CoverLg,
	}

	if err := download.Download(lgParams); err != nil {
		return err
	}

	mdParams := download.DownloadParams{
		Path: path.Join(config.Env.ImagesPath, "covers", "medium", fmt.Sprintf("%d.jpeg", params.Id)),
		Url:  params.CoverMd,
	}

	if err := download.Download(mdParams); err != nil {
		return err
	}

	smParams := download.DownloadParams{
		Path: path.Join(config.Env.ImagesPath, "covers", "small", fmt.Sprintf("%d.jpeg", params.Id)),
		Url:  params.CoverSm,
	}

	if err := download.Download(smParams); err != nil {
		return err
	}

	return nil
}
