package parser

import (
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/model"
	"github.com/valyala/fastjson"
	"log"
	"regexp"
	"strconv"
)

const jsonRe = `<script>window.__INITIAL_STATE__=({.+});`

// parse user profile
func ParseProfile(content []byte) engine.ParseResult {
	re := regexp.MustCompile(jsonRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		objectString := string(m[1])
		var p fastjson.Parser
		v, err := p.Parse(objectString)
		if err != nil {
			log.Fatal(err)
		}

		// 用户信息对象
		userInfoObject := v.GetObject("objectInfo")
		// 基础信息列表
		basicInfo, _ := userInfoObject.Get("basicInfo").Array()
		// 详细信息列表
		detailInfo, _ := userInfoObject.Get("detailInfo").Array()

		// new profile object
		profile := model.Profile{}
		// name
		profile.Name = userInfoObject.Get("nickname").String()
		// nickname
		profile.Nickname = userInfoObject.Get("nickname").String()
		// gender
		profile.Gender = userInfoObject.Get("genderString").String()
		// age
		profile.Age = userInfoObject.Get("age").GetInt()
		// income
		profile.Income = userInfoObject.Get("salaryString").String()
		// marriage
		profile.Marriage = userInfoObject.Get("marriageString").String()
		// education
		profile.Education = userInfoObject.Get("educationString").String()
		// city
		profile.City = userInfoObject.Get("workCityString").String()
		// height
		heightString := userInfoObject.Get("heightString").String()
		heightWithoutUnit := heightString[1:4]
		height, _ := strconv.Atoi(heightWithoutUnit)
		profile.Height = height

		if len(basicInfo) >= 8 {
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

		if len(detailInfo) >= 8 {
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

		result.Items = append(result.Items, profile)
		result.Requests = append(
			result.Requests, engine.Request{})
	}
	return result
}
