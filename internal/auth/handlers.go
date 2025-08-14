package auth

import (
	"net/http"
	"open_gamba/internal/user"
	"open_gamba/internal/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body LoginRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	foundUser, err := user.GetWithUsername(body.Username)
	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	isValid, err := util.IsPasswordMatch(foundUser.Password, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Error validating password",
		})
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Password don't match",
		})
		return
	}

	generetateTokenAndSaveAsCookie(c, foundUser)

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User logged in",
	})
}

func SignUp(c *gin.Context) {
	var body SignUpRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Check if username already exists
	isPresent, err := user.IsUserPresent(body.Username, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error checking username availability",
		})
		return
	}

	if isPresent {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Username or email already taken",
		})
		return
	}

	// Hash password
	hash, err := util.EncryptPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to process password",
		})
		return
	}

	// Create user object
	newUser := user.User{
		Email:       body.Email,
		Password:    string(hash),
		Username:    body.Username,
		Birthday:    &body.Birthday,
		AccessRoles: []user.AccessRole{user.UserRole},
	}

	// Save user to database
	savedUser, err := user.SaveUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create user",
		})
		return
	}
	generetateTokenAndSaveAsCookie(c, savedUser)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}

func generetateTokenAndSaveAsCookie(c *gin.Context, user *user.User) {
	// Generate authentication token
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "User created but failed to generate token. Please login manually.",
		})
		return
	}

	// Set secure cookie
	c.SetCookie(
		"Authorization",                      // name
		token,                                // value
		int((30 * 24 * time.Hour).Seconds()), // maxAge in seconds
		"/",                                  // path
		"",                                   // domain (empty for current domain)
		c.Request.TLS != nil,                 // secure (true if HTTPS)
		true,                                 // httpOnly
	)
}
