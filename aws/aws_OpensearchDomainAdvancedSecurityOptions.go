// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainAdvancedSecurityOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#enabled OpensearchDomain#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#internal_user_database_enabled OpensearchDomain#internal_user_database_enabled}.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#master_user_options OpensearchDomain#master_user_options}
	MasterUserOptions *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

