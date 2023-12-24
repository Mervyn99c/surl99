package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"starry/dao/surl/model"
	"starry/internal/config"
	"starry/internal/middleware"
)

type ServiceContext struct {
	Config        config.Config
	UrlModel      model.URLMapping
	BizRedis      *redis.Redis
	DbEngine      *gorm.DB
	CTRMiddleWare rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "",
		SingularTable: true,
	}})
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(db)

	//conn := sqlx.NewMysql(c.DataSource)
	//mysql := orm.NewMysql()
	return &ServiceContext{
		Config:        c,
		BizRedis:      redis.MustNewRedis(c.BizRedis),
		DbEngine:      db,
		CTRMiddleWare: middleware.NewCTRMiddleWareMiddleware().Handle,
	}
}
