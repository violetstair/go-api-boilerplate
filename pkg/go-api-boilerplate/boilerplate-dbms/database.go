package dbms

import (
	"fmt"
	"log/syslog"
	"runtime"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

// Storage : Storage
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
		DBAccount,
		DBPassword,
		DBServer,
		DBPort,
		DBTable,
		DBSetting,
	)
	if runtime.GOOS == "darwin" {
		DbURI = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			DBDevAccount,
			DBDevPassword,
			DBDevServer,
			DBDevPort,
			DBDevTable,
			DBDevSetting,
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

func syncKeyDB(engine *xorm.Engine) {
	if err := engine.Sync2(new(Storage)); err != nil {
		panic(err)
	}
}
