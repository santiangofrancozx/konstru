package get_project_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserProjectsByTokenService() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := Adapters.GetUserIdByToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		}
		projects, errI := Adapters.GetUserProjectsByUSerIdAdapter(user.ID)
		if errI != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": errI.Error()})
		}

		context.JSON(http.StatusOK, projects)

	}
}
