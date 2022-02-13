package cmd

import (
	"fmt"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

var addCollabCmd = &cobra.Command{
	Use:   "add-collab",
	Short: "add collaborator for a file",
	Long:  `add collaborator for a file`,
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		allocationID := cmd.Flag("allocation").Value.String()
		if len(allocationID) == 0 {
			return util.LogFatalErr("Error: allocation flag is missing")
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		if len(remotepath) == 0 {
			return util.LogFatalErr("Error: remotepath flag is missing")
		}

		collabID := cmd.Flag("collabid").Value.String()
		if len(collabID) == 0 {
			return util.LogFatalErr("Error: collabid flag is missing")
		}

		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			return util.LogFatalErrf("Error fetching the allocation: %s", err)
		}

		if err := allocationObj.AddCollaborator(remotepath, collabID); err != nil {
			return util.LogFatalErrf("%s", err.Error())
		}
		fmt.Printf("Collaborator %s added successfully for the file %s \n", collabID, remotepath)
		return nil
	},
}

var deleteCollabCmd = &cobra.Command{
	Use:   "delete-collab",
	Short: "delete collaborator for a file",
	Long:  `delete collaborator for a file`,
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		allocationID := cmd.Flag("allocation").Value.String()
		if len(allocationID) == 0 {
			return util.LogFatalErr("Error: allocation flag is missing")
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		if len(remotepath) == 0 {
			return util.LogFatalErr("Error: remotepath flag is missing")
		}

		collabID := cmd.Flag("collabid").Value.String()
		if len(collabID) == 0 {
			return util.LogFatalErr("Error: collabid flag is missing")
		}

		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			return util.LogFatalErrf("Error fetching the allocation: %s", err)
		}

		err = allocationObj.RemoveCollaborator(remotepath, collabID)
		if err != nil {
			return util.LogFatalErr(err.Error())
		}
		fmt.Printf("Collaborator %s removed successfully for the file %s \n", collabID, remotepath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCollabCmd)
	addCollabCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	addCollabCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	addCollabCmd.PersistentFlags().String("collabid", "", "Collaborator's clientID")
	addCollabCmd.MarkFlagRequired("allocation")
	addCollabCmd.MarkFlagRequired("remotepath")
	addCollabCmd.MarkFlagRequired("collabid")

	rootCmd.AddCommand(deleteCollabCmd)
	deleteCollabCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	deleteCollabCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	deleteCollabCmd.PersistentFlags().String("collabid", "", "Collaborator's clientID")
	deleteCollabCmd.MarkFlagRequired("allocation")
	deleteCollabCmd.MarkFlagRequired("remotepath")
	deleteCollabCmd.MarkFlagRequired("collabid")
}
