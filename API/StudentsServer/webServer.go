package StudentsServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grate Grate  `json:"grate"`
}

type Grate struct {
	ID int
}

var students = []Student{
	{ID: 1, Name: "Ivan", Grate: Grate{ID: 1}},
	{ID: 2, Name: "Petro", Grate: Grate{ID: 2}},
	{ID: 3, Name: "Mykola", Grate: Grate{ID: 3}},
}

var teachers = map[string]bool{
	"Teacher1": true,
	"Teacher2": true,
	"Teacher3": true,
}

func isTeacherAuthorized(username string) bool {
	return teachers[username]
}
func GetStudentInfo(context *gin.Context) {
	username := context.GetHeader("Authorization")
	studentIDStr := context.Param("id")

	if !isTeacherAuthorized(username) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Wrong ID"})
		return
	}

	var foundStudent Student
	for _, student := range students {
		if studentID == student.ID {
			_ = student
			break
		}
	}

	if foundStudent.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Student is not found"})
		return
	}

	context.JSON(http.StatusOK, foundStudent)
}
