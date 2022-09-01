package codebuild


type CodebuildProjectBuildBatchConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#service_role CodebuildProject#service_role}.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#combine_artifacts CodebuildProject#combine_artifacts}.
	CombineArtifacts interface{} `field:"optional" json:"combineArtifacts" yaml:"combineArtifacts"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#restrictions CodebuildProject#restrictions}
	Restrictions *CodebuildProjectBuildBatchConfigRestrictions `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#timeout_in_mins CodebuildProject#timeout_in_mins}.
	TimeoutInMins *float64 `field:"optional" json:"timeoutInMins" yaml:"timeoutInMins"`
}

