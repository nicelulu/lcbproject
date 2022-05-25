package main

import (
	"fmt"
)

func main() {
	// 01 res
	Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
	// 02 只有简单概述，且有疑问点，如果可以望简单告知，感谢

	//03 res
	res := rand13()
	fmt.Println(res)


}
