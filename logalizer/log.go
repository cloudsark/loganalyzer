package logalizer

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadFile ... Read access.log file to preapre for analysis
func ReadFile(fileName string) []string {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	return fileTextLines
}

func findMatch(log []string, exp *regexp.Regexp) []string {
	var matches []string
	for _, e := range log {
		ips := exp.FindAllString(e, -1)
		for _, element := range ips {
			matches = append(matches, element)
		}
	}
	return matches
}

// ParseLog ... find all matches in the log file
func ParseLog(log []string) (ipMatch, dateMatch, requestMatch, responseMatch, sizeMatch, referrerMatch, agentMatch, fullReqMatch []string) {
	ip := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	date := regexp.MustCompile(`\[\s*(\d+/\D+/.*?)\]`)
	request := regexp.MustCompile(`^GET|POST|UPDATE|HEAD|DELETE`)
	response := regexp.MustCompile(`\d+ (\d+)`)
	size := regexp.MustCompile(`([0-9]{1,3}) \"`)
	referrer := regexp.MustCompile(`"([^"]*)" \"`)
	agent := regexp.MustCompile(`" "([^"]*)"`)
	fullReq := regexp.MustCompile(`(GET|POST|UPDATE|HEAD|DELETE) \/.*?\.[\w:]+`)
	ipMatch = findMatch(log, ip)
	dateMatch = findMatch(log, date)
	requestMatch = findMatch(log, request)
	responseMatch = findMatch(log, response)
	sizeMatch = findMatch(log, size)
	referrerMatch = findMatch(log, referrer)
	agentMatch = findMatch(log, agent)
	fullReqMatch = findMatch(log, fullReq)

	return
}

// TopOccurr ... top occurrence
func TopOccurr(field []string) map[string]int {
	var d = map[string]int{}

	for _, item := range field {
		d[item]++
	}
	return d
}

// TotalBandwidth ... calculate total bandwidt
func TotalBandwidth(slice []string) int {
	nums := slice
	sum := 0

	for _, num := range nums {
		x := strings.Trim(num, "\"")
		x = strings.TrimSpace(x)
		i1, _ := strconv.Atoi(x)
		sum += i1

	}

	return sum
}
