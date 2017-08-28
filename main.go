package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c2h5oh/datasize"
	"github.com/sirupsen/logrus"
	"github.com/vlad-s/hcpxread/helpers"
	"github.com/vlad-s/hcpxread/structs"
)

var (
	capture = flag.String("capture", "", "The HCCAPX `file` to read")
	debug   = flag.Bool("debug", false, "Show additional, debugging info")
)

const BANNER = ` _                                       _
| |__   ___ _ ____  ___ __ ___  __ _  __| |
| '_ \ / __| '_ \ \/ / '__/ _ \/ _` + "` |/ _`" + ` |
| | | | (__| |_) >  <| | |  __/ (_| | (_| |
|_| |_|\___| .__/_/\_\_|  \___|\__,_|\__,_|
           |_|
`

var (
	log       = helpers.Logger
	Instances structs.HccapxInstances
)

func init() {
	flag.Parse()
	log.SetLevel(logrus.DebugLevel)

	if *capture == "" {
		fmt.Println(BANNER)
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("debug pre-set", helpers.Debug())
	helpers.SetDebugging(*debug)
	fmt.Println("debug post-set", helpers.Debug())
}

func main() {
	if !helpers.Debug() {
		helpers.ClearScreen()
	}

	stat, err := os.Stat(*capture)
	if err != nil {
		log.WithError(err).Fatal("Error stating the file")
	}

	content, err := ioutil.ReadFile(*capture)
	if err != nil {
		log.WithError(err).Fatal("Error reading the file")
	}

	if len(content) < 393 {
		if helpers.Debug() {
			log.WithField("size", len(content)).Debug("File too small")
		}
		log.WithField("bytes", len(content)).Fatal("File too small for a single HCPX structure")
	}

	fileSize := datasize.ByteSize(stat.Size()).HumanReadable()
	log.WithFields(logrus.Fields{"name": stat.Name(), "size": fileSize}).Info("Opened file for reading")

	fileHeader := content[0:4]
	correctHeader := bytes.Equal(fileHeader, structs.HcpxHeader)

	if !correctHeader {
		log.WithField("header", string(fileHeader)).Fatal("Wrong file header")
	}

	log.Info("Searching for HCPX headers...")
	indexes := helpers.SearchHeaders(content)
	log.WithField("indexes", len(indexes)).Info("Finished searching for headers")

	for _, i := range indexes {
		j := i + 393
		h := helpers.ParseHccapx(content[i:j])
		Instances = append(Instances, h)
	}

	log.Infof("Summary: %d networks, %d WPA/%d WPA2, %d unique APs",
		len(Instances), Instances.WPANum(), Instances.WPA2Num(), Instances.UniqueAPs())

	var choice int
	for {
		Instances.Print()
		fmt.Printf("\nnetwork > ")

		_, err := fmt.Fscanf(os.Stdin, "%d", &choice)
		if err != nil {
			if !helpers.Debug() {
				helpers.ClearScreen()
			}
			log.Error(err)
			continue
		}

		if choice <= 0 {
			log.Info("Exiting, goodbye")
			os.Exit(0)
		}

		if choice > len(Instances) {
			if !helpers.Debug() {
				helpers.ClearScreen()
			}
			log.Warn("Invalid index")
			continue
		}

		if !helpers.Debug() {
			helpers.ClearScreen(true)
		}
		Instances[choice-1].Print()
	}
}
