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

	sm "par/internal/handler/submissionhandler"
	rm "par/internal/repo/reposubmission"
	lsm "par/internal/service/submissionservice"

	// srd "par/internal/handler/submissiondireksi"
	// rrd "par/internal/repo/repodireksi"
	// lsd "par/internal/service/submissiondireksi"

	// sra "par/internal/handler/submissionadmin"
	// rra "par/internal/repo/repoadmin"
	// lsa "par/internal/service/submissionadmin"
	middlewares "par/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// s
func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpm := ru.NewRepoUser(db)
	ucmhsw := us.NewServiceUser(rpm)
	hndlmhs := uh.NewHandleUser(ucmhsw)
	Usergrup := e.Group("/user")
	Usergrup.POST("/adduser", hndlmhs.Register)

	rpl := rl.NewRepoLogin(db)
	servicelogin := ls.NewServiceLogin(rpl, rpm)
	handlelogin := lh.NewHandlLogin(servicelogin)
	logingrup := e.Group("/login")
	logingrup.POST("", handlelogin.Login)

	rpd := rd.NewRepoDepartments(db)
	servicedepart := ds.NewServiceDepartments(rpd)
	handledepart := dh.NewHandlesDepartment(servicedepart)
	departgrup := e.Group("/department")
	departgrup.POST("/adddepartment", handledepart.AddDepartment, middlewares.JWTMiddleware())
	departgrup.GET("", handledepart.AllDepartment, middlewares.JWTMiddleware())
	departgrup.PUT("/:id", handledepart.UpdatedDepartment, middlewares.JWTMiddleware())
	departgrup.DELETE("/:id", handledepart.DeletedDepartment, middlewares.JWTMiddleware())

	rpsm := rm.NewRepoSubmission(db)
	servicesubmission := lsm.NewServiceSubmission(rpsm, rpd, rpm)
	handlesubmmission := sm.NewHandlesSubmission(servicesubmission)
	submissiongrup := e.Group("/submission")
	submissiongrup.POST("/addsubmission", handlesubmmission.AddSubmissionManager, middlewares.JWTMiddleware())
	submissiongrup.GET("/manager", handlesubmmission.GetAllSubmissionManager, middlewares.JWTMiddleware())
	submissiongrup.GET("/direksi", handlesubmmission.GetAllSubmissionDireksi, middlewares.JWTMiddleware())
	submissiongrup.GET("/admin", handlesubmmission.GetAllSubmissionAdmin, middlewares.JWTMiddleware())
	submissiongrup.PUT("/admin/:id", handlesubmmission.UpdateSubmissionAdmin, middlewares.JWTMiddleware())
	submissiongrup.GET("/vicepresident", handlesubmmission.GetAllSubmissionPresident, middlewares.JWTMiddleware())
	submissiongrup.PUT("/vicepresident/:id", handlesubmmission.UpdateSubmissionPresident, middlewares.JWTMiddleware())
	submissiongrup.PUT("/direksi/:id", handlesubmmission.UpdateSubmissionDireksi, middlewares.JWTMiddleware())

}
