package main

import (
	"ag5/cmd/dependancy_injection"
	"ag5/internal/presentation/grpc"
	"ag5/internal/presentation/http"
	_ "ag5/pkg"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start the application
	dependancy_injection.DB()
	http.StartHttp()
	grpc.Start()

	// wait for `Ctrl+c` or docker stop/restart signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch

	// Stop the application
	grpc.Stop()
}
