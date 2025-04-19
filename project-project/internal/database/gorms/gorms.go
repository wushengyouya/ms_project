package gorms

import (
	"context"
	"fmt"

	"github.com/wushengyouya/project-project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	username := config.AppConf.MysqlConfig.UserName
	password := config.AppConf.MysqlConfig.Password
	host := config.AppConf.MysqlConfig.Host
	port := config.AppConf.MysqlConfig.Port
	dbName := config.AppConf.MysqlConfig.Db

	// 创建数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败,err: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	DB *gorm.DB
	tx *gorm.DB
}

func New() *GormConn {
	return &GormConn{DB: GetDB()}
}
func NewTran() *GormConn {
	return &GormConn{DB: GetDB(), tx: GetDB()}
}

// Default 用于基于现有连接 创建一个新的独立会话。新会话会继承原始连接的配置（如数据库类型、连接池设置）
// 但允许覆盖特定参数。
// &gorm.Session{Context: ctx}：配置会话参数，将传入的上下文 ctx 注入到新会话中。
// 这使得通过该会话执行的所有数据库操作（如查询、事务）都关联到该上下文
func (g *GormConn) Default(ctx context.Context) *gorm.DB {
	// gorm.Session 实现了 上下文绑定 和 会话隔离 的核心功能，适用于需要精细化控制数据库操作的场景。
	return g.DB.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Begin() {
	g.tx.Begin()
}
func (g *GormConn) Rollback() {
	g.tx.Rollback()
}
func (g *GormConn) Commit() {
	g.tx.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
