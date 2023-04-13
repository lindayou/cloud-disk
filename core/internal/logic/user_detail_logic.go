package logic

import (
	"cloud-disk/model"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = new(types.UserDetailResponse)
	detail := new(model.UserDetail)
	detail.Sex = req.Sex
	detail.Tall = req.Tall
	detail.Hobbit = req.Hobbit
	detail.Pid = req.Id
	detail.Email = req.Email
	err = l.svcCtx.DB.Create(detail).Error
	if err != nil {
		return nil, errors.New("创建个人信息失败")
	}
	resp.Message = "创建个人信息成功"
	return
}
