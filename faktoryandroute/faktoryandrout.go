package faktoryandroute

import (
	"par/config"
	uh "par/internal/handler/userhandler"
	ru "par/internal/repo/repouser"
	us "par/internal/service/userservice"
	"par/uploadgambar"

	lh "par/internal/handler/loginhandler"
	rl "par/internal/repo/repologin"
	ls "par/internal/service/loginservice"

	dh "par/internal/handler/departmenthandler"
	rd "par/internal/repo/repodepartment"
	ds "par/internal/service/departmentservice"

	sm "par/internal/handler/submissionhandler"
	rm "par/internal/repo/reposubmission"
	lsm "par/internal/service/submissionservice"

	kh "par/internal/handler/kandidathandler"
	rk "par/internal/repo/repokandidat"
	ks "par/internal/service/kandidatservice"

	shl "par/internal/handler/soalhandler"
	rsl "par/internal/repo/reposoal"
	ssl "par/internal/service/soalservice"

	ih "par/internal/handler/interviewhandler"
	ri "par/internal/repo/repointerview"
	is "par/internal/service/interviewservice"

	ph "par/internal/handler/proseshandler"
	rp "par/internal/repo/repoproses"
	ps "par/internal/service/prosesservice"

	pth "par/internal/handler/posisihandler"
	rpt "par/internal/repo/repoposisi"
	pts "par/internal/service/posisiservice"

	hf "par/internal/handler/fpthandler"
	rf "par/internal/repo/repofpt"
	sf "par/internal/service/fptservice"

	middlewares "par/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// s
func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	cfg := config.GetConfig()
	cld := uploadgambar.NewCloud(cfg)
	rpm := ru.NewRepoUser(db)
	ucmhsw := us.NewServiceUser(rpm)
	hndlmhs := uh.NewHandleUser(ucmhsw)
	Usergrup := e.Group("/user")
	Usergrup.POST("/adduser", hndlmhs.Register)
	Usergrup.GET("/allmanager/:roles", hndlmhs.NamaManager, middlewares.JWTMiddleware())

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
	submissiongrup.POST("/addsubmission", handlesubmmission.AddSubmission, middlewares.JWTMiddleware())
	submissiongrup.GET("/user", handlesubmmission.GetAllSubmissionUser, middlewares.JWTMiddleware())
	submissiongrup.GET("/admin", handlesubmmission.GetAllSubmissionAdmin, middlewares.JWTMiddleware())
	submissiongrup.PUT("/admin/:id", handlesubmmission.UpdateSubmissionAdmin, middlewares.JWTMiddleware())
	submissiongrup.PUT("/vicepresident/:id", handlesubmmission.UpdateSubmissionPresident, middlewares.JWTMiddleware())
	submissiongrup.PUT("/direksi/:id", handlesubmmission.UpdateSubmissionDireksi, middlewares.JWTMiddleware())
	submissiongrup.GET("/manager/:nama", handlesubmmission.GetNamaManager, middlewares.JWTMiddleware())
	submissiongrup.GET("", handlesubmmission.GetCode, middlewares.JWTMiddleware())

	rpk := rk.NewRepoKandidat(db)
	servicekandidat := ks.NewServiceKandidat(rpk, rpsm, rpd, rpm)
	handlekandiat := kh.NewHandlesKandidat(servicekandidat, cld)
	kandidatgrup := e.Group("/kandidat")
	kandidatgrup.POST("/addformulir", handlekandiat.AddFormulirKandidat, middlewares.JWTMiddleware())
	kandidatgrup.GET("", handlekandiat.GetCodeKandidat, middlewares.JWTMiddleware())

	rps := rsl.NewReposoal(db)
	servicesoal := ssl.NewServiceSoal(rps)
	handlesoal := shl.NewHandlesSoal(servicesoal)
	soalgrup := e.Group("/soal")
	soalgrup.POST("/addsoal", handlesoal.Addsoal, middlewares.JWTMiddleware())
	soalgrup.GET("/datasoal", handlesoal.AllSoal, middlewares.JWTMiddleware())
	soalgrup.GET("", handlesoal.KategoriSoal, middlewares.JWTMiddleware())
	soalgrup.PUT("", handlesoal.UpdatedSoal, middlewares.JWTMiddleware())
	soalgrup.DELETE("", handlesoal.Deletedsoal, middlewares.JWTMiddleware())

	rpi := ri.NewRepoInterview(db)
	rpp := rp.NewRepoproses(db)
	rfp := rf.NewRepofpt(db)
	serviceinteview := is.NewServiceinterview(rpi, rpk, rpd, rpm, rps, rpsm, rpp, rfp)
	handleinterview := ih.NewHandlesInterview(serviceinteview)
	interviewgrup := e.Group("/interview")
	interviewgrup.POST("/addinterview", handleinterview.AddFormulirInterview, middlewares.JWTMiddleware())
	interviewgrup.GET("", handleinterview.GetallInterview, middlewares.JWTMiddleware())
	interviewgrup.POST("/addfpt", handleinterview.AddFormulirInterviewFpt, middlewares.JWTMiddleware())

	// rpp := rp.NewRepoproses(db)
	serviceproses := ps.NewServiceprocess(rpp, rpi, rpk, rpm, rps, rpsm, rfp)
	handleproses := ph.NewHandlesProcess(serviceproses)
	prosesgrup := e.Group("/proses")
	prosesgrup.POST("/addproses", handleproses.AddProcess, middlewares.JWTMiddleware())
	prosesgrup.GET("", handleproses.GetallDetail, middlewares.JWTMiddleware())
	prosesgrup.GET("/manager", handleproses.GetallDetailManager, middlewares.JWTMiddleware())
	prosesgrup.PUT("", handleproses.Updatedetail, middlewares.JWTMiddleware())
	prosesgrup.PUT("/admin", handleproses.UpdatedDetailAdmin, middlewares.JWTMiddleware())
	prosesgrup.PUT("/direksi", handleproses.UpdatedDetaildireksi, middlewares.JWTMiddleware())

	rpt := rpt.NewRepoposisi(db)
	serviceposisi := pts.NewServiceposisi(rpt, rpm)
	handleposis := pth.NewHandlesPosisi(serviceposisi)
	posisigrup := e.Group("/posisi")
	posisigrup.POST("/add", handleposis.AddPosisi, middlewares.JWTMiddleware())

	servicefpt := sf.NewServicefpt(rfp)
	handlefpt := hf.NewHandlesFpt(servicefpt)
	fptgrup := e.Group("/fpt")
	fptgrup.POST("/addsoal", handlefpt.Addsoal, middlewares.JWTMiddleware())
	fptgrup.GET("", handlefpt.AllSoal, middlewares.JWTMiddleware())
	fptgrup.GET("", handlefpt.KategoriSoal, middlewares.JWTMiddleware())
	fptgrup.PUT("", handlefpt.UpdatedSoal, middlewares.JWTMiddleware())
	fptgrup.DELETE("", handlefpt.Deletedsoal, middlewares.JWTMiddleware())
}
