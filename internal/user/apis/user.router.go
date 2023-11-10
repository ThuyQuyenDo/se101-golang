package usrapis

import (
	"go-ecommerce/db"
	usrhandler "go-ecommerce/internal/user/apis/handler"
	usrrepo "go-ecommerce/internal/user/repository"
	"go-ecommerce/provider/hasher"
	utils "go-ecommerce/utils/validator"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine, dbInstance *db.Database) {
	userFinder := usrrepo.NewUserFinderImpl(*dbInstance)
	userWriter := usrrepo.NewUserWriterImpl(*dbInstance)

	userRepo := usrrepo.NewUserRepository(userFinder, userWriter)

	hasher := hasher.NewMd5Hash()

	validatorRequest := utils.NewValidator()

	handler := usrhandler.NewUserHandler(validatorRequest, *userRepo, hasher)

	initRouter(engine, *handler)
}

func initRouter(engine *gin.Engine, handler usrhandler.UserHandler) {
	engine.POST("/user/create", handler.HandlecreateUser)
}
