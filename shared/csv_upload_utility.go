package shared

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// maxConcurrentUploads controls the number of concurrent Firestore write operations
const maxConcurrentUploads = 100

// UploadCaCountyCsvToFirestore uploads California county tax rate data from a CSV file to Firestore.
//
// The CSV file should be located at "./extras/ca_county_details.csv" and must have the following structure:
//   - Expected columns: [irrelevant,..., irrelevant, county, city, irrelevant, irrelevant, rate, ...]
//     where 'county' is at index 3, 'city' at index 4, and 'rate' at index 6.
//
// Behavior:
//   - Reads the CSV file from disk.
//   - Parses each record, converting the tax rate to a float.
//   - Uploads each record as a document to the "tax_rates" collection in Firestore.
//   - Uses concurrency with a limit of maxConcurrentUploads to optimize upload performance.
//   - Logs progress and errors encountered during the process.
//
// Logs:
//   - Logs start and completion of upload.
//   - Logs individual record errors if upload fails.
//
// Panics:
//   - If the CSV file cannot be opened or read, or if any tax rate cannot be converted to float.
func UploadTaxRatesInBulkToFirestore() {
	file, err := os.Open("./extras/ca_county_details.csv")
	if err != nil {
		log.Fatal("Error occurred reading the file, please try again")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
		return
	}

	log.Print("Uploading the county CSV data to Firestore...")

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentUploads)

	// Iterate over the CSV records (skipping the header row).
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}

		wg.Add(1)
		sem <- struct{}{} // Acquire a slot in the semaphore for concurrency control

		go func(row []string) {
			defer wg.Done()
			defer func() { <-sem }() // Release the semaphore slot

			// Convert the tax rate (string) to float64.
			rate, err := strconv.ParseFloat(row[6], 64)
			if err != nil {
				log.Fatalf("Error converting rate to float for row (%s - %s): %v", row[3], row[4], err)
			}

			countyObject := map[string]any{
				"state":  "CALIFORNIA",
				"county": row[3],
				"city":   row[4],
				"rate":   rate,
			}

			// Upload the document to Firestore.
			_, _, err = FirestoreClient.Collection("tax_rates").Add(context.Background(), countyObject)
			if err != nil {
				log.Printf("Error uploading record (%s - %s): %v", row[3], row[4], err)
			}
		}(row)
	}

	wg.Wait()
	log.Print("Added all the county data to Firestore.")
}

func UploadPropertiesDetailsInBulkToFirestore(){
	file, err := os.Open("./extras/property_details.csv")
	if err != nil {
		log.Fatal("Error occurred reading the file, please try again")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
		return
	}

	log.Print("Uploading the property details to firestore...")

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentUploads)

	// Iterate over the CSV records (skipping the header row).
	for i, row := range records {
		if i == 0 {
			continue // Skip the first row 
		}

		wg.Add(1)
		sem <- struct{}{} // Acquire a slot in the semaphore for concurrency control

		go func(row []string) {
			defer wg.Done()
			defer func() { <-sem }() // Release the semaphore slot

			property := map[string]any{
				"name":  strings.ToUpper(row[0]),
				"street_address": strings.ToUpper(row[2]),
				"city": strings.ToUpper(row[3]),
				"state": "CALIFORNIA", 
				"country": "UNITED STATES", 
				"postal_code": row[6],
			}
			// Upload the document to Firestore.
			_, _, err = FirestoreClient.Collection("properties").Add(context.Background(), property)
			if err != nil {
				log.Printf("Error uploading the record: %s", err)
			}
		}(row)
	}

	wg.Wait()
	log.Print("Added all the properties have been uploaded to Firestore.")
}