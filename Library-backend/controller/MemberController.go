package controller

import (
	"Library-backend/models"
	"Library-backend/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
	memberService services.MemberService
}

func NewMemberController(member services.MemberService) *MemberController {
	return &MemberController{memberService: member}
}

func (mc *MemberController) CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.memberService.CreateMember(context.Background(), &member); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Member created successfully"})
}

func (mc *MemberController) GetAllMember(c *gin.Context) {
	members, err := mc.memberService.GetAllMembers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

func (mc *MemberController) GetMemberByID(c *gin.Context) {
	id := c.Param("id")
	member, err := mc.memberService.GetMemberByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, member)
}

func (mc *MemberController) UpdateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := mc.memberService.UpdateMember(context.Background(), &member); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member updated successfully"})
}

func (mc *MemberController) DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if err := mc.memberService.DeleteMember(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
