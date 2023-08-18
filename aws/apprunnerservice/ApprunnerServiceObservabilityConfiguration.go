package apprunnerservice


type ApprunnerServiceObservabilityConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/apprunner_service#observability_enabled ApprunnerService#observability_enabled}.
	ObservabilityEnabled interface{} `field:"required" json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/apprunner_service#observability_configuration_arn ApprunnerService#observability_configuration_arn}.
	ObservabilityConfigurationArn *string `field:"optional" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

