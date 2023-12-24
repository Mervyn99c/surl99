package admin

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/pkg/xerr"
	"starry/dao/surl/model"
	"starry/internal/svc"
	"starry/internal/types"
	su "starry/pkg/surl"
)

type GenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLogic {
	return &GenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateLogic) Generate(req *types.GenerateRequest) (resp *types.GenerateResponse, err error) {

	shortcode := su.GetShortCode()
	println(shortcode)
	var um = model.URLMapping{
		Surl:          shortcode,
		Lurl:          req.Lurl,
		EffectiveDays: req.EffectiveDays,
	}

	result := l.svcCtx.DbEngine.Create(&um)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", um, err)
	}

	prefixKey := "SMS_BATCH_SERVICE" + ":" + "GLOBAL" + ":"
	key := prefixKey + shortcode
	_, err = l.svcCtx.BizRedis.SetnxCtx(l.ctx, key, req.Lurl)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", um, err)
	}
	filter := bloom.New(l.svcCtx.BizRedis, prefixKey, 64)
	err = filter.Add([]byte(key))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", um, err)
	}

	return &types.GenerateResponse{Surl: shortcode}, nil
}
