package account

import (
	"github.com/bhmy-shm/gofks"
	"github.com/gin-gonic/gin"
	"web/wire"
)

type AccountCase struct {
	*wire.ServiceWire `inject:"-"`
}

func NewAccountController() *AccountCase {
	return &AccountCase{}
}

func (s *AccountCase) Build(gofk *gofks.Gofk) {
	Account := gofk.Group("account")
	Account.Handle("POST", "/login", s.accountLogin)

	org := Account.Group("org")
	org.Handle("POST", "/addOrg", s.accountOrgAdd)

	user := Account.Group("user")
	user.Handle("POST", "/addUser", s.accountUserAdd)
}

func (s *AccountCase) Name() string {
	return "AccountCase"
}

func (s *AccountCase) Wire() *wire.ServiceWire {
	return s.ServiceWire
}

func (s *AccountCase) accountLogin(gin *gin.Context) {}

func (s *AccountCase) accountOrgAdd(gin *gin.Context) {}

func (s *AccountCase) accountUserAdd(gin *gin.Context) {}
