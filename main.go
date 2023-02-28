package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

type BlockHeader struct {
	Timestamp uint64
}

const blockInterval = 10 // block interval in seconds

var lastFetched uint64

func main() {
	app := &cli.App{
		Name:  "synced",
		Usage: "tells if thor is synced",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "thor-rest",
				Value:   "http://127.0.0.1:8669",
				Usage:   "Thor node API address",
				Aliases: []string{"rest"},
			},
			&cli.IntFlag{
				Name:    "port",
				Value:   8000,
				Usage:   "Synced API service listening port",
				Aliases: []string{"p"},
			},
			&cli.IntFlag{
				Name:    "tolerable-diff",
				Value:   5,
				Usage:   "tolerable left behind block amount",
				Aliases: []string{"t"},
			},
		},
		Action: func(ctx *cli.Context) error {
			url := ctx.String("thor-rest") + "/blocks/best"
			listen := ":" + strconv.Itoa(ctx.Int("port"))
			diff := ctx.Int("tolerable-diff")

			router := mux.NewRouter()
			router.HandleFunc("/synced", func(w http.ResponseWriter, req *http.Request) {
				if err := (func() error {
					now := time.Now().Unix()
					current := uint64(now - now%blockInterval)

					if lastFetched >= current && lastFetched <= current+blockInterval {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("ok"))
					} else {
						resp, err := http.Get(url)
						if err != nil {
							return err
						}
						defer resp.Body.Close()

						body, err := io.ReadAll(resp.Body)
						if err != nil {
							return err
						}

						var b BlockHeader
						if err := json.Unmarshal(body, &b); err != nil {
							return err
						}

						if b.Timestamp+uint64(diff)*blockInterval >= current {
							w.WriteHeader(http.StatusOK)
							w.Write([]byte("ok"))
							lastFetched = b.Timestamp
						} else {
							w.WriteHeader(http.StatusServiceUnavailable)
							fmt.Fprintf(w, "syncing(%d blocks)", (uint64(now)-b.Timestamp)/blockInterval)
						}
					}

					return nil
				})(); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			})

			listener, err := net.Listen("tcp", listen)
			if err != nil {
				log.Fatal(errors.Wrap(err, "listening tcp"))
			}
			log.Printf("Server listening at http://%s", listener.Addr().String())

			return http.Serve(listener, router)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
