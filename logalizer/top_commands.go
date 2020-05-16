package logalizer

import "strings"

// TopIpCmd .. returns top 10 IPs
func TopIpCmd(filename *string) {
	ips := ParseLog(*filename, "ip")
	ip := TopOccurr(ips)
	TablePrint("IP", "Count", ip)
}

// TopIP2LocCmd returns top 10 IPs with their locations
func TopIP2LocCmd(filename *string) {
	ips := ParseLog(*filename, "ip")
	ip := TopOccurr(ips)
	TablePrintIP2Loc("IP", "Count", "Location", ip)
}

// TopMehodCmd ... returns top 10 methods
func TopMethodCmd(filename *string) {
	var list []string
	methods := ParseLog(*filename, "method")
	for _, method := range methods {
		list = append(list, method)
	}
	result := TopOccurr(list)
	TablePrint("Method", "Count", result)
}

// TopRequestCmd ... returns top 10 requests
func TopRequestCmd(filename *string) {
	var list []string
	requests := ParseLog(*filename, "request")
	for _, request := range requests {
		list = append(list, request)
	}
	result := TopOccurr(list)
	TablePrint("Request", "Count", result)
}

// TopReferrerCmd returns top 10 referrers
func TopReferrerCmd(filename *string) {
	var list []string
	referrers := ParseLog(*filename, "referrer")
	for _, referrer := range referrers {
		referrer = strings.Trim(referrer, "\" ")
		list = append(list, referrer)
	}
	result := TopOccurr(list)
	TablePrint("Referrer", "Count", result)
}

// TopStatusCmd ... returns top 10 status codes
func TopStatusCmd(filename *string) {
	var list []string
	response := ParseLog(*filename, "response")
	for _, code := range response {
		list = append(list, code[:3])
	}

	result := TopOccurr(list)
	TablePrint("Status", "Count", result)
}

// TopAgentCmd returns top 10 agents
func TopAgentCmd(filename *string) {
	var list []string
	agents := ParseLog(*filename, "agent")
	for _, agent := range agents {
		agent = strings.Trim(agent, "\" ")
		// table width problem if length of agent is too long
		// limit width to 90 still show importatn data
		if len(agent) >= 100 {
			list = append(list, agent[:90])
		} else {
			list = append(list, agent)
		}
	}
	result := TopOccurr(list)
	TablePrint("Agent", "Count", result)
}
