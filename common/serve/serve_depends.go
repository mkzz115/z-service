package serve

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/idgenutil"
)

type DependsServer interface {
	GenId() (int64, error)
	DbInstance() *sqlx.DB
}

func NewDepServe(cfg *confutil.DependServeConfig) DependsServer {
	ds := &dependsDerve{
		cfg: cfg,
	}
	ds.idgen = idgenutil.NewIdWorker()
	return ds
}

type dependsDerve struct {
	cfg   *confutil.DependServeConfig
	idgen idgenutil.IdGenerator
}

func (d *dependsDerve) GenId() (int64, error) {
	return d.idgen.GenId()
}

func (d *dependsDerve) DbInstance() *sqlx.DB {
	if len(d.cfg.MysqlConfig.Dsn) == 0 {
		return nil
	}
	db, err := sqlx.Open("mysql", d.cfg.MysqlConfig.Dsn)
	if err != nil {
		panic(err)
	}
	return db
}
