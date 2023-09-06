package dal

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"minitok/cmd/user/constant"
	"minitok/cmd/user/dal/cache"
	"minitok/cmd/user/dal/db"
	"minitok/cmd/user/global"
	"minitok/model"
	"minitok/pkg/errno"
)

// QueryUserInfoById 查询用户信息: 包含基本信息 (目前的场景为只读信息, 而实际情况应该为读多写少) 和计数信息 (经常会变更的数据);
// 数据查询顺序为先查Redis, 当Redis未命中时查MySQL, MySQL也未命中时使用空值缓存应对Redis缓存穿透问题;
// 当Redis命中时就更新一次过期时间, 防止出现Redis缓存雪崩的问题
func QueryUserInfoById(ctx context.Context, uid int64) (*model.User, error) {
	// 1. 空值缓存查询
	if err := cache.GetUserInfoNullKey(ctx, uid); err == nil {
		return nil, errno.UserNotRegisterErr
	} else {
		if !errors.Is(err, redis.Nil) {
			klog.Errorf("redis query error: %v\n", err)
			return nil, errno.ServiceErr.WithMessage(err.Error())
		}
	}
	// 2. 基础信息查询
	userInfo, err := cache.GetUserInfo(ctx, uid)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			klog.Errorf("redis query error: %v\n", err)
		}
		// 缓存未命中 or 缓存查询失败 -> 查询MySQL
		userInfo, err = db.QueryUserInfoByID(ctx, uid, constant.UserBaseInfoQueryString)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 未找到该UID的用户, 设置空值缓存防止Redis缓存穿透
				_ = cache.NewUserInfoNullKey(ctx, uid, global.ExpireDurationNullKey)
				return nil, errno.UserNotRegisterErr
			}
			klog.Errorf("gorm query error: %v\n", err)
			return nil, errno.ServiceErr.WithMessage(err.Error())
		} else {
			// 添加缓存信息
			if err = cache.NewUserInfos(ctx, []*model.User{userInfo}, global.ExpireDurationUserBaseInfo); err != nil {
				klog.Errorf("redis add error: %v\n", err)
			}
		}
	}
	if len(userInfo.Avatar) == 0 {
		userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
	}
	if len(userInfo.BackgroundImage) == 0 {
		userInfo.BackgroundImage = global.Configs.StaticResource.DefaultBackgroundImage
	}
	// 3. 计数信息查询
	userCounter, err := cache.GetUserCounter(ctx, uid)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			klog.Errorf("redis query error: %v\n", err)
		}
		userCounter, err = db.QueryUserInfoByID(ctx, uid, constant.UserCounterInfoQueryString)
		if err != nil {
			return nil, errno.ServiceErr.WithMessage(err.Error())
		} else {
			// 添加缓存信息（忽略错误信息）
			if err = cache.NewUserCounters(ctx, []*model.User{userCounter}); err != nil {
				klog.Errorf("redis add error: %v\n", err)
			}
		}
	}
	// 4. 合并基础信息和计数信息
	userInfo.FollowCount = userCounter.FollowCount
	userInfo.FollowerCount = userCounter.FollowerCount
	userInfo.FavoriteCount = userCounter.FavoriteCount
	userInfo.WorkCount = userCounter.WorkCount
	userInfo.TotalFavorited = userCounter.TotalFavorited
	return userInfo, nil
}
