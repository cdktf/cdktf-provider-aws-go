package apprunner


type ApprunnerServiceObservabilityConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#observability_configuration_arn ApprunnerService#observability_configuration_arn}.
	ObservabilityConfigurationArn *string `field:"required" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#observability_enabled ApprunnerService#observability_enabled}.
	ObservabilityEnabled interface{} `field:"required" json:"observabilityEnabled" yaml:"observabilityEnabled"`
}

