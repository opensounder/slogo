package slogo

//Header represents the log file header
type Header struct {
	Format    int16
	Version   int16
	Blocksize int16
	Reserved1 int16
}
