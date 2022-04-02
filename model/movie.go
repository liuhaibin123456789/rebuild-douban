package model

import "errors"

//Movie 数据表的映射结构
type Movie struct {
	Id               int    `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`                         //影视身份标识
	Form             string `json:"form" form:"form" gorm:"type:VARCHAR(30);NOT NULL"`                    //影视的形式：电视剧、电影、动漫等
	Name             string `json:"name" form:"name" gorm:"type:VARCHAR(40); NOT NULL; UNIQUE"`           //影视名字
	Picture          string `json:"picture" form:"picture" gorm:"type:VARCHAR(100);NOT NULL;UNIQUE"`      //影视的图片
	Director         string `json:"director" form:"director" gorm:"type:VARCHAR(500);NOT NULL"`           //影视导演
	Writer           string `json:"writer" form:"writer" gorm:"type:VARCHAR(500);NOT NULL"`               //影视编剧
	LeadRole         string `json:"lead_role" form:"lead_role" gorm:"type:VARCHAR(500);NOT NULL"`         //影视主演
	Type             string `json:"type" form:"type" gorm:"type:VARCHAR(500);NOT NULL"`                   //影视类型
	Special          string `json:"special" form:"special" gorm:"type:VARCHAR(500);DEFAULT:''"`           //影视特色：搞笑、励志、魔幻、感人等
	ProduceWhere     string `json:"produce_where" form:"produce_where" gorm:"type:VARCHAR(100);NOT NULL"` //制片国家/地区
	Language         string `json:"language" form:"language" gorm:"type:VARCHAR(25);NOT NULL"`
	ReleaseTime      string `json:"release_time" form:"release_time" gorm:"type:VARCHAR(25);NOT NULL"`       //上映日期
	Duration         int    `json:"duration" form:"duration" gorm:"type:INT;NOT NULL"`                       //片长
	AName            string `json:"a_name" form:"a_name" gorm:"type:VARCHAR(100);NOT NULL;DEFAULT:''"`       //又名
	Imdb             string `json:"imdb" form:"imdb" gorm:"type:VARCHAR(40);DEFAULT:''"`                     //世界电影数据库索引编号
	PlotIntroduction string `json:"plot_introduction" form:"plot_introduction" gorm:"type:TEXT;NOT NULL"`    //剧情简介 todo string储存超大数据？？？
	OneStarNum       int    `json:"one_star_num" form:"one_star_num" gorm:"type:INT;NOT NULL;DEFAULT:0"`     //评1颗星人数
	TwoStarNum       int    `json:"two_star_num" form:"two_star_num" gorm:"type:INT;NOT NULL;DEFAULT:0"`     //评2颗星人数
	ThreeStarNum     int    `json:"three_star_num" form:"three_star_num" gorm:"type:INT;NOT NULL;DEFAULT:0"` //评3颗星人数
	FourStarNum      int    `json:"four_star_num" form:"four_star_num" gorm:"type:INT;NOT NULL;DEFAULT:0"`   //评4颗星人数
	FiveStarNum      int    `json:"five_star_num" form:"five_star_num" gorm:"type:INT;NOT NULL;DEFAULT:0"`   //评5颗星人数
}

func (Movie) TableName() string {
	return "movie"
}

//OfMovie 影视部分数据载体
type OfMovie struct {
	Id           int    `json:"id" form:"id"`                       //影视身份标识
	Name         string `json:"name" form:"name"`                   //影视名字
	Picture      string `json:"picture" form:"picture"`             //影视的图片
	Director     string `json:"director" form:"director"`           //影视导演
	LeadRole     string `json:"lead_role" form:"lead_role"`         //影视主演
	ProduceWhere string `json:"produce_where" form:"produce_where"` //制片国家/地区
	ReleaseTime  string `json:"release_time" form:"release_time"`   //上映日期
	Duration     int    `json:"duration" form:"duration"`           //片长
	Score        int    `json:"score" form:"score"`                 //电影评分
}

type MovieStaff struct {
	MvId    int `json:"mv_id" gorm:"type:INT;NOT NULL"`
	StaffId int `json:"staff_id" gorm:"type:INT;NOT NULL"`
}

func (MovieStaff) TableName() string {
	return "movie_staff"
}

type UserMovie struct {
	MvId  int    `json:"mv_id" gorm:"type:INT; NOT NULL"`
	Phone string `json:"phone" gorm:"type:VARCHAR(11);NOT NULL"`
	Label string `json:"label" gorm:"type:VARCHAR(1);NOT NULL;DEFAULT:'1'"`
}

func (UserMovie) TableName() string {
	return "user_movie"
}

//Discuss 讨论
type Discuss struct {
	Id           int    `json:"id"  form:"id"`                        //讨论唯一标识
	FromPhone    string `json:"from_phone" form:"from_phone"`         //发表用户
	FromUserName string `json:"from_user_name" form:"from_user_name"` //用户名
	FromAvatar   string `json:"from_avatar" form:"from_avatar"`       //用户头像名
	FromMvId     int    `json:"from_mv_id" form:"from_mv_id"`         //所属电影id
	Title        string `json:"title" form:"title"`                   //标题
	Content      string `json:"content" form:"content"`               //内容
	DiscussNum   int    `json:"discuss_num" form:"discuss_num"`       //讨论数量
	DateTime     string `json:"date_time" form:"date_time"`           //时间
}

// MvDiscuss 表的映射
type MvDiscuss struct {
	Id         int    `json:"id"  form:"id" gorm:"primaryKey;autoIncrement"`                 //讨论唯一标识
	FromPhone  string `json:"from_phone" form:"from_phone" gorm:"type:VARCHAR(11);NOT NULL"` //发表用户
	FromMvId   int    `json:"from_mv_id" form:"from_mv_id" gorm:"type:INT;NOT NULL"`         //所属电影id
	Title      string `json:"title" form:"title" gorm:"type:VARCHAR(100);not null"`          //标题
	Content    string `json:"content" form:"content" gorm:"type:VARCHAR(1000);not null"`     //内容
	DiscussNum int    `json:"discuss_num" form:"discuss_num" gorm:"type:int;default:1"`      //讨论数量
	DateTime   string `json:"date_time" form:"date_time" gorm:"type:VARCHAR(25);not null"`   //时间
}

func (MvDiscuss) TableName() string {
	return "movie_discuss"
}

//UnderDiscuss 讨论下的回应
type UnderDiscuss struct {
	Id           int    `json:"id"  form:"id"`                        //某个讨论下的回应标识，依赖于讨论id
	FromPhone    string `json:"from_phone" form:"from_phone"`         //发表用户
	ToPhone      string `json:"to_phone" form:"to_phone"`             //回复用户
	FromUserName string `json:"from_user_name" form:"from_user_name"` //用户名
	FromAvatar   string `json:"from_avatar" form:"from_avatar"`       //用户头像名
	FromMvId     int    `json:"from_mv_id" form:"from_mv_id"`         //所属电影id
	Content      string `json:"content" form:"content"`               //内容
	UsedNum      int    `json:"used_num" form:"used_num"`             //赞成数
	DateTime     string `json:"date_time" form:"date_time"`           //时间
	Label        string `json:"label" form:"label"`                   //最后一条评论的标记
}

// MvUnderDiscuss 表的映射
type MvUnderDiscuss struct {
	Id        int    `json:"id"  form:"id" gorm:"type:INT;NOT NULL"`                         //讨论唯一标识
	FromPhone string `json:"from_phone" form:"from_phone" gorm:"type:VARCHAR(11);NOT NULL"`  //发表用户
	FromMvId  int    `json:"from_mv_id" form:"from_mv_id" gorm:"type:INT;NOT NULL"`          //所属电影id
	ToPhone   string `json:"to_phone" form:"to_phone" gorm:"type:VARCHAR(11);NOT NULL"`      //回复用户
	Content   string `json:"content" form:"content" gorm:"type:text;not null"`               //内容
	UsedNum   int    `json:"used_num" form:"used_num" gorm:"type:int;default:0"`             //赞成数
	DateTime  string `json:"date_time" form:"date_time" gorm:"type:VARCHAR(25);not null"`    //时间
	Label     string `json:"label" form:"label" gorm:"type:VARCHAR(1);NOT NULL;DEFAULT:'0'"` //最后一条评论的标记
}

func (MvUnderDiscuss) TableName() string {
	return "discuss_under"
}

//UDsQueue 用于储存讨论下的回应
type UDsQueue struct {
	maxSize int              //队列长度
	Uds     [10]UnderDiscuss //数组模拟队列
	front   int              //头指针：指向首部队列元素的前一个位置，不含元素,空队列从-1开始
	rear    int              //尾指针：指向尾部队列元素，含有元素，空队列从-1开始
}

func (uds *UDsQueue) add(underDiscuss UnderDiscuss) error {
	if uds.rear == uds.maxSize-1 {
		return errors.New("queue full")
	}
	uds.rear++
	uds.Uds[uds.rear] = underDiscuss
	return nil
}

func (q *UDsQueue) get() (UnderDiscuss, error) {
	if q.front == q.rear {
		return UnderDiscuss{}, errors.New("queue is nil")
	}
	q.front++
	return q.Uds[q.front], nil
}
