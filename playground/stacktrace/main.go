package main

import (
	"context"
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	if err := hierarchy1(); err != nil {
		fmt.Printf("%+v\n", err)
	}
	if err := hierarchy1(); err != nil {
		wraperr := xerrors.Errorf("main: %w", err)
		fmt.Printf("%+v\n", wraperr)
	}
	if err := hierarchy1WithFrame(); err != nil {
		wraperr := xerrors.Errorf("main: %w", err)
		fmt.Printf("%+v\n", wraperr)
	}
}

func hierarchy1() error {
	return hierarchy2()
}

func hierarchy2() error {
	return hierarchy3()
}

func hierarchy3() error {
	err := context.Canceled
	return xerrors.Errorf("hierarchy3: %w", err)
}

func hierarchy1WithFrame() error {
	return xerrors.Errorf("hierarchy1WithFrame: %w", hierarchy2())
}
