package logic

import (
	"cloud-disk/model"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserInfoResponse)
	userDetail := new(model.UserDetail)
	err = l.svcCtx.DB.Where("pid = ?", req.Id).First(userDetail).Error
	if err != nil {
		return nil, errors.New("查询失败")
	}
	resp.Sex = userDetail.Sex
	resp.Hobbit = userDetail.Hobbit
	resp.Tall = userDetail.Tall
	resp.Eamil = userDetail.Email
	return
}
