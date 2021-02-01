package main

import (
	"fmt"
	"gin-vue-admin/utils"
	uuid "github.com/satori/go.uuid"
)

func main()  {
	p := utils.MD5V([]byte("niu123"))
	ul := uuid.NewV4().String()
	fmt.Printf(p)
	fmt.Printf("\n" + ul)
}
