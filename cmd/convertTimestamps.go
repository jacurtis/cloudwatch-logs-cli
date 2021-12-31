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

// convertTimestampsCmd represents the convertTimestamps command
var convertTimestampsCmd = &cobra.Command{
	Use:   "convert-timestamps",
	Short: "Convert the timestamps column in cloudwatch logs into a human readable date",
	Long: `CloudWatch Log exports will by default provide the first column with UNIX Timestamps
	instead of dates. Since UNIX timestamps are not human readable, this command will convert
	into something that is easier to parse for humans. 
	For example:

Example here`,
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
				log.Fatal(err)
			}
			tm := time.UnixMilli(i)
			timezone, _ := cmd.Flags().GetString("timezone")
			tz, _ := time.LoadLocation(timezone)
			fmt.Println(tm.In(tz).Format(time.RFC1123))
		}

		// for {
		// 	row, err := reader.Read()

		// 	if err == io.EOF {
		// 		break
		// 	}

		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(row[0])
		// for col := range row {
		// 	fmt.Printf("%s\n", row[col])
		// }
		// Loop through each line
		// Convert the timestamp to date
		// Save the file
		// Print something to the user
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
	convertTimestampsCmd.Flags().StringP("timezone", "z", "America/Chicago", "Timezone to convert into")
}
