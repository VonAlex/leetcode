package leetcode

/**
 * LeetCode 题1 两数之和
 * 给定一个整数数组 nums 和一个目标值 target，
 * 请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
 */
// 这里的数组应该是不含相同元素的数组
// TwoSum 暴力法
// 时间复杂度 O(n^2)，空间复杂度 O(1)
func TwoSum(nums []int, target int) []int {
	res := []int{0, 0}
	len := len(nums)
	for i, num := range nums {
		for j := i + 1; j < len; j++ {
			if num+nums[j] != target {
				continue
			}
			res[0] = i
			res[1] = j
			return res
		}
	}
	return nil
}

// TwoSum2 两遍 hash 表法
// 时间复杂度 O(n)，空间复杂度 O(n)
func TwoSum2(nums []int, target int) []int {
	numIdx := make(map[int]int)
	// 如果说数组中有相同的数字，那么使用 hash 是会覆盖掉相同元素
	for idx, num := range nums {
		numIdx[num] = idx
	}

	for i, num := range nums {
		j, ok := numIdx[target-num]
		// 防止索引到自身，如 6=3+3
		// 如果是数组中是 3，但是给出的 taget 是 6，这样是有问题的
		if ok && i != numIdx[j] {
			return []int{i, j}
		}
	}
	return nil
}

// TwoSum3 一遍 hash 表法
// 时间复杂度 O(n)，空间复杂度 O(n)
// 由于相同元素的覆盖，一遍遍历 hash 表和两遍遍历，获得的结果是不同的
func TwoSum3(nums []int, target int) []int {
	res := []int{0, 0}
	numIdx := make(map[int]int)
	for i, num := range nums {
		j, ok := numIdx[target-num]
		if !ok {
			numIdx[num] = i
			continue
		}
		res[0] = j
		res[1] = i
		return res
	}
	return nil
}

/**
 * 剑指 offer 面试题03 找出数组中重复的数字
 * 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
 * https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
 *
 * 输入：[2, 3, 1, 0, 2, 5, 3]
 * 输出：2 或 3
 */

// FindRepeatNumber 利用 map 结构
// 时间复杂度和空间复杂度 O(n)
func FindRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sets := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := sets[num]; ok {
			return num
		}
		sets[num] = struct{}{}
	}
	return -1
}

// FindRepeatNumber2 强调了长度为 n 的数组内数字范围都是 0-n-1。
// 可以想到，当数组内的数字经过排序后，那么下标 i，就应该等于数字 num，第一个不相等的，就是题解。
// 可以先排序，时间复杂度 O(nlogn)，再求解
// 也可以通过交换数字达到目的。时间复杂度 O(n),空间复杂度 O(1)
func FindRepeatNumber2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i, num := range nums {
		if i == num { // 下标 == 数字，表示该数字已经在最早位置上了
			continue
		}
		if num == nums[num] { // 数字 num 排序后应该在序号为 num 的位置，但是这个位置上已经有一个数字了，等于 num，那么这就是题解
			return num
		}
		nums[i], nums[num] = nums[num], nums[i]
	}
	return -1
}

/**
 * 面试题 10.01 合并排序的数组
 * https://leetcode-cn.com/problems/sorted-merge-lcci/
 *
 * 输入:
 * 		A = [1,2,3,0,0,0], m = 3
 * 		B = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 */

// 解法1 双指针
// 充分利用 A 和 B 已经排好序的特点
// 时间复杂度和空间复杂度都是 O(m+n)
func Merge(A []int, m int, B []int, n int) {
	C := make([]int, m+n) // 准备一个新数组，足以容纳 A + B
	var iA, iB, iC int
	for ; iC < m+n; iC++ {
		if iA == m || iB == n { // 遍历完 A 或者 B 为止
			break
		}
		if A[iA] < B[iB] {
			C[iC] = A[iA]
			iA++
		} else {
			C[iC] = B[iB]
			iB++
		}
	}
	// 遍历剩下的 A
	for ; iA < m; iA++ {
		C[iC] = A[iA]
		iC++
	}
	// 遍历剩下的 B
	for ; iB < n; iB++ {
		C[iC] = B[iB]
		iC++
	}
	copy(A, C)
}

/*
 * LeetCode 题 1103 分糖果 II
 * https://leetcode-cn.com/problems/distribute-candies-to-people/
 */

// 解法 1 暴力解法，不断循环
func DistributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	loop := 0
	candy := 0
	for candies != 0 {
		for i := 0; i < num_people; i++ {
			// i 从 0 开始
			// 第 i 位小朋友应该分到的糖果数量是（小朋友序号i+1）+ 总人数 * 第 n 轮（n 从 1 开始）
			candy = (i + 1) + num_people*loop
			// 判断剩下的糖果是否足够给这个小朋友
			if candies > candy {
				res[i] += candy
				candies -= candy
			} else {
				res[i] += candies
				candies = 0
			}
		}
		loop++
	}
	return res
}

// 解法 2 等差数列求和
// 参考 https://leetcode-cn.com/problems/distribute-candies-to-people/solution/fen-tang-guo-ii-by-leetcode-solution/

/*
 * 剑指 offer 面试题21. 调整数组顺序使奇数位于偶数前面
 * https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
 * 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
 * 示例：
 *     输入：nums = [1,2,3,4]
 *     输出：[1,3,2,4] （注：[3,1,2,4] 也是正确的答案之一）
 */
// 前后指针法
// 时间复杂度O(n),空间复杂度O(1)
// 扩展：判断指针是否做前移个后移的条件可以抽象成一个函数，这样就可以支持更多的情况，比如所有的正数放后面，负数放前面
func Exchange(nums []int) []int {
	low := 0
	high := len(nums) - 1
	for low < high {
		if nums[low]&1 == 1 { //  low 为奇数，满足最终条件，往后移动一个
			low++
			continue
		}
		if nums[high]&1 != 1 { // high 为偶数，满足最终条件，往前移动一个
			high--
			continue
		}
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
	return nums
}
