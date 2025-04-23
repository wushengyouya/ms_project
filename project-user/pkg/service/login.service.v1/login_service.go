package login_service_v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/encrypts"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-common/jwts"
	"github.com/wushengyouya/project-grpc/user/login"
	"github.com/wushengyouya/project-user/config"
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
	login.UnimplementedLoginServiceServer
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

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {
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
	return &login.CaptchaResponse{Code: int32(2000)}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	// 1.校验验证码
	redisValue, err := ls.Cache.Get("REGISTER_" + msg.Mobile)
	if err != nil {
		zap.L().Error("Register redis search faild", zap.Any("REGISTER_"+msg.Mobile, redisValue), zap.Error(err))
		return new(login.RegisterResponse), errs.GrpcError(model.RedisError)
	}
	if redisValue != msg.Captcha {
		return new(login.RegisterResponse), errs.GrpcError(model.CaptchaNotExist)
	}
	// 2.校验是否已经注册
	c := context.Background()
	exist, err := ls.MemberRepo.GetMemberByAccount(c, msg.Name)
	if err != nil {
		zap.L().Error("Register GetMemberByAccount db fail", zap.Error(err))
		return new(login.RegisterResponse), errs.GrpcError(model.AccountExist)
	}
	if exist {
		return new(login.RegisterResponse), errs.GrpcError(model.AccountExist)
	}
	// 3.接收数据插入数据库
	password := encrypts.Sha256(msg.Password)
	mem := &member.Member{
		Account:       msg.Name,
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
	return new(login.RegisterResponse), err
}

func (ls *LoginService) Login(ctx context.Context, in *login.LoginMessage) (*login.LoginResponse, error) {
	// 接收到登录数据
	c := context.Background()
	mem, err := ls.MemberRepo.FindMember(c, in.Account, in.Password)
	if err != nil {
		zap.L().Error("login db FindMember error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.AccountAndPwdError)
	}
	// 查询用户下的组织
	memMsg := new(login.MemberMessage)
	err = copier.Copy(&memMsg, mem)
	if err != nil {
		zap.L().Error("login Copy mem error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyErr)
	}
	orgs, err := ls.OrganizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("login db FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	// 对org id进行加密
	for _, v := range orgsMessage {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
	}
	if err != nil {
		zap.L().Error("login Copy orgs error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyErr)
	}
	// 生成token, refreshToken
	memIdStr := strconv.FormatInt(mem.Id, 10)
	// 加密memId
	memMsg.Code, _ = encrypts.Encrypt(memIdStr, model.AESKey)
	exp := time.Duration(config.AppConf.JwtConfig.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.AppConf.JwtConfig.RefreshExp*3600*24) * time.Second

	token := jwts.CreateToken(memIdStr, config.AppConf.JwtConfig.AccessSecret, config.AppConf.JwtConfig.RefreshSecret, exp, rExp)

	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessTokenExp: token.AccessExp,
		TokenType:      "bearer",
	}

	return &login.LoginResponse{
		Member:           memMsg,
		OrganizationList: orgsMessage,
		TokenList:        tokenList,
	}, nil
}

func (ls *LoginService) TokenVerify(ctx context.Context, in *login.LoginMessage) (*login.LoginResponse, error) {
	token := in.Token
	if strings.Contains(token, "bearer") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}

	// 解析token信息
	ParseToken, err := jwts.ParseToken(token, config.AppConf.JwtConfig.AccessSecret)
	if err != nil {
		zap.L().Error("Login TokenVerify error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	// 数据库查询 登录之后 把用户信息存储起来
	id, _ := strconv.ParseInt(ParseToken, 10, 64)
	mem, err := ls.MemberRepo.FindMemberById(context.Background(), id)
	if err != nil {
		zap.L().Error("TokenVerify db FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memberMsg := &login.MemberMessage{}
	err = copier.Copy(&memberMsg, mem)
	if err != nil {
		zap.L().Error("MyOrgList copy err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyErr)
	}
	memberMsg.Code, _ = encrypts.EncryptInt64(memberMsg.Id, model.AESKey)

	// 返回结果
	return &login.LoginResponse{Member: memberMsg}, nil
}
func (ls *LoginService) MyOrgList(ctx context.Context, in *login.UserMessage) (*login.OrgListResponse, error) {
	memId := in.MemId
	orgs, err := ls.OrganizationRepo.FindOrganizationByMemId(ctx, memId)
	if err != nil {
		zap.L().Error("MyOrgList FindOrganizationByMemId error ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	if err != nil {
		zap.L().Error("MyOrgList copy err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyErr)
	}
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgsMessage}, nil
}
