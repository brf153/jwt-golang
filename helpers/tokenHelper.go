package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brf153/jwt-golang.git/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims
}


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastname string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:       email,
		First_name:  firstname,
		Last_name:   lastname,
		Uid:         uid,
		User_type:   userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 168).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	return token, refreshToken, nil
}

	
	func ValidateToken(signedToken string) (claims *SignedDetails, msg string){
		token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token)(interface{}, error){
			return []byte(SECRET_KEY), nil
		})

		if err!= nil{
			msg = err.Error()
			return 
		}

		claims, ok := token.Claims.(*SignedDetails)
		if !ok {
			msg = fmt.Sprintf("the token is invalid")
			msg = err.Error()
			return 
		}

		if claims.ExpiresAt < time.Now().Local().Unix(){
			msg = fmt.Sprintf("token is expired")
			msg = err.Error()
			return 
		}
		return claims, msg 
	}

	func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
	
		var updateObj primitive.D
	
		updateObj = append(updateObj, primitive.E{Key: "token", Value: signedToken})
		updateObj = append(updateObj, primitive.E{Key: "refresh_token", Value: signedRefreshToken})
		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, primitive.E{Key: "updated_at", Value: updatedAt})
	
		upsert := true
		filter := bson.M{"user_id": userId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
	
		_, err := userCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
	
		if err != nil {
			log.Panic(err)
			return
		}
	}
	
