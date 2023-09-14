package main

import (
	"fmt"

	"github.com/lijunchen/gomod/character"
	"github.com/lijunchen/gomod/character/space"
	"github.com/lijunchen/gomod/character/symbol"

	"github.com/lijunchen/gomod/hello"
	"github.com/lijunchen/gomod/world"
)

func main() {
	s := character.FirstToUpper(hello.Hello()) + symbol.Comma() + space.Space() + world.World() + symbol.Exclamation()
	fmt.Println(s)
}
