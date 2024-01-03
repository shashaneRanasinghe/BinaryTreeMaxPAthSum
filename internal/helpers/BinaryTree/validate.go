package helpers

import (
	"fmt"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/pkg/consts"
)

// ValidateRootNode validates if the given root node is not empty and exists in the node list
func ValidateRootNode(tree models.Tree) (models.Node, error) {
	if tree.Root == "" {
		return models.Node{}, fmt.Errorf("%v", consts.RootNodeNotProvided)
	}

	var rootNode models.Node
	for _, node := range tree.Nodes {
		if node.ID == tree.Root {
			rootNode = node
		}
	}
	if rootNode.ID == "" {
		return models.Node{}, fmt.Errorf("%v", consts.RootNodeNotFound)
	}
	return rootNode, nil
}
