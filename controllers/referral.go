package controllers

import (
	"net/http"
	"referral-system-2/config"
	"referral-system-2/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContributionInput struct {
	Email string `json:"email" binding:"required,email"`
}

func CreateContribution(c *gin.Context) {
	var input ContributionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	referralLink := c.Param("referralLink")

	// Check if the referral link is valid and not expired
	var user models.User
	if err := config.DB.Where("referral_link = ?", referralLink).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid referral link"})
		return
	}

	// Check if the link is expired
	if time.Now().Sub(user.CreatedAt) > 7*24*time.Hour {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referral link has expired"})
		return
	}

	contribution := models.Contribution{
		Email:        input.Email,
		ReferralLink: referralLink,
	}

	if err := config.DB.Create(&contribution).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"contribution": contribution})
}

func GenerateNewLink(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	newLink := uuid.New().String()
	user.ReferralLink = newLink

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"new_referral_link": newLink})
}
