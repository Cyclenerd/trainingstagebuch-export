package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter (Copy & Paste) your session key from https://trainingstagebuch.org/login/sso: ")
	ssoKey, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Trim whitespace and newlines
	ssoKey = strings.TrimSpace(ssoKey)

	fmt.Println("Starting exporting process. Please wait...")

	// Download CSV list from URL
	csvURL := fmt.Sprintf("https://trainingstagebuch.org/workouts/tschuess?sso=%s", ssoKey)
	csvData, err := downloadFile(csvURL)
	if err != nil {
		fmt.Printf("Error downloading CSV list: %v\n", err)
		os.Exit(1)
	}

	// Parse CSV data
	reader = bufio.NewReader(strings.NewReader(csvData))
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';' // Set semicolon as delimiter

	// Skip header row if it exists
	_, err = csvReader.Read()
	if err != nil {
		fmt.Printf("Error reading CSV header: %v\n", err)
		os.Exit(1)
	}

	// Loop over the CSV list and extract IDs
	var ids []string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading CSV row: %v\n", err)
			continue
		}

		if len(record) > 0 && record[0] != "" {
			ids = append(ids, record[0])
		}
	}

	fmt.Printf("Found %d IDs to process\n", len(ids))

	// Download GPX and CSV files for each ID
	for i, id := range ids {
		fmt.Printf("Processing ID %s (%d of %d)\n", id, i+1, len(ids))

		// Download GPX file
		gpxFilename := fmt.Sprintf("%s.gpx", id)
		if fileExists(gpxFilename) {
			fmt.Printf("Skipping download of %s (file already exists)\n", gpxFilename)
		} else {
			gpxURL := fmt.Sprintf("https://trainingstagebuch.org/map/export/%s?sso=%s", id, ssoKey)
			fmt.Printf("Downloading %s...\n", gpxFilename)
			gpxContent, err := downloadFile(gpxURL)
			if err != nil {
				fmt.Printf("Error downloading GPX file: %v\n", err)
			} else {
				err = saveToFile(gpxFilename, gpxContent)
				if err != nil {
					fmt.Printf("Error saving GPX file: %v\n", err)
				} else {
					fmt.Printf("Successfully downloaded %s\n", gpxFilename)
				}
			}
			// Wait between downloads
			waitTime := randomWaitTime(10, 25)
			fmt.Printf("Waiting %d seconds before next download...\n", waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)
		}

		// Download CSV file
		csvFilename := fmt.Sprintf("%s.csv", id)
		if fileExists(csvFilename) {
			fmt.Printf("Skipping download of %s (file already exists)\n", csvFilename)
		} else {
			csvURL := fmt.Sprintf("https://trainingstagebuch.org/file/csv/%s?sso=%s", id, ssoKey)
			fmt.Printf("Downloading %s...\n", csvFilename)
			csvContent, err := downloadFile(csvURL)
			if err != nil {
				fmt.Printf("Error downloading CSV file: %v\n", err)
			} else {
				err = saveToFile(csvFilename, csvContent)
				if err != nil {
					fmt.Printf("Error saving CSV file: %v\n", err)
				} else {
					fmt.Printf("Successfully downloaded %s\n", csvFilename)
				}
			}
			// Wait between downloads if not the last ID
			if i < len(ids)-1 {
				waitTime := randomWaitTime(10, 25)
				fmt.Printf("Waiting %d seconds before next download...\n", waitTime)
				time.Sleep(time.Duration(waitTime) * time.Second)
			}
		}
	}

	fmt.Println("Export process completed!")
}

// Helper function to download file content as string
func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close() // Explicitly ignoring the error
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// Helper function to save content to file
// Helper function to save content to file
func saveToFile(filename string, content string) error {
	file, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}

	_, writeErr := file.WriteString(content)

	// Close the file and check for errors
	closeErr := file.Close()

	// Return the first error encountered
	if writeErr != nil {
		return writeErr
	}
	return closeErr
}

// Helper function to check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Helper function to generate random wait time between min and max seconds
func randomWaitTime(min, max int) int {
	return min + rand.Intn(max-min+1)
}
