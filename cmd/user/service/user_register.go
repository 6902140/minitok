package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"minitok/cmd/user/constant"
	"minitok/cmd/user/dal/cache"
	"minitok/cmd/user/dal/db"
	"minitok/cmd/user/global"
	"minitok/kitex_gen/douyin/user"
	"minitok/model"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

func (s *UserRegisterService) UserRegister(req *user.RegisterRequest) (int64, string, error) {
	// 1. 判断用户是否已经注册, 未注册将返回`gorm.ErrRecordNotFound`错误
	res, err := db.QueryFirstUserInfoByUsername(s.ctx, req.Username, constant.UserRegisterInfoQueryString)
	if res != nil {
		// 查询到用户信息
		return 0, "", errno.UserAlreadyExistErr
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 错误信息不是ErrRecordNotFound
		klog.Errorf("gorm query error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	// 2. 对用户密码进行加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		klog.Errorf("bcrypt error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(fmt.Sprintf("密码加密错误, %s", err.Error()))
	}
	// 3. 创建用户, 并且写入数据库
	newUserInfo := &model.User{
		//UUID:     uuid.New(),
		Username: req.Username,
		Password: string(bytes),
		Nickname: req.Username,
	}
	if err = db.CreateUserInfos(s.ctx, []*model.User{newUserInfo}); err != nil {
		klog.Errorf("gorm create error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	// 4. 若该UID或Username存在空值缓存, 则将其删除
	if err = cache.DelUserInfoNullKey(s.ctx, newUserInfo.Id); err != nil {
		klog.Errorf("redis delete error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	if err = cache.DelUserLoginNullKey(s.ctx, req.Username); err != nil {
		klog.Errorf("redis delete error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	// 5. 获得用户id, 并且颁发token
	claims, err := jwt.BuildCustomClaims(
		newUserInfo.Id, global.Configs.JWT.ExpiresTime, global.Configs.JWT.Issuer, global.Configs.JWT.Subject)
	if err != nil {
		klog.Errorf("jwt error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	token, err := jwt.NewJWT(global.Configs.JWT.SigningKey).CreateToken(claims)
	if err != nil {
		klog.Errorf("jwt error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	// 5. 返回用户id和token
	return newUserInfo.Id, token, nil
}
