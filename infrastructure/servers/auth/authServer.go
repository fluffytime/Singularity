package servers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fluffytime/singularity/infrastructure/database/models"
	"github.com/fluffytime/singularity/infrastructure/servers/auth/handlers"
	auth "github.com/fluffytime/singularity/infrastructure/service/auth"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AuthServer Represents Authentication server provider
type AuthServer struct {
	DB *mgo.Database
}

func (srv *AuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	c, _ := srv.DB.C("users").Find(bson.M{"email": req.Email}).Limit(1).Count()
	if c > 0 {
		return nil, grpc.Errorf(codes.AlreadyExists, "User %s already exists", req.Email)
	}
	salt := 666
	saltArray := []byte(strconv.Itoa(salt))

	data, err := scrypt.Key([]byte(req.Password), saltArray, 16384, 8, 1, 32)
	pass := hex.EncodeToString(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	usr := dbmodels.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Groups:    nil,
		Id:        bson.NewObjectId(),
		Password:  pass,
	}

	srv.DB.C("users").Insert(usr)
	uj, _ := json.Marshal(usr)
	fmt.Printf("%s", uj)

	return &auth.RegisterResponse{}, nil
}

//Login user authentication via email and passwordHash
func (srv *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	h := loginHandlers.CreateLoginChain(srv.DB)
	c, err := h.Handle(req)
	return c, err
}
