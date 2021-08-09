package list

import (
	"encoding/json"
	"fmt"

	"github.com/devopstoday11/sigrun/pkg/controller"

	"github.com/tidwall/pretty"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists metadata about sigrun repos that have been added",
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			cont, err := controller.GetController()
			if err != nil {
				return err
			}

			repoInfo, err := cont.List()
			if err != nil {
				return err
			}

			encodedRepoInfo, err := json.Marshal(repoInfo)
			if err != nil {
				return err
			}

			fmt.Println("Sigrun-repos:\n" + string(pretty.Pretty(encodedRepoInfo)))
			return nil
		},
	}
}
