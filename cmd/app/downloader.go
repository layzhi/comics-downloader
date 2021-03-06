package app

import (
	"strings"

	"github.com/Girbons/comics-downloader/pkg/config"
	"github.com/Girbons/comics-downloader/pkg/detector"
	"github.com/Girbons/comics-downloader/pkg/loader"
	log "github.com/sirupsen/logrus"
)

func init() {
	// use log INFO Level
	log.SetLevel(log.InfoLevel)
}

func Run(link, format, country string) {
	conf := new(config.ComicConfig)
	if err := conf.LoadConfig(); err != nil {
		log.Warning(err)
	}

	// link is required
	if link == "" {
		log.Fatal("url parameter is required")
	}

	if !strings.HasSuffix(link, ",") {
		link = link + ","
	}

	for _, u := range strings.Split(link, ",") {
		if u != "" {
			// check if the link is supported
			source, check := detector.DetectComic(u)
			if !check {
				log.WithFields(log.Fields{"site": u}).Error("This site is not supported :(")
				continue
			}

			log.WithFields(log.Fields{
				"url": u,
			}).Info("Downloading...")
			// in case the link is supported
			// setup the right strategy to parse a comic
			comic, err := loader.LoadComicFromSource(conf, source, u, country)
			if err != nil {
				log.Error(err)
				continue
			}
			comic.SetFormat(format)
			comic.MakeComic()
		}
	}
}
