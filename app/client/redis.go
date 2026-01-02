package client

import (
	"encoding/json"

	"github.com/core-wallet/app/config"
	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/utils"
)

func GetUserSession(token string, useCase string) (dto.UserLoginResponse, error) {
	var user dto.UserLoginResponse
	data, err := config.RedisClient.Get(config.Ctx, "auth_token:"+token).Bytes()
	if err != nil {
		return user, err
	}
	utils.InfoLog("Get user session from redis for use case: " + useCase)
	err = json.Unmarshal(data, &user)
	return user, err
}
