package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParseResult {
		re := regexp.MustCompile(cityRe)
		matchs := re.FindAllSubmatch(contents, -1)

		result := engine.ParseResult{}
		for _, m := range matchs {
			result.Items = append(result.Items, "User " + string(m[2]))
			result.Request = append(result.Request, engine.Request{
				Url: string(m[1]),
				ParseFunc: engine.NilParser,
			})
		}
		return result
	}

