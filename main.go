package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/sslime336/bot-486/bot"
	"github.com/sslime336/bot-486/config"
	"github.com/sslime336/bot-486/db"
	"github.com/sslime336/bot-486/handler"
	"github.com/sslime336/bot-486/logging"
	"github.com/sslime336/bot-486/service"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"go.uber.org/zap"
)

func init() {
	config.Init()
	logging.Init()
	db.Init()
	handler.Init()
	service.Init()
}

func main() {
	conf := config.App.Bot
	token := token.BotToken(conf.AppId, conf.Token)

	var api openapi.OpenAPI
	if os.Getenv("SUBARU_MODE") == "release" {
		api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	} else {
		api = botgo.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	}

	bot.BuildClient(api)

	botgo.SetLogger(logging.Logger().Sugar())

	ws, err := api.WS(context.Background(), nil, "")
	if err != nil {
		logging.Fatalf("%+v, error:%v", ws, err)
	}

	intent := websocket.RegisterHandlers(handler.Get())
	intent |= 1 << 25 // 注册群和私聊消息

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		logging.Info("bot was exited by Ctrl C", zap.Any("signal", <-sig))
		os.Exit(0)
	}()

	if err := botgo.NewSessionManager().Start(ws, token, &intent); err != nil {
		logging.Fatal("unexpected exit", zap.Error(err))
	}
}
