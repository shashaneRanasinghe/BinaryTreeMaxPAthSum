package models

type Tree struct {
	Nodes []Node `json:"nodes"`
	Root  string `json:"root"`
}

type Node struct {
	ID    string `json:"ID"`
	Left  string `json:"left"`
	Right string `json:"right"`
	Value int    `json:"value"`
}

type FindMaxPathSumRequest struct {
	Tree Tree `json:"tree"`
}

type BinaryTreeMaxPathSumResponse struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	MaxPathSum int `json:"maxPathSum"`
}
