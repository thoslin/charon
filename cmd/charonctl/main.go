package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/mnemosyne"
	"github.com/piotrkowalczuk/ntypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var config configuration

func init() {
	config.init()
}

func main() {
	config.parse()

	switch config.cmd() {
	case "register":
		registerUser(config)
	default:
		fmt.Printf("unknown command %s", config.cmd())
	}
}

func client() (client charon.RPCClient, ctx context.Context) {
	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client = charon.NewRPCClient(conn)
	ctx = context.Background()
	if !config.noauth {
		resp, err := client.Login(context.Background(), &charon.LoginRequest{
			Username: config.username,
			Password: config.password,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ctx = metadata.NewContext(ctx, metadata.Pairs(mnemosyne.AccessTokenMetadataKey, string(resp.AccessToken.Encode())))
	}

	return
}

func registerUser(config configuration) {
	c, ctx := client()
	resp, err := c.CreateUser(ctx, &charon.CreateUserRequest{
		Username:      config.register.username,
		PlainPassword: config.register.password,
		FirstName:     config.register.firstName,
		LastName:      config.register.lastName,
		IsSuperuser:   &ntypes.Bool{Bool: config.register.superuser, Valid: true},
	})
	if err != nil {
		fmt.Printf("charonctl: registration failure: %s", grpc.ErrorDesc(err))
		os.Exit(1)
	}

	if config.register.superuser {
		fmt.Printf(`superuser "%s" has been created`, resp.User.Username)
	} else {
		fmt.Printf(`user "%s" has been created`, resp.User.Username)
	}
}
