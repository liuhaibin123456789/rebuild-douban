package model

//Staff 影视演职员数据表结构映射
type Staff struct {
	Id            string `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name" form:"name" gorm:"type:VARCHAR(40);NOT NULL"`
	Sex           string `json:"sex" form:"sex" gorm:"type:VARCHAR(3);NOT NULL"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"type:VARCHAR(200);DEFAULT:'默认头像.jpg'"`  //头像名
	Constellation string `json:"constellation" form:"constellation" gorm:"type:VARCHAR(12)"`        //星座
	Birthday      string `json:"birthday" form:"birthday" gorm:"type:VARCHAR(25)"`                  //出生日期
	Birthplace    string `json:"birthplace" form:"birthplace" gorm:"type:VARCHAR(100);DEFAULT:'无'"` //出生地
	Jobs          string `json:"jobs" form:"jobs" gorm:"type:VARCHAR(500);DEFAULT:'无'"`             //从事过的工作
	ACName        string `json:"ac_name" form:"ac_name" gorm:"type:VARCHAR(100);DEFAULT:'无'"`       //中文别名
	AEName        string `json:"ae_name" form:"ae_name" gorm:"type:VARCHAR(100);DEFAULT:'无'"`       //英文别名
	Family        string `json:"family" form:"family" gorm:"type:VARCHAR(100);DEFAULT:'无'"`         //家庭成员
	Imdb          string `json:"imdb" form:"imdb" gorm:"type:VARCHAR(100)"`                         //世界最大电影数据库编号
	Introduction  string `json:"introduction" form:"introduction" gorm:"type:TEXT;NOT NULL"`        //介绍
}

func (Staff) TableName() string {
	return "staff"
}

//OfStaff 影职员部分数据
type OfStaff struct {
	Id     string `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Avatar string `json:"avatar" form:"avatar"` //头像名
	Jobs   string `json:"jobs" form:"jobs"`     //从事过的工作
}
