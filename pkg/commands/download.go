package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type downloadArgs struct {
	uri    string
	errors []error
}

func (da *downloadArgs) Error() string {
	var s strings.Builder
	for _, err := range da.errors {
		fmt.Fprintf(&s, "%v\n", err)
	}
	return s.String()
}

func (da *downloadArgs) Validate() error {
	errs := []error{}

	if vErrs := validateURI(da.uri); len(vErrs) > 0 {
		errs = append(errs, vErrs...)
	}

	if len(errs) > 0 {
		da.errors = errs
		return da
	}

	return nil
}

var downloadCMD = &cobra.Command{
	Use:   "download",
	Short: "Download a blob from the storage",
	Long:  `Download a blob from the storage by specifying its unique identifier.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		downloadArgs := &downloadArgs{}
		downloadArgs.uri, _ = cmd.Flags().GetString("uri")
		err := downloadArgs.Validate()
		if err != nil {
			fmt.Println("Validation errors:")
			fmt.Println(err)
			return err
		}

		fmt.Println("Download command executed")

		return nil
	},
}

func InitDownload() *cobra.Command {
	downloadCMD.Flags().StringP("uri", "u", "", "URI of the blob to download")
	downloadCMD.MarkFlagRequired("uri")

	return downloadCMD
}
