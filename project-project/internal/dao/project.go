package dao

import (
	"context"

	"github.com/wushengyouya/project-project/internal/data/project"
	"github.com/wushengyouya/project-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	var pms []*project.ProjectAndMember
	session := p.conn.Seesion(ctx)
	index := (page - 1) * size
	raw := session.Raw("select * from ms_project a,ms_project_member b where a.id = b.project_code and b.member_code = ? limit ?,?", memId, index, size)

	raw.Scan(&pms)
	var total int64
	err := session.Model(&project.ProjectMember{}).Where("member_code = ?", memId).Count(&total).Error
	return pms, total, err

}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}
