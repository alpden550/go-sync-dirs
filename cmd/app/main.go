package main

import (
	"context"
	"flag"
	"go-sync-dirs/internal/filesync"
	"go-sync-dirs/internal/logging"
	"time"
)

func main() {
	initial := flag.String("initial", "", "Initial directory")
	target := flag.String("target", "", "Target directory to copy files from initial")
	interval := flag.Int("interval", 5, "Interval in minutes between checking initial directory")
	flag.Parse()

	logger := logging.NewLogger()
	ctx := context.WithValue(context.Background(), "logger", logger)

	for {
		filesync.SyncDirs(ctx, *initial, *target)
		time.Sleep(time.Duration(int64(*interval)) * time.Minute)
	}
}
