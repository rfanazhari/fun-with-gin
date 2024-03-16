package login_token

import (
	"fmt"
	"fun-with-gin/domain/service"
)

func (s redisLoginToken) key(key service.RedisKey, dynamicKey string) string {
	return fmt.Sprintf("%s_%s", key, dynamicKey)
}
