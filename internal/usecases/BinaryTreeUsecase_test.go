package usecases

import (
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"github.com/tryfix/log"
	"go.uber.org/mock/gomock"
	"testing"
)

var (
	testTree = models.Tree{
		Nodes: []models.Node{
			{
				ID:    "1",
				Left:  "2",
				Right: "3",
				Value: 1,
			},
			{
				ID:    "3",
				Left:  "6",
				Right: "7",
				Value: 3,
			},
			{
				ID:    "7",
				Left:  "null",
				Right: "null",
				Value: 7,
			},
			{
				ID:    "6",
				Left:  "null",
				Right: "null",
				Value: 6,
			},
			{
				ID:    "2",
				Left:  "4",
				Right: "5",
				Value: 2,
			},
			{
				ID:    "5",
				Left:  "null",
				Right: "null",
				Value: 5,
			},
			{
				ID:    "4",
				Left:  "null",
				Right: "null",
				Value: 4,
			},
		},
		Root: "1",
	}
)

func TestBinaryTreeUsecase_FindMaxPathSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	binaryTreeUsecase := NewBinaryTreeUsecase()

	type test struct {
		rootNode *models.Node
		Tree     *models.Tree
		expected int
	}

	tests := []test{
		{
			rootNode: &testTree.Nodes[0],
			Tree:     &testTree,
			expected: 18,
		},
	}

	for _, test := range tests {
		actual := binaryTreeUsecase.FindMaxPathSum(test.rootNode, test.Tree)
		if actual != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkFindMaxPathSum(b *testing.B) {

	binaryTreeUsecase := NewBinaryTreeUsecase()

	type test struct {
		rootNode *models.Node
		Tree     *models.Tree
		expected int
	}

	tests := []test{
		{
			rootNode: &testTree.Nodes[0],
			Tree:     &testTree,
			expected: 18,
		},
	}
	for i := 0; i < b.N; i++ {
		actual := binaryTreeUsecase.FindMaxPathSum(tests[0].rootNode, tests[0].Tree)
		if actual != tests[0].expected {
			return
		}
	}
}

func TestFindNodeByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		ID       string
		Tree     *models.Tree
		expected *models.Node
	}

	tests := []test{
		{
			ID:       "1",
			Tree:     &testTree,
			expected: &testTree.Nodes[0],
		},
		{
			ID:       "11",
			Tree:     &testTree,
			expected: nil,
		},
		{
			ID:       "null",
			Tree:     &testTree,
			expected: nil,
		},
		{
			ID:       "",
			Tree:     &testTree,
			expected: nil,
		},
	}

	for _, test := range tests {
		actual := FindNodeByID(test.ID, test.Tree)
		if actual != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkFindNodeByID(b *testing.B) {
	type test struct {
		ID       string
		Tree     *models.Tree
		expected *models.Node
	}

	tests := []test{
		{
			ID:       "1",
			Tree:     &testTree,
			expected: &testTree.Nodes[0],
		},
	}
	for i := 0; i < b.N; i++ {
		actual := FindNodeByID(tests[0].ID, tests[0].Tree)
		if actual != tests[0].expected {
			return
		}
	}
}
