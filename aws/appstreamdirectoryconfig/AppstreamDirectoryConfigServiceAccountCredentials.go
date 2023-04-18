package appstreamdirectoryconfig


type AppstreamDirectoryConfigServiceAccountCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/appstream_directory_config#account_name AppstreamDirectoryConfig#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/appstream_directory_config#account_password AppstreamDirectoryConfig#account_password}.
	AccountPassword *string `field:"required" json:"accountPassword" yaml:"accountPassword"`
}

