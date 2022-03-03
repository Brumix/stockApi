package services

import (
	"fmt"
	"golang.org/x/net/context"
	"stockApi/services/core"
)

func ChangePassowrd(request *core.ChangePasswordRequest) error {
	response, err := grpcServer.Server.ChangePassword(context.Background(), request)
	if err != nil {
		return err
	}
	if response.Status {
		return nil
	}
	return fmt.Errorf("error changing the passowrd")
}
func DeleteUser(request *core.DeleteUserRequest) error {

	response, err := grpcServer.Server.DeleteUser(context.Background(), request)
	if err != nil {
		return err
	}
	if response.Status {
		return nil
	}
	return fmt.Errorf("error deleting the user")
}
