package controllers

import (
	"fmt"
	//"io/ioutil"
	"os"
	"time"

	"tugas-akhir-2/common"
	"tugas-akhir-2/middlewares"
	"tugas-akhir-2/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User is alias for models.User
type User = models.User

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(data common.JSON) (string, error) {

	//  token is valid for 7days
	date := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	// get path from root dir
	/*pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}*/
	tokenString, err := token.SignedString([]byte(os.Getenv("RANDOM_STRING")))
	return tokenString, err
}

//UserSignUp controller
func UserSignUp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Email     string  `json:"email" binding:"required"`
		Name      string  `json:"name" binding:"required"`
		Password  string  `json:"password" binding:"required"`
		Birthdate string  `json:"birthdate" binding:"required"`
		Gender    string  `json:"gender" binding:"required"`
		Weight    float64 `json:"weight" binding:"required"`
		Height    float64 `json:"height" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var exists User
	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.AbortWithStatus(409)
		return
	}

	hash, hashErr := hash(body.Password)
	if hashErr != nil {
		c.AbortWithStatus(500)
		return
	}

	// create user
	user := User{
		Email:          body.Email,
		Name:           body.Name,
		HashedPassword: hash,
		Birthdate:      body.Birthdate,
		Gender:         body.Gender,
		Weight:         body.Weight,
		Height:         body.Height,
	}

	db.NewRecord(user)
	db.Create(&user)

	serialized := user.Serialize()
	token, _ := generateToken(serialized)
	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, common.JSON{
		"user":  user.SSerialize(),
		"token": token,
	})
}

//UserLogin controller
func UserLogin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var user User
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404) // user not found
		return
	}

	if !checkHash(body.Password, user.HashedPassword) {
		c.AbortWithStatus(401)
		return
	}

	serialized := user.Serialize()
	token, _ := generateToken(serialized)

	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, common.JSON{
		"user":  user.SSerialize(),
		"token": token,
	})
}

//UserRetrieve controller
func UserRetrieve(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	c.JSON(200, common.JSON{
		"user": user.SSerialize(),
	})
}

// UserRenewToken check API will renew token when token life is less than 3 days, otherwise, return null for token
func UserRenewToken(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	tokenExpire := int64(c.MustGet("token_expire").(float64))
	now := time.Now().Unix()
	diff := tokenExpire - now

	//fmt.Println(diff)
	if diff < 60*60*24*3 {
		// renew token
		token, _ := generateToken(user.Serialize())
		c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)
		c.JSON(200, common.JSON{
			"token": token,
			//"user":  user.Serialize(),
		})
		return
	}

	c.JSON(200, common.JSON{
		"token": nil,
		//"user":  user.SSerialize(),
	})
}
