package logalizer

// TotlaBandwidthCmd ... returns total bandwidth in KB , MB and GB
func TotalBandwidthCmd(filename *string) {
	F := ReadFile(*filename)
	_, _, _, _, size, _, _, _ := ParseLog(F)
	totalsize := TotalBandwidth(size)
	TabelBandwidth(totalsize)
}

// TopIpCmd .. returns top 10 IPs
func TopIpCmd(filename *string) {
	F := ReadFile(*filename)
	ips, _, _, _, _, _, _, _ := ParseLog(F)
	ip := TopOccurr(ips)
	TablePrint("IP", "Count", ip)
}

// TopMehodCmd ... returns top 10 methods
func TopMethodCmd(filename *string) {
	var list []string
	F := ReadFile(*filename)
	_, _, methods, _, _, _, _, _ := ParseLog(F)
	for _, method := range methods {
		list = append(list, method)
	}
	result := TopOccurr(list)
	TablePrint("Method", "Count", result)
}

// TopRequestCmd ... returns top 10 requests
func TopRequestCmd(filename *string) {
	var list []string
	F := ReadFile(*filename)
	_, _, _, _, _, _, _, requests := ParseLog(F)
	for _, request := range requests {
		list = append(list, request)
	}
	result := TopOccurr(list)
	TablePrint("Request", "Count", result)
}

// TopStatusCmd ... returns top 10 status codes
func TopStatusCmd(filename *string) {
	var list []string
	F := ReadFile(*filename)
	_, _, _, response, _, _, _, _ := ParseLog(F)
	for _, code := range response {
		list = append(list, code[:3])
	}

	result := TopOccurr(list)
	TablePrint("Status", "Count", result)
}
