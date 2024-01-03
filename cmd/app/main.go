package main

import (
	"github.com/shashaneRanasinghe/BinaryTreeMaxPathSum/pkg/server"
	"github.com/tryfix/log"
)

func main() {
	closeChannel := server.Serve()
	<-closeChannel

	log.Info("Service Stopped")
}
