package server

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	. "simple-server/db"
	d "simple-server/data"
	"fmt"
)

type jsonResponse map[string]interface{}

func home(db DbAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		response := jsonResponse{}
		response["status"] = 200
		response["time"] = time.Now()
		c.JSON(http.StatusOK, response)
	}
}

func login(db DbAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		response := jsonResponse{}
		email := c.Query("email")

		if email == "" {
			response["errors"] = []string{"Provide user email"}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		user, err := db.SearchByEmail(email)

		if err != nil {
			response["errors"] = []string{fmt.Sprint(err)}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response["user"] = user
		c.JSON(http.StatusOK, response)		
	}
}

func userList(db DbAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		response := jsonResponse{}
		users, err := db.UserList()

		if err != nil {
			response["errors"] = []string{fmt.Sprint(err)}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response["users"] = users
		c.JSON(http.StatusOK, response)		
	}
}

func createUser(db DbAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {

		response := jsonResponse{}

		user, validationErrors := createUserFromRequest(c)

		if len(validationErrors) > 0 {
			response["errors"] = validationErrors
			c.JSON(http.StatusBadRequest, response)
			return
		}

		err := db.CreateUser(user)

		response["errors"] = []string{fmt.Sprint(err)}
		c.JSON(http.StatusOK, response)		
	}
}

func createUserFromRequest(c *gin.Context) (d.User, []string){
	c.Request.ParseForm()
	form := c.Request.PostForm

	var user d.User

	user.FirstName = form.Get("first_name")
	user.LastName = form.Get("last_name")
	user.Email = form.Get("email")
	user.Gender = form.Get("gender")

	password := form.Get("password")
	passwordConfirmation := form.Get("password_confirmation")

	user.SetPassword(password)
	user.SetPasswordConfirmation(passwordConfirmation)

	errors := validateUser(user, password, passwordConfirmation)

	return user, errors
}

func validateUser(u d.User, pass, passConf string) []string {
	var errors []string

	if u.FirstName == "" {
		errors = append(errors, "First Name can't left blank")
	}

	if u.LastName == "" {
		errors = append(errors, "Last Name can't left blank")
	}

	if u.Email == "" {
		errors = append(errors, "Email can't left blank")
	}

	if u.Gender == "" {
		errors = append(errors, "Gender can't left blank")
	}

	if pass == "" {
		errors = append(errors, "Password can't left blank")
	} else if len(pass) < 8 {
		errors = append(errors, "Password must be 8 character long")
	}

	if passConf == "" {
		errors = append(errors, "Password Confirmation can't left blank")
	}

	if pass != "" && passConf != "" && pass != passConf {
		errors = append(errors, "Password and Password Confirmation did not match")
	}

	return errors
}