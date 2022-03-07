package services

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"stockApi/models"
	"stockApi/services/core"
)

var grpcServer *models.Grpc

func init() {
	initConnection()
}

func initConnection() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+os.Getenv("GRPCPORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("GRPC CONECTION FAILD: %s", err)
	}
	c := core.NewAuthServiceClient(conn)

	grpcServer = new(models.Grpc)
	grpcServer.Connection = conn
	grpcServer.Server = c
}
func RegisterUser(request core.RegisterRequest) error {

	_, err := grpcServer.Server.Register(context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(request *core.LoginRequest) (*core.LoginResponse, error) {

	t, err := grpcServer.Server.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func ValidateCookie(cookie string) (string, error) {
	reponse, err := grpcServer.Server.ValidateToken(context.Background(), &core.ValidateRequest{Token: cookie})
	if err != nil {
		return "", err
	}

	return reponse.Username, nil
}
