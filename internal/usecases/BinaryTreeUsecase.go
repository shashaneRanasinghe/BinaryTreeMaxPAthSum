package usecases

import (
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"math"
)

type BinaryTree interface {
	FindMaxPathSum(rootNode *models.Node, tree *models.Tree) int
}

type binaryTreeUsecase struct {
}

var maxPathSum float64

func NewBinaryTreeUsecase() BinaryTree {
	return &binaryTreeUsecase{}
}

func (m *binaryTreeUsecase) FindMaxPathSum(rootNode *models.Node, tree *models.Tree) int {
	maxPathSum = math.Inf(-1)
	findPathSum(rootNode, tree)
	return int(maxPathSum)
}

// findPathSum will calculate the maximum path sum in a tree by recursively starting from the node given.
func findPathSum(node *models.Node, tree *models.Tree) float64 {
	if node == nil {
		return 0
	}

	leftMax := math.Max(findPathSum(FindNodeByID(node.Left, tree), tree), 0)
	rightMax := math.Max(findPathSum(FindNodeByID(node.Right, tree), tree), 0)
	maxPathSum = math.Max(maxPathSum, leftMax+rightMax+float64(node.Value))
	return math.Max(leftMax, rightMax) + float64(node.Value)
}

// FindNodeByID will find the node based on the given node ID from Node list of the tree
func FindNodeByID(id string, tree *models.Tree) *models.Node {
	if id == "null" || id == "" {
		return nil
	}
	for i := range tree.Nodes {
		if id == tree.Nodes[i].ID {
			return &tree.Nodes[i]
		}
	}
	return nil
}
