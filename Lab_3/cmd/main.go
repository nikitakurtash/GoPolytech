package main

import (
	"flag"
	"fmt"
	"github.com/nikitakurtash/matchLib/pkg/calculations"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	var logFlag = flag.Bool("log", false, "Enable detailed logging with logrus")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: mymodule <number> [-log]")
		os.Exit(1)
	}

	n, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil && *logFlag {
		logrus.WithError(err).Fatal("Error converting argument to int64")
	} else if err != nil {
		os.Exit(1)
	}

	if *logFlag {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
		})
		logrus.SetLevel(logrus.InfoLevel)

		result := calculations.Calculate(n, true)
		logrus.WithFields(logrus.Fields{
			"result": result,
		}).Info("Calculation result")
	} else {
		result := calculations.Calculate(n, false)
		fmt.Println(result)
	}
}
