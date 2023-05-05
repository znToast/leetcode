package interview

import (
	"fmt"
	"leetcode/array"
	"math"
	"math/big"
	"sort"
	"strings"
	"time"
)

// 给定一个数组，数组全是整数。要求将奇数全部按顺序排在数组前面，偶数排在后面，不能新建其它数组
// 题解一、
func Array1(list []int) []int {
	length := 0
	for i := 0; i < len(list); i++ {
		if list[i]%2 == 1 {
			length++
		}
	}
	left := 0
	right := length
	for left < length && right < len(list) {
		if list[left]%2 == 0 && list[right]%2 == 1 {
			aa := list[left]
			list[left] = list[right]
			list[right] = aa
			left++
			right++
			continue
		}
		if list[left]%2 == 1 {
			left++
		}
		if list[right]%2 == 0 {
			right++
		}
	}
	array.Sort(list[:length])
	array.Sort(list[length:])
	return list
}

// 题解二、
func Array(list []int) []int {
	for i := 0; i < len(list); i++ {

	}
	return nil
}

// 给定一个数组1表示种有花,0表示没有。花与花之间不能紧挨着种。n表示可以种植的数量，可以则返回true，否则false
func Flower(list []int, n int) bool {
	left := 0
	right := 0
	m := 0
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			left = i - 1
			right = i + 1
			if left < 0 {
				if list[right] == 0 {
					m++
					i++
				}
			} else if left >= 0 && right < len(list) {
				if list[left] == 0 && list[right] == 0 {
					m++
					i++
				}
			} else if right >= len(list) {
				if list[left] == 0 {
					m++
					i++
				}
			}
		}
	}
	fmt.Println(m)
	return m >= n
}

// 找出两个字符串的最大公共交集
func FindString(s1, s2 string) string {
	str := ""
	if len(s1) < len(s2) {
		str = do(s1, s2)
	} else {
		str = do(s2, s1)
	}
	return str
}
func do(min, max string) string {
	len1 := len(min)
	str := ""
	for i := 0; i < len1; i++ {
		s := string(min[i])
		index := i
		for strings.Contains(max, s) {
			if len(str) < len(s) {
				str = s
			}
			if index < len1-1 {
				s = s + string(min[index+1])
			} else {
				break
			}
			index++
		}
	}
	return str
}

/*
斐波那契数列
*/
func FibonacciSequence(num int) {
	start := time.Now()
	nums := make([]*big.Int, num)
	for i := 0; i < num; i++ {
		if i <= 1 {
			nums[i] = big.NewInt(1)
		} else {
			temp := new(big.Int)
			nums[i] = temp.Add(nums[i-1], nums[i-2])
		}
		fmt.Printf("数位第%v位: %v\n", i, nums[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行所耗: %v\n", delta)
}

/*
快速排序
*/
func FastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//mid中间值(假定mid = values[0] 为中间值)	, i下标
	mid, i := nums[0], 1
	head, end := 0, len(nums)-1
	for head < end {
		// 把 小于mid 的值放到 head前面 反则亦然
		if nums[i] > mid {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	// 中间值替换成mid
	nums[head] = mid
	FastSort(nums[:head])
	FastSort(nums[head+1:])
}

/*
冒泡排序
*/
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	time.Sleep(time.Nanosecond * 100)
}

/*
插入排序
*/
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1], nums[j] = nums[j], key
			j--
		}
	}
}

/*
合并排序 将两个或多个有序数组进行排序
*/
func MergeSort(a []int, left, right int) {
	if left == right {
		return
	}
	if left < right {
		//拆分成n份进行排序
		mid := (left + right) / 2
		MergeSort(a, left, mid)
		MergeSort(a, mid+1, right)
		Merge(a, left, right)
	}
}

// 合并数组
func Merge(a []int, left, right int) {
	m := right - left + 1
	//临时数组b
	b := make([]int, m)
	left0 := left
	// 数组下标
	i := 0
	//中间数组的中间值
	mid := (left + right) / 2
	k := mid + 1
	for left <= mid && k <= right {
		if a[left] < a[k] {
			b[i] = a[left]
			i++
			left++
		} else {
			b[i] = a[k]
			i++
			k++
		}
	}
	if left > mid {
		for k <= right {
			b[i] = a[k]
			i++
			k++
		}
	}
	if k > right {
		for left <= mid {
			b[i] = a[left]
			left++
			i++
		}
	}
	for j := 0; j < m; j++ {
		a[left0] = b[j]
		left0++
	}
}

/*
统计n位数以内有多少个素数,不计入 0 ,1 ,埃筛法
*/
func Eratosthenes(n int) int {
	if n < 2 {
		fmt.Println("n 不能小于2")
		return n
	}
	count := 0
	// 素数存储
	isPrims := make([]bool, n)
	for i := 2; i < n; i++ {
		//判断是否为素数 ,是则进入
		if !isPrims[i] {
			count++
			for j := 2 * i; j < n; j += i {
				//将非素数改为true
				isPrims[j] = true
			}
		}
	}
	return count
}

/*
删除有序数组的重复项,每个元素只能出现一次,返回删除后数组的长度
考察双指针
*/
func RemoveDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

/*
寻找数组的中心下标,使得左边数组的和跟右边数组的和相等
*/
func PivotIndex(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		if total == num {
			return i
		}
		num -= nums[i]
	}
	return -1
}

/*
在不使用官方包情况, 算x的平方根 取整数
二分查找
*/
func BinarySearch(x int) int {
	if x < 2 {
		return 1
	}
	index, left, right := -1, 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			index = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}

type ListNote struct {
	Val  int
	Next *ListNote
}

/*
反转链表,递归实现
*/
func Recursion(head *ListNote) *ListNote {
	if head == nil || head.Next == nil {
		return head
	}
	node := Recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

/*
数组三个数的最大乘积
*/
func Sort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}

/*
两数之和，无序数组,获取两个数的下标(用数组返回)
*/
func Solution(nums []int, tager int) []int {
	isExist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := tager - nums[i]
		if _, ok := isExist[a]; ok {
			return []int{isExist[a], i}
		}
		isExist[nums[i]] = i
	}
	return nil
}

/*
排列硬币,摆成阶梯,第k行有k枚硬币 .n枚硬币能摆多少行(完成的)
*/
func ArrangeCoins(n int) (int, int) {
	if n < 1 {
		return 0, 0
	}
	count := 0
	for k := 1; k < n; k++ {
		count += k
		if ((k+1)*k/2 <= n) && ((k+2)*(k+1)/2 > n) {
			return k, count
		}
	}
	return -1, 0
}

/*
寻找符合要求的字符串
*/
func FindStrings(a, b string) int {
	count := 0
	l := len(b)
	for i := l; i <= len(a); i++ {
		if a[i-l:i] == b {
			count += 1
		}
	}
	return count
}

/*
华为第一题
*/
func test1() {
	n := 0
	ans := -1
	values := map[int][]int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		values[x] = append(values[x], i)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for _, v := range values {
		sort.Ints(v)
		if len(v) >= 2 {
			ans = max(ans, v[len(v)-1]-v[0])
		}
	}
	fmt.Printf("%d\n", ans)
}

/*
华为面试第二题 (未完成)
*/
func test2() {
	n := 0
	ans := ""
	j := 0
	res := map[int]string{}
	invert := func(s string) string {
		result := ""
		for k := range s {
			result = result + string(s[k])
		}
		return result
	}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := ""
		fmt.Scan(&x)
		if x == " " {
			// 反转
			ans += invert(res[j])
			j++
		}
		res[j] += x
	}
	fmt.Printf("%s\n", ans)
}

/*
华为面试第三题 Excel单元格数值统计
*/
func test3() {
	// table := map[int][]string{}
	// a := 0
	// b := 0
	// for {
	// 	n, _ := fmt.Scan(&a, &b)
	// 	if n == 0 {
	// 		break
	// 	} else {
	// 		fmt.Printf("%d\n", a+b)
	// 	}
	// }
}

/*
面试题67. 把字符串转换成整数
*/
func StrToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	ans := 0
	isSigna := 0
	sign := 1
	//删除前面的空格
	for str[0] == ' ' && len(str) >= 2 {
		str = str[1:]
	}
	if str == " " {
		return 0
	}
	for i := 0; i < len(str); i++ {
		// 传入的字符 不是数值
		if str[i] < '0' || str[i] > '9' {
			// 是否为+号
			isAdd := str[i] == '+'
			// 是否为-号
			isMinus := str[i] == '-'
			// 不是加号也不是减号
			if !isAdd && !isMinus {
				break
			}
		}
		if str[i] == '-' || str[i] == '+' {
			isSigna++
			if str[i] == '-' {
				sign = -1
			}
			if isSigna > 1 {
				return 0
			}
			continue
		}
		ans = ans*10 + int((str[i] - '0'))
	}
	ans = ans * sign
	if ans < math.MinInt32 {
		return math.MinInt32
	} else if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}
