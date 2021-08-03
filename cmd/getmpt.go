package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var getMptKeyCommand = &cobra.Command{
	Use:   "get-mpt",
	Short: "Directly view blockchain data",
	Long:  `Directly view blockchain data from MPT key`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("key") == false {
			log.Fatal("Required Mpt key missing\n")
		}
		key := cmd.Flag("key").Value.String()
		jsonBytes, err := sdk.GetMptData(key)
		if err != nil {
			log.Fatalf("Failed to get Mpt key: %v\n", err)
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, jsonBytes, "", "\t")
		if err != nil {
			log.Fatalf("Result %s baddly formated: %v\n", string(jsonBytes), err)
		}

		fmt.Println(key, ": ", prettyJSON.String())
		return
	},
}

func init() {
	rootCmd.AddCommand(getMptKeyCommand)
	getMptKeyCommand.PersistentFlags().String("key", "", "Key into MPT datastore")
	getMptKeyCommand.MarkFlagRequired("key")
}
