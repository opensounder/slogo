package slogo

import (
	"encoding/binary"
	"log"
	"math"
	"os"
)

const float64Epsilon = 1e-9

func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func almostEqual32(a, b, epsilon float32) bool {
	return almostEqual(float64(a), float64(b), float64(epsilon))
}

func pointLatLng(lat float64, lng float64) Point {
	return Point{
		YMerc: merc_y(lat),
		XMerc: merc_x(lng),
	}
}

func logOffset(file *os.File) {
	offset, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current Offset = %v", offset)
}

func logSize(t interface{}, msg string) {
	size := int64(binary.Size(t))
	log.Printf("%s binary.Size()=%v\n", msg, size)
}

func logFloatAtOffset(file *os.File, offset int) {
	var data float32
	file.Seek(int64(offset), os.SEEK_SET)
	log.Printf("Float-")
	err := binary.Read(file, binary.LittleEndian, &data)
	if err != nil {
		log.Println("!!ERROR logAtOffset", err)
		return
	}
	log.Printf("Offset: %d (% x) = %f ", offset, offset, data)
}

func logLongAtOffset(file *os.File, offset int) {
	var data uint32
	file.Seek(int64(offset), os.SEEK_SET)
	log.Printf("Long-")
	err := binary.Read(file, binary.LittleEndian, &data)
	if err != nil {
		log.Println("!!ERROR logAtOffset", err)
		return
	}
	log.Printf("Offset: %d (% x) = %d (% x)", offset, offset, data, data)
}

func logShortAtOffset(file *os.File, offset int) {
	var data uint16
	file.Seek(int64(offset), os.SEEK_SET)
	log.Printf("Short-")
	err := binary.Read(file, binary.LittleEndian, &data)
	if err != nil {
		log.Println("!!ERROR logAtOffset", err)
		return
	}
	log.Printf("Offset: %d (% x) = %d (% x)", offset, offset, data, data)
}
