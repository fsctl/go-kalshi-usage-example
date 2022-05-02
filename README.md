# go-kalshi Usage Example

![build workflow](https://github.com/fsctl/go-kalshi-usage-example/actions/workflows/makefile.yml/badge.svg)

Example cli tool that uses [`fsctl/go-kalshi`](https://github.com/fsctl/go-kalshi/) module to list markets using Kalshi API

## Build

```
make
```

## Run

Create a `.env` file like this in the repo root:

```
KALSHI_USERNAME="your kalshi login (email)"
KALSHI_PASSWORD="your kalshi password"
```

Now run the test program:

```
./run.sh list-markets
```

## Notes

#### How do I update the dependency version here when fsctl/go-kalshi repo is updated?

Assuming a newer one exists, `go get -u` it to _update the go.mod file_ and fetch it:
    `go get -u github.com/fsctl/go-kalshi@main`

