package main

import (
	"flag"
	"fmt"
	"github.com/remicaumette/norminette/pkg/norminette"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	log.SetFlags(0)

	versionFlag := flag.Bool("version", false, "get the norminette version")
	credentialsFlag := flag.String("credentials", "amqp://guest:guest@norminette.le-101.fr/", "change credentials")
	flag.Parse()

	norm, err := norminette.New(*credentialsFlag)
	if err != nil {
		log.Fatalf("unable to connect to rabbitmq: %v\n", err)
	}
	defer norm.Connection.Close()

	if *versionFlag {
		version, err := norm.Version()
		if err != nil {
			log.Fatalf("unable to get norminette version: %v\n", err)
		}
		log.Printf("version: %v\n", version.Version)
	} else {
		files := make([]string, 0)

		for _, file := range flag.Args() {
			stat, err := os.Stat(file)
			if err != nil {
				log.Printf("[ERR]  %v: %v\n", file, err)
				continue
			}
			if stat.IsDir() {
				dir, err := ioutil.ReadDir(file)
				if err != nil {
					log.Printf("[ERR]  %v: %v\n", file, err)
					continue
				}
				for _, dfile := range dir {
					files = append(files, path.Join(file, dfile.Name()))
				}
			} else {
				files = append(files, file)
			}
		}

		response, err := norm.CheckFiles(files...)
		if err != nil {
			log.Fatalf("unable to use norminette: %v\n", err)
		}
		for _, file := range response {
			if file.Display == "" || strings.Contains(file.Display, "Norminette can't check this file.") {
				fmt.Printf("[OK]   %v\n", file.Filename)
			} else {
				fmt.Printf("[FAIL] %v\n%v\n", file.Filename, file.Display)
			}
		}
	}
}
