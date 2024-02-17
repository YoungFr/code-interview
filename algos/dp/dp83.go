package dp

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	f := make([]int, n)
	f[0] = nums[0]
	f[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		f[i] = max(nums[i]+f[i-2], f[i-1])
	}
	return f[n-1]
}
