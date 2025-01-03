package validator

import (
	"Banner_Infrastructure/internal/domain/request"
	"errors"
	"fmt"
	"reflect"
)

func CheckGetRequest(req *request.GetUserRequest) error {
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId")
	}
	if req.TagId < 0 {
		return errors.New("incorrect tagId")
	}
	return nil
}

// CheckPostRequest checks if featureId>-1,len(tagIds[])>0, and
// removing duplicates from tagIds
// optionally if the second parameter is true a full tag revision(all(tagId) >-1) can be performed.
func CheckPostRequest(req *request.CreateRequest, advanced bool) error {
	if reflect.TypeOf(req.FeatureId) != reflect.TypeOf(int64(0)) {
		return errors.New("incorrect featureId type")
	}
	if reflect.TypeOf(req.TagIds) != reflect.TypeOf([]int64{}) {
		return errors.New("incorrect tagIds type")
	}
	if req.FeatureId < 0 {
		return errors.New("incorrect featureId: negative")
	}
	if len(req.TagIds) == 0 || req.TagIds == nil {
		return errors.New("incorrect tags: tags cannot be empty")
	}
	req.TagIds = removeDuplicates(req.TagIds)
	if advanced {
		for i, tag := range req.TagIds {
			if tag < 0 {
				return fmt.Errorf("incorrect tagId #%v", i)
			}
		}
	}
	return nil
}

// CheckUpdateRequest checks if featureId>-1,len(tagIds[])>0, and
// removing duplicates from tagIds
// optionally if the second parameter is true a full tag revision(all(tagId) >-1) can be performed.
func CheckUpdateRequest(req *request.UpdateRequest, advanced bool) error {
	if req.BannerId < 1 {
		return errors.New("incorrect bannerId: bannerId must be >1")
	}
	req.TagIds = removeDuplicates(req.TagIds)
	if advanced {
		for i, tag := range req.TagIds {
			if tag < 0 {
				return fmt.Errorf("incorrect tagId #%v", i)
			}
		}
	}
	return nil
}

func CheckDeleteRequest(req *request.DeleteRequest) error {
	if req.BannerId < 0 && req.FeatureId < 0 && req.TagId < 0 {
		return errors.New("bad request! at least one parameter required")
	}
	if (req.BannerId > -1 && req.FeatureId > -1) || (req.BannerId > -1 && req.TagId > -1) {
		return errors.New("bad request! you should choose: Id xor tag|feature")
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
