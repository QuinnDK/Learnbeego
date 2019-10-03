package controllers

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

const (
	DEFAULT_DB = "default"

	// 设置数据库 driver
	MYSQL_DRIVER = "mysql"
)

//func RegisterDB() error {
//	dbUser := beego.AppConfig.String("mysqluser")
//	dbPass := beego.AppConfig.String("mysqlpass")
//	dbHost := beego.AppConfig.String("mysqlhost")
//	dbPort := beego.AppConfig.String("mysqlport")
//	dbName := beego.AppConfig.String("mysqldb")
//}

type Users struct {
	Id   int
	Name string
	Pwd  string
	Age  int
	Sex  string
}

//初始化
func init() {
	//注册数据驱动
	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	// mysql / sqlite3 / postgres 这三种是beego默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:03231412@tcp(127.0.0.1:3306)/HelloBeego")
	//注册模型
	orm.RegisterModel(new(Users))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", true, true)
}

//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)
// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
// 参数2        数据库驱
// 参数3        对应的连接字符串
// 参数4(可选)  设置最大空闲连接
// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
