package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/brf153/jwt-golang.git/database"
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

func SignUp()

func Login()

func GetUsers()

func GetUser()