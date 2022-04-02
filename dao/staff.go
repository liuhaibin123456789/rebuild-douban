package dao

import (
	"douban/global"
	"douban/model"
	"douban/tool"
	"fmt"
)

func SelectStaffInfo(id int) (model.Staff, error) {
	staff := model.Staff{}
	//sql := "select * from staff where staff_id=?"
	//err := tool.DB.QueryRow(sql, id).Scan(
	//	&staff.Id,
	//	&staff.Name,
	//	&staff.Sex,
	//	&staff.Avatar,
	//	&staff.Constellation,
	//	&staff.Birthday,
	//	&staff.Birthplace,
	//	&staff.Jobs,
	//	&staff.ACName,
	//	&staff.AEName,
	//	&staff.Family,
	//	&staff.Imdb,
	//	&staff.Introduction,
	//)
	//拼接图片路径
	//staff.Avatar = global.PerfPicturePath + fmt.Sprintf("%s", staff.Avatar)
	//return staff, err
	if err := tool.GDb.Model(&staff).Where("id=?", id).Find(&staff).Error; err != nil {
		return model.Staff{}, err
	}
	//拼接图片路径
	staff.Avatar = global.PerfPicturePath + fmt.Sprintf("%s", staff.Avatar)
	return staff, nil
}

func SelectStaffsOfMv(mvId int) ([]model.OfStaff, error) {
	ids := make([]int, 0)
	var staff model.OfStaff
	staffsOfMv := make([]model.OfStaff, 0)
	////从movie_staff数据表获取对应影视下的所有演职员id
	//sql1 := "select staff_id from `movie_staff` where mv_id=?;"
	//rows, err := tool.DB.Query(sql1, mvId)
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
	////根据id遍历staff数据表获取演职员详细数据
	//sql2 := "select staff_id,staff_name,staff_avatar,staff_jobs from `staff` where staff_id=?"
	//for _, id := range ids {
	//	err = tool.DB.QueryRow(sql2, id).Scan(
	//		&staff.Id,
	//		&staff.Name,
	//		&staff.Avatar,
	//		&staff.Jobs,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	//拼接路径
	//	staff.Avatar = global.PerfPicturePath + fmt.Sprintf("%s", staff.Avatar)
	//	staffsOfMv = append(staffsOfMv, staff)
	//}
	//return staffsOfMv, nil
	if err := tool.GDb.Model(&model.MovieStaff{}).Select("staff_id").Where("mv_id=?", mvId).Find(&ids).Error; err != nil {
		return []model.OfStaff{}, err
	}
	for _, id := range ids {
		if err := tool.GDb.Model(&model.Staff{}).Select("id", "name", "avatar", "jobs").Where("id=?", id).Find(&staff).Error; err != nil {
			return []model.OfStaff{}, err
		}
		//拼接路径
		staff.Avatar = global.PerfPicturePath + fmt.Sprintf("%s", staff.Avatar)
		staffsOfMv = append(staffsOfMv, staff)
	}
	return staffsOfMv, nil
}

func SelectMvsOfStaff(sId int) ([]model.OfMovie, error) {
	var ids = make([]int, 0)
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	////从movie_staff数据表获取对应影视下的所有演职员id
	//sql1 := "select mv_id from `movie_staff` where staff_id =?;"
	//rows, err := tool.DB.Query(sql1, sId)
	//if err != nil {
	//	log.Println("1 ", err)
	//	return nil, err
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err = rows.Scan(&id)
	//	if err != nil {
	//		log.Println("2 ", err)
	//		return nil, err
	//	}
	//	ids = append(ids, id)
	//}
	//log.Println(ids)
	////根据id遍历movie数据表获取电影数据
	//var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum int
	//sql2 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_id=?"
	//for _, id := range ids {
	//	err = tool.DB.QueryRow(sql2, id).Scan(
	//		&mv.Id,
	//		&mv.Name,
	//		&mv.Picture,
	//		&mv.Director,
	//		&mv.LeadRole,
	//		&mv.ProduceWhere,
	//		&mv.ReleaseTime,
	//		&mv.Duration,
	//		&oneStarNum,
	//		&twoStarNum,
	//		&threeStarNum,
	//		&fourStarNum,
	//		&fiveStarNum,
	//	)
	//	if err != nil {
	//		log.Println("3 ", err)
	//		return nil, err
	//	}
	//	//防止除0错误
	//	if oneStarNum+twoStarNum+threeStarNum+fourStarNum+fiveStarNum == 0 {
	//		mv.Score = 0
	//	} else {
	//		//合并
	//		mv.Score = (oneStarNum + twoStarNum*2 + threeStarNum*3 + 4*fourStarNum + 5*fiveStarNum) / (oneStarNum + twoStarNum + threeStarNum + fourStarNum + fiveStarNum)
	//	}
	//	//拼接路径
	//	mv.Picture = global.MvPicturePath + fmt.Sprintf("%s", mv.Picture)
	//	mvs = append(mvs, mv)
	//}
	//return mvs, nil
	if err := tool.GDb.Model(&model.MovieStaff{}).Select("mv_id").Where("staff_id=?", sId).Find(&ids).Error; err != nil {
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
