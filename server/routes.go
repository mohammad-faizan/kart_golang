package server

import(
	"github.com/gin-gonic/gin"
	"simple-server/db"
)

func addRouting(s *gin.Engine, db db.DbAdapter){
	s.GET("/", home(db))
	s.GET("/login", login(db))
	s.GET("/users", userList(db))
	s.POST("/users", createUser(db))
}