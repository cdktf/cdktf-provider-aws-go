package codebuild


type CodebuildProjectSourceGitSubmodulesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#fetch_submodules CodebuildProject#fetch_submodules}.
	FetchSubmodules interface{} `field:"required" json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

