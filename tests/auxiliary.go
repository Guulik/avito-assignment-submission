package tests

import (
	"Banner_Infrastructure/internal/domain/model"
	"Banner_Infrastructure/tests/client"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"slices"

	"github.com/brianvoe/gofakeit"
	"github.com/hashicorp/go-set"
)

type ReqBody struct {
	Feature  int64                  `json:"feature_id"`
	Tags     []int64                `json:"tag_ids"`
	Content  map[string]interface{} `json:"content"`
	IsActive bool                   `json:"is_active"`
}

var GetAllURL = client.BaseURL + "/banner"

func GetFirstInactive() *model.Banner {
	banners := GetAll()
	for _, b := range banners {
		if !b.IsActive {
			return &b
		}
	}
	return nil
}

func GetRandomActive() *model.Banner {
	banners := GetAll()
	for {
		b := banners[rand.Intn(len(banners))]
		if b.IsActive {
			return &b
		}
	}
}

func FeaturesSet() *set.Set[int64] {
	banners := GetAll()
	var featuresList []int64
	for _, b := range banners {
		featuresList = append(featuresList, b.FeatureId)
	}

	features := set.From[int64](featuresList)
	return features
}

func TagsSet() *set.Set[int64] {
	banners := GetAll()
	var tagsList []int64
	for _, b := range banners {
		tagsList = append(tagsList, b.TagIds...)
	}
	tags := set.From[int64](tagsList)
	return tags
}

func RandomTags() []int64 {
	var tags []int64
	for i := 0; i < 1+rand.Intn(3); i++ {
		tag := int64(gofakeit.Uint32())
		if !slices.Contains(tags, tag) && tag > 0 {
			tags = append(tags, tag)
		}
	}
	return tags
}

func RandomContent() map[string]interface{} {
	// just random silly values
	fields := []string{gofakeit.CurrencyShort(), gofakeit.Extension(), gofakeit.HackerVerb(), gofakeit.SSN(), gofakeit.Month()}
	values := []interface{}{gofakeit.BeerName(), gofakeit.Name(), gofakeit.Color(), gofakeit.Country(), gofakeit.Gender(), gofakeit.Int8()}

	content := make(map[string]interface{})
	for i := 0; i < 1+rand.Intn(5); i++ {
		content[fields[rand.Intn(len(fields))]] = values[rand.Intn(len(values))]
	}
	return content
}

func RandomBody(isActive bool) *ReqBody {
	return &ReqBody{
		Feature:  int64(gofakeit.Uint32()),
		Tags:     RandomTags(),
		Content:  RandomContent(),
		IsActive: isActive,
	}
}

func GetAll() []model.Banner {
	c := &http.Client{}
	req := client.FormRequest(http.MethodGet, GetAllURL, nil, client.AdminToken)
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var banners []model.Banner
	err = json.Unmarshal(respBytes, &banners)
	if err != nil {
		log.Fatal(err)
	}
	return banners
}

func GetPostedId(body *ReqBody) int64 {
	c := &http.Client{}
	url := GetAllURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req := client.FormRequest(http.MethodGet, url, nil, client.AdminToken)
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	resBytes, _ := io.ReadAll(resp.Body)
	var postedBanner []model.Banner
	err = json.Unmarshal(resBytes, &postedBanner)
	if err != nil {
		log.Fatal(err)
	}
	return postedBanner[0].ID
}
