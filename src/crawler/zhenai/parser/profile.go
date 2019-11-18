package parser

import (
	"crawler/engine"
	"crawler/module"
	"regexp"
)

const ageRe = `(<div class="m-btn purple" [^>]+)>([^<]+)</div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := module.Profile{}
	profile.Name = name
	re := regexp.MustCompile(ageRe)
	matchs := re.FindAllSubmatch(contents, -1)

	profile.Age = string(matchs[0][2])

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result

}
