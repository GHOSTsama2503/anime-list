package anilist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const ApiUrl = "https://graphql.anilist.co"

func SearchAnime(title string, page uint, perPage uint) ([]SearchResult, error) {
	result := make([]SearchResult, 0)

	vars := SearchRequestVariables{

		Search:  title,
		Page:    page,
		PerPage: perPage,
	}

	reqBody := SearchRequest{
		Query:     QuerySearch,
		Variables: vars,
	}

	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return result, err
	}

	res, err := http.Post(ApiUrl, "application/json", bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return result, err
	}

	if res.StatusCode != http.StatusOK {
		msg := "anilist response with status code"

		content, err := io.ReadAll(res.Body)
		if err != nil {
			return result, fmt.Errorf("%s %d and unreadable response body: %v", msg, res.StatusCode, err)
		}

		return result, fmt.Errorf("%s %d: %v", msg, res.StatusCode, string(content))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("error reading anilist search response: %v", err)
	}

	var data SearchResponse
	if err := json.Unmarshal(resBody, &data); err != nil {
		return result, fmt.Errorf("error unmarshing anilist search response: %v", err)
	}

	for _, mediaElement := range data.Data.Page.Media {
		element := SearchResult{}
		element.Id = mediaElement.Id
		element.Title = mediaElement.Title.Romaji
		element.Image = mediaElement.CoverImage

		result = append(result, element)
	}

	return result, nil
}

func GetAnime(id uint) (Anime, error) {
	anime := Anime{}

	data := GetAnimeRequest{}
	data.Query = QueryInfo
	data.Variables.Id = id

	reqBody, err := json.Marshal(data)
	if err != nil {
		return anime, err
	}

	res, err := http.Post(ApiUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return anime, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return anime, err
	}

	resData := &GetAnimeResponse{}

	if err := json.Unmarshal(resBody, resData); err != nil {
		return anime, err
	}

	if resData.Data.Media.Id == 0 {
		return anime, errors.New("anime not found")
	}

	return resData.Data.Media, nil
}
