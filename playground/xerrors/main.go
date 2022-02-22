package main

import (
	"context"
	"fmt"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	_ = helloXerrors()
	_ = helloXerrors2()
	_ = handleContextCanceled()
}

func helloXerrors() error {
	baseErr := xerrors.New("base error")
	err := xerrors.Errorf("error in main: %w", baseErr)
	fmt.Printf("%v\n", xerrors.Is(err, baseErr))
	return err
}

func helloXerrors2() error {
	err := xerrors.Errorf("extauth: signup by ExternalAuthentication: repository: signup: pf: create signed token: rpc error: code = Canceled desc = %w", context.Canceled)
	fmt.Printf("check context canceled: %v\n", xerrors.Is(err, context.Canceled))
	return err
}

func handleContextCanceled() error {
	err := xerrors.New("extauth: signup by ExternalAuthentication: repository: signup: pf: create signed token: rpc error: code = Canceled desc = context canceled")
	fmt.Println(status.Code(err) == codes.Canceled)
	fmt.Println(status.Code(err) == codes.Canceled)
	return err
}
