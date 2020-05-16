package logalizer

// TotlaBandwidthCmd ... returns total bandwidth in KB , MB and GB
func TotalBandwidthCmd(filename *string) {
	size := ParseLog(*filename, "size")
	totalsize := TotalBandwidth(size)
	TabelBandwidth(totalsize)
}
