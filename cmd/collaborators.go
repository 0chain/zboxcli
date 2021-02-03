package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var addCollabCmd = &cobra.Command{
	Use:   "add-collab",
	Short: "add collaborator for a file",
	Long:  `add collaborator for a file`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		allocationID := cmd.Flag("allocation").Value.String()
		if len(allocationID) == 0 {
			PrintError("Error: allocation flag is missing")
			os.Exit(1)
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		if len(remotepath) == 0 {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		collabID := cmd.Flag("collabid").Value.String()
		if len(collabID) == 0 {
			PrintError("Error: collabid flag is missing")
			os.Exit(1)
		}

		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}

		err = allocationObj.AddCollaborator(remotepath, collabID)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Collaborator %s added successfully for the file %s \n", collabID, remotepath)
		return
	},
}

var deleteCollabCmd = &cobra.Command{
	Use:   "delete-collab",
	Short: "delete collaborator for a file",
	Long:  `delete collaborator for a file`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		allocationID := cmd.Flag("allocation").Value.String()
		if len(allocationID) == 0 {
			PrintError("Error: allocation flag is missing")
			os.Exit(1)
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		if len(remotepath) == 0 {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		collabID := cmd.Flag("collabid").Value.String()
		if len(collabID) == 0 {
			PrintError("Error: collabid flag is missing")
			os.Exit(1)
		}

		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}

		err = allocationObj.RemoveCollaborator(remotepath, collabID)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Collaborator %s removed successfully for the file %s \n", collabID, remotepath)
		return
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
