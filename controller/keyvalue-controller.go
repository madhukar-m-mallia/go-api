package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/madhukar-m-mallia/go-api/entity"
	"github.com/madhukar-m-mallia/go-api/service"
)

type KeyValueController interface {
	Set(ctx *gin.Context) (entity.KeyValue, error)
	FindOne(string) (string, error)
	Search(c *gin.Context) (string, error)
}

type controller struct {
	service service.KeyValueService
}

func New(service service.KeyValueService) KeyValueController {
	return controller{
		service: service,
	}
}

func (c controller) Set(ctx *gin.Context) (entity.KeyValue, error) {
	var keyVal entity.KeyValue
	ctx.BindJSON(&keyVal)
	keyValRes, err := c.service.Set(keyVal)

	if err != nil {
		return entity.KeyValue{}, err
	}
	return keyValRes, nil
}

func (c controller) FindOne(key string) (string, error) {
	a, err := c.service.FindOne(key)
	return a, err
}

func (c controller) Search(ctx *gin.Context) (string, error) {
	queryParams := ctx.Request.URL.Query()
	keys := []string{}
	for k := range queryParams {
		keys = append(keys, k)
	}
	var (
		filterVal string   = ""
		matched   []string = []string{}
	)
	switch keys[0] {
	case "prefix":
		filterVal = queryParams.Get("prefix")
		matched, _ = c.service.Search(filterVal, "prefix")
	case "suffix":
		filterVal = queryParams.Get("suffix")
		matched, _ = c.service.Search(filterVal, "suffix")
	default:
		matched = []string{}
	}
	if len(matched) > 0 {
		return strings.Join(matched, ","), nil
	} else {
		return "", errors.New("No String with provided filters")
	}
}
