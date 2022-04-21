package xkcd

import (
	"encoding/json"
	"net/http"
)

func SearchComic(numComic string) (*DataComic, error){
	linkComic := Link + numComic + "/info.0.json"
	resp,er1:=http.Get(linkComic)
	if er1!=nil{
		return nil,er1
	}
	defer resp.Body.Close()
	var d DataComic
	er2:=json.NewDecoder(resp.Body).Decode(&d)
	if er2!=nil{
		return nil,er2
	}
	return &d,nil
}
