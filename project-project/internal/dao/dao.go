package dao

import (
	"github.com/wushengyouya/project-project/internal/database"
	"github.com/wushengyouya/project-project/internal/database/gorms"
)

// 通过模板方法 Action 抽象通用逻辑，业务代码只需传入操作函数 f。
// 业务代码只需关注核心逻辑，减少事务样板代码。
type TransactionImpl struct {
	conn database.DbConn
}

func (t *TransactionImpl) Action(f func(conn database.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.NewTran(),
	}
}
