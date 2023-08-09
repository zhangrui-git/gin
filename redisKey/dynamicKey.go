package redisKey

import "fmt"

func UserModel(id uint) string {
	return fmt.Sprintf("UserModel:id:%d", id)
}
