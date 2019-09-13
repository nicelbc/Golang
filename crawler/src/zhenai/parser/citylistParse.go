package parser

import (
	"engine"
	"regexp"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) (engine.ParseResult){
		re := regexp.MustCompile(citylistRe)
		matchs := re.FindAllSubmatch(contents, -1)

		result := engine.ParseResult{}
		for _, m := range matchs {
			result.Items = append(result.Items, string(m[2]))
			result.Request = append(result.Request, engine.Request{
				Url: string(m[1]),
				ParseFunc: engine.NilParser,
			})
		}
		return result
	}
