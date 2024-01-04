package handlers

import (
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/mocks"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	binaryTreeRequest = models.FindMaxPathSumRequest{
		Tree: models.Tree{
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
			},
			Root: "1",
		}}

	errorNode1 = models.Node{}
)

func TestRoutes_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mux.NewRouter()

	binaryTreeUsecase := mocks.NewMockBinaryTree(ctrl)
	binaryTreeUsecase.EXPECT().FindMaxPathSum(&binaryTreeRequest.Tree.Nodes[0], &binaryTreeRequest.Tree).
		Return(18)

	binaryTreeHandler := BinaryTreeHandler{binaryTree: binaryTreeUsecase}
	r.HandleFunc("/binaryTree/findMaxPathSum", binaryTreeHandler.FindMaxPathSum).
		Methods("POST")

	testCases := []struct {
		name           string
		url            string
		method         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get Max Path Sum",
			url:            "/binaryTree/findMaxPathSum",
			method:         "POST",
			requestBody:    `{ "tree": { "nodes": [ { "id": "1", "left": "2", "right": "3", "value": 1 }, { "id": "3", "left": "6", "right": "7", "value": 3 } ], "root": "1" } }`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"maxPathSum":18},"message":"Max Path Sum Calculated Successfully"}`,
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest(test.method, test.url, strings.NewReader(test.requestBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != test.expectedStatus {
			t.Errorf("Test %s : Expected status code %d, but got %d", test.name,
				test.expectedStatus, w.Code)
		}

		if w.Body.String() != test.expectedBody {
			t.Errorf("Test %s : Expected response body %s, but got %s", test.name,
				test.expectedBody, w.Body.String())
		}
	}
}

func TestRoutes_ErrorPath1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mux.NewRouter()

	binaryTreeUsecase := mocks.NewMockBinaryTree(ctrl)

	binaryTreeHandler := BinaryTreeHandler{binaryTree: binaryTreeUsecase}
	r.HandleFunc("/binaryTree/findMaxPathSum", binaryTreeHandler.FindMaxPathSum).
		Methods("POST")

	testCases := []struct {
		name           string
		url            string
		method         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Root Node Not Provided",
			url:            "/binaryTree/findMaxPathSum",
			method:         "POST",
			requestBody:    `{ "tree": { "nodes": [ { "id": "1", "left": "2", "right": "3", "value": 1 }, { "id": "3", "left": "6", "right": "7", "value": 3 } ]} }`,
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":{"maxPathSum":0},"message":"Validation Error: Root Node not Provided"}`,
		},
		{
			name:           "Unknown Root Node ID",
			url:            "/binaryTree/findMaxPathSum",
			method:         "POST",
			requestBody:    `{ "tree": { "nodes": [ { "id": "1", "left": "2", "right": "3", "value": 1 }, { "id": "3", "left": "6", "right": "7", "value": 3 } ], "root": "11"} }`,
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":{"maxPathSum":0},"message":"Validation Error: Provided Root Node ID does not match with the provided nodes"}`,
		},
		{
			name:           "Invalid JSON",
			url:            "/binaryTree/findMaxPathSum",
			method:         "POST",
			requestBody:    `{ "tree": { "nodes": [ { "id": "1", "left": "2", "right": "3", "value": 1 }, { "id": "3", "left": "6", "right": "7", "value": 3 } ],  }`,
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":{"maxPathSum":0},"message":"Error While Getting the Binary Tree Input"}`,
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest(test.method, test.url, strings.NewReader(test.requestBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != test.expectedStatus {
			t.Errorf("Test %s : Expected status code %d, but got %d", test.name,
				test.expectedStatus, w.Code)
		}

		if w.Body.String() != test.expectedBody {
			t.Errorf("Test %s : Expected response body %s, but got %s", test.name,
				test.expectedBody, w.Body.String())
		}
	}
}
