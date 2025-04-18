package dao

import (
	"context"

	"github.com/wushengyouya/project-user/internal/data/member"
	"github.com/wushengyouya/project-user/internal/database"
	"github.com/wushengyouya/project-user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}

// SaveMember 保存Member到数据库
func (m *MemberDao) SaveMember(conn database.DbConn, ctx context.Context, member *member.Member) error {
	m.conn = conn.(*gorms.GormConn)
	return m.conn.Tx(ctx).Create(member).Error
}

// GetMemberByAccount 根据account查询是否存在member
func (m *MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("account = ?", account).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("mobile = ?", mobile).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) FindMember(ctx context.Context, account, password string) (*member.Member, error) {
	mem := new(member.Member)
	err := m.conn.Default(ctx).Model(&member.Member{}).Where(&member.Member{Account: account, Password: password}).First(mem).Error
	return mem, err
}
