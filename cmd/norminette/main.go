package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/remicaumette/norminette/pkg/norminette"
	flag "github.com/spf13/pflag"
)

func recursiveAdd(file string, files *[]string) {
	stat, err := os.Stat(file)
	if err != nil {
		return
	}
	if stat.IsDir() {
		dir, err := ioutil.ReadDir(file)
		if err != nil {
			return
		}
		for _, dfile := range dir {
			recursiveAdd(path.Join(file, dfile.Name()), files)
		}
	} else if filepath.Ext(file) == ".c" || filepath.Ext(file) == ".h" {
		*files = append(*files, file)
	}
}

func main() {
	log.SetFlags(0)

	versionFlag := flag.BoolP("version", "v", false, "get the norminette version")
	credentialsFlag := flag.String("credentials", "amqp://guest:guest@norminette.le-101.fr/", "change credentials")
	disabledRulesFlag := flag.StringArrayP("rules", "R", []string{}, "Rule to disable")
	flag.Parse()

	norm, err := norminette.New(*credentialsFlag, *disabledRulesFlag)
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
		args := flag.Args()
		if len(args) == 0 {
			args = append(args, ".")
		}
		for _, file := range args {
			recursiveAdd(file, &files)
		}
		if len(files) == 0 {
			return
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
