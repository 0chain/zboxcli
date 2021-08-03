package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var getMptKeyCommand = &cobra.Command{
	Use:   "get-mpt",
	Short: "directly view blockchain data",
	Long:  `directly view blockchain data from MPT key`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		key := cmd.Flag("allocation").Value.String()
		jsonBytes, err := sdk.GetMptData(key)
		if err != nil {
			log.Fatalf("Failed to get Mpt key: %v\n", err)
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, jsonBytes, "", "\t")
		if err != nil {
			log.Fatalf("Result %s baddly formated: %v\n", string(b), err)
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
