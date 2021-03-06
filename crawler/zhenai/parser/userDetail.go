package parser

import (
	"github.com/KarenLKL/studygolang/crawler/model"
	"log"
	"strings"
)

var marriageStatusArray = []string{"离异", "未婚", "丧偶", ""}

func ParseUserDetail(userInfo interface{}) *model.UserDetail {

	if user, ok := userInfo.(model.UserInfo); ok {
		if len(user.PersonalData) < 1 {
			return nil
		}
		userDetail := &model.UserDetail{
			UserName:       user.Name,
			MarriageStatus: getTargetValueIfContain(user.PersonalData, 0, "", marriageStatusArray),
			Age:            getTargetValueIfContain(user.PersonalData, 1, "岁", nil),
			Constellation:  getTargetValueIfContain(user.PersonalData, 2, "座", nil),
			Height:         getTargetValueIfContain(user.PersonalData, 3, "cm", nil),
			Weight:         getTargetValueIfContain(user.PersonalData, 4, "kg", nil),
			Workplace:      getTargetValueIfContain(user.PersonalData, 5, "工作地:", nil),
			IncomeOfMonth:  getTargetValueIfContain(user.PersonalData, 6, "月收入", nil),
			Occupation:     getTargetValueIfExit(user.PersonalData, 7, ""),
			Education:      getTargetValueIfExit(user.PersonalData, 8, ""),
		}
		userDetail.PartnerCondition = &model.PartnerCondition{
			Age:            getTargetValueIfContain(user.PartnerCondition, 0, "岁", nil),
			Height:         getTargetValueIfContain(user.PartnerCondition, 1, "cm", nil),
			Workplace:      getTargetValueIfContain(user.PartnerCondition, 2, "工作地", nil),
			IncomeOfMonth:  getTargetValueIfContain(user.PersonalData, 3, "月薪:", nil),
			MarriageStatus: getTargetValueIfContain(user.PersonalData, 4, "", marriageStatusArray),
			Shape:          getTargetValueIfContain(user.PersonalData, 5, "体型", nil),
			DrinkAble:      getTargetValueIfContain(user.PersonalData, 6, "喝酒", nil),
			SmokingAble:    getTargetValueIfContain(user.PersonalData, 7, "吸烟", nil),
			HasChild:       getTargetValueIfContain(user.PersonalData, 8, "小孩", nil),
			Child:          getTargetValueIfContain(user.PersonalData, 9, "是否想要孩子", nil),
		}
		//log.Printf("ParseUserDetail: %v+%", userDetail)
		return userDetail
	} else {
		log.Println("interface convert to UserInfo exception!")
		return nil
	}
}

func getTargetValueIfExit(user []string, position int, contain string) string {
	//log.Printf("len(user)=%d position=%d", len(user), position)
	if len(user) == 0 || position >= len(user) {
		return ""
	}
	if len(contain) > 0 {
		if strings.Contains(user[position], contain) {
			return user[position]
		}
		return ""
	} else {
		return user[position]
	}
}

func getTargetValueIfContain(user []string, position int, contain string, from []string) string {
	// scene 1、从多种情况中选择，是否存在
	if len(from) > 0 {
		return getSameValueFromArrays(from, user)
	}
	// scene 2、根据某段字符串判断，是否存在包含的属性
	if len(contain) > 0 {
		return getContainValue(contain, user)
	}
	// scene 3、根据position取值
	if position >= len(user) {
		return ""
	}
	return user[position]

}

func getSameValueFromArrays(from, user []string) string {
	for _, value := range from {
		if strings.EqualFold(value, getEqualedValue(value, user)) {
			return value
		}
	}
	return ""
}

func getEqualedValue(target string, arras []string) string {
	for _, value := range arras {
		if strings.EqualFold(target, value) {
			return value
		}
	}
	return ""
}

func getContainValue(containVal string, arrays []string) string {
	for _, value := range arrays {
		if strings.Contains(value, containVal) {
			return value
		}
	}
	return ""
}
