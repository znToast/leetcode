package easy

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
1005. K 次取反后最大化的数组和
*/
func LargestSumAfterKNegations(nums []int, k int) int {
	min_index := 0
	res := 0
	// 排序
	sort.Ints(nums)
	for i := range nums {
		//  让负数进行反转
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
		// 获取最小值的下标
		if nums[min_index] > nums[i] {
			min_index = i
		}
	}
	//k还剩余时全部用于最小的那个数进行反转
	if k%2 == 1 {
		nums[min_index] *= -1
	}
	// 计算数组总和
	for _, v := range nums {
		res += v
	}
	return res
}

/*
LC 13. 罗马数字转整数
*/
func RomanToInt(s string) int {
	//罗马数字对应的阿拉伯数字
	symbolValues := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	ans := 0
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		//下一个字符对应的数值大于当前字符对应数值
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

/*
LC 35.搜索插入位置
*/
func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			continue
		}
		return i
	}
	return len(nums)
}

/*
LC 58.最后一个单词的长度
*/
func LengthOfLastWord(s string) int {
	lenght := len(s)
	ans := 0
	for i := lenght - 1; i >= 0; i-- {
		if s[i] == ' ' && ans == 0 {
			continue
		}
		if s[i] == ' ' {
			return ans
		}
		ans++
	}
	return ans
}

/*
LC 67.二进制求和
*/
func AddBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

/*
LC 69.x的平方根
*/
func MySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l < r {
		// 二分查找
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 链
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
LC 83.删除排序链表中的重复元素
*/
func DeleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC 100.相同的树
*/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	ans := true
	// 深度优先遍历
	dfs := func(p *TreeNode, q *TreeNode) {}
	dfs = func(p, q *TreeNode) {
		if p == nil && q == nil {
			return
		} else if p != nil && q != nil {
			if p.Val != q.Val {
				ans = false
			}
			dfs(p.Left, q.Left)
			dfs(p.Right, q.Right)
		} else {
			ans = false
		}
	}
	dfs(p, q)
	return ans
}

/*
LC 88.合并两个有序数组
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n
	for n > 0 {
		// 从后面开始排序
		if m > 0 && nums1[m-1] > nums2[n-1] {
			i--
			m--
			nums1[i] = nums1[m]
		} else {
			i--
			n--
			nums1[i] = nums2[n]
		}
	}
}

/*
LC 101. 对称二叉树
*/
func IsSymmetric(root *TreeNode) bool {
	isMirror := func(l, r *TreeNode) bool { return false }
	isMirror = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
	}
	return isMirror(root, root)
}

/*
LC 104. 二叉树的最大深度
深度优先算法
*/
func MaxDepth(root *TreeNode) int {
	ans := 0
	dfs := func(r *TreeNode, depth int) {}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		ans = max(depth, ans)
		dfs(r.Left, depth+1)
		dfs(r.Right, depth+1)
	}
	dfs(root, 1)
	return ans
}

/*
LC 108.将有序数组转换为二叉搜索树
*/
func SortedArrayToBST(nums []int) *TreeNode {
	dfs := func(nums []int, l, r int) *TreeNode { return nil }
	dfs = func(nums []int, l, r int) *TreeNode {
		if l < r {
			return nil
		}
		mid := l + (r-l)/2
		node := &TreeNode{nums[mid], nil, nil}
		node.Left = dfs(nums, l, mid-1)
		node.Right = dfs(nums, mid+1, r)
		return node
	}
	return dfs(nums, 0, len(nums)-1)
}

/*
LC 110.平衡二叉树
*/
func IsBalanced(root *TreeNode) bool {
	return false
}

/*
LC 111.二叉树的最小深度
*/
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(MinDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(MinDepth(root.Right), minD)
	}
	return minD + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
LC 112.路径总和
*/
func HasPathSum(root *TreeNode, targetSum int) bool {
	result := false
	nums := make([]int, 0)
	dfs := func(node *TreeNode, count int) {}
	dfs = func(node *TreeNode, count int) {
		if node == nil {
			return
		}
		// 到达叶子节点的时候
		if node.Left == nil && node.Right == nil {
			count += node.Val
			// 判断下路径总和是否等于tag
			if count == targetSum {
				result = true
			}
			nums = append(nums, count)
		}
		if node.Left != nil {
			dfs(node.Left, count+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, count+node.Val)
		}

	}
	dfs(root, 0)
	fmt.Printf("nums: %v\n", nums)
	return result
}

/*
Lc 118.杨辉三角
*/
func Generate(numRows int) [][]int {
	nums := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			// 当横坐标在0 或者 numRows - 1的位置 赋值为1
			if j == 0 || j == i {
				nums[i] = append(nums[i], 1)
				continue
			}
			nums[i] = append(nums[i], nums[i-1][j]+nums[i-1][j-1])
		}
	}
	return nums
}

/*
Lc 119.杨辉三角II
*/
func GetRow(rowIndex int) []int {
	nums := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			// 当横坐标在0 或者 numRows - 1的位置 赋值为1
			if j == 0 || j == i {
				nums[i] = append(nums[i], 1)
				continue
			}
			nums[i] = append(nums[i], nums[i-1][j]+nums[i-1][j-1])
		}
	}
	return nums[rowIndex]
}

/*
lC 141.环形链表
*/
func HasCycle(head *ListNode) bool {
	//使用map记录
	/* node := head*/
	/*temp := make(map[*ListNode]bool)*/
	/*for node != nil && node.Next != nil {*/
	/*if ok := temp[node]; !ok {*/
	/*temp[node] = true*/
	/*}*/
	/*if ok := temp[node.Next]; ok {*/
	/*return true*/
	/*}*/
	/*node = node.Next*/
	/*}*/
	/*return false*/

	// 将经过的节点的指针指向自己, 等到后面指向这个节点时只需要判断 head == head.Next
	for head != nil {
		temp := head.Next
		if head != head.Next {
			head.Next = head
		} else {
			return true
		}
		head = temp
	}
	return false
}

/*
LC 168.Excel表列名称
*/
func ConvertToTitle(columnNumber int) string {
	ans := ""
	// 相当于26进制转换
	for columnNumber > 0 {
		columnNumber--
		ans = string(rune(columnNumber%26+'A')) + ans
		columnNumber /= 26
	}
	return ans
}

/*
LC 169.多数元素
*/
func MajorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v] += 1
	}
	for k, v := range numMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

/*
LC 171.Excel 表列序号
*/
func TitleToNumber(columnTitle string) int {
	ans := 0
	for k := range columnTitle {
		ans = ans*26 + int((columnTitle[k] - 'A' + 1))
		fmt.Printf("ans: %v\n", ans)
	}
	return ans
}

/*
LC 190. 颠倒二进制位
*/
func ReverseBits(num uint32) uint32 {
	return 0
}

/*
LC 191.位1的个数
*/
func HammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		fmt.Println(1 << i & num)
		if 1<<i&num > 0 {
			ans++
		}
	}
	return ans
}

/*
LC 219.存在重复元素II
*/
func ContainsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 判断是否value是否存在，如果存在了 则计算 j-i <= k
		if _, ok := indexMap[nums[i]]; ok && (i-indexMap[nums[i]]) <= k {
			return true
		} else {
			// 存储下标
			indexMap[nums[i]] = i
		}
	}
	return false
}

/*
LC 203. 移除链表元素
*/
func RemoveElements(head *ListNode, val int) *ListNode {
	newNode := &ListNode{}
	newNode.Next = head
	node := newNode
	for node.Next != nil {
		for node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return newNode.Next
}

/*
LC 225.用队列实现栈
*/
type MyStack struct {
	nums []int
}

func Constructor() MyStack {
	return MyStack{}
}
func (this *MyStack) Push(x int) {
	this.nums = append(this.nums, x)
	return
}
func (this *MyStack) Pop() int {
	re_len := len(this.nums)
	if re_len == 0 {
		return 0
	}
	result := this.nums[re_len-1]
	this.nums = this.nums[0 : re_len-1]
	return result
}
func (this *MyStack) Top() int {
	re_len := len(this.nums)
	if re_len == 0 {
		return 0
	}
	result := this.nums[re_len-1]
	return result
}
func (this *MyStack) Empty() bool {
	re_len := len(this.nums)
	if re_len == 0 {
		return true
	}
	return false
}

/*
LC 226.翻转二叉树
*/
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Right
	// 左右节点互换
	root.Right = InvertTree(root.Left)
	root.Left = InvertTree(temp)
	return root
}

/*
LC 228.汇总区间
*/
func SummaryRanges(nums []int) []string {
	left := 0
	ans := []string{}
	for i := 1; i < len(nums); i++ {
		// 判断是否连续
		if nums[i-1]+1 == nums[i] {
			continue
		} else {
			if left == i-1 {
				ans = append(ans, fmt.Sprintf("%v", nums[left]))
			} else {
				ans = append(ans, fmt.Sprintf("%v->%v", nums[left], nums[i-1]))
			}
			left = i
		}
	}
	return ans
}

/*
LC 231. 2 的幂
*/
func IsPowerOfTwo(n int) bool {
	i := 0
	for (2 << 0) <= n {
		fmt.Println(2 << i)
		if 2<<i == n {
			return true
		}
		i++
	}
	return false
}

/*
LC 258.各位相加
*/
func AddDigits(num int) int {
	ans := 0
	if num < 10 {
		return num
	}
	for num >= 10 {
		temp := num
		total := 0
		for temp > 0 {
			total += temp % 10
			temp /= 10
		}
		if total < 10 {
			ans = total
			break
		} else {
			num = total
		}
	}
	return ans
}

/*
LC 263.丑数
*/
func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else if n == 1 {
			return true
		} else {
			return false
		}
	}
	return true
}

/*
LC 268.丢失的数字
*/
func MissingNumber(nums []int) int {
	//sort.Ints(nums)
	//for k, v := range nums {
	//if k != v {
	//return k
	//}
	//}
	//return len(nums)
	n := len(nums)
	total := n * (n + 1) / 2
	num := 0
	for _, v := range nums {
		num += v
	}
	return total - num
}

/*
LC 290.单词规律
*/
func WordPattern(pattern string, s string) bool {
	word2ch := map[string]byte{}
	ch2word := map[byte]string{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}

/*
lC 292. Nim 游戏
*/
func CanWinNim(n int) bool {
	// 如果被4整除必输
	if n%4 == 0 {
		return false
	}
	return true
}

/*
LC 326.3的幂
*/
func IsPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

/*
LC 338. 比特位计数
*/
func CountBits(n int) []int {
	// Brian Kernighan 算法 统计 二进制有多少个1 ,每次执行会把最右边的1删除
	result := make([]int, n+1)
	count := func(n int) (count int) {
		for n > 0 {
			count++
			n &= n - 1
		}
		return
	}
	for i := range result {
		result[i] = count(i)
	}
	return result
}

func Count(n int) int {
	ans := 0
	for n > 0 {
		ans++
		n &= (n - 1)
	}
	return ans
}

/*
LC 374.猜数字大小
*/
func GuessNumber(n int) int {
	left := 0
	right := n
	ans := 0
	for left <= right {
		mid := (right-left)/2 + left
		fmt.Printf("mid: %v\n", mid)
		switch guese(mid) {
		// pick > mid
		case 1:
			{
				left = mid + 1
			}
		case -1:
			{
				right = mid - 1
			}
		default:
			{
				ans = mid
				return ans
			}
		}
	}
	return ans
}
func guese(i int) int {
	if i > 6 {
		return 1
	} else if i < 6 {
		return -1
	} else {
		return 6
	}
}

/*
LC 575. 分糖果
*/
func DistributeCandies(candyType []int) int {
	ans := 0
	mid := len(candyType) / 2
	type_map := make(map[int]int, 0)
	for _, v := range candyType {
		type_map[v] = 1
	}
	ans = len(type_map)
	if ans > mid {
		return mid
	} else {
		return ans
	}
}

/*
LC 485. 最大连续 1 的个数
*/
func FindMaxConsecutiveOnes(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	cut := 0
	maxAns := 0
	for _, v := range nums {
		if v == 1 {
			cut++
		} else {
			maxAns = max(cut, maxAns)
			cut = 0
		}

	}
	maxAns = max(maxAns, cut)
	return maxAns
}

/*
LC 492. 构造矩形
*/
func ConstructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}
