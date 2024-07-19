package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/http/controllers"
)

func SetupApiRoutes(router *gin.RouterGroup) {

	var MenuController = controllers.NewMenuController()
	menuRoutes := router.Group("/menus")
	{
		menuRoutes.GET("/menubyuser/:id", MenuController.GetMenuByUserid)
		menuRoutes.GET("/datatable", MenuController.DataTable)
		menuRoutes.GET("/edit/:id", MenuController.Edit)
		menuRoutes.POST("/update/:id", MenuController.Update)
		menuRoutes.POST("/insert", MenuController.Create)
		menuRoutes.GET("/parentnot/:id", MenuController.GetParentNotId)
	}

	var UserController = controllers.NewUserController()
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/datatable", UserController.DataTable)
		userRoutes.POST("/update/:id", UserController.Update)
		userRoutes.POST("/updateotp/:id", UserController.UpdateOTP)
	}

	var RoleController = controllers.NewRolesController()
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.POST("/datatable", RoleController.DataTable)
		roleRoutes.GET("/edit/:id", RoleController.Edit)
		roleRoutes.POST("/update/:id", RoleController.Update)
		roleRoutes.POST("/insert", RoleController.Create)
		roleRoutes.GET("/test", RoleController.Test)
	}

	var SchoolController = controllers.NewSchoolController()
	schoolRoutes := router.Group("/school")
	{
		schoolRoutes.GET("/datatable", SchoolController.DataTable)
		schoolRoutes.POST("/update/:id", SchoolController.Update)
	}
	var BranchController = controllers.NewBranchController()
	branchRoutes := router.Group("/branch")
	{
		branchRoutes.POST("/datatable", BranchController.DataTable)
		branchRoutes.GET("/dropdown", BranchController.Dropdown)
		branchRoutes.POST("/update/:id", BranchController.Update)
	}

	var ChildsController = controllers.NewChildsController()
	childRoutes := router.Group("/child")
	{
		childRoutes.POST("/datatable", ChildsController.DataTable)
		childRoutes.GET("/child/:id", ChildsController.GetChildByUserid)
	}

	var EducatorController = controllers.NewEducatorController()
	educatorRoutes := router.Group("/educator")
	{
		educatorRoutes.POST("/datatable", EducatorController.DataTable)
	}

	var MilestoneController = controllers.NewMilestoneController()
	milestoneRoutes := router.Group("/milestone")
	{
		milestoneRoutes.POST("/report", MilestoneController.Edit)
	}

	var LoginController = controllers.NewLoginController()
	loginROute := router.Group("/login")
	{
		loginROute.GET("/dologin", LoginController.DoLogin)
	}

	var TimesheetController = controllers.NewTimesheetController()
	timesheetRoutes := router.Group("/timesheet")
	{
		timesheetRoutes.POST("/report", TimesheetController.Edit)
	}

	var TimesheetEducatorController = controllers.NewTimesheetEducatorController()
	timesheetEducatorRoutes := router.Group("/timesheeteducator")
	{
		timesheetEducatorRoutes.POST("/report", TimesheetEducatorController.Edit)
	}
	var ParentsController = controllers.NewParentsController()
	parentRoutes := router.Group("/parent")
	{
		parentRoutes.POST("/datatable", ParentsController.DataTable)

	}
	var MediaController = controllers.NewMediaController()
	mediaRoutes := router.Group("/media")
	{
		mediaRoutes.GET("/:id", MediaController.Edit)
	}

	var InvoiceController = controllers.NewInvoiceController()
	invoiceController := router.Group("/invoice")
	{
		invoiceController.POST("/generate", InvoiceController.Generate)
	}

	var EducatorBranchController = controllers.NewEducatorBranchController()
	educatorbranchController := router.Group("/educatorbranch")
	{
		educatorbranchController.GET("/branch/:id", EducatorBranchController.GetBranch)
	}
}
