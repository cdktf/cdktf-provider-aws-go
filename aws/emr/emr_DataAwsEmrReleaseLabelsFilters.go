package emr


type DataAwsEmrReleaseLabelsFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels#application DataAwsEmrReleaseLabels#application}.
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels#prefix DataAwsEmrReleaseLabels#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

