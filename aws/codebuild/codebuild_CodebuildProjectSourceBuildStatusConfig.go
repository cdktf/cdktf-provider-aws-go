package codebuild


type CodebuildProjectSourceBuildStatusConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#context CodebuildProject#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#target_url CodebuildProject#target_url}.
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

