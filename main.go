package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

//LucasOneZ
var validTLDs = []string{
	"com", "org", "net", "edu", "gov", "io", "co", "us", "in", "info", "biz", "online", "xyz",
}

func isValidTLD(word string) bool {
	for _, tld := range validTLDs {
		if strings.EqualFold(word, tld) {
			return true
		}
	}
	return false
}

func cleanDomain(domain, addOn string, removeCount int, force bool) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 1 {
		if force || removeCount > 0 || !isValidTLD(parts[len(parts)-1]) {
			if removeCount > 0 && removeCount < len(parts) {
				// Remove specified number of words from the end
				parts = parts[:len(parts)-removeCount]
			} else {
				// Remove the last part if it's not a valid TLD or if forced
				parts = parts[:len(parts)-1]
			}
		}
	}

	cleanedDomain := strings.Join(parts, ".")
	if addOn != "" && !strings.HasSuffix(cleanedDomain, "."+addOn) {
		cleanedDomain += "." + addOn
	}

	return cleanedDomain
}

func main() {
	fileFlag := flag.String("file", "", "File containing domains or subdomains")
	addOnFlag := flag.String("add", "", "String to add to the cleaned domain if missing")
	outputFlag := flag.String("output", "", "File to write only the cleaned domains")
	removeCountFlag := flag.Int("removecount", 0, "Number of words to remove after the last dot (Recommended: 2)")
	forceFlag := flag.Bool("force", false, "Forcibly remove words even if they seem to be valid TLDs")
	verboseFlag := flag.Bool("verbose", false, "Print both original and cleaned domains to the console")
	flag.Parse()

	if *fileFlag == "" {
		fmt.Println("Usage: go run main.go -file=<filename> [-add=<string>] [-output=<outputfile>] [-removecount=<count>] [-force=true] [-verbose=true]")
		return
	}

	inputFile, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	domainSet := make(map[string]bool)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		domain := scanner.Text()
		cleaned := cleanDomain(domain, *addOnFlag, *removeCountFlag, *forceFlag)
		domainSet[cleaned] = true
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var sortedDomains []string
	for domain := range domainSet {
		sortedDomains = append(sortedDomains, domain)
	}
	sort.Strings(sortedDomains)

	var outputFile *os.File
	if *outputFlag != "" {
		outputFile, err = os.Create(*outputFlag)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()
	}

	for _, domain := range sortedDomains {
		if *outputFlag != "" {
			outputFile.WriteString(domain + "\n")
		}

		if *verboseFlag || *outputFlag == "" {
			fmt.Println("Cleaned:", domain)
		}
	}
}

