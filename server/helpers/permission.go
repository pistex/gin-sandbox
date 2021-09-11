package helpers

import (
	"kwanjai/libraries"
	"kwanjai/models"

	"github.com/gin-gonic/gin"
)

// GetUsername fucntion returns username (string).
func GetUsername(ginContext *gin.Context) string {
	panic("implement me")
}

// ExceedProjectLimit
func ExceedProjectLimit(ginContext *gin.Context) bool {
	panic("implement me")
}

// ExceedBoardLimit
func ExceedBoardLimit(ginContext *gin.Context, currentBoard int) bool {
	panic("implement me")
}

// IsSuperUser fucntion returns superuser status (bool).
func IsSuperUser(ginContext *gin.Context) bool {
	panic("implement me")
}

// IsProjectMember return membership status (bool) and error.
func IsProjectMember(username string, projectID string) bool {
	if projectID == "" {
		return false
	}
	project := new(models.Project)
	getProject, _ := libraries.FirestoreFind("projects", projectID) // projectID != "" ensures no error.
	getProject.DataTo(project)
	_, found := libraries.Find(project.Members, username)
	return found
}

// IsOwner return ownership status (bool) and error.
func IsOwner(username string, objectType string, objectID string) bool {
	if objectID == "" || objectType == "" {
		return false
	}
	getObject, _ := libraries.FirestoreFind(objectType, objectID) // objectID != "", objectType != "" ensures no error.
	if !getObject.Exists() {
		return false
	}
	return getObject.Data()["User"].(string) == username
}
