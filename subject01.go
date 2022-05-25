package main


// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
	// TODO
	alen := len(a)
	if i == alen-1 {
		f(a)
	}
	onlyMap := map[rune]bool{}
	for index := i; index < alen; index++ {
		if _, ok := onlyMap[a[index]]; ok { // map中存在 跳过
			continue
		}
		onlyMap[a[index]] = true
		a[index], a[i] = a[i], a[index] // 交换 a[i] 固定在a[index]
		perm(a, f, i+1)
		a[index], a[i] = a[i], a[index] // 恢复交换
	}
}
