package cache

import (
	"context"
	"minitok/cmd/feed/global"
)

func GetFollowState(ctx context.Context, uid, followUid int64) (bool, error) {
	key := getFollowInfoKey(uid)
	return global.RedisClient.SIsMember(ctx, key, followUid).Result()
}
