package controller

import (
	"Library-backend/models"
	"Library-backend/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.accountService.CreateAccount(context.Background(), &account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully!"})
}

func (ac *AccountController) GetAllAccount(c *gin.Context) {
	accounts, err := ac.accountService.GetAllAccounts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (ac *AccountController) GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	account, err := ac.accountService.GetAccountByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) UpdateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.accountService.UpdateAccount(context.Background(), &account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully!"})
}

func (ac *AccountController) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := ac.accountService.DeleteAccount(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully!"})
}
