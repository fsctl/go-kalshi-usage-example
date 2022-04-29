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

Assuming a newer one exists, `go get -u` it to _update the go.mod file_ and fetch it:
    `go get -u github.com/fsctl/go-kalshi@main`

