package web

import (
	"basic-project/internal/domain"
	"basic-project/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

const (
	EmailRegexPatter = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	PasswordPatter   = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func (h *UserHandler) RegisterRouters(server *gin.Engine) {
	server.POST("user/signup", h.SignUp)
	server.POST("user/login", h.Login)
	server.POST("user/edit", h.Edit)
	server.POST("user/profile", h.Profile)
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(EmailRegexPatter),
		passwordRexExp: regexp.MustCompile(PasswordPatter),
		svc:            svc,
	}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {

	type SignupReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignupReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	re := regexp.MustCompile(EmailRegexPatter)

	if !re.MatchString(req.Email) {
		ctx.JSON(http.StatusOK, gin.H{"message": "邮箱格式异常,匹配失败"})
		return
	}
	err := h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	switch err {
	case nil:
		ctx.JSON(http.StatusOK, gin.H{
			"message": "signup success",
		})
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

func (h *UserHandler) Login(ctx *gin.Context) {

	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	fmt.Println("u", u)
	switch err {
	case nil:
		ctx.String(http.StatusOK, "登录成功")

	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或者密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {

}
