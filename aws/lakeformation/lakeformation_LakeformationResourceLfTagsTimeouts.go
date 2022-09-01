package lakeformation


type LakeformationResourceLfTagsTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource_lf_tags#create LakeformationResourceLfTags#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_resource_lf_tags#delete LakeformationResourceLfTags#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

