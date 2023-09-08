package lib1

import (
	"fmt"

	"golang.org/x/xerrors"
)

func Say() {
	fmt.Println("gov lib1 say: hi: commit 0.0.11")

	x := xerrors.New("cc")

	fmt.Println(x)

}

func Say2() {
	fmt.Println("say2")
}
