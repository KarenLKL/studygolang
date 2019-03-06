package persist

import (
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"testing"
)

func TestSave(t *testing.T) {
	detail := model.UserDetail{
		UserName:       "张三",
		MarriageStatus: "未婚",
		Age:            "25岁",
		Constellation:  "双鱼座",
		Height:         "175cm",
		Weight:         "65kg",
		Workplace:      "昆明",
		IncomeOfMonth:  "8000元/月",
		Occupation:     "程序员",
		Education:      "本科",
		PartnerCondition: &model.PartnerCondition{
			Age:            "30岁",
			Height:         "165cm",
			Workplace:      "昆明",
			IncomeOfMonth:  "10000元",
			MarriageStatus: "未婚",
			Shape:          "体型",
			DrinkAble:      "可以喝酒",
			SmokingAble:    "可以",
			Child:          "想要孩子",
			HasChild:       "没有小孩",
		},
	}
	row := engine.Item{
		Url:      "http://album.zhenai.com/u/1726316985",
		Id:       "1726316985",
		Type:     "zhenai",
		UserInfo: detail,
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	err = save(client, row, "dating_test")
	if err != nil {
		t.Error(err)
	}
	userDetail, err := get(client, "dating_test", row.Id)
	if err != nil {
		t.Errorf("get user from elastc error:%s", err.Error())
	}
	if userDetail != detail {
		t.Errorf("detail not equaled userDetail")
	}
	log.Printf("detail: %+v", detail)
	log.Printf("userDetail: %+v", userDetail)
}
