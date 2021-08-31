package routers

import (
	"github.com/Dasha-Kinsely/leaveswears/controllers"
	"github.com/gin-gonic/gin"
)

func AdminSignin(c *gin.Context) {
	controllers.AdminSigninControllers(c)
}

func AdminViewProducts(c *gin.Context){
	controllers.AdminViewProductsControllers(c)
}

func AdminViewBundles(c *gin.Context){
	controllers.AdminViewBundlesControllers(c)
}

func AdminViewMisc(c *gin.Context){
	controllers.AdminViewMiscControllers(c)
}

func AdminViewStatistics(c *gin.Context){
	controllers.AdminViewStatisticsControllers(c)
}

func AdminViewUsers(c *gin.Context){
	controllers.AdminViewUsersControllers(c)
}

func AdminAddProduct(c *gin.Context){
	controllers.AdminAddProductControllers(c)
}
	
func AdminAddBundle(c *gin.Context){
	controllers.AdminAddBundleControllers(c)
}

func AdminAddMisc(c *gin.Context){
	controllers.AdminAddMiscControllers(c)
}

func AdminDeleteUser(c *gin.Context){
	controllers.AdminDeleteUserControllers(c)
}

func AdminDeleteBundle(c *gin.Context){
	controllers.AdminDeleteBundleControllers(c)
}

func AdminDeleteMisc(c *gin.Context){
	controllers.AdminDeleteMiscControllers(c)
}

func AdminDeleteProduct(c *gin.Context){
	controllers.AdminDeleteProductControllers(c)
}

func AdminUpdateProduct(c *gin.Context){
	controllers.AdminUpdateProductControllers(c)
}

func AdminUpdateBundle(c *gin.Context){
	controllers.AdminUpdateBundleControllers(c)
}

func AdminRoutes(router *gin.RouterGroup) {
	router.POST("/signin", AdminSignin)
	// Admin Required routes
	// Dashboard view functions are grouped into one controller file for simplicity since they are asscessed by form data
	router.GET("/view/product", AdminViewProducts)
	router.GET("/view/bundle", AdminViewBundles)
	router.GET("/view/misc", AdminViewMisc)
	router.GET("/view/statistic", AdminViewStatistics)
	router.GET("/view/user", AdminViewUsers)
	// Adding and editing will require form data
	router.POST("/add/product", AdminAddProduct)
	router.POST("/add/bundle", AdminAddBundle)
	router.POST("/add/misc/:field", AdminAddMisc)
	router.PUT("/update/product/:id", AdminUpdateProduct)
	router.PUT("/update/bundle/:id", AdminAddBundle)
	// Dashboard delete functions are grouped into one controller file for simplicity since they are accessed by form data
	router.DELETE("/delete/product/:id", AdminDeleteProduct)
	router.DELETE("/delete/bundle/:id", AdminDeleteBundle)
	router.DELETE("/delete/misc/:field/:id", AdminDeleteMisc)
	router.DELETE("/delete/user/:id", AdminDeleteUser)
}