package logic

import (
	"cloud-disk/model"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest) (resp *types.UserUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserUpdateResponse)
	updUser := new(model.UserDetail)
	updUser.Hobbit = req.Hobbit
	updUser.Tall = req.Tall
	updUser.Sex = req.Sex

	err = l.svcCtx.DB.Model(&model.UserDetail{}).Where("pid =?", req.Id).Updates(updUser).Error
	if err != nil {
		return nil, errors.New("更新失败")
	}
	resp.Message = "更新成功"
	return
}
