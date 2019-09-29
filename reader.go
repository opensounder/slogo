package slogo

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//ReadLogfile does just that
func ReadLogfile(path string) error {
	fmt.Println("Will read", path)
	logfile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer logfile.Close()

	header := Header{}
	err = binary.Read(logfile, binary.LittleEndian, &header.Format)
	check(err)
	err = binary.Read(logfile, binary.LittleEndian, &header.Version)
	check(err)
	err = binary.Read(logfile, binary.LittleEndian, &header.Blocksize)
	check(err)
	err = binary.Read(logfile, binary.LittleEndian, &header.Reserved1)
	check(err)

	log.Println("header", header)
	return nil
}
