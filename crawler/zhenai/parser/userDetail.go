package parser

import (
	"github.com/KarenLKL/studygolang/crawler/model"
	"log"
	"strings"
)

type UserDetail struct {
	UserName       string `json:"user_name"`       // 昵称
	MarriageStatus string `json:"marriage_status"` // 婚姻状态
	Age            string `json:"age"`             // 年龄
	Constellation  string `json:"constellation"`   //星座
	Height         string `json:"height"`          //身高
	Weight         string `json:"weight"`          //体重
	Workplace      string `json:"workplace"`       //工作地
	IncomeOfMonth  string `json:"income_of_month"` //月收入
	Occupation     string `json:"occupation"`      //职位
	Education      string `json:"education"`       //学历
}

func ParseUserDetail(userInfo interface{}) *UserDetail {
	if user, ok := userInfo.(model.UserInfo); ok {
		if len(user.PersonalData) < 1 {
			return nil
		}
		userDetail := &UserDetail{
			UserName:       user.Name,
			MarriageStatus: getTargetValueIfExit(user.PersonalData, 0, ""),
			Age:            getTargetValueIfExit(user.PersonalData, 1, "岁"),
			Constellation:  getTargetValueIfExit(user.PersonalData, 2, "座"),
			Height:         getTargetValueIfExit(user.PersonalData, 3, "cm"),
			Weight:         getTargetValueIfExit(user.PersonalData, 4, "kg"),
			Workplace:      getTargetValueIfExit(user.PersonalData, 5, "工作地:"),
			IncomeOfMonth:  getTargetValueIfExit(user.PersonalData, 6, "月收入"),
			Occupation:     getTargetValueIfExit(user.PersonalData, 7, ""),
			Education:      getTargetValueIfExit(user.PersonalData, 8, ""),
		}
		return userDetail
	} else {
		log.Println("interface convert to UserInfo exception!")
		return nil
	}
}

func getTargetValueIfExit(user []string, position int, contain string) string {
	log.Printf("len(user)=%d position=%d", len(user), position)
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
