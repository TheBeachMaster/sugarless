package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"com.thebeachmaster/sugarless/client"
	"github.com/echovault/sugardb/sugardb"
)

func main() {
	fmt.Println("Hello World")
	port := 16500
	bindAddr := "127.0.0.1"
	conndAddr := fmt.Sprintf("%s:%d", bindAddr, port)

	conf := sugardb.DefaultConfig()
	conf.ServerID = "ServerInstance1"
	conf.RestoreAOF = true
	conf.DataDir = "./disk"
	conf.BindAddr = bindAddr
	conf.Port = 16500

	server, err := sugardb.NewSugarDB(
		sugardb.WithConfig(conf),
	)

	if err != nil {
		log.Fatal(err)
	}

	// (Optional): Listen for TCP connections on this SugarDB instance.
	go func() {
		server.Start()
	}()

	time.Sleep(5 * time.Second)

	redisClient, err := client.NewCacheDBClient(conndAddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.GetServerInfo().Id

	fmt.Printf("Server ID ...%s\n", srv)

	if err := redisClient.Set(context.Background(), "key", "Hello, SugarDB!", 5*time.Minute).Err(); err != nil {
		log.Fatalf("set error: %s", err.Error())
	}

	v, _ := redisClient.Get(context.Background(), "key").Result()
	if err != nil {
		log.Fatalf("get error: %s", err.Error())
	}
	fmt.Println(v) // Hello, SugarDB!

	/*

		if err := data.CreateData(redisClient); err != nil {
			log.Fatal(err)
		}

		exts, err := data.FetchMarketplaceNativeModulesCache(redisClient)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found %d items\n", len(*exts))
	*/

	quitApplication := make(chan struct{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	close(quitApplication)

	<-quitApplication
	fmt.Println("Stopping...")
}
