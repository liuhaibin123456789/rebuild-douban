package service

import (
	"context"
	"douban/dao"
	"douban/global"
	"douban/model"
	"douban/proto"
	"douban/tool"
	"errors"
	"log"
)

func Register(user model.User) (string, error) {
	//调用微服务
	resToken, err := tool.GrpcClient.Register(context.Background(), &proto.UserInfo{
		UserName: user.UserName,
		Password: user.Password,
		Phone:    user.Phone,
		Question: user.Question,
		Answer:   user.Answer,
	})
	if err != nil {
		return "", err
	}

	return resToken.GetToken(), nil
}

func CreateIntroduction(phone, introduction string) error {
	_, err := tool.GrpcClient.CreateIntroduction(
		context.Background(),
		&proto.ReqUser{One: &proto.ReqUser_Introduction{Introduction: introduction}},
	)
	return err
}

func CreateSign(phone, sign string) error {
	_, err := tool.GrpcClient.CreateSign(
		context.Background(),
		&proto.ReqUser{One: &proto.ReqUser_Sign{Sign: sign}},
	)
	return err
}

func UploadUserAvatar(phone, filename string) error {
	//todo 删除原有头像文件

	_, err := tool.GrpcClient.CreateAvatar(
		context.Background(),
		&proto.ReqUser{One: &proto.ReqUser_Avatar{Avatar: filename}},
	)
	return err
}

func GetInfo(phone string) (model.UserSide, error) {
	u := model.UserSide{}

	user, err := tool.GrpcClient.GetUser(
		context.Background(),
		&proto.ReqUser{Phone: phone},
	)
	if err != nil {
		return u, err
	}
	u.UserId = user.UserId
	u.Phone = user.Phone
	u.Avatar = global.UserAvatarPath + user.Avatar
	u.UserName = user.UserName
	u.UserIntroduction = user.UserIntroduction
	u.UserSign = user.UserSign
	return u, nil
}

func GetOInfo(phone string) (model.UserSide, error) {
	u := model.UserSide{}

	user, err := tool.GrpcClient.GetUser(
		context.Background(),
		&proto.ReqUser{Phone: phone},
	)
	if err != nil {
		return u, err
	}
	u.UserId = user.UserId
	u.Phone = user.Phone
	u.Avatar = global.UserAvatarPath + user.Avatar
	u.UserName = user.UserName
	u.UserIntroduction = user.UserIntroduction
	u.UserSign = user.UserSign
	return u, nil
}

func GetRandCode(phone string) (string, error) {
	resCode, err := tool.GrpcClient.GetCode(context.Background(), &proto.ReqUser{Phone: phone})
	if err != nil {
		return "-1", err
	}
	return resCode.Code, nil
}

func LoginByPwd(phone, password string) (string, error) {

	token, err := tool.GrpcClient.Login(context.Background(), &proto.ReqUser{
		Phone: phone,
		One:   &proto.ReqUser_Password{Password: password},
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token.GetToken(), nil
}

func UpdatePwd(phone, oldPassword, newPassword string) error {
	_, err := tool.GrpcClient.UpdatePwd(context.Background(), &proto.ReqUser{
		Phone: phone,
		One:   &proto.ReqUser_Password{Password: oldPassword},
		Two:   &proto.ReqUser_OldPassword{OldPassword: newPassword},
	})
	return err
}

func GetAnswer(phone string) (string, error) {
	resUser, err := tool.GrpcClient.GetAnswer(context.Background(), &proto.ReqUser{Phone: phone})
	if err != nil {
		return "", err
	}
	return resUser.GetAnswer(), nil
}

func GetQuestion(phone string) (string, error) {
	resUser, err := tool.GrpcClient.GetQuestion(context.Background(), &proto.ReqUser{Phone: phone})
	if err != nil {
		return "", err
	}
	return resUser.GetQuestion(), nil
}

func GetWODMvs(label string, phone string) ([]model.OfMovie, error) {
	if label != "0" && label != "1" {
		return nil, errors.New("the label must be '0' or '1'")
	}
	wodMVs, err := dao.SelectWODMVs(label, phone)
	return wodMVs, err
}

func GetLComments(phone string) ([]model.LongComment, error) {
	lComments, err := dao.SelectLComments(phone)
	return lComments, err
}
