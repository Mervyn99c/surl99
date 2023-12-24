package surl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/pkg/xerr"
	"starry/dao/surl/model"
	"starry/internal/svc"
	"starry/internal/types"
)

type SurlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSurlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SurlLogic {
	return &SurlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SurlLogic) Router(req *types.SurlRequest) (lurl string, err error) {

	prefixKey := "SMS_BATCH_SERVICE" + ":" + "GLOBAL" + ":"
	filter := bloom.New(l.svcCtx.BizRedis, prefixKey, 64)

	key := prefixKey + req.Surl

	bexi, err := filter.Exists([]byte(key))
	fmt.Println(bexi)
	// 不存在一定不存在，直接返回
	if !bexi {
		// todo call_back_url
		return "404", errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", "um", err)
	}
	// 存在不一定存在，需要去缓存和数据库检查
	rexi, err := l.svcCtx.BizRedis.Get(key)
	if err == nil {
		return rexi, err
	} else {
		// 去查数据库
		if err != nil {
			mapping := &model.URLMapping{}
			result := l.svcCtx.DbEngine.Where("surl = ?", key).First(&mapping)
			if result.Error != nil {
				return "404", errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", "um", err)
			}
			// 重新缓存redis和布隆过滤器
			f := bloom.New(l.svcCtx.BizRedis, key, 64)
			f.Add([]byte(key))
			_, err = l.svcCtx.BizRedis.SetnxCtx(l.ctx, key, mapping.Lurl)
			if err != nil {
				return "404", errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", "um", err)
			}
			return mapping.Lurl, nil
		} else {
			return "404", errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUrlMappinig Database Exception : %+v , err: %v", "um", err)
		}
	}

}
