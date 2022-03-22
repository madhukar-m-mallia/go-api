package service

import (
	"testing"

	"github.com/madhukar-m-mallia/go-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	testService := New()
	_, err := testService.Set(entity.KeyValue{})

	assert.NotNil(t, err)
	assert.Equal(t, "Improper data sent", err.Error())
}

func TestFindOne(t *testing.T) {
	testService := New()
	_, err := testService.FindOne("key_not_present")

	assert.NotNil(t, err)
	assert.Equal(t, "Key not found", err.Error())
}

func TestSearch(t *testing.T) {
	testService := New()
	_, err := testService.Search("", "")

	assert.NotNil(t, err)
	assert.Equal(t, "Empty key passed", err.Error())
}
