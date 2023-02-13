package main

import (
	"Go-Learning/familyaccount/utils"
	"fmt"
)

func main() {
	fmt.Println("这个是面向对象的方式完成")
	utils.NewFamilyAccount().MainMenu()
}
