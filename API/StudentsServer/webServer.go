package StudentsServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade Grade  `json:"grade"`
}

type Grade struct {
	ID      int
	Teacher string
}

var students = []Student{
	{ID: 1, Name: "Ivan", Grade: Grade{ID: 1, Teacher: "Teacher1"}},
	{ID: 2, Name: "Petro", Grade: Grade{ID: 2, Teacher: "Teacher2"}},
	{ID: 3, Name: "Mykola", Grade: Grade{ID: 3, Teacher: "Teacher3"}},
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

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Wrong ID"})
		return
	}

	var foundStudent Student
	for _, student := range students {
		if studentID == student.ID {
			foundStudent = student
			break
		}
	}

	if foundStudent.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Student is not found"})
		return
	}

	if !isTeacherAuthorized(username) || username != foundStudent.Grade.Teacher {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	context.JSON(http.StatusOK, foundStudent)
}
