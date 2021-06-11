package config

import (
	"fmt"
	"github.com/0chain/zboxcli/model"
	"github.com/spf13/cobra"
	"log"
)

//GetMigrationConfig returns user defined configs to migrate data from s3
func GetMigrationConfig(cmd *cobra.Command) (*model.MigrationConfig, error) {
	config := model.MigrationConfig{}

	fflags := cmd.Flags()                 // fflags is a *flag.FlagSet
	if fflags.Changed("bucket") == true { // set bucket
		buckets, err := fflags.GetStringSlice("bucket")
		if err != nil {
			log.Println("bucket list not provided, backing up all buckets in this region:")

		} else {
			//log.Println("buckets from cmd:", buckets)
			config.Buckets = buckets
		}
	}

	if fflags.Changed("region") == true { // check if the flag "region" is set
		config.Region, _ = fflags.GetString("region")
	}
	config.Prefix, _ = fflags.GetString("prefix")
	config.Concurrency, _ = fflags.GetInt("concurrency")
	config.DeleteSource, _ = cmd.Flags().GetBool("delete-source")

	if fflags.Changed("allocation") == false { // check if the flag "allocation" is set
		return nil, fmt.Errorf("allocation flag is missing")
	}

	config.AppConfig.AllocationID = cmd.Flag("allocation").Value.String()

	config.AppConfig.Encrypt, _ = cmd.Flags().GetBool("encrypt")
	config.AppConfig.Commit, _ = cmd.Flags().GetBool("commit")
	if fflags.Changed("attr-who-pays-for-reads") {
		wps, err := fflags.GetString("attr-who-pays-for-reads")
		if err != nil {
			return nil, fmt.Errorf("getting 'attr-who-pays-for-reads' flag: %v", err)
		}
		config.AppConfig.WhoPays = wps
	}

	return &config, nil
}
