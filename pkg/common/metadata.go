package common

import (
	"errors"
	"fun-with-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type MetaHttpToken struct {
	Id    uint
	Name  string
	Email string
}

func GetMetaHttpToken(ctx *gin.Context) (*MetaHttpToken, error) {
	userId := 0
	userName := ""
	userEmail := ""

	userClaims, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("unauthorized")
	}

	userId = int(userClaims.(*utils.JwtPayload).Id)
	userName = userClaims.(*utils.JwtPayload).Name
	userEmail = userClaims.(*utils.JwtPayload).Email

	if userId == 0 {
		return nil, errors.New("invalid user ID")
	}
	if userEmail == "" {
		return nil, errors.New("invalid user email")
	}
	if userName == "" {
		return nil, errors.New("invalid user name")
	}
	return &MetaHttpToken{
		Id:    uint(userId),
		Name:  userName,
		Email: userEmail,
	}, nil
}
