package boilerplate_dbms

import (
	"fmt"
	"time"
	"runtime"
	"log/syslog"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	Index     int       `xorm:"index pk autoincr unique" json:"index"`
	CreatedAt time.Time `xorm:"created"                  json:"created_at"`
	UpdatedAt time.Time `xorm:"updated"                  json:"updated_at"`
}

func init() {
	engine := connectKeyDB()
	syncKeyDB(engine)
}

func connectKeyDB() *xorm.Engine {
	var err error
	var engine *xorm.Engine

	DbURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		DB_ACCOUNT,
		DB_PASSWORD,
		DB_SERVER,
		DB_PORT,
		DB_TABLE,
		DB_SETTING,
	)
	if runtime.GOOS == "darwin" {
		DbURI = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			DB_DEV_ACCOUNT,
			DB_DEV_PASSWORD,
			DB_DEV_SERVER,
			DB_DEV_PORT,
			DB_DEV_TABLE,
			DB_DEV_SETTING,
		)
	}

	logWriter, err := syslog.New(syslog.LOG_DEBUG, "key-db")
	if err != nil {
		panic(err)
	}

	xormLog := xorm.NewSimpleLogger(logWriter)
	xormLog.ShowSQL(true)
	xormLog.SetLevel(core.LOG_DEBUG)
	xormLog.IsShowSQL()

	engine, err = xorm.NewEngine("mysql", DbURI)
	engine.SetLogger(xormLog)
	engine.TZLocation, _ = time.LoadLocation("Asia/Seoul")

	return engine
}

func syncKeyDB( engine *xorm.Engine ) {
	if err := engine.Sync2(new(Storage)); err != nil {
		panic(err)
	}
}
