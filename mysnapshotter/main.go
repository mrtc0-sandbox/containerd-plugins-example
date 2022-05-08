package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	snapshotsapi "github.com/containerd/containerd/api/services/snapshots/v1"
	"github.com/containerd/containerd/contrib/snapshotservice"
	"github.com/containerd/containerd/snapshots/native"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("invalid args: usage: %s <unix addr> <root>\n", os.Args[0])
		os.Exit(1)
	}

	go func() {
		rpc := grpc.NewServer()

		sn, err := native.NewSnapshotter(os.Args[2])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		service := snapshotservice.FromSnapshotter(sn)

		snapshotsapi.RegisterSnapshotsServer(rpc, service)

		l, err := net.Listen("unix", os.Args[1])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		defer l.Close()

		if err := rpc.Serve(l); err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)

  signal.Notify(quit, os.Interrupt)
  <-quit

	os.Remove(os.Args[1])
}
