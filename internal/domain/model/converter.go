package model

import (
	"encoding/json"
	"strconv"
	"strings"
)

func BannersDbToService(bannersDB []BannerDB) ([]Banner, error) {
	result := make([]Banner, len(bannersDB))

	for i, bDB := range bannersDB {
		b, err := toBanner(bDB)
		if err != nil {
			return nil, err
		}

		result[i] = b
	}
	return result, nil
}

func toBanner(bDB BannerDB) (Banner, error) {
	b := Banner{
		ID:        bDB.ID,
		FeatureId: bDB.FeatureId,
		IsActive:  bDB.IsActive,
		CreatedAt: bDB.CreatedAt,
		UpdatedAt: bDB.UpdatedAt,
	}

	var err error
	b.TagIds, err = ParseIntArrayFromString(bDB.TagIds)
	if err != nil {
		return Banner{}, err
	}

	err = json.Unmarshal(bDB.Content, &b.Content)
	if err != nil {
		return Banner{}, err
	}
	return b, nil
}

func ParseIntArrayFromString(str string) ([]int64, error) {
	str = strings.Trim(str, "{}")
	parts := strings.Split(str, ",")
	var result []int64

	for _, part := range parts {
		val, _ := strconv.ParseInt(part, 10, 64)
		result = append(result, val)
	}

	return result, nil
}
