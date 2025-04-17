package repo

import (
	"context"

	"github.com/wushengyouya/project-user/internal/data/member"
	"github.com/wushengyouya/project-user/internal/database"
)

type MemberRepo interface {
	SaveMember(conn database.DbConn, ctx context.Context, member *member.Member) error
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
}
