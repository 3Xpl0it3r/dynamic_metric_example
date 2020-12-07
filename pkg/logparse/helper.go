package logparse

import (
	"strings"
)

func getShortUri(uri string)string{
	uriSplitSpace := strings.Split(uri, " ")
	if len(uriSplitSpace) == 0 {
		return "/"
	}else if  len(uriSplitSpace) == 1 {
		if strings.HasPrefix(uriSplitSpace[0], "/"){
			return strings.Split(uriSplitSpace[0], "?")[0]
		} else {
			return "/"
		}
	}
	if strings.HasPrefix(uriSplitSpace[0], "/") {
		return strings.Split(uriSplitSpace[0], "?")[0]
	} else {
		return strings.Split(uriSplitSpace[1], "?")[0]
	}
}


func tidParse(tid string){

}