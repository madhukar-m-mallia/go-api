package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/madhukar-m-mallia/go-api/entity"
)

type KeyValueService interface {
	Set(entity.KeyValue) (entity.KeyValue, error)
	FindOne(string) (string, error)
	Search(string, string) ([]string, error)
}

type keyValueService struct {
	keyValues sync.Map
}

func New() KeyValueService {
	return &keyValueService{}
}

func (service *keyValueService) Set(keyVal entity.KeyValue) (entity.KeyValue, error) {
	if keyVal.Key == "" || keyVal.Value == "" {
		return entity.KeyValue{}, errors.New("Improper data sent")
	} else {
		service.keyValues.Store(keyVal.Key, keyVal.Value)
		return keyVal, nil
	}
}

func (service *keyValueService) FindOne(key string) (string, error) {
	if key != "" {
		value, err := service.keyValues.Load(key)
		if err != true {
			return "", errors.New("Key not found")
		}

		return fmt.Sprintf("%v", value), nil
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

func getValueFromCache(keyValues sync.Map, key string) (string, error) {
	value, err := keyValues.Load(key)
	if err != true {
		return "", errors.New("Key not found")
	}
	return fmt.Sprint(value), nil
}

func getAllValueFromCacheStartsWith(keyValues sync.Map, key string) []string {
	keys := []string{}
	keyValues.Range(func(k, v interface{}) bool {
		if strings.HasPrefix(fmt.Sprint(k)+"-"+fmt.Sprint(v), key) {
			keys = append(keys, fmt.Sprint(k)+"-"+fmt.Sprint(v))
		}
		return true
	})
	return keys
}

func getAllValueFromCacheEndsWith(keyValues sync.Map, key string) []string {
	keys := []string{}
	keyValues.Range(func(k, v interface{}) bool {
		if strings.HasSuffix(fmt.Sprint(k)+"-"+fmt.Sprint(v), key) {
			keys = append(keys, fmt.Sprint(k)+"-"+fmt.Sprint(v))
		}
		return true
	})
	return keys
}
