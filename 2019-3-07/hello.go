package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/BootyDog/golang-learn/2019-3-07/utils"
)

func main() {
	who := "World"
	if len(os.Args)>1 {
		who = strings.Join(os.Args[1:],"")
	}
	fmt.Println("Hello",who)
	utils.Printhello()
}
