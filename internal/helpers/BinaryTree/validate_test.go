package helpers

import (
	"errors"
	"fmt"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/pkg/consts"
	"github.com/tryfix/log"
	"go.uber.org/mock/gomock"
	"testing"
)

var (
	testTree1 = models.Tree{
		Nodes: []models.Node{
			{
				ID:    "1",
				Left:  "2",
				Right: "3",
				Value: 1,
			},
		},
		Root: "1",
	}
	testTree2 = models.Tree{
		Nodes: []models.Node{
			{
				ID:    "1",
				Left:  "2",
				Right: "3",
				Value: 1,
			},
		},
	}

	testTree3 = models.Tree{
		Nodes: []models.Node{
			{
				ID:    "1",
				Left:  "2",
				Right: "3",
				Value: 1,
			},
		},
		Root: "null",
	}
)

func TestValidateRootNode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		ID           string
		Tree         models.Tree
		expectedNode models.Node
		expectedErr  error
	}

	tests := []test{
		{
			ID:           "1",
			Tree:         testTree1,
			expectedNode: testTree1.Nodes[0],
			expectedErr:  nil,
		},
		{
			ID:           "2",
			Tree:         testTree2,
			expectedNode: models.Node{},
			expectedErr:  fmt.Errorf("%v", consts.RootNodeNotProvided),
		},
		{
			ID:           "3",
			Tree:         testTree3,
			expectedNode: models.Node{},
			expectedErr:  fmt.Errorf("%v", consts.RootNodeNotFound),
		},
	}

	for _, test := range tests {
		actual, actualErr := ValidateRootNode(test.Tree)
		if actual != test.expectedNode || errors.Is(actualErr, test.expectedErr) {
			if actualErr == nil && actualErr != test.expectedErr {
				log.Info(fmt.Sprintf("Test %v", test.ID))
				log.Info(fmt.Sprintf("Expected Node : %v, Got : %v ", test.expectedNode, actual))
				log.Info(fmt.Sprintf("Expected Error : %v, Got : %v ", test.expectedErr, actualErr))
				t.Fail()
			}
		}
	}
}
