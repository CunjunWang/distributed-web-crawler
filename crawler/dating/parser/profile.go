package parser

import (
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/model"
	"github.com/valyala/fastjson"
	"log"
	"regexp"
	"strconv"
)

var jsonRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=({.+});`)

var expandRe = regexp.MustCompile(
	`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[^"]+)">`)

// parse user profile
func ParseProfile(content []byte) engine.ParseResult {

	matches := jsonRe.FindAllSubmatch(content, -1)

	// log.Println("matches: ", len(matches))

	// new profile object
	profile := model.Profile{}

	for _, m := range matches {

		objectString := string(m[1])
		var p fastjson.Parser
		v, err := p.Parse(objectString)
		if err != nil {
			log.Fatal(err)
		}

		// 用户信息对象
		userInfoObject := v.GetObject("objectInfo")

		if userInfoObject != nil {

			// name
			if userInfoObject.Get("nickname") != nil {
				profile.Name = userInfoObject.Get("nickname").String()
			}
			// nickname
			if userInfoObject.Get("nickname") != nil {
				profile.Nickname = userInfoObject.Get("nickname").String()
			}
			// gender
			if userInfoObject.Get("genderString") != nil {
				profile.Gender = userInfoObject.Get("genderString").String()
			}
			// age
			if userInfoObject.Get("age") != nil {
				profile.Age = userInfoObject.Get("age").GetInt()
			}
			// income
			if userInfoObject.Get("salaryString") != nil {
				profile.Income = userInfoObject.Get("salaryString").String()
			}
			// marriage
			if userInfoObject.Get("marriageString") != nil {
				profile.Marriage = userInfoObject.Get("marriageString").String()
			}
			// education
			if userInfoObject.Get("educationString") != nil {
				profile.Education = userInfoObject.Get("educationString").String()
			}
			// city
			if userInfoObject.Get("workCityString") != nil {
				profile.City = userInfoObject.Get("workCityString").String()
			}
			// height
			if userInfoObject.Get("heightString") != nil {
				heightString := userInfoObject.Get("heightString").String()
				heightWithoutUnit := heightString[1:4]
				height, _ := strconv.Atoi(heightWithoutUnit)
				profile.Height = height
			}

			// 基础信息列表
			basicInfo, _ := userInfoObject.Get("basicInfo").Array()
			// 详细信息列表
			detailInfo, _ := userInfoObject.Get("detailInfo").Array()

			if basicInfo != nil && len(basicInfo) >= 8 {
				// weight
				weightString := basicInfo[4].String()
				weightWithoutUnit := weightString[1:3]
				weight, _ := strconv.Atoi(weightWithoutUnit)
				profile.Weight = weight
				// occupation
				profile.Occupation = basicInfo[len(basicInfo)-2].String()
				// constellation
				profile.Constellation = basicInfo[2].String()
			}

			if detailInfo != nil && len(detailInfo) >= 8 {
				// nationality
				profile.Nationality = detailInfo[0].String()
				// smoke
				profile.Smoke = detailInfo[3].String()
				// drink
				profile.Drink = detailInfo[4].String()
				// house
				profile.House = detailInfo[5].String()
				// car
				profile.Car = detailInfo[6].String()
			}
		}
	}

	result := engine.ParseResult{}
	result.Items = append(result.Items, profile)

	expandMatches := expandRe.FindAllSubmatch(content, -1)
	for _, m := range expandMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
