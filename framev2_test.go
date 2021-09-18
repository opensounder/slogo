package slogo

import (
	"io"
	"testing"
)

func Test_FrameV2_Many(t *testing.T) {

	tests := []struct {
		filename  string
		count     int
		center    Point
		distance  float64
		indexDiff int
		wantErr   bool
	}{
		{"testdata/sample-data-lowrance/Elite_4_Chirp/small.sl2", 4017, Point{8147302, 1383678}, 800, 1, false},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/version-1.sl2", 7, Point{8179735, 1372428}, 800, 1, false},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/bigger.sl2", 16885, Point{8180800, 1373465}, 800, 50, false},
		// {"testdata/sample-data-lowrance/Elite_4_Chirp/Chart 05_11_2018 [0].sl2", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			stream, decoder, err := OpenDecoder(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDecoder error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer stream.Close()
			f := FrameF2{}
			var offset uint32 = 0
			var index uint32 = 0
			count := 0

			for err == nil {
				err = decoder.Next(&f)
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Errorf("[%v] Next() error = %v", count, err)
					break
				}
				count += 1

				if f.Offset < offset {
					t.Errorf("[%v] offset got %v, want > %v", count, f.Offset, offset)
					break
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

		})
	}
}
