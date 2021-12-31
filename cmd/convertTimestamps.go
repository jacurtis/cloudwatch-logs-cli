/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// convertTimestampsCmd represents the convert-timestamps command
var convertTimestampsCmd = &cobra.Command{
	Use:   "convert-timestamps",
	Short: "Convert the timestamps column in cloudwatch logs into a human readable date",
	Long: `CloudWatch Log exports will by default provide the first column with UNIX Timestamps
	instead of dates. Since UNIX timestamps are not human readable, this command will convert
	into something that is easier to parse for humans. 

	Usage:

	cloudwatch-logs convert-timestamps [file] -z [timezone]
	cloudwatch-logs convert-timestamps myfile.csv -z America/Chicago
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Process CSV
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		// Skip first line
		if _, err := reader.Read(); err != nil {
			log.Fatal(err)
		}

		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		for _, record := range records {
			i, err := strconv.ParseInt(record[0], 10, 64)
			if err != nil {
				fmt.Println("Could not parse timestamps. Are they already converted?")
				fmt.Println("Check the file and run again.")
				os.Exit(1)
			}
			tm := time.UnixMilli(i)
			timezone, _ := cmd.Flags().GetString("timezone")
			tz, _ := time.LoadLocation(timezone)
			record[0] = tm.In(tz).Format(time.RFC1123)
		}
		fmt.Printf("Successfully converted %s records\n", strconv.Itoa(len(records)))
		fmt.Println("Saving to file...")

		newFilename, _ := cmd.Flags().GetString("rename")
		if newFilename == "" {
			newFilename = args[0]
		}

		// Save the file
		newFile, err := os.Create(newFilename)
		defer newFile.Close()
		if err != nil {
			fmt.Println("Failed to create new file")
			log.Fatal(err)
		}
		writer := csv.NewWriter(newFile)
		err = writer.WriteAll(records)
		if err != nil {
			log.Fatal(err)
		}

		// Print something to the user
		fmt.Println("File is updated and ready to share!")
	},
}

func init() {
	rootCmd.AddCommand(convertTimestampsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertTimestampsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	convertTimestampsCmd.Flags().StringP("rename", "n", "", "Rename the file to this name (add .csv to end)")
	convertTimestampsCmd.Flags().StringP("timezone", "z", "America/Chicago", "Timezone to convert into")
}
