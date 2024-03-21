package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Right, f)
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
}
