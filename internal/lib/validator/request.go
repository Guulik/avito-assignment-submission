package validator

import (
	"Avito_trainee_assignment/internal/domain/request"
	"errors"
	"fmt"
)

func CheckGetUserRequest(req *request.GetUserRequest) error {
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId")
	}
	if req.TagId < 0 {
		return errors.New("incorrect tagId")
	}
	return nil
}

// CheckPostRequest checks if featureId>-1,len(tagIds[])>0,
// optionally if the second parameter is true a full tag revision(all(tagId) >-1) and
// removing duplicates from tagIds can be performed
func CheckPostRequest(req *request.CreateRequest, advanced bool) error {
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId")
	}
	if len(req.TagIds) == 0 || req.TagIds == nil {
		return errors.New("incorrect tags: tags cannot be empty")
	}
	if advanced {
		req.TagIds = removeDuplicates(req.TagIds)
		for i, tag := range req.TagIds {
			if tag < 0 {
				return errors.New(fmt.Sprintf("incorrect tagId #%v", i))
			}
		}
	}
	return nil
}

// CheckUpdateRequest checks if bannerId>1, featureId>-1,len(tagIds[])>0,
// optionally if the second parameter is true a full tag revision(all(tagId) >-1) and
// removing duplicates from tagIds can be performed
func CheckUpdateRequest(req *request.UpdateRequest, advanced bool) error {
	if req.BannerId < 1 {
		return errors.New("incorrect bannerId: bannerId must be >1")
	}
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId")
	}
	if len(req.TagIds) == 0 || req.TagIds == nil {
		return errors.New("incorrect tags: tags cannot be empty")
	}
	if advanced {
		req.TagIds = removeDuplicates(req.TagIds)
		for i, tag := range req.TagIds {
			if tag < 0 {
				return errors.New(fmt.Sprintf("incorrect tagId #%v", i))
			}
		}
	}
	return nil
}

func CheckDeleteRequest(req *request.DeleteRequest) error {
	if req.BannerId < 1 {
		return errors.New("incorrect bannerId: bannerId must be >1")
	}
	return nil
}

func removeDuplicates(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	slice := []int64{}

	for _, entry := range intSlice {
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			slice = append(slice, entry)
		}
	}
	return slice
}
