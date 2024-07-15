package codeExam

import (
	"fmt"
	"math"
)

type TreeNode2 struct {
	Value int
	Left  *TreeNode2
	Right *TreeNode2
}

func InitTree2(value int) *TreeNode2 {
	return &TreeNode2{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func InsertTree2(root *TreeNode2, value int) {
	if root == nil {
		return
	}

	if value < root.Value {
		if root.Left == nil {
			root.Left = &TreeNode2{Value: value}
		} else {
			InsertTree2(root.Left, value)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode2{Value: value}
		} else {
			InsertTree2(root.Right, value)
		}
	}
}

func printSpace(n float64, removed *TreeNode2) {
	for ; n > 0; n-- {
		fmt.Print("\t")
	}
	if removed == nil {
		fmt.Print(" ")
	} else {
		fmt.Print(removed.Value)
	}
}

func heightOfTree(root *TreeNode2) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(heightOfTree(root.Left)), float64(heightOfTree(root.Right))))
}

func printBinaryTree(root *TreeNode2) {
	treeLevel := []*TreeNode2{root}
	temp := []*TreeNode2{}
	counter := 0
	height := heightOfTree(root) - 1
	numberOfElements := math.Pow(2, float64(height+1)) - 1

	for counter <= height {
		removed := treeLevel[0]
		treeLevel = treeLevel[1:]

		if len(temp) == 0 {
			printSpace(numberOfElements/math.Pow(2, float64(counter+1)), removed)
		} else {
			printSpace(numberOfElements/math.Pow(2, float64(counter)), removed)
		}

		if removed == nil {
			temp = append(temp, nil)
			temp = append(temp, nil)
		} else {
			temp = append(temp, removed.Left)
			temp = append(temp, removed.Right)
		}

		if len(treeLevel) == 0 {
			fmt.Println()
			fmt.Println()
			treeLevel = temp
			temp = []*TreeNode2{}
			counter++
		}
	}
}

func TestTreeTree2() {
	// Example usage
	root := InitTree2(5)
	InsertTree2(root, 3)
	InsertTree2(root, 7)
	InsertTree2(root, 2)
	InsertTree2(root, 4)
	InsertTree2(root, 10)
	InsertTree2(root, 6)
	InsertTree2(root, 8)
	InsertTree2(root, 11)

	// Print the tree values using inorder traversal
	printBinaryTree(root)
}