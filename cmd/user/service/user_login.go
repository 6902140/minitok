package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"minitok/cmd/user/constant"
	"minitok/cmd/user/dal/cache"
	"minitok/cmd/user/dal/db"
	"minitok/cmd/user/global"
	"minitok/kitex_gen/douyin/user"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}

func (s *UserLoginService) UserLogin(req *user.LoginRequest) (int64, string, error) {
	// 1. 查询该Username是否在存在空值缓存
	if err := cache.GetUserLoginNullKey(s.ctx, req.Username); err == nil {
		return 0, "", errno.UserNotRegisterErr
	} else {
		if !errors.Is(err, redis.Nil) {
			// 非致命错误，无需返回
			klog.Errorf("redis query error: %v\n", err)
		}
	}
	// 2. 判断用户是否存在
	userInfo, err := db.QueryFirstUserInfoByUsername(s.ctx, req.Username, constant.UserLoginInfoQueryString)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户未注册, 添加空值缓存
			if err = cache.NewUserLoginNullKey(
				s.ctx, req.Username, global.Configs.CacheExpire.ParseNullKeyExpireDuration()); err != nil {
				klog.Errorf("redis add error: %v\n", err)
			}
			return 0, "", errno.UserNotRegisterErr
		} else {
			// 其他错误
			klog.Errorf("gorm query error: %v\n", err)
			return 0, "", errno.ServiceErr.WithMessage(err.Error())
		}
	}
	// 3. 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password)); err != nil {
		return 0, "", errno.AuthorizationFailedErr
	}
	// 4. 根据用户id颁发token
	claims, err := jwt.BuildCustomClaims(
		userInfo.Id, global.Configs.JWT.ExpiresTime, global.Configs.JWT.Issuer, global.Configs.JWT.Subject)
	if err != nil {
		klog.Errorf("jwt error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	token, err := jwt.NewJWT(global.Configs.JWT.SigningKey).CreateToken(claims)
	if err != nil {
		klog.Errorf("jwt error: %v\n", err)
		return 0, "", errno.ServiceErr.WithMessage(err.Error())
	}
	return userInfo.Id, token, nil
}
