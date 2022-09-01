// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#create OpensearchDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#delete OpensearchDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#update OpensearchDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

