package tran

import "github.com/wushengyouya/project-project/internal/database"

// Transaction 事务的操作 一定跟数据库有关， 注入数据库的链接 gorm.db
type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
