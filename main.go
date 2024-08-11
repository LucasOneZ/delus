package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

// A simple list of common TLDs. You can extend this list as needed.
var validTLDs = []string{
	"com", "org", "net", "edu", "gov", "io", "co", "us", "in", "info", "biz", "online", "xyz",
}

// Function to check if a string is a valid TLD
func isValidTLD(word string) bool {
	for _, tld := range validTLDs {
		if strings.EqualFold(word, tld) {
			return true
		}
	}
	return false
}

// Function to clean the domain or subdomain by removing parts based on flags
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

	// Ensure the domain ends with the specified add-on string
	cleanedDomain := strings.Join(parts, ".")
	if addOn != "" && !strings.HasSuffix(cleanedDomain, "."+addOn) {
		cleanedDomain += "." + addOn
	}

	return cleanedDomain
}

func main() {
	// Define command-line flags
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

	// Open the input file
	inputFile, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	// Map to store cleaned domains to handle uniqueness
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

	// Sort and write or print the cleaned domains
	var sortedDomains []string
	for domain := range domainSet {
		sortedDomains = append(sortedDomains, domain)
	}
	sort.Strings(sortedDomains)

	// Prepare output file if specified
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

