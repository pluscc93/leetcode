package main

import (
	"strings"
	"fmt"
)

func main() {
	s := strings.Split(strings.Replace(",,",",,",",",-1), "")
	fmt.Println(s, len(s))
}
