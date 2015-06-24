package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func printEnvs() error {
	environments, err := loadEnvMap()
	if err != nil {
		return err
	}

	for _, eModel := range environments {
		envString := "- " + eModel.Key + ": " + eModel.Value
		fmt.Println(envString)
		if eModel.IsExpand == false {
			expandString := "  " + "isExpand" + ": " + "false"
			fmt.Println(expandString)
		}
	}

	//log.Info("EnvStore:", environments)
	return nil
}

func printCmd(c *cli.Context) {
	log.Info("Work path:", currentEnvStoreFilePath)

	err := printEnvs()
	if err != nil {
		log.Fatal("Failed to print:", err)
	}
}