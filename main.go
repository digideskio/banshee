// Copyright 2015 Eleme Inc. All rights reserved.

package main

import (
	"flag"
	"runtime"

	"github.com/eleme/banshee/config"
	"github.com/eleme/banshee/detector"
	"github.com/eleme/banshee/storage"
	"github.com/eleme/banshee/util/log"
)

func main() {
	// Arguments
	fileName := flag.String("c", "config.json", "config file")
	debug := flag.Bool("d", false, "debug mode")
	flag.Parse()
	// Logging
	if *debug {
		log.SetLevel(log.DEBUG)
	}
	log.Debug("using %s, max cpus: %d", runtime.Version(), runtime.GOMAXPROCS(-1))
	// Config
	cfg := config.New()
	if flag.NFlag() == 1 && *debug == false {
		err := cfg.UpdateWithJSONFile(*fileName)
		if err != nil {
			log.Fatal("failed to load %s: %s", *fileName, err)
		}
	} else {
		log.Warn("no config file specified, using default..")
	}
	// Storage
	options := &storage.Options{
		NumGrid: cfg.Period[0],
		GridLen: cfg.Period[1],
	}
	db, err := storage.Open(cfg.Storage.Path, options)
	if err != nil {
		log.Fatal("failed to open %s: %v", cfg.Storage.Path, err)
	}
	// Detector
	detector := detector.New(cfg, db)
	detector.Start()
}
