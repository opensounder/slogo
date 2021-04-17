package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/opensounder/slogo"
)

var (
	count int
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type point struct {
	lon uint32
	lat uint32
}

func main() {
	var err error
	flag.IntVar(&count, "count", 0, "number of frames to parse. 0=all")
	flag.Parse()

	path := flag.Arg(0)

	// log.Println("Will read", path)
	logfile, header, err := slogo.OpenLog(path)
	check(err)
	defer logfile.Close()
	log.Printf("header: %+v\n", header)

	d := slogo.NewDecoder(logfile, header.Version, header.Blocksize)
	var f slogo.FrameV2
	check(err)
	fmt.Printf("Getting %d frames\n\n", count)
	fc := 0
	last_point := slogo.Point{}
	for err == nil && fc < count {
		err = d.DecodeV2(&f)
		check(err)
		p := f.Location()
		lo, la := p.GeoLatLon()
		if p != last_point {
			fmt.Printf("%d\t%f\t%f\t%d\t%d\t%f\t%f\t https://maps.google.com/maps?q=@%f,%f&z=14 \n",
				f.Time,
				f.GpsSpeed,
				slogo.KnotsToKph(f.GpsSpeed),
				f.LatEncoded, f.LonEncoded,
				la, lo,
				la, lo,
			)
			fc++
			last_point = p
		}
	}

	log.Println("Done!")

}
