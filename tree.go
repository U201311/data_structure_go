package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*102  dfs*/
func DFSlevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	dfs(root, 0, &result)

	return result
}
func dfs(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	var temp []int
	if len(*result) < level+1 {
		*result = append(*result, temp)
	}
	(*result)[level] = append((*result)[level], root.Val)
	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)

}

/*bfs*/
func BFSlevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{}
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		level_size := len(queue)
		current_level := []int{}
		for i := 0; i < level_size; i++ {
			node := queue[i]
			current_level = append(current_level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[level_size:]
		result = append(result, current_level)

	}
	return result
	// result := [][]int{}
	// if root == nil {
	//     return result
	// }
	// queue := []*TreeNode{}
	// queue = append(queue, root)
	// for len(queue) > 0 {
	//     level_size := len(queue)
	//     current_level := []int{}
	//     for i := 0; i < level_size; i++ {
	//         node := queue[i]
	//         current_level = append(current_level, node.Val)
	//         if node.Left != nil {
	//             queue = append(queue, node.Left)
	//         }
	//         if node.Right != nil {
	//             queue = append(queue, node.Right)
	//         }
	//     }
	//     queue = queue[level_size:]
	//     result = append(result, current_level)
	// }
	// return result

}

/*242*/
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	res = []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right

	}
	return res
}

//leetcode 20
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// func isValid(s string) bool {
//     n := len(s)
//     if n % 2 == 1 {
//         return false
//     }
//     pairs := map[byte]byte{
//         ')': '(',
//         ']': '[',
//         '}': '{',
//     }
//     stack := []byte{}
//     for i := 0; i < n; i++ {
//         if pairs[s[i]] > 0 {
//             if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
//                 return false
//             }
//             stack = stack[:len(stack)-1]
//         } else {
//             stack = append(stack, s[i])
//         }
//     }
//     return len(stack) == 0
// }

func quickSort(nums []int, start, end int) {
	if start > end {
		return
	}
	i, j := start, end
	mid := nums[(start+end)/2]
	for i <= j {
		for nums[i] < mid && i <= j {
			i++
		}
		for nums[j] > mid && i <= j {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	// if i > start {
	quickSort(nums, start, j)
	// }
	// if j < end {
	quickSort(nums, i, end)
	// }
	// return

}
