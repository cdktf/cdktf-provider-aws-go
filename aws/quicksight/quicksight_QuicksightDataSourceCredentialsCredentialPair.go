package quicksight


type QuicksightDataSourceCredentialsCredentialPair struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#password QuicksightDataSource#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#username QuicksightDataSource#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

