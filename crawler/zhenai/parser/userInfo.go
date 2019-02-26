package parser

import (
	"fmt"
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/model"
	"regexp"
)

var userNameRe = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]+)</span>`)
var baseUserInfoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]*)</div>`)
var partnerConditionsRe = regexp.MustCompile(`<div class="m-btn"[^>]*>([^<]+)</div>`)

func ParseUserInfo(contents []byte) engine.ParseResult {
	//fmt.Printf("%s \n",contents)
	var result []string
	userName := regexpString(contents, userNameRe, result)
	baseUserInfos := regexpString(contents, baseUserInfoRe, result)
	partnerConditions := regexpString(contents, partnerConditionsRe, result)
	userInfo := model.UserInfo{Name: string(userName[0]), PersonalData: baseUserInfos, PartnerCondition: partnerConditions}
	fmt.Printf("%v \n", userInfo)
	return engine.ParseResult{}
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
