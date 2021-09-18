package slogo

import (
	"io"
	"testing"
)

func Test_FrameF3_First(t *testing.T) {
	stream, decoder, err := OpenDecoder("testdata/sample-data-lowrance/other/sonar-log-api-testdata.sl3")
	if err != nil {
		t.Errorf("OpenDecoder error = %v", err)
		return
	}
	defer stream.Close()
	f := FrameF3{}
	err = decoder.Next(&f)
	if err != nil {
		t.Errorf("Next() = %v", err)
		return
	}

	if f.Offset != 8 {
		t.Errorf("Offset %v, want 8", f.Offset)
	}
	if f.Framesize != 2732 {
		t.Errorf("Blocksize = %v, want 2732", f.Framesize)
	}

	if f.Channel != 9 {
		t.Errorf("Channel =%v, want 9", f.Channel)
	}
	if !almostEqual32(f.UpperLimit, 0, 1e-2) {
		t.Errorf("LowerLimit = %v, want %v", f.UpperLimit, 98.4)
	}
	if !almostEqual32(f.LowerLimit, 98.4, 1e-2) {
		t.Errorf("LowerLimit = %v, want %v", f.UpperLimit, 98.4)
	}
	if f.TimeStamp != 1477636650 {
		t.Errorf("TimeStamp = %v, want 1477636650", f.TimeStamp)
	}
	if f.Payloadsize != 3072 {
		t.Errorf("PacketSize = %v, want 3072", f.Payloadsize)
	}
	if f.WaterDepth != 38.501 {
		t.Errorf("WaterDepth = %v, want 38.5", f.WaterDepth)
	}
	if f.WaterDepth.ToMeters() != 11.735105 {
		t.Errorf("WaterDepth.ToMeters() = %v, want 11.74", f.WaterDepth.ToMeters())
	}
	if f.Frequency != 0 {
		t.Errorf("Frequency = %v, want 0", f.Frequency)
	}
	if f.GpsSpeed != 2.2548594 {
		t.Errorf("GPS Speed = %v, want 2.2548594", f.GpsSpeed)
	}
	if f.Temperature != 2.3099976 {
		t.Errorf("Temperature = %v, want 2.3099976", f.Temperature)
	}
	if f.XMerc != 4048018 {
		t.Errorf("XMerc = %v, want 4048018", f.XMerc)
	}
	if f.YMerc != 7652086 {
		t.Errorf("YMerc = %v, want 7652086", f.YMerc)
	}
	if f.WaterSpeed != 0 {
		t.Errorf("WaterSpeed = %v, want 0", f.WaterSpeed)
	}
	if f.COG != 1.186824 {
		t.Errorf("Course = %v, want 1.186824", f.COG)
	}
	if f.Altitude != 394.35693 {
		t.Errorf("Altitude = %v, want 394.3", f.Altitude)
	}
	if f.Heading != 1.3785058 {
		t.Errorf("Heading = %v, want 1.38", f.Heading)
	}
	if f.Flags != 950 {
		t.Errorf("Flags = %v, want 950", f.Flags)
	}
	if f.Time1 != 269151 {
		t.Errorf("Time1 = %v, want 269151", f.Time1)
	}

	// fmt.Printf("%+v\n", f)
}

func Test_FrameF3_Many(t *testing.T) {

	tests := []struct {
		filename       string
		count          int
		center         Point
		distance       float64
		allowIndexDiff int
		wantErr        bool
	}{
		{"testdata/sample-data-lowrance/other/sonar-log-api-testdata.sl3", 10124, Point{4048018, 7652086}, 500, 0, false},
		// {"testdata/sample-data-lowrance/other/format3_version2.sl3", 4017, Point{4048018, 7652086}, 800, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			stream, decoder, err := OpenDecoder(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDecoder error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer stream.Close()
			f := FrameF3{}
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
				if tt.allowIndexDiff > 0 {
					diff := int(f.Frameindex) - int(index)
					if diff > tt.allowIndexDiff {
						t.Errorf("[%v] index got %v, want ~= %v which was last, %v", count, f.Frameindex, index, diff)
						break
					}
				}
				loc := f.Location()
				dist := pointDistance(loc, tt.center)
				if dist > tt.distance {
					t.Errorf("[%v] loc %v to far away (>%v m) from %v. got %v.", count, loc, tt.distance, tt.center, dist)
					break
				}
				offset = f.Offset
				index = f.Frameindex
			}
			if count != tt.count {
				t.Errorf("count got, %v wants %v", count, tt.count)
			}
			// fmt.Printf("F3: %+v", f)

		})
	}
}
