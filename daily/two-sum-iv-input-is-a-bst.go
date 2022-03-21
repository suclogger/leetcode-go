package daily

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	dfs(root, &m)
	for key, _ := range m {
		if key == k-key {
			continue
		}
		if _, ok := m[k-key]; ok {
			return true
		}
	}
	return false
}

func dfs(root *TreeNode, m *map[int]int) {
	if root == nil {
		return
	}
	(*m)[root.Val] = 1
	dfs(root.Left, m)
	dfs(root.Right, m)
}
