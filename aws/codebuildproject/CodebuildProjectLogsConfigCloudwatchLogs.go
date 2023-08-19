package codebuildproject


type CodebuildProjectLogsConfigCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_project#group_name CodebuildProject#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_project#status CodebuildProject#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_project#stream_name CodebuildProject#stream_name}.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

