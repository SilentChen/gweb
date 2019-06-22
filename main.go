package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web/packs/gin"
	"web/packs/session"
	"web/packs/util"
)

var globalSessions *session.Manager

func init() {
	// log path setting
	webRootPath, err := os.Getwd()
	util.CheckErr(err)
	logPath := webRootPath + "/logs/access.log"
	logFile, _ := os.Create(logPath)

	// request log out put, file and terminate stdout
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	globalSessions, err = session.NewSessionManager("memory", "GOSESSIONID", 10)

	if err != nil {
		log.Println(err)
	}

	go globalSessions.GC()
}

func main() {
	if "dev" == util.Gapp_mode {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	r := LoadRouters()

	r.Delims("{{", "}}")

	r.SetFuncMap(template.FuncMap{
		"echo"		:	fmt.Sprintf,
		"date"		:	util.DateFormat,
		"str2html"	:	util.Str2html,
		"unix2time"	:	util.Unix2time,
		"unix2date"	:	util.Unix2date,
		"date2unix"	:	util.Date2unix,
	})

	r.LoadHTMLGlob("views/**/**/*")

	r.Static("/static", "static")

	port :=  util.Gapp_port

	s := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown  Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
