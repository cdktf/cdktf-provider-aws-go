package provider


type AwsProviderIgnoreTags struct {
	// Resource tag key prefixes to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs#key_prefixes AwsProvider#key_prefixes}
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Resource tag keys to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs#keys AwsProvider#keys}
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

