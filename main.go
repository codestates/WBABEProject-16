package main

import (
	adminCtl "codestates_lecture/WBABEProject-16/admincontroller"
	conf "codestates_lecture/WBABEProject-16/config"
	"codestates_lecture/WBABEProject-16/logger"
	"codestates_lecture/WBABEProject-16/model"
	orderCtl "codestates_lecture/WBABEProject-16/ordercontroller"
	rt "codestates_lecture/WBABEProject-16/router"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)


var (
	g errgroup.Group
)

func main() {
	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)
	
	// 로그 초기화
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")

	if mod, err := model.NewModel(); err != nil {
		errors.New("model Error")
	} else if  admincontroller, err := adminCtl.NewController(mod); err != nil {
		errors.New("controller Error")
	}else if  ordercontroller, err := orderCtl.NewController(mod); err != nil {
		errors.New("controller Error")
	}else if rt, err := rt.NewRouter(admincontroller,ordercontroller); err != nil {
		errors.New("router error")
	}else {
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}

		logger.Info("Server exiting")
	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}

	