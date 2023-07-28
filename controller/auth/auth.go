package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	m "smapurv1_api/models"
	s "smapurv1_api/setup"
	u "smapurv1_api/utils"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/mileusna/useragent"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (au *AuthController) AuthenticateByUsername(c *gin.Context) {
	var payload *m.UsernameLoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user m.Users

	result := au.DB.First(&user, "nama_user = ?", strings.ToLower(payload.Username))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid username"})
		return
	}

	if err := u.CheckPasswordHash(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid password"})
		return
	}

	config, _ := s.LoadConfig(".")

	now := time.Now()

	expiresAt := now.Add(config.Expiration * time.Minute)

	newSession := m.Session{
		KD_User:    user.ID_User.String(),
		Expired:    expiresAt,
		IP:         u.FetchUserIP(c),
		Info:       "new session",
		Status:     "1",
		Created_At: now,
		Creator:    "System",
		Updated_At: now,
		Updater:    "System",
	}

	resultSession := au.DB.Create(&newSession)

	if resultSession.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": resultSession.Error.Error()})
		return
	}

	session_token := m.GetSessionAfterLoginResponse{
		Session_ID: newSession.Session_ID,
		KD_User:    newSession.KD_User,
		Expired:    newSession.Expired,
	}

	session := sessions.Default(c)
	session.Set("session_tokens", session_token)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.Set("sessions", session)

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawUa := c.Request.Header.Get("User-Agent")
	ua := useragent.Parse(rawUa)

	newLoginHistory := m.LoginHistory{
		KD_User:    user.ID_User.String(),
		Login_Time: now,
		IP:         u.FetchUserIP(c),
		Browser:    ua.Name + " v. " + ua.Version,
		Host_Name:  host,
		Info:       "new login autivity by username",
		Status:     "1",
		Created_At: now,
		Creator:    "System",
		Updated_At: now,
		Updater:    "System",
	}

	historyResult := au.DB.Create(&newLoginHistory)
	if historyResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": historyResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": session_token})
}

func (au *AuthController) AuthenticateByNIK(c *gin.Context) {
	var payload *m.NIKLoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user m.Users

	result := au.DB.First(&user, "nik = ?", strings.ToLower(payload.NIK))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid NIK"})
		return
	}

	if err := u.CheckPasswordHash(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid password"})
		return
	}

	config, _ := s.LoadConfig(".")

	now := time.Now()

	expiresAt := now.Add(config.Expiration * time.Minute)

	newSession := m.Session{
		KD_User:    user.ID_User.String(),
		Expired:    expiresAt,
		IP:         u.FetchUserIP(c),
		Info:       "new session",
		Status:     "1",
		Created_At: now,
		Creator:    "System",
		Updated_At: now,
		Updater:    "System",
	}

	resultSession := au.DB.Create(&newSession)

	if resultSession.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": resultSession.Error.Error()})
		return
	}

	session_token := m.GetSessionAfterLoginResponse{
		Session_ID: newSession.Session_ID,
		KD_User:    newSession.KD_User,
		Expired:    newSession.Expired,
	}

	session := sessions.Default(c)
	session.Set("session_tokens", session_token)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.Set("sessions", session)

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawUa := c.Request.Header.Get("User-Agent")
	ua := useragent.Parse(rawUa)

	newLoginHistory := m.LoginHistory{
		KD_User:    user.ID_User.String(),
		Login_Time: now,
		IP:         u.FetchUserIP(c),
		Browser:    ua.Name + " v. " + ua.Version,
		Host_Name:  host,
		Info:       "new login autivity by NIK",
		Status:     "1",
		Created_At: now,
		Creator:    "System",
		Updated_At: now,
		Updater:    "System",
	}

	historyResult := au.DB.Create(&newLoginHistory)
	if historyResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": historyResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": session_token})
}

func (au *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	getSession := session.Get("session_tokens")

	id := reflect.ValueOf(getSession).Elem().Field(0).Interface()

	var updateSessionLogout m.Session
	result := au.DB.First(&updateSessionLogout, "session_id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Session ID was not found"})
		return
	}

	sessionLogout := m.Session{
		Info:       "inautive session",
		Status:     "0",
		Updated_At: time.Now(),
		Updater:    "System",
	}

	if err := au.DB.Model(&updateSessionLogout).Updates(sessionLogout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
		return
	}

	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	userId := reflect.ValueOf(getSession).Elem().Field(1).Interface()

	var updateUserLogout m.Users
	res := au.DB.First(&updateUserLogout, "kd_user = ?", userId)

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User ID was not found"})
		return
	}

	userLogout := m.Users{
		Last_Logout: time.Now(),
	}

	if err := au.DB.Model(&updateUserLogout).Updates(userLogout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
		return
	}

	c.Set("currentUser", "")

	c.JSON(http.StatusOK, gin.H{"message": "Logged Out"})
}

func (au *AuthController) VerifyEmail(c *gin.Context) {
	getCode := c.Params.ByName("verificationCode")
	verificationCode := u.EncodeBase64(getCode)

	var updatedUser m.Users
	result := au.DB.First(&updatedUser, "verification_code = ?", verificationCode)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exist"})
		return
	}

	if updatedUser.EmailVerified {
		c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified!"})
		return
	}

	updatedUser.VerificationCode = ""
	updatedUser.EmailVerified = true
	au.DB.Save(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email verified successfully!"})
}

func (au *AuthController) ForgotPassword(ctx *gin.Context) {
	var userCredential *m.ForgotPasswordRequest

	if err := ctx.ShouldBindJSON(&userCredential); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	message := "You will receive a reset email if user with that email exist"

	var user m.Users
	result := au.DB.First(&user, "email = ?", strings.ToLower(userCredential.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email"})
		return
	}

	//Implement verified user email after testing

	// if !user.Verified {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "account not verified"})
	// 	return
	// }

	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := u.EncodeBase64(resetToken)
	user.PasswordResetToken = passwordResetToken
	user.PasswordResetExpiry = time.Now().Add(time.Minute * 15)
	au.DB.Save(&user)

	var firstName = user.Fullname

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := u.EmailData{
		URL:       config.ClientOrigin + "/resetpassword/" + resetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	u.SendEmail(&user, &emailData, "resetPassword.html")

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

func (au *AuthController) ResetPassword(c *gin.Context) {
	var payload *m.UpdatePasswordRequest
	resetToken := c.Params.ByName("resetToken")

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.Password_Confirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Password doesn't match"})
		return
	}

	hashedPassword, _ := u.HashingPassword(payload.Password)

	passwordResetToken := u.EncodeBase64(resetToken)

	var updatedUser m.Users
	result := au.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The reset token is invalid or has expired"})
		return
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	updatedUser.Updated_At = time.Now()
	updatedUser.Updater = "Reset Pass Token"
	au.DB.Save(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password successfully updated"})
}
