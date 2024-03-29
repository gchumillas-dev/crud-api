package manager

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	auth "github.com/gchumillas/crud-api/auth"
	"golang.org/x/crypto/bcrypt"
)

// User manages users.
type User struct {
	ID       string
	Username string
}

type userClaims struct {
	UserID string
	jwt.StandardClaims
}

// NewUser creates a user.
func NewUser(userID ...string) *User {
	id := ""
	if len(userID) > 0 {
		id = userID[0]
	}

	return &User{ID: id}
}

// NewToken creates a new token based on a privateKey.
func (user *User) NewToken(privateKey string, expiration time.Duration) string {
	claims := userClaims{UserID: user.ID}
	claims.ExpiresAt = time.Now().Add(expiration).Unix()

	return auth.NewToken(privateKey, claims)
}

// ReadUser reads a user.
func (user *User) ReadUser(db *sql.DB, ID string) (found bool) {
	stmt, err := db.Prepare(`
		select id, username
		from user where id = ?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	switch err := stmt.QueryRow(ID).Scan(&user.ID, &user.Username); {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	return true
}

// ReadUserByCredentials reads a user by username and password.
func (user *User) ReadUserByCredentials(db *sql.DB, uname string, upass string) (found bool) {
	stmt, err := db.Prepare(`
		select id, username, password
		from user where username = ?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var hashedPassword []byte
	switch err := stmt.QueryRow(uname).Scan(&user.ID, &user.Username, &hashedPassword); {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(upass)); err != nil {
		return false
	}

	return true
}

// ReadUserByToken reads a user by token.
func (user *User) ReadUserByToken(db *sql.DB, privateKey, signedToken string) (found bool) {
	claims := &userClaims{UserID: user.ID}
	if _, err := auth.ParseToken(privateKey, signedToken, claims); err != nil {
		return false
	}

	return user.ReadUser(db, claims.UserID)
}
