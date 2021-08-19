package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

// scConfig shows SC configurations
var scConfig = &cobra.Command{
	Use:   "sc-config",
	Short: "Show storage SC configuration.",
	Long:  `Show storage SC configuration.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")
		var conf, err = sdk.GetStorageSCConfig()
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}
		if doJSON {
			util.PrintJSON(conf)
			return
		}

		printStorageSCConfig(conf)
	},
}

func printStorageSCConfig(conf *sdk.InputMap) {
	keys := make([]string, 0, len(conf.Fields))
	for k := range conf.Fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, "\t", conf.Fields[k])
	}

}

func init() {
	rootCmd.AddCommand(scConfig)
	scConfig.Flags().Bool("json", false, "pass this option to print response as json data")
}
