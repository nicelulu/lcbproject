package main

/**
* 两次rand5 获取25个数，取前13个
 */
func rand13() int {
	num := (rand5()-2)*5 + rand5()
	for num > 13 {
		num = (rand5()-2)*5 + rand5()
	}
	return 1 + num%13
}

func rand5() (res int) {
	return res
}
