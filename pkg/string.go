package pkg

const (
	inf = 0x3f3f3f3f
)

func GetMinimalEditDistance(a, b string) int {
	n, m := len(a), len(b)
	f := make([][]int, n+1)
	a = " " + a
	b = " " + b
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			f[i][j] = inf
		}
	}
	for i := 0; i <= n; i++ {
		f[i][0] = i
	}
	for i := 0; i <= m; i++ {
		f[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if a[i] == b[j] {
				f[i][j] = Min(f[i][j], f[i-1][j-1])
			} else {
				f[i][j] = Min(f[i-1][j], f[i][j-1]) + 1
				f[i][j] = Min(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	return f[n][m]
}

func GetDF(ids []int) int {
	res := 0
	used := make(map[int]bool, len(ids))
	for _, id := range ids {
		if _, has := used[id]; !has {
			res += 1
			used[id] = true
		}
	}
	return res
}

func GetTF(need int, ids []int) int {
	cnt := 0
	for _, id := range ids {
		if id == need {
			cnt++
		}
	}
	return cnt
}
