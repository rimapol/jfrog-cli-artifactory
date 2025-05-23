package repocreate

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
)

var Usage = []string{"rt rc <template path>"}

func GetDescription() string {
	return "Create a new repository in Artifactory."
}

func GetArguments() []components.Argument {
	return []components.Argument{
		{
			Name: "template path",
			Description: "Specifies the local file system path for the template file to be used for the repository creation. " +
				"The template can be created using the \"" + coreutils.GetCliExecutableName() + " rt rpt\" command.",
		},
	}
}
