# Kalshi Order Tools (kot)

**WIP - Experimental**

Kalshi order tools:  cli tools that use [`fsctl/go-kalshi`](https://github.com/fsctl/go-kalshi/) module to do various operations using Kalshi API

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

Now run the test programs:

```
./run.sh list-markets
```

```
./run.sh kot
```

## Notes

#### How do I update the dependency version here when fsctl/go-kalshi repo is updated?

1.  Check for a newer go-kalshi than what's in `go.mod`:
    `go list -m -u github.com/fsctl/go-kalshi`
If there's a newer one, you'll see brackets:
    `github.com/fsctl/go-kalshi v0.0.0-20220428182036-1a91897e9826 [v0.0.0-20220428223602-07b90d5f45c7]`

2.  Assuming a newer one exists, `go get` it @latest to _update the go.mod file_ and fetch it:
    `go get github.com/fsctl/go-kalshi@latest`

