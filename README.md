# go-kalshi Usage Example

![build workflow](https://github.com/fsctl/go-kalshi-usage-example/actions/workflows/makefile.yml/badge.svg)

This repo is a simple example of how you would use the [`fsctl/go-kalshi`](https://github.com/fsctl/go-kalshi/) module in a CLI tool.

## Build

```
make
```

## Run

Create a `.env` file modeled after `.env.example` in the repo root.  Then run the program:

```
./run.sh list-markets
```

## Notes

#### How do I update the dependency version here when fsctl/go-kalshi repo is updated?

Assuming a newer one exists, use `go get -u` to fetch `go-kalshi` and update the go.mod file:  `go get -u github.com/fsctl/go-kalshi@main`

#### Repo forks

If you fork this repo, you will need to create two Github Repository Secrets in your fork's repo to get CI to work. They must be called `KALSHI_USERNAME` and `KALSHI_PASSWORD` and contain Kalshi account credentials.  Used by `.github/workflows/makefile.yml`.