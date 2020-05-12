package logalizer

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/oschwald/geoip2-golang"
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

// find matches for a certain field in access log
func ParseLog(field string, log []string) []string {

	var fieldRegex *regexp.Regexp

	switch field {
	case "ip":
		fieldRegex = regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	case "date":
		fieldRegex = regexp.MustCompile(`\[\s*(\d+/\D+/.*?)\]`)
	case "method":
		fieldRegex = regexp.MustCompile(`^GET|POST|UPDATE|HEAD|DELETE`)
	case "response":
		fieldRegex = regexp.MustCompile(`\d+ (\d+)`)
	case "size":
		fieldRegex = regexp.MustCompile(`([0-9]{1,3}) \"`)
	case "referrer":
		fieldRegex = regexp.MustCompile(`"([^"]*)" \"`)
	case "agent":
		fieldRegex = regexp.MustCompile(`" "([^"]*)"`)
	case "request":
		fieldRegex = regexp.MustCompile(`(GET|POST|UPDATE|HEAD|DELETE) \/.*?\.[\w:]+`)
	default:
		err := errors.New("field not supported!")
		fmt.Printf("%v \n", err)
		os.Exit(1)
	}
	return findMatch(log, fieldRegex)
}

// Parse a custom field from the user
func ParseCustom(regexString string, log []string) []string {
	fieldRegex := regexp.MustCompile(regexString)
	return findMatch(log, fieldRegex)
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

// IP2Geo translate ip to Country name
func IP2Geo(IP string) string {
	db, err := geoip2.Open("db/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	ip := net.ParseIP(IP)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names["en"]
}
