package projectservicev1

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/wushengyouya/project-common/encrypts"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-grpc/project"
	"github.com/wushengyouya/project-project/internal/dao"
	"github.com/wushengyouya/project-project/internal/data/menu"
	"github.com/wushengyouya/project-project/internal/database/tran"
	"github.com/wushengyouya/project-project/internal/repo"
	"github.com/wushengyouya/project-project/pkg/model"
	"go.uber.org/zap"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	MenuRepo    repo.MenuRepo
	ProjecRepo  repo.ProjectRepo
	Transaction tran.Transaction
}

// 初始化service
func New() *ProjectService {
	return &ProjectService{
		cache:       dao.RC,
		MenuRepo:    dao.NewMenuDao(),
		ProjecRepo:  dao.NewProjectDao(),
		Transaction: dao.NewTransaction(),
	}
}

func (p *ProjectService) Index(ctx context.Context, in *project.IndexMessage) (*project.IndexResponse, error) {
	c := context.Background()
	pms, err := p.MenuRepo.FindMenus(c)
	if err != nil {
		zap.L().Error("Index db FindMenus error: ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	childs := menu.CovertChild(pms)
	var mms []*project.MenuMessage
	copier.Copy(&mms, childs)
	return &project.IndexResponse{Menus: mms}, nil
}

func (p *ProjectService) FindProjectByMemId(ctx context.Context, msg *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {

	memberId := msg.MemberId
	page := msg.Page
	pageSize := msg.PageSize
	pms, total, err := p.ProjecRepo.FindProjectByMemId(context.Background(), memberId, page, pageSize)
	if err != nil {
		zap.L().Error("project FindProjectByMemId error ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if pms == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}
	var pmm []*project.ProjectMessage
	copier.Copy(&pmm, pms)

	for _, v := range pmm {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}
