package main

import (
	"context"

	"github.com/kunlun-qilian/confx"
	"github.com/kunlun-qilian/gin-demo/cmd/demo-docker/apis"
	"github.com/kunlun-qilian/gin-demo/cmd/demo-docker/global"
	"github.com/spf13/cobra"
)

func main() {
	confx.Execute(func(cmd *cobra.Command, args []string) {
		apis.NewRooter(global.Config.Server.Engine())
		if err := global.Config.Server.Serve(context.Background()); err != nil {
			panic(err)
		}
	})
}
