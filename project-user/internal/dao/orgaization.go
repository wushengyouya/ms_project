package dao

import (
	"context"

	"github.com/wushengyouya/project-user/internal/data/organization"
	"github.com/wushengyouya/project-user/internal/database"
	"github.com/wushengyouya/project-user/internal/database/gorms"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}
func (o *OrganizationDao) SaveOrganization(conn database.DbConn, ctx context.Context, org *organization.Organization) error {
	o.conn = conn.(*gorms.GormConn)
	return o.conn.Tx(ctx).Create(org).Error
}
func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error) {
	orgs := make([]*organization.Organization, 0)
	err := o.conn.Default(ctx).Where("member_id = ?", memId).Find(&orgs).Error
	return orgs, err
}
