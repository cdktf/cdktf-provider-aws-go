package ivschatloggingconfiguration


type IvschatLoggingConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ivschat_logging_configuration#create IvschatLoggingConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ivschat_logging_configuration#delete IvschatLoggingConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ivschat_logging_configuration#update IvschatLoggingConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

