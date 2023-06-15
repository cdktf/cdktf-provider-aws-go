package codebuildproject


type CodebuildProjectSourceAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/codebuild_project#resource CodebuildProject#resource}.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

