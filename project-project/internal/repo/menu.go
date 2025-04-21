package repo

import (
	"context"

	"github.com/wushengyouya/project-project/internal/data/menu"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*menu.ProjectMenu, error)
}
