package handlers

import (
	"encoding/json"
	"fmt"
	helpers "github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/helpers/BinaryTree"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/models"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/internal/usecases"
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/pkg/consts"
	"github.com/tryfix/log"
	"io"
	"net/http"
)

type BinaryTreeHandler struct {
	binaryTree usecases.BinaryTree
}

func NewBinaryTreeHandler() *BinaryTreeHandler {
	binaryTreeUsecase := usecases.NewBinaryTreeUsecase()
	return &BinaryTreeHandler{
		binaryTree: binaryTreeUsecase,
	}
}
func (handler *BinaryTreeHandler) FindMaxPathSum(w http.ResponseWriter, r *http.Request) {

	var respModel models.BinaryTreeMaxPathSumResponse
	var reqBody models.FindMaxPathSumRequest

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetBinaryTreeError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(consts.RequestBodyCloseError, err)
		}
	}(r.Body)

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetBinaryTreeError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}

	//ValidateRootNode validates if the given root node is not empty and exists in the node list
	rootNode, err := helpers.ValidateRootNode(reqBody.Tree)
	if err != nil {
		log.Error(consts.ValidationError, err)

		respModel.Status = consts.Error
		respModel.Message = fmt.Sprintf("%s: %s", consts.ValidationError, err)

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	maxPathSum := handler.binaryTree.FindMaxPathSum(&rootNode, &reqBody.Tree)
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data.MaxPathSum = maxPathSum
	respModel.Message = consts.MaxPathSumCalculated

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}
