package slogo

import (
	"os"
	"reflect"
	"testing"
)

func Test_slDecoder_Decode(t *testing.T) {
	type fields struct {
		r         *os.File
		version   uint16
		blocksize uint16
	}
	filename := "./test-fixtures/sample-data-lowrance/Elite_4_Chirp/bigger.sl2"
	logfile, header, err := OpenLog(filename)

	defer logfile.Close()

	if err != nil {
		t.Errorf("Could not open file %s", filename)
	}

	if header.Format != 2 {
		t.Errorf("header.Format = %v, want %v", header.Format, 2)
	}
	if header.Blocksize != 1970 {
		t.Errorf("header.Blocksize = %v, wants %v", header.Blocksize, 1970)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Frame
		wantLon float64
		wantLat float64
		wantCog float64
		wantAlt float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				version:   header.Version,
				blocksize: header.Blocksize,
				r:         logfile,
			},
			wantErr: false,
			want: &Frame{
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
				r:         tt.fields.r,
				version:   tt.fields.version,
				blocksize: tt.fields.blocksize,
			}
			got, err := d.Decode()
			if (err != nil) != tt.wantErr {
				t.Errorf("slDecoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// logOffset(logfile)
			// logLongAtOffset(logfile, 8+140)
			// log.Printf("COG in degrees %f", RadToDeg(got.COG))
			// log.Printf("Altitude in meters %f", FeetToMeter(got.Altitude))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slDecoder.Decode() =\n %+v,\n want\n %+v", got, tt.want)
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
