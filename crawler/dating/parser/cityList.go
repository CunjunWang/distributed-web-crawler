package parser

import (
	"distributed-web-crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	// TODO:
	// 测试时控制一下city数量
	limit := 5

	for _, m := range matches {
		result.Items = append(
			result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
