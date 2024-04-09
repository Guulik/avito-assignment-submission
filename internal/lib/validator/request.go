package validator

import (
	"Avito_trainee_assignment/internal/domain/request"
	"errors"
	"fmt"
)

func CheckGetUserRequest(req request.GetUserRequest) error {
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId: featureId is not negative")
	}
	if req.TagId < 0 {
		return errors.New("incorrect featureId: featureId is not negative")
	}
	return nil
}

func CheckPostRequest(req request.CreateRequest, fullCheck bool) error {
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId: featureId is not negative")
	}
	if len(req.TagIds) == 0 || req.TagIds == nil {
		return errors.New("incorrect tags: tags cannot be empty")
	}
	if fullCheck {
		for i, tag := range req.TagIds {
			if tag < 0 {
				return errors.New(fmt.Sprintf("incorrect tagId #%v: tagId is not negative", i))
			}
		}
	}
	return nil
}

func CheckDeleteRequest(req request.DeleteRequest) error {
	if req.BannerId < 1 {
		return errors.New("incorrect bannerId: bannerId must be >1")
	}
	return nil
}
