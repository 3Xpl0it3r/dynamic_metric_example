package logparse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T){
	originLogData := `223.70.153.106 - - [07/Dec/2020:11:18:53 +0800] "GET /api/standardObject/getListLayout?standardCollectionId=72897&resourceId=118345&requestType=code&resourceType=menu HTTP/2.0" 200 2297 "https://app.lab.clickpaas.com/" "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36" "-" 0.180 - 1.38140106366.16073111348020004`
	logDataGram := LogParse(originLogData)
	assert.Equal(t, logDataGram.Status, "200")
	assert.Equal(t, getShortUri(logDataGram.Request), "/api/standardObject/getListLayout")
	assert.Equal(t ,logDataGram.RemoteAddr, "223.70.153.106")

}
