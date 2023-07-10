package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/brf153/jwt-golang.git/database"
	"github.com/brf153/jwt-golang.git/models"
	helper "github.com/brf153/jwt-golang/helpers"
	"github.com/brf153/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.com/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func SignUp()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, )
		}
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
		if err:= helper.MatchUserTypeToUid(c , userId); err != nil{
			c.JSON(http.StatusBadRequest, gin.H("error":err.Error()))
			return 
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.M{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}