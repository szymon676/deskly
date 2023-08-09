package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/szymon676/betterdocker/mysql"
)

func init() {
	opts := &mysql.MySQLContainerOptions{
		RootPassword: "1234",
		Database:     "main",
	}
	container := mysql.NewMySQLContainer(opts)
	container.Run()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
	container.Stop()
}

func main() {
	storage, err := NewStorage("root:1234@tcp(127.0.0.1:3306)/main")
	if err != nil {
		panic(err)
	}

	as := &apiserver{storage: storage}
	s := as.SetupServer()

	go s.Listen(":4000")
	storage.WatchBookings()
}
