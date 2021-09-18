package slogo

import (
	"fmt"
	"io"
	"testing"
)

func Test_FrameF2_First(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		want     *FrameF2
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			"small", "./testdata/sample-data-lowrance/Elite_4_Chirp/small.sl2",
			&FrameF2{
				Offset:          8,
				PreviousPrimary: 8,
				Framesize:       3216,
				Packetsize:      3072,
				LowerLimit:      19.6, //feet
				Frequency:       8,
				WaterDepth:      6.622, //feet
				GpsSpeed:        2.5853128,
				Temperature:     19.350006,
				XMerc:           1383678,
				YMerc:           8147302,
				COG:             3.7873645,
				Altitude:        118.89766,
				Flags:           702,
				Time1:           5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream, d, err := OpenDecoder(tt.filename)
			if err != nil {
				t.Errorf("error %w opening file", err)
				return
			}
			defer stream.Close()

			var got FrameF2
			err = d.Next(&got)
			if (err != nil) != tt.wantErr {
				t.Errorf("slDecoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// logOffset(logfile)
			// logLongAtOffset(logfile, 8+140)
			// log.Printf("COG in degrees %f", RadToDeg(got.COG))
			// log.Printf("Altitude in meters %f", FeetToMeter(got.Altitude))
			gots := fmt.Sprintf("%+v", &got)
			wants := fmt.Sprintf("%+v", tt.want)
			if gots != wants {
				t.Errorf("slDecoder.Decode() =\n %+v,\n want\n %+v", gots, wants)
			}

		})
	}
}

func Test_FrameF2_Many(t *testing.T) {

	tests := []struct {
		filename  string
		count     int
		center    Point
		distance  float64
		indexDiff int
		wantErr   int
	}{
		{"testdata/sample-data-lowrance/Elite_4_Chirp/small.sl2", 4017, Point{8147302, 1383678}, 800, 1, -1},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/version-1.sl2", 7, Point{8179735, 1372428}, 800, 1, 7},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/bigger.sl2", 16885, Point{8180800, 1373465}, 800, 50, -1},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/Chart 05_11_2018 [0].sl2", 27458, Point{8163659, 1373761}, 800, 50, -1},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			stream, decoder, err := OpenDecoder(tt.filename)
			if err != nil {
				t.Errorf("OpenDecoder error = %v", err)
				return
			}
			defer stream.Close()
			stat, _ := stream.Stat()
			size := stat.Size()
			f := FrameF2{}
			var offset uint32 = 0
			var index uint32 = 0
			count := 0

			for err == nil {
				here, _ := stream.Seek(0, io.SeekCurrent)
				err = decoder.Next(&f)
				if err == io.EOF {
					break
				}
				if err != nil {
					if tt.wantErr > -1 && count == tt.wantErr {
						break
					} else {
						t.Errorf("[%v] Next() error = %v", count, err)
						if here != size {
							t.Errorf("here %v is not the expected %v", here, size)
						}
					}
					break
				}
				count += 1

				if f.Offset < offset {
					t.Errorf("[%v] offset got %v, want > %v", count, f.Offset, offset)
					break
				}
				if (f.Offset - offset) > 3500 {
					t.Errorf("[%v] offset got %v, expected < %v", count, f.Offset, offset+3500)
				}
				diff := int(f.Frameindex) - int(index)
				if diff > tt.indexDiff {
					t.Errorf("[%v] index got %v, want ~= %v which was last, %v", count, f.Frameindex, index, diff)
					// break
				}
				loc := f.Location()
				dist := pointDistance(loc, tt.center)
				if dist > tt.distance {
					t.Errorf("[%v] loc %v to far away from %v. got %v, max %v", count, loc, tt.center, dist, tt.distance)
				}

				offset = f.Offset
				index = f.Frameindex
			}
			if count != tt.count {
				t.Errorf("count got, %v wants %v", count, tt.count)
			}
			here, _ := stream.Seek(0, io.SeekCurrent)
			if here != size {
				t.Errorf("bad end position %v. wants %v", here, size)
			}

		})
	}
}
