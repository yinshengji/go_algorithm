package smallestFactorization
/*
 * leetcode: 625. 最小因式分解
 * 给定一个正整数 a，找出最小的正整数 b 使得 b 的所有数位相乘恰好等于 a。
 * 如果不存在这样的结果或者结果不是 32 位有符号整数，返回 0。
 * 48 -> 68
 */
func smallestFactorization(a int) int {
	if a < 2 {
		return a;
	}

	mul, res := 1, 0
	for i := 9; i >= 2; i-- {
		for a % i == 0 {
			a /= i
			res = mul * i + res
			mul = mul * 10
		}
	}
	if a == 1 && res <= (1 << 31 -1) {
		return res
	}

	return 0
}