// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AwsProviderIgnoreTags struct {
	// Resource tag key prefixes to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#key_prefixes AwsProvider#key_prefixes}
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Resource tag keys to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#keys AwsProvider#keys}
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

