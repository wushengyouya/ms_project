package repo

import (
	"context"

	"github.com/wushengyouya/project-project/internal/data/project"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error)
}
