package tests

import (
	"Avito_trainee_assignment/internal/domain/model"
	"Avito_trainee_assignment/tests/client"
	"encoding/json"
	"github.com/brianvoe/gofakeit"
	"github.com/hashicorp/go-set"
	"io"
	"log"
	"math/rand"
	"net/http"
	"slices"
)

type ReqBody struct {
	Feature  int64                  `json:"feature_id"`
	Tags     []int64                `json:"tag_ids"`
	Content  map[string]interface{} `json:"content"`
	IsActive bool                   `json:"is_active"`
}

var (
	GetAllURL = client.BaseURL + "/banner"
)

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
		for _, t := range b.TagIds {
			tagsList = append(tagsList, t)
		}
	}
	tags := set.From[int64](tagsList)
	return tags
}

func SpareFeature() int64 {
	usedFeatures := FeaturesSet().Slice()
	if len(usedFeatures) > 200 {
		return int64(gofakeit.Uint32())
	}
	feature := int64(gofakeit.Uint32())
	for {
		if !slices.Contains(usedFeatures, feature) {
			return feature
		}
		feature = int64(gofakeit.Uint32())
	}
}

func SpareTags() []int64 {
	usedTags := TagsSet().Slice()

	var tags []int64
	for i := 0; i < 1+rand.Intn(3); i++ {
		tag := int64(gofakeit.Uint32())
		if len(usedTags) > 200 {
			tags = append(tags, tag)
		}
		if !slices.Contains(usedTags, tag) && tag > 0 {
			tags = append(tags, tag)
		}
	}
	return tags
}

func randomTags() []int64 {
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
	//just random silly values
	fields := []string{gofakeit.CurrencyShort(), gofakeit.Extension(), gofakeit.HackerVerb(), gofakeit.SSN(), gofakeit.Month()}
	values := []interface{}{gofakeit.BeerName(), gofakeit.Name(), gofakeit.Color(), gofakeit.Country(), gofakeit.Gender(), gofakeit.Int8()}

	content := make(map[string]interface{})
	for i := 0; i < 1+rand.Intn(5); i++ {
		content[fields[rand.Intn(len(fields))]] = values[rand.Intn(len(values))]
	}
	return content
}

func RandomBody() *ReqBody {
	return &ReqBody{
		Feature:  int64(gofakeit.Uint32()),
		Tags:     randomTags(),
		Content:  RandomContent(),
		IsActive: gofakeit.Bool(),
	}
}

func GetAll() []model.Banner {
	c := &http.Client{}
	req := client.FormRequest(http.MethodGet, GetAllURL, nil, client.AdminToken)
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	var banners []model.Banner
	err = json.Unmarshal(respBytes, &banners)
	if err != nil {
		log.Fatal(err)
	}
	return banners
}
