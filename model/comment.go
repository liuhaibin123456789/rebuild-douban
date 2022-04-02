package model

//ShortComment 短评
type ShortComment struct {
	Id            int    `json:"id"  form:"id"`                           //短评唯一标识
	FromPhone     string `json:"from_phone" form:"from_phone" `           //用户唯一标识
	FromUserName  string `json:"from_user_name" form:"from_user_name"`    //用户名
	FromAvatar    string `json:"from_avatar" form:"from_avatar"`          //用户头像名
	FromMvId      int    `json:"from_mv_id" form:"from_mv_id" `           //短评所属电影id
	WantOrWatched string `json:"want_or_watched" form:"want_or_watched" ` //想看or不想看
	MvStar        int    `json:"mv_star" form:"mv_star" `                 //电影评分
	Tag           string `json:"tag" form:"tag"`                          //短评标签
	Content       string `json:"content" form:"content"`                  //短评内容
	DateTime      string `json:"date_time" form:"date_time"`              //短评时间
	UsedNum       int    `json:"used_num" form:"used_num"`                //赞成数
}

// MvShortComment 表的映射
type MvShortComment struct {
	Id            int    `json:"id"  form:"id" gorm:"primaryKey;autoIncrement"`                          //短评唯一标识
	FromPhone     string `json:"from_phone" form:"from_phone" gorm:"type:VARCHAR(11);NOT NULL"`          //用户唯一标识
	FromMvId      int    `json:"from_mv_id" form:"from_mv_id" gorm:"type:INT;NOT NULL"`                  //短评所属电影id
	WantOrWatched string `json:"want_or_watched" form:"want_or_watched" gorm:"type:VARCHAR(1) NOT NULL"` //想看or不想看
	MvStar        int    `json:"mv_star" form:"mv_star" gorm:"type:int"`                                 //电影评分
	Tag           string `json:"tag" form:"tag" gorm:"type:VARCHAR(100)"`                                //短评标签
	Content       string `json:"content" form:"content" gorm:"type:VARCHAR(350)"`                        //短评内容
	DateTime      string `json:"date_time" form:"date_time" gorm:"type:VARCHAR(25);not null"`            //短评时间
	UsedNum       int    `json:"used_num" form:"used_num" gorm:"type:INT;DEFAULT:0"`                     //赞成数
}

func (MvShortComment) TableName() string {
	return "movie_short_comment"
}

//LongComment 影评
type LongComment struct {
	Id           int    `json:"id"  form:"id"`                        //影评唯一标识
	FromPhone    string `json:"from_phone" form:"from_phone"`         //用户唯一标识
	FromUserName string `json:"from_user_name" form:"from_user_name"` //用户名
	FromAvatar   string `json:"from_avatar" form:"from_avatar"`       //用户头像名
	FromMvId     int    `json:"from_mv_id" form:"from_mv_id"`         //影评所属电影id
	MvStar       int    `json:"mv_star" form:"mv_star"`               //电影评分
	Title        string `json:"title" form:"title"`                   //影评标题
	Content      string `json:"content" form:"content"`               //影评内容
	DateTime     string `json:"date_time" form:"date_time"`           //影评时间
	UsedNum      int    `json:"used_num" form:"used_num"`             //赞成数
	UnusedNum    int    `json:"unused_num" form:"unused_num"`         //不赞成数
}

// MvLongComment 表的映射
type MvLongComment struct {
	Id        int    `json:"id"  form:"id" gorm:"primaryKey;autoIncrement"`                 //影评唯一标识
	FromPhone string `json:"from_phone" form:"from_phone" gorm:"type:VARCHAR(11);NOT NULL"` //用户唯一标识
	FromMvId  int    `json:"from_mv_id" form:"from_mv_id" gorm:"type:INT;NOT NULL"`         //影评所属电影id
	MvStar    int    `json:"mv_star" form:"mv_star" gorm:"type:int"`                        //电影评分
	Title     string `json:"title" form:"title" gorm:"type:VARCHAR(140);not null"`          //影评标题
	Content   string `json:"content" form:"content" gorm:"type:TEXT;NOT NULL"`              //影评内容
	DateTime  string `json:"date_time" form:"date_time" gorm:"type:VARCHAR(25);NOT NULL"`   //影评时间
	UsedNum   int    `json:"used_num" form:"used_num" gorm:"type:INT;DEFAULT:0"`            //赞成数
	UnusedNum int    `json:"unused_num" form:"unused_num" gorm:"type:INT;DEFAULT:0"`        //不赞成数
}

func (MvLongComment) TableName() string {
	return "movie_long_comment"
}

//UnderLongComment 影评下的留言及讨论
type UnderLongComment struct {
	Id           int    `json:"id"  form:"id"`                        //影评唯一标识
	FromPhone    string `json:"from_phone" form:"from_phone"`         //发表用户
	ToPhone      string `json:"to_phone" form:"to_phone"`             //回复用户
	FromUserName string `json:"from_user_name" form:"from_user_name"` //用户名
	FromAvatar   string `json:"from_avatar" form:"from_avatar"`       //用户头像名
	FromMvId     int    `json:"from_mv_id" form:"from_mv_id"`         //所属电影id
	Content      string `json:"content" form:"content"`               //内容
	DateTime     string `json:"date_time" form:"date_time"`           //时间
}

//MvUnderLongComment 表的映射
type MvUnderLongComment struct {
	Id        int    `json:"id"  form:"id" gorm:"type:INT;NOT NULL"`                        //影评唯一标识
	FromPhone string `json:"from_phone" form:"from_phone" gorm:"type:VARCHAR(11);NOT NULL"` //发表用户
	ToPhone   string `json:"to_phone" form:"to_phone" gorm:"type:VARCHAR(11);NOT NULL"`     //回复用户
	FromMvId  int    `json:"from_mv_id" form:"from_mv_id" gorm:"type:INT;NOT NULL"`         //所属电影id
	Content   string `json:"content" form:"content" gorm:"type:TEXT;NOT NULL"`              //内容
	DateTime  string `json:"date_time" form:"date_time" gorm:"type:VARCHAR(25);NOT NULL"`   //时间
}

func (MvUnderLongComment) TableName() string {
	return "under_long_comment"
}
