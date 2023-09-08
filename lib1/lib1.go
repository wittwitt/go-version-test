package lib1

import (
	"fmt"

	"golang.org/x/xerrors"
)

func Say() {
	fmt.Println("gov lib1 say: hi: commit 0.0.4")

	err := xerrors.New("x")
	fmt.Println("add xerr", err)
}

func Say2() {
	fmt.Println("say2")
}
