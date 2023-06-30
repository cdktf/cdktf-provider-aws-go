package codebuildproject


type CodebuildProjectBuildBatchConfigRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/codebuild_project#compute_types_allowed CodebuildProject#compute_types_allowed}.
	ComputeTypesAllowed *[]*string `field:"optional" json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/codebuild_project#maximum_builds_allowed CodebuildProject#maximum_builds_allowed}.
	MaximumBuildsAllowed *float64 `field:"optional" json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

