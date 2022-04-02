package dao

import (
	"douban/global"
	"douban/model"
	"douban/tool"
	"fmt"
)

//InsertUser 同时插入隐私表、非隐私表数据
func InsertUser(user1 model.User, user2 model.UserSide) error {
	//tx, err := tool.DB.Begin()
	//if err != nil {
	//	return err
	//}
	//insertStr1 := "insert into `user_info` (user_name,password,phone,question,answer)values(?,?,?,?,?)"
	//_, err1 := tx.Exec(insertStr1, user1.UserName, user1.Password, user1.Phone, user1.Question, user1.Answer)
	//insertStr2 := "insert into `user_side` (user_name,phone,avatar,user_introduction,user_sign,register_time)values(?,?,default ,default ,default,? )"
	//_, err2 := tx.Exec(insertStr2, user2.UserName, user2.Phone, user2.RegisterTime)
	//if err1 == nil && err2 == nil {
	//	err = tx.Commit()
	//	return err
	//}
	//err = tx.Rollback()
	//return err2
	//
	//
	tx := tool.GDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Model(&model.User{}).Create(&user1).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.UserSide{}).Create(&user2).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateUserPwd(phone, userPwd string) error {
	//updateStr := "update `user_info` set password=? where phone=?"
	//_, err := tool.DB.Exec(updateStr, userPwd, phone)
	//return err
	return tool.GDb.Model(&model.User{}).Where("phone=?", phone).Update("password", userPwd).Error

}

func SelectUserPwd(phone string) (string, error) {
	//var pwd string
	//selectStr := "select password from `user_info` where phone =?"
	//err := tool.DB.QueryRow(selectStr, phone).Scan(&pwd)
	//return pwd, err
	var pwd string
	if err := tool.GDb.Model(&model.User{}).Select("password").Where("phone=?", phone).Find(&pwd).Error; err != nil {
		return "", err
	}
	return pwd, nil
}

func SelectUserPhone(phone string) (string, error) {
	var un string
	//selectStr := "select user_name from `user_info` where phone=?"
	//err := tool.DB.QueryRow(selectStr, phone).Scan(&un)
	//return un, err

	if err := tool.GDb.Model(&model.UserSide{}).Select("user_name").Where("phone=?", phone).Find(&un).Error; err != nil {
		return "", err
	}
	return un, nil
}

func SelectUserAnswer(phone string) (string, error) {
	var answer string
	//selectStr := "select answer from `user_info` where phone =?"
	//err := tool.DB.QueryRow(selectStr, phone).Scan(&answer)
	//return answer, err
	if err := tool.GDb.Model(&model.User{}).Select("answer").Where("phone=?", phone).Find(&answer).Error; err != nil {
		return "", err
	}
	return answer, nil

}

func SelectUserQuestion(phone string) (string, error) {
	var question string
	//selectStr := "select question from `user_info` where phone =?"
	//err := tool.DB.QueryRow(selectStr, phone).Scan(&question)
	//return question, err

	if err := tool.GDb.Model(&model.User{}).Select("question").Where("phone=?", phone).Find(&question).Error; err != nil {
		return "", err
	}
	return question, nil
}

func InsertUserIntroduction(phone, introduction string) error {
	//str := "update `user_side` set user_introduction =? where phone=?"
	//_, err := tool.DB.Exec(str, introduction, phone)
	//return err
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("user_introduction", introduction).Error
}

func InsertUserSign(phone, sign string) error {
	//str := "update `user_side` set user_sign = ? where phone=?"
	//_, err := tool.DB.Exec(str, sign, phone)
	//return err
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("user_sign", sign).Error
}

func InsertUserAvatar(phone, filename string) error {
	//str := "update `user_side` set avatar = ? where phone=?"
	//_, err := tool.DB.Exec(str, filename, phone)
	//return err
	return tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Update("avatar", filename).Error
}

func SelectUserSide(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	//str := "select * from `user_side` where phone=?"
	//err := tool.DB.QueryRow(str, phone).Scan(
	//	&us.UserId,
	//	&us.UserName,
	//	&us.Phone,
	//	&us.Avatar,
	//	&us.UserIntroduction,
	//	&us.UserSign,
	//	&us.RegisterTime,
	//)
	////拼接路径
	//fileName := us.Avatar
	//dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	//us.Avatar = dst
	//return us, err
	if err := tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Find(&us).Error; err != nil {
		return model.UserSide{}, err
	}
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, nil
}

func SelectOInfo(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	//str := "select * from `user_side` where phone=?"
	//err := tool.DB.QueryRow(str, phone).Scan(
	//	&us.UserId,
	//	&us.UserName,
	//	&us.Phone,
	//	&us.Avatar,
	//	&us.UserIntroduction,
	//	&us.UserSign,
	//	&us.RegisterTime,
	//)
	////拼接路径
	//fileName := us.Avatar
	//dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	//us.Avatar = dst
	//return us, err
	if err := tool.GDb.Model(&model.UserSide{}).Where("phone=?", phone).Find(&us).Error; err != nil {
		return model.UserSide{}, err
	}
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath + fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, nil
}

func SelectWODMVs(label, phone string) ([]model.OfMovie, error) {
	var ids = make([]int, 0)
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	////查询user_movie表获取想看或看过的电影id
	//sql1 := "select mv_id from `user_movie` where label=? and phone=?"
	//rows, err := tool.DB.Query(sql1, label, phone)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err = rows.Scan(&id)
	//	if err != nil {
	//		return nil, err
	//	}
	//	ids = append(ids, id)
	//}
	////依据id,查询影视具体数据
	//var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum float32
	//sql2 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_id=?"
	//for _, id := range ids {
	//	err = tool.DB.QueryRow(sql2, id).Scan(
	//		&ofMv.Id,
	//		&ofMv.Name,
	//		&ofMv.Picture,
	//		&ofMv.Director,
	//		&ofMv.LeadRole,
	//		&ofMv.ProduceWhere,
	//		&ofMv.ReleaseTime,
	//		&ofMv.Duration,
	//		&oneStarNum,
	//		&twoStarNum,
	//		&threeStarNum,
	//		&fourStarNum,
	//		&fiveStarNum,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	//四舍五入
	//	sc := (oneStarNum+twoStarNum*2+threeStarNum*3+4*fourStarNum+5*fiveStarNum)/(oneStarNum+twoStarNum+threeStarNum+fourStarNum+fiveStarNum) + 0.5
	//	ofMv.Score = int(sc)
	//	//拼接图片路径
	//	ofMv.Picture = global.MvPicturePath + fmt.Sprintf("%s", ofMv.Picture)
	//	ofMvs = append(ofMvs, ofMv)
	//}
	//return ofMvs, nil
	if err := tool.GDb.Model(&model.UserMovie{}).Select("mv_id").Where("label=? and phone=?", label, phone).Find(&ids).Error; err != nil {
		return []model.OfMovie{}, err
	}
	for _, id := range ids {
		mv := model.Movie{}
		if err := tool.GDb.Model(&model.Movie{}).Where("id=?", id).Find(&mv).Error; err != nil {
			return []model.OfMovie{}, err
		}
		//数据迁移
		sc := (mv.OneStarNum + mv.TwoStarNum*2 + mv.ThreeStarNum*3 + 4*mv.FourStarNum + 5*mv.FiveStarNum) / (mv.OneStarNum + mv.TwoStarNum + mv.ThreeStarNum + mv.FourStarNum + mv.FiveStarNum)
		ofMv.Score = sc
		ofMv.Id = mv.Id
		ofMv.Name = mv.Name
		ofMv.Director = mv.Director
		ofMv.LeadRole = mv.LeadRole
		ofMv.Duration = mv.Duration
		ofMv.ProduceWhere = mv.ProduceWhere
		ofMv.ReleaseTime = mv.ReleaseTime
		//拼接图片路径
		ofMv.Picture = global.MvPicturePath + fmt.Sprintf("%s", mv.Picture)
		ofMvs = append(ofMvs, ofMv)
	}
	return ofMvs, nil
}

func SelectLComments(phone string) ([]model.LongComment, error) {
	var lcs = make([]model.LongComment, 0)
	var mlcs = make([]model.MvLongComment, 0)
	//sql0 := "select user_name,avatar from `user_side` where phone= ?;"
	//sql1 := "select * from `movie_long_comment` where from_phone = ?;"
	//rows, err := tool.DB.Query(sql1, phone)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err = rows.Scan(
	//		&lc.Id,
	//		&lc.FromPhone,
	//		&lc.FromMvId,
	//		&lc.MvStar,
	//		&lc.Title,
	//		&lc.Content,
	//		&lc.DateTime,
	//		&lc.UsedNum,
	//		&lc.UnusedNum,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	err = tool.DB.QueryRow(sql0, lc.FromPhone).Scan(&lc.FromUserName, &lc.FromAvatar)
	//	if err != nil {
	//		return nil, err
	//	}
	//	//拼接图片路径
	//	lc.FromAvatar = global.UserAvatarPath + fmt.Sprintf("%s", lc.FromAvatar)
	//
	//	lcs = append(lcs, lc)
	//}
	//return lcs, nil
	if err := tool.GDb.Model(&model.MvLongComment{}).Where("from_phone=?", phone).Find(&mlcs).Error; err != nil {
		return []model.LongComment{}, err
	}
	for _, v := range mlcs {
		lc := model.LongComment{}
		if err := tool.GDb.Model(&model.UserSide{}).Select("user_name", "avatar").Where("phone=?", phone).Find(&lc.FromUserName, &lc.FromAvatar).Error; err != nil {
			return []model.LongComment{}, err
		}
		//数据迁移
		lc.FromAvatar = global.UserAvatarPath + fmt.Sprintf("%s", lc.FromAvatar)
		lc.Id = v.Id
		lc.FromPhone = v.FromPhone
		lc.Content = v.Content
		lc.DateTime = v.DateTime
		lc.FromMvId = v.FromMvId
		lc.MvStar = v.MvStar
		lc.Title = v.Title
		lc.UsedNum = v.UsedNum
		lc.UnusedNum = v.UnusedNum
		lcs = append(lcs, lc)
	}
	return lcs, nil
}
