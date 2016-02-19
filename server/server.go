package server

import(
	"github.com/gin-gonic/gin"
	"simple-server/db"
)


func NewServer(db db.DbAdapter) *gin.Engine {
	s := gin.Default()
	addRouting(s, db)
	return s
}