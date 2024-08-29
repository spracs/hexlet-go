package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spracs/hexlet-go/greeting"
)

func main() {
	fmt.Println(color.GreenString(greeting.Hello()))
}
