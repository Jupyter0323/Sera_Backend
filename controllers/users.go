package controllers

import (
    "example/task/database"
    // "fmt"
	"net/http"
	"example/task/model"
	"github.com/gin-gonic/gin"
)

func SignUpUser(c *gin.Context) {
   var input model.User

   if err := c.Bind(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/SignUpUser",
            "status": "Server Error !",
            "msg":    "Server Error !",
            "data":   nil,
        })
        c.Abort()
        return
    } 
 
    user, err := model.FindUserByUsername(input.Email)
    if len(user.Email) != 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 409,
            "api_version": "v1",
            "endpoint": "/SignUpUser",
            "status": "SignUp Failure!",
            "msg":    "Conflict Email",
         })
        return
    }

    wallet, err := model.IsValidWallet(input.Wallet_address)

    if len(wallet.Wallet_address) != 0  {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 409,
            "api_version": "v1",
            "endpoint": "/SignUpUser",
            "status": "SignUp Failure!",
            "msg":    "Conflict Wallet_address",
        })
        return
    }

    savedUser, err := input.Save()

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/SignUpUser",
            "status": "SignUp Failure!",
            "msg":    "Internal Server Error!",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "status_code": 201,
        "api_version": "v1",
        "endpoint": "/SignUpUser",
        "status": "SignUp Success!",
        "msg":    "Welcome to SignUp",
        "data": savedUser, 
    })
}

func SignInUser(c *gin.Context) {
    var input model.User

    if err := c.Bind(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/SignInUser",
            "status": "Server Error !",
            "msg":    "Server Error !",
            "data":   nil,
        })
        c.Abort()
        return
    } 

    user, err := model.FindUserByUsername(input.Email)

    if len(user.Email) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 409,
            "api_version": "v1",
            "endpoint": "/SignInUser",
            "status": "Wrong Email!",
            "msg":    "There is a no exist person!",
            "data": err,
        })
        return
    }

    password, err := model.IsValidPassword(input.Password)
    if len(password.Email) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 409,
            "api_version": "v1",
            "endpoint": "/SignInUser",
            "status": "Wrong Password!",
            "msg":    "There is a no exist person!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status_code": 200,
        "api_version": "v1",
        "endpoint": "/SignInUser",
        "status": "Login Success!",
        "msg":    "Welcome to Login",
        "data": user,
    })
}

func GetListUser(c *gin.Context) {
    var input []model.User

    if err := database.Database.Find(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/SignInUser",
            "status": "Server Error !",
            "msg":    "Server Error !",
            "data":   err.Error(),
        })
        return
    }

    if len(input) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "status_code": 204,
            "api_version": "v1",
            "endpoint": "/GetListUser",
            "status": "No Content!",
        })
        return
    }

    c.JSON(http.StatusOK,  gin.H{
        "status_code": 200,
        "api_version": "v1",
        "endpoint": "/GetListUser",
        "data": input,
    })
}


