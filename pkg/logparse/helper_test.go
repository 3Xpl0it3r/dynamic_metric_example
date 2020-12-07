package logparse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelper(t *testing.T){
	urlTestData := map[string]string{
		"/api/standardObject/getListData":"/api/standardObject/getListData",
		"/api/standardObject/getListLayout?standardCollectionId=72897&resourceId=118345&requestType=code&resourceType=menu": "/api/standardObject/getListLayout",
		"/": "/",
		"/?=xxx?v=xxxx": "/",
	}

	for k,v := range urlTestData{
		uri := getShortUri(k)
		assert.Equal(t, uri, v)
	}
}
