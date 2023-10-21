package faktoryandroute

import (
	uh "par/internal/handler/userhandler"
	ru "par/internal/repo/repouser"
	us "par/internal/service/userservice"

	lh "par/internal/handler/loginhandler"
	rl "par/internal/repo/repologin"
	ls "par/internal/service/loginservice"

	dh "par/internal/handler/departmenthandler"
	rd "par/internal/repo/repodepartment"
	ds "par/internal/service/departmentservice"
	middlewares "par/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpm := ru.NewRepoUser(db)
	ucmhsw := us.NewServiceUser(rpm)
	hndlmhs := uh.NewHandleUser(ucmhsw)
	Usergrup := e.Group("/user")
	Usergrup.POST("/adduser", hndlmhs.Register)

	rpl := rl.NewRepoLogin(db)
	servicelogin := ls.NewServiceLogin(rpl)
	handlelogin := lh.NewHandlLogin(servicelogin)
	logingrup := e.Group("/login")
	logingrup.POST("", handlelogin.Login)

	rpd := rd.NewRepoDepartments(db)
	servicedepart := ds.NewServiceDepartments(rpd)
	handledepart := dh.NewHandlesDepartment(servicedepart)
	departgrup := e.Group("/department")
	departgrup.POST("/adddepartment", handledepart.AddDepartment, middlewares.JWTMiddleware())
}
