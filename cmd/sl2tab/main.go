package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/opensounder/slogo"
)

var (
	count int
	zoom  int
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	var err error
	flag.IntVar(&count, "count", 0, "number of frames to parse. 0=all")
	flag.IntVar(&zoom, "zoom", 14, "zoom level in url. 1-17")
	flag.Parse()

	if count < 0 || count > math.MaxInt32 {
		check(fmt.Errorf("invalid count flag value"))
	}

	if zoom < 1 || zoom > 17 {
		check(fmt.Errorf("invalid zoom flag value"))
	}

	path := flag.Arg(0)

	// log.Println("Will read", path)
	logfile, header, err := slogo.OpenLog(path)
	check(err)
	defer logfile.Close()
	fmt.Printf("header: %+v\n", header)

	d := slogo.NewDecoder(logfile, header)

	if count > 0 {
		fmt.Printf("Getting %d frames\n\n", count)
	} else {
		fmt.Printf("Getting all frames\n\n")
	}
	fc := 0
	skipped := 0
	last_point := slogo.Point{}

	w := tabwriter.NewWriter(os.Stdout, 5, 4, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	record := []string{"Offset", "Time", "Diff", "Skipped", "Kts", "Kph", "Feet", "Meter", "COG", "Latitude", "Longitude", "URL", ""}
	fmt.Fprintln(w, strings.Join(record, "\t"))
	var last_time uint32
	var f slogo.FrameF2
	for err == nil && (count == 0 || fc < count) {
		err = d.Next(&f)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error %v", err)
			break
		}
		p := f.Location()
		lo, la := p.GeoLatLon()
		if p != last_point { //only print if moved
			if last_time == 0 {
				last_time = f.Time1
			}
			_, err = fmt.Fprintf(w, "%d\t%d\t%d\t%d\t%.2f\t%.2f\t%.2f\t%.2f\t%.0f\t%f\t%f\t%s\t\n",
				f.Offset,
				f.Time1,
				f.Time1-last_time,
				skipped, //skipped number of frames
				f.GpsSpeed,
				f.GpsSpeed.ToKph(),
				f.WaterDepth,
				f.WaterDepth.ToMeters(),
				f.COG.ToDeg(),
				la, lo,
				p.ToGMapsURL(byte(zoom)),
			)
			check(err)
			fc++
			last_time = f.Time1
			last_point = p
			skipped = 0
		} else {
			skipped++
		}
	}

	w.Flush()

	log.Println("Done!")
}
