package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/refaldyrk/ttdl/constant"
	"github.com/refaldyrk/ttdl/dto"
)

type tiktok struct {
	Token string
	Id    string
}

type methodTiktok struct {
	tiktok *tiktok
}

func New() *tiktok {
	return &tiktok{}
}

func (t *tiktok) Get(uri string) *methodTiktok {
	var result dto.TikTokData

	urls := make(url.Values)
	urls.Set("url", uri)

	resp, err := http.PostForm(constant.API_BASE, urls)
	if err != nil {
		fmt.Println(err)

		return &methodTiktok{}
	}

	body := resp.Body

	defer resp.Body.Close()

	//Read Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)

		return &methodTiktok{}
	}

	//Unmarshal Json
	err = json.Unmarshal(byteBody, &result)
	if err != nil {
		fmt.Println(err)
		return &methodTiktok{}
	}

	//Set
	if result.ID != "" && result.Token != "" {
		t.Id = result.ID
		t.Token = result.Token
	}

	return &methodTiktok{tiktok: t}
}

func (m *methodTiktok) GetLink() []string {
	if m.tiktok.Id == "" && m.tiktok.Token == "" {
		return []string{}
	}

	linkDownload := []string{fmt.Sprintf("%s%s/%s.mp4?hd=1", constant.API_DOWNLOAD, m.tiktok.Token, m.tiktok.Id), fmt.Sprintf("%s%s/%s.mp4", constant.API_DOWNLOAD, m.tiktok.Token, m.tiktok.Id)}

	return linkDownload
}
