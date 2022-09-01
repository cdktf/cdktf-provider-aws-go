// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AwsProviderDefaultTags struct {
	// Resource tags to default across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#tags AwsProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

