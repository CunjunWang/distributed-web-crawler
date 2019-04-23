package parser

import (
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/model"
)

// TODO:
// 1. get the json in script
// 2. deserialize the json
// 3. fill the profile object

// parse user profile
func ParseProfile(content []byte) engine.ParseResult {
	profile := model.Profile{}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}
