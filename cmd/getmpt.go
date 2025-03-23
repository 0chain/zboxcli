package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var getMptKeyCommand = &cobra.Command{
	Use:   "get-mpt",
	Short: "Directly view blockchain data",
	Long:  `Directly view blockchain data from MPT key`,
	Args:  cobra.MinimumNArgs(0),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("key") == false {
			Fatal("Required Mpt key missing\n")
		}
		key := cmd.Flag("key").Value.String()
		jsonBytes, err := sdk.GetMptData(key)
		if err != nil {
			Fatalf("Failed to get Mpt key: %v\n", err)
		}

		var indented bytes.Buffer
		err = json.Indent(&indented, jsonBytes, "", "\t")
		if err != nil {
			Fatalf("Result %s baddly formated: %v\n", string(jsonBytes), err)
		}

		noBackSlash := strings.Replace(indented.String(), "\\", "", -1)
		fmt.Println(key, ": ", noBackSlash)
		return
	},
}

func init() {
	rootCmd.AddCommand(getMptKeyCommand)
	getMptKeyCommand.PersistentFlags().String("key", "", "Key into MPT datastore")
	getMptKeyCommand.MarkFlagRequired("key")
}
