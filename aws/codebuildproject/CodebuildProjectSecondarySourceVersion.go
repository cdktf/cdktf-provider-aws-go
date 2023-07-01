package codebuildproject


type CodebuildProjectSecondarySourceVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/codebuild_project#source_identifier CodebuildProject#source_identifier}.
	SourceIdentifier *string `field:"required" json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/codebuild_project#source_version CodebuildProject#source_version}.
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
}

