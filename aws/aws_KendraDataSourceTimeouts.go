// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraDataSourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#create KendraDataSource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#delete KendraDataSource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#update KendraDataSource#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

