package logic

import (
	"cloud-disk/define"
	"cloud-disk/helper"
	"cloud-disk/model"
	"context"
	"errors"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	user := new(model.UserBasic)
	resp = new(types.LoginResponse)
	l.svcCtx.DB.Where("name= ? and password =?", req.Name, req.Password).First(user)
	if user.Id > 0 {
		token, err := helper.GenerateJWTToken(define.JwtKey, user.Name, user.Id)
		if err != nil {
			log.Println("Token Generate Failed")
		}
		resp.Token = token
	} else {
		return nil, errors.New("用户名或密码错误")
	}
	return
}
