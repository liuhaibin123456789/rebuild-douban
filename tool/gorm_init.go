package tool

import (
	"douban/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDb *gorm.DB //gorm的db对象

func LinkMysql() error {
	//连接数据库，关闭默认启动事务
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/gorm_db?charset=utf8mb4&loc=Local&parseTime=true"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return err
	}
	GDb = db
	return nil
}

// CreateTables 初始化表格
func CreateTables() error {
	//校验表是否已经存在，存在就不创建
	tx := GDb.Begin()

	//遇到panic时，回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//用户隐私数据表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.User{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//用户非隐私数据表
	if !tx.Migrator().HasTable(&model.UserSide{}) {
		err := tx.AutoMigrate(&model.UserSide{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//用户想看或看过的影视表
	if !tx.Migrator().HasTable(&model.UserMovie{}) {
		err := tx.AutoMigrate(&model.UserMovie{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//电影表
	if !tx.Migrator().HasTable(&model.Movie{}) {
		err := tx.AutoMigrate(&model.Movie{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//演职员表
	if !tx.Migrator().HasTable(&model.Staff{}) {
		err := tx.AutoMigrate(&model.Staff{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//影人作品对照表表
	if !tx.Migrator().HasTable(&model.MovieStaff{}) {
		err := tx.AutoMigrate(&model.MovieStaff{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//影视短评表
	if !tx.Migrator().HasTable(&model.MvShortComment{}) {
		err := tx.AutoMigrate(&model.MvShortComment{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//影视长评表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.MvLongComment{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//影评下的讨论表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.MvUnderLongComment{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//电影讨论帖子表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.MvDiscuss{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//电影讨论下对话表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.MvUnderDiscuss{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
