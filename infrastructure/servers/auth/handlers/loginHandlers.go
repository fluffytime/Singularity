package loginHandlers

import (
	"bytes"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fluffytime/singularity/app"
	"github.com/fluffytime/singularity/infrastructure/service/auth"
	"github.com/fluffytime/singularity/infrastructure/utils/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LoginHandler interface {
	Handle(*auth.LoginRequest) (*auth.LoginResponse, error)
}

type ValidateLoginParameters struct {
	next LoginHandler
}

type CheckUserExistence struct {
	db   *mgo.Database
	next LoginHandler
}

type TokenCreator struct {
	next LoginHandler
}

func (vlp *ValidateLoginParameters) Handle(lr *auth.LoginRequest) (*auth.LoginResponse, error) {
	if lr.Email != "" && lr.Password != "" && vlp.next != nil {
		return vlp.next.Handle(lr)
	}
	return nil, grpc.Errorf(codes.InvalidArgument, "Email or password are empty")
}

func (cue *CheckUserExistence) Handle(lr *auth.LoginRequest) (*auth.LoginResponse, error) {

	pass := scrypt.Encode(bytes.NewBufferString(lr.Password), bytes.NewBufferString(app.Config.ScryptSalt))
	c, _ := cue.db.C("users").Find(bson.M{"email": lr.Email, "password": pass}).Limit(1).Count()
	if c > 0 {
		return cue.next.Handle(lr)
	}
	return nil, grpc.Errorf(codes.Unauthenticated, "Invalid email or password")
}

func (tc *TokenCreator) Handle(lr *auth.LoginRequest) (*auth.LoginResponse, error) {
	claims := jwt.StandardClaims{
		Subject:   lr.Email,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().AddDate(0, 0, 14).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)
	tokenStr, err := token.SignedString([]byte(app.Config.JWTSigningKey))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	return &auth.LoginResponse{Success: true, Token: tokenStr}, nil
}

func CreateLoginChain(db *mgo.Database) LoginHandler {
	validator := new(ValidateLoginParameters)
	userChecker := new(CheckUserExistence)
	tokenGen := new(TokenCreator)

	validator.next = userChecker
	userChecker.next = tokenGen
	userChecker.db = db
	return validator
}
