package boilerplate

import (
	"fmt"
	"time"
	"log/syslog"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	engine := connectDB()
	syncDB(engine)
}

func connectDB() *xorm.Engine {
	var err error
	var engine *xorm.Engine

	var (
		DBMS_ID = GetEnv("DBMS_ID", "account")
		DBMS_PW = GetEnv("DBMS_PW", "password")
		CONNECT = GetEnv("CONNECT", "localhost")
		TABLE   = GetEnv("TABLE",   "owdin")
		DbURI  = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?%s",
			DBMS_ID,
			DBMS_PW,
			CONNECT,
			TABLE,
			"charset=utf8&parseTime=True&loc=Local",
		)

	)

	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-api-log")
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

func syncDB( engine *xorm.Engine ) {
	if err := engine.Sync2(new(EOSInfo)); err != nil {
		panic(err)
	}
}