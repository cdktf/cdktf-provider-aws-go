package glacier


type GlacierVaultNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glacier_vault#events GlacierVault#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glacier_vault#sns_topic GlacierVault#sns_topic}.
	SnsTopic *string `field:"required" json:"snsTopic" yaml:"snsTopic"`
}

