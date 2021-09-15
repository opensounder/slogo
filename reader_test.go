package slogo

import "testing"

func TestReadLogfile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				path: "",
			},
			wantErr: true,
		},
		{
			args: args{
				path: "./testdata/sample-data-lowrance/Elite_4_Chirp/bigger.sl2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadLogfile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ReadLogfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
