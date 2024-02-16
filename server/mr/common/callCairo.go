package common

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// could call a bash script
// or run the cairo code directly (current implementation)

// runs main cairo program in a given directory, returns list of intermediate files
// errs if there is no data provided or if running has error
// outputs the result of Cairo program to intermedite.json
// CallCairoMap runs the Cairo program with specified gas and outputs the results to a file.
func CallCairoMap(mapJobNumber int, dst string) []string {
	var filenames []string // Slice to store the names of the intermediate files

	// Generate the filename based on mapjob and partition.
	// temp just make reducer num the same as mapper
	filename := fmt.Sprintf("mr-%d-%d", mapJobNumber, mapJobNumber)

	// Define the directory where the file will be saved.
	outputDir := dst                     //created to debug
	executionDir := "/app/cairo/map/src" // working

	// Ensure the output directory exists.
	err := os.MkdirAll(outputDir, 0755) // 0755 is commonly used permission for directories
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// Create the full path for the file.
	fullPath := filepath.Join(outputDir, filename)

	// Append the fullPath of the created file to the filenames slice
	filenames = append(filenames, filename)

	// Create the output file.
	outFile, err := os.Create(fullPath)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outFile.Close()

	// Prepare the command to run the Cairo program. Set the execution and output.
	cmd := exec.Command("bash", "exe_map.sh")
	cmd.Dir = executionDir
	cmd.Stdout = outFile

	// Run the command.
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	fmt.Println("Executed Cairo program successfully, output saved to", fullPath)

	return filenames
}

// runs main cairo program in a given directory
// errs if there is no data provided or if running has error
// outputs the result of Cairo program to mr-out-NUMBER.json
func CallCairoReduce(jobid string, dst string) {
	//Name of the reduce file
	filename := fmt.Sprintf("mr-out-%s", jobid)

	// Define the directory where the file will be saved.
	outputDir := dst //created to debug
	executionDir := "/app/cairo/reducer/src"
	// Ensure the output directory exists.
	err := os.MkdirAll(outputDir, 0755) // 0755 is commonly used permission for directories
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// Create the full path for the file.
	fullPath := filepath.Join(outputDir, filename)

	// Create the output file.
	outFile, err := os.Create(fullPath)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer outFile.Close()

	// Prepare the command to run the Cairo program. Set the execution and output.
	cmd := exec.Command("bash", "exe_reduce.sh", jobid)
	cmd.Dir = executionDir
	cmd.Stdout = outFile

	// Run the command.
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	fmt.Println("Executed Cairo program successfully, output saved to", fullPath)
}
