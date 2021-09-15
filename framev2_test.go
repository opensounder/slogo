package slogo

import (
	"io"
	"testing"
)

func Test_FrameV2_Many(t *testing.T) {

	tests := []struct {
		filename string
		count    int
		wantErr  bool
	}{
		{"testdata/sample-data-lowrance/Elite_4_Chirp/small.sl2", 4017, false},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/version-1.sl2", 7, false},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/bigger.sl2", 5, false},
		{"testdata/sample-data-lowrance/Elite_4_Chirp/Chart 05_11_2018 [0].sl2", 5, false},
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
					t.Errorf("Next() error = %v", err)
					break
				}
				count += 1
				if f.Offset < offset {
					t.Errorf("offset got %v, want > %v", f.Offset, offset)
					break
				}
				if f.Frameindex < index {
					t.Errorf("index got %v, want >= %v", f.Frameindex, index)
					break
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
