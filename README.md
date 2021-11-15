# mdcat

A simple markdown viewer.

## Building

For Production,

```
go build -ldflags="-X 'main.Version=$(git describe --tags $(git rev-list --tags --max-count=1) || echo "dev")' -X 'main.BuildDate=$(date -u --rfc-2822)' -X 'main.CommitHash=$(git rev-parse HEAD)'"
```

## Usage

```
mdcat [OPTIONS] [FILE]
```

## License

Copyright (c) 2021, (@mrinjamul)[https://github.com/mrinjamul]
