package main

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/utils"
)

var root = &cobra.Command{
	Use:                    "run",
	Aliases:                []string{"start"},
	SuggestFor:             []string{"what?"},
	Short:                  "",
	Long:                   "启动项目",
	Example:                "run",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "0.1",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
	Run: func(cmd *cobra.Command, args []string) {

	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}

func main() {
	utils.NoError(root.Execute())

}

func init() {
	//root.Flags().StringP("version", "v", "", "version")
}
