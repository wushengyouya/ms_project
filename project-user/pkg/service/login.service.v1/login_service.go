package login_service_v1

import (
	"context"
	"fmt"
	"time"

	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/encrypts"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-user/internal/dao"
	"github.com/wushengyouya/project-user/internal/data/member"
	"github.com/wushengyouya/project-user/internal/data/organization"
	"github.com/wushengyouya/project-user/internal/database"
	"github.com/wushengyouya/project-user/internal/database/tran"
	"github.com/wushengyouya/project-user/internal/repo"
	"github.com/wushengyouya/project-user/pkg/model"
	"go.uber.org/zap"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	Cache            repo.Cache
	MemberRepo       repo.MemberRepo
	OrganizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		Cache:            dao.RC,
		MemberRepo:       dao.NewMemberDao(),
		OrganizationRepo: dao.NewOrganizationDao(),
		transaction:      dao.NewTransaction(),
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1.获取参数
	mobile := msg.Mobile
	// 2.校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 3.生成验证码（调用第三平台，放入go协程中执行）
	code := "123456"
	go func() {
		// 模拟
		err := ls.Cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Info("验证码存入redis发生错误,cause by : ", zap.Error(err))
		}
		zap.L().Info("短信平台调用成功，发送短信")
		zap.L().Debug("短信平台调用成功，发送短信 debug")
		zap.L().Warn("短信平台调用成功，发送短信warn")
		zap.L().Error("短信平台调用成功，发送短信error")
		// 4.存储验证码到redis中过期时间15分钟
		zap.L().Info(fmt.Sprintf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, code))

	}()
	return &CaptchaResponse{Code: int32(2000)}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *RegisterMessage) (*RegisterResponse, error) {
	// 1.校验验证码
	redisValue, err := ls.Cache.Get("REGISTER_" + msg.Mobile)
	if err != nil {
		zap.L().Error("Register redis search faild", zap.Error(err))
		return new(RegisterResponse), errs.GrpcError(model.RedisError)
	}
	if redisValue != msg.Captcha {
		return new(RegisterResponse), errs.GrpcError(model.CaptchaNotExist)
	}
	// 2.校验是否已经注册
	c := context.Background()
	exist, err := ls.MemberRepo.GetMemberByAccount(c, msg.Name)
	if err != nil {
		zap.L().Error("Register GetMemberByAccount db fail", zap.Error(err))
		return new(RegisterResponse), errs.GrpcError(model.AccountExist)
	}
	if exist {
		return new(RegisterResponse), errs.GrpcError(model.AccountExist)
	}
	// 3.接收数据插入数据库
	password := encrypts.Sha256(msg.Password)
	mem := &member.Member{
		Mobile:        msg.Mobile,
		Name:          msg.Name,
		Password:      password,
		Email:         msg.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        1,
	}

	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.MemberRepo.SaveMember(conn, c, mem)
		if err != nil {
			zap.L().Error("register save member db err", zap.Error(err))
			return errs.GrpcError(model.RegisterErr)
		}
		// 4.插入组织
		org := &organization.Organization{
			Name:       mem.Name + "个人项目",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   1,
			Avatar:     "https://www.baidu.com/img/flexible/logo/pc/index.png",
		}
		err = ls.OrganizationRepo.SaveOrganization(conn, c, org)
		if err != nil {
			zap.L().Error("register save organization db err", zap.Error(err))
			return errs.GrpcError(model.RegisterErr)
		}
		return nil
	})
	return new(RegisterResponse), err
}
