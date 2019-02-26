package parser

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/model"
	"regexp"
)

//var userNameRe = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]+)</span>`)
var baseUserInfoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]*)</div>`)
var partnerConditionsRe = regexp.MustCompile(`<div class="m-btn"[^>]*>([^<]+)</div>`)

func ParseUserInfo(contents []byte, userName string) engine.ParseResult {
	//fmt.Printf("%s \n",contents)
	var result []string
	baseUserInfos := regexpString(contents, baseUserInfoRe, result)
	partnerConditions := regexpString(contents, partnerConditionsRe, result)
	userInfo := model.UserInfo{Name: userName, PersonalData: baseUserInfos, PartnerCondition: partnerConditions}
	return engine.ParseResult{Items: []interface{}{userInfo}}
}

func regexpString(contents []byte, regexp *regexp.Regexp, result []string) []string {
	result = []string{}
	submatchs := regexp.FindAllSubmatch(contents, -1)
	for _, submatch := range submatchs {
		if len(submatch) < 2 {
			continue
		}
		result = append(result, string(submatch[1]))
	}
	return result
}
