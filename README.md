# Go Sync

Simple sync files realization with goroutines between two directories.

## How to prepare

To create exec file, type make command

```bash
make build
```

## How start

To see available command arguments:

```bash
./gosync -help
```

```bash
Usage of ./gosync:
  -initial string
        Initial directory
  -interval int
        Interval in minutes between checking initial directory (default 5)
  -target string
        Target directory to copy files from initial

```

To start sync initial folder:

```bash
make build
```
and
```bash
./gosync -initial=/Users/username/Desktop/first-directory -target=/Users/username/Desktop/target -interval=10
```

or

```bash
go run cmd/app/main.go -initial=/Users/username/Desktop/first-directory -target=/Users/username/Desktop/target -interval=10
```