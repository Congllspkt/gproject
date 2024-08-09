package initialize

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

type User struct {
	Username string
	Password string
}

var users = []User{}

func TryJWT() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/user", authMiddleware(), getUser)

	r.Run(":8080")
}

func register(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			claims := &jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Minute * 1).Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": c.Keys["user"]})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil && !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)

		username := claims["username"].(string)
		for _, u := range users {
			if u.Username == username {
				c.Keys = map[string]any{}
				c.Keys["user"] = u
				break
			}
		}

		c.Next()
	}
}

/*

<!DOCTYPE html>
<html>
<head>
    <title>AJAX Test</title>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
</head>
<body>
    <h1>AJAX Test</h1>

    <button id="registerBtn">Register User</button>
    <button id="loginBtn">Login</button>
    <button id="getUserBtn">Get User Data</button>

    <script>
        const baseUrl = 'http://localhost:8080';

        $('#registerBtn').click(function() {
            $.ajax({
                type: 'POST',
                url: baseUrl + '/register',
                contentType: 'application/json',
                data: JSON.stringify({ username: 'new_user', password: 'new_password' }),
                success: function(response) {
                    console.log('User registered successfully:', response);
                }
            });
        });

        $('#loginBtn').click(function() {
            $.ajax({
                type: 'POST',
                url: baseUrl + '/login',
                contentType: 'application/json',
                data: JSON.stringify({ username: 'new_user', password: 'new_password' }),
                success: function(response) {
                    console.log('Login successful. Token:', response.token);
                }
            });
        });

        $('#getUserBtn').click(function() {
            const token =
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxNzY5MDQsInVzZXJuYW1lIjoibmV3X3VzZXIifQ.eQGZKBZrixCG_KF4eLYQKSVBBkM7pmpPdRmPy0WRnyk"
			; // Replace with your actual JWT token
            $.ajax({
                type: 'GET',
                url: baseUrl + '/user',
                headers: { 'Authorization': 'Bearer ' + token },
                success: function(response) {
                    console.log('User data:', response.user);
                }
            });
        });
    </script>
</body>
</html>

*/
