package slogo

import (
	"fmt"
	"math"
	"os"
	"testing"
)

const float64Epsilon = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64Epsilon
}

func Test_slDecoder_Decode(t *testing.T) {
	type fields struct {
		r      *os.File
		header *Header
	}
	filename := "./testdata/sample-data-lowrance/Elite_4_Chirp/bigger.sl2"
	logfile, header, err := OpenLog(filename)

	if err != nil {
		t.Errorf("Could not open file %s", filename)
	}
	defer logfile.Close()

	if header.Format != 2 {
		t.Errorf("header.Format = %v, want %v", header.Format, 2)
	}
	if header.Blocksize != 1970 {
		t.Errorf("header.Blocksize = %v, wants %v", header.Blocksize, 1970)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *FrameF2
		wantLon float64
		wantLat float64
		wantCog float32
		wantAlt float32
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				header: &header,
				r:      logfile,
			},
			wantErr: false,
			want: &FrameF2{
				Offset:      8,
				Primary:     8,
				Blocksize:   2064,
				Packetsize:  1920,
				LowerLimit:  53.4,      //feet
				WaterDepth:  21.488001, //feet
				Temperature: 11.473407,
				LonEncoded:  1373465,
				LatEncoded:  8180800,
				COG:         3.2986722,
				Altitude:    244.71785,
				Flags:       542,
				Time:        28769,
			},
			wantLon: 12.379552312136807,
			wantLat: 59.12899916049587,
			wantCog: 1189.0,
			wantAlt: 74.590004,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &slDecoder{
				r:      tt.fields.r,
				header: tt.fields.header,
			}
			var got FrameF2
			err := d.Next(&got)
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
			lon := Longitude(got.LonEncoded)
			if lon != tt.wantLon {
				t.Errorf("Longitude() = %v, wants %v", lon, tt.wantLon)
			}
			lat := Latitude(got.LatEncoded)
			if lat != tt.wantLat {
				t.Errorf("Latitude() = %v, wants %v", lat, tt.wantLat)
			}
			//TODO: test cog and alt conversion
		})
	}
}

func TestLongitudeDD(t *testing.T) {
	type args struct {
		lon int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "a",
			args: args{1373465},
			want: 12.379552312136807,
		},
		{
			name: "b",
			args: args{0},
			want: 0,
		},
		{
			name: "c",
			args: args{1},
			want: 9.013372974292616e-06,
		},
		{
			name: "max",
			args: args{math.MaxInt32},
			want: 19356.071066605145,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Longitude(tt.args.lon); !almostEqual(got, tt.want) {
				t.Errorf("LongitudeDD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatitude(t *testing.T) {
	type args struct {
		lat int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			args: args{0},
			want: 0,
		},
		{
			args: args{1},
			want: 9.013372970355275e-06,
		},
		{
			args: args{8180800},
			want: 59.12899916049587,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Latitude(tt.args.lat); !almostEqual(got, tt.want) {
				t.Errorf("Latitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
