# Synced

A simple program that reflects the syncing status of [vechain thor node](https://github.com/vechain/thor). Return `HTTP 200` when it's synced `HTTP 503` when it's syncing.

## Get Started

### Build from source

Clone the source:

```shell
git clone https://github.com/libotony/synced.git
cd synced
```

Build:

```
make synced
```

Start:
```shell
bin/synced --rest http://127.0.0.1:8669
```

### Docker

TBD

## Run

``` shell
bin/synced -h
NAME:
   synced - tells if thor is synced

USAGE:
   synced [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --thor-rest value, --rest value   Thor node API address (default: "http://127.0.0.1:8669")
   --port value, -p value            Synced API service listening port (default: 8000)
   --tolerable-diff value, -t value  tolerable left behind block amount (default: 5)
   --help, -h    
```