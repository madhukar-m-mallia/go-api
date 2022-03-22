package service

import (
	"errors"
	"strings"

	"github.com/madhukar-m-mallia/go-api/entity"
)

type KeyValueService interface {
	Set(entity.KeyValue) (entity.KeyValue, error)
	FindOne(string) (string, error)
	Search(string, string) ([]string, error)
}

type keyValueService struct {
	keyValues []entity.KeyValue
}

func New() KeyValueService {
	return &keyValueService{}
}

func (service *keyValueService) Set(keyVal entity.KeyValue) (entity.KeyValue, error) {
	if keyVal.Key == "" || keyVal.Value == "" {
		return entity.KeyValue{}, errors.New("Improper data sent")
	} else {
		service.keyValues = append(service.keyValues, keyVal)
		return keyVal, nil
	}
}

func (service *keyValueService) FindOne(key string) (string, error) {
	if key != "" {
		for _, keyVal := range service.keyValues {
			if keyVal.Key == key {
				return keyVal.Key + "-" + keyVal.Value, nil
			}
		}
		return "", errors.New("Key not found")
	} else {
		return "", errors.New("Empty key sent")
	}

}

func (service *keyValueService) Search(key string, searchType string) ([]string, error) {
	if key == "" {
		return []string{}, errors.New("Empty key passed")
	} else {
		switch searchType {
		case "prefix":
			return getAllValueFromCacheStartsWith(service.keyValues, key), nil
		case "suffix":
			return getAllValueFromCacheEndsWith(service.keyValues, key), nil
		default:
			return []string{}, errors.New("Improper prefix")
		}
	}
}

func getValueFromCache(keyValues []entity.KeyValue, key string) (string, error) {
	for _, keyVal := range keyValues {
		if keyVal.Key == key {
			return keyVal.Key + "-" + keyVal.Value, nil
		}
	}
	return "", errors.New("Key not found")
}

func getAllValueFromCacheStartsWith(keyValues []entity.KeyValue, key string) []string {
	keys := []string{}
	for _, keyVal := range keyValues {
		if strings.HasPrefix(keyVal.Key+"-"+keyVal.Value, key) {
			keys = append(keys, keyVal.Key+"-"+keyVal.Value)
		}
	}
	return keys
}

func getAllValueFromCacheEndsWith(keyValues []entity.KeyValue, key string) []string {
	keys := []string{}
	for _, keyVal := range keyValues {
		if strings.HasSuffix(keyVal.Key+"-"+keyVal.Value, key) {
			keys = append(keys, keyVal.Key+"-"+keyVal.Value)
		}
	}
	return keys
}
