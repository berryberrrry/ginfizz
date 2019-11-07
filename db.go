/*
 * @Author: berryberry
 * @since: 2019-05-16 21:11:41
 * @lastTime: 2019-05-31 15:54:16
 * @LastAuthor: Do not edit
 */
package ginfizz

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func initDB() {
	// v0版本不支持数据库操作
	// switch strings.ToLower(FizzConfig.App.DB.DBType) {
	// case DBTypeMysql:
	// 	db, err = sql.Open(DBTypeMysql,
	// 		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
	// 			FizzConfig.App.DB.Username,
	// 			FizzConfig.App.DB.Password,
	// 			FizzConfig.App.DB.Host,
	// 			FizzConfig.App.DB.Port,
	// 			FizzConfig.App.DB.DBName,
	// 			FizzConfig.App.DB.Charset,
	// 		),
	// 	)
	// 	if err != nil {
	// 		Logger.Errorf("[GINFIZZ][initDB] error: %s", err)
	// 		panic(err)
	// 	}
	// 	err = db.Ping()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// case DBTypeMongo:
	// 	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// }
}
