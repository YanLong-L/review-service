package data

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"review-service/internal/conf"
	"review-service/internal/data/query"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewReviewRepo)

// Data .
type Data struct {
	query *query.Query
	log   *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// 非常重要!为GEN生成的query代码设置数据库连接对象
	query.SetDefault(db)

	return &Data{query: query.Q, log: log.NewHelper(logger)}, cleanup, nil
}

func NewDB(cfg *conf.Data) (*gorm.DB, error) {
	switch strings.ToLower(cfg.Database.GetDriver()) {
	case "mysql":
		return gorm.Open(mysql.Open(cfg.Database.GetSource()))
	case "sqlite":
		return gorm.Open(sqlite.Open(cfg.Database.GetSource()))
	}
	return nil, errors.New("connect db fail unsupported db driver")
}
