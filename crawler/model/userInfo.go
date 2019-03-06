package model

type UserInfo struct {
	Name             string
	PersonalData     []string
	PartnerCondition []string
}

type UserDetail struct {
	UserName         string            `json:"user_name"`         // 昵称
	MarriageStatus   string            `json:"marriage_status"`   // 婚姻状态
	Age              string            `json:"age"`               // 年龄
	Constellation    string            `json:"constellation"`     //星座
	Height           string            `json:"height"`            //身高
	Weight           string            `json:"weight"`            //体重
	Workplace        string            `json:"workplace"`         //工作地
	IncomeOfMonth    string            `json:"income_of_month"`   //月收入
	Occupation       string            `json:"occupation"`        //职位
	Education        string            `json:"education"`         //学历
	PartnerCondition *PartnerCondition `json:"partner_condition"` // 择偶条件
}

type PartnerCondition struct {
	Age            string `json:"age"`             // 年龄
	Height         string `json:"height"`          //身高
	Workplace      string `json:"workplace"`       //工作地
	IncomeOfMonth  string `json:"income_of_month"` //月收入
	MarriageStatus string `json:"marriage_status"` // 婚姻状态
	Shape          string `json:"shape"`           // 体型
	DrinkAble      string `json:"drink_able"`      //是否可以喝酒
	SmokingAble    string `json:"smoking_able"`    //是否允许抽烟
	Child          string `json:"child"`           //是否要孩子
	HasChild       string `json:"has_child"`       //是否有小孩
}
