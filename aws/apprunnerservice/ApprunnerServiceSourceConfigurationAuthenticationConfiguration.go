package apprunnerservice


type ApprunnerServiceSourceConfigurationAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/apprunner_service#access_role_arn ApprunnerService#access_role_arn}.
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/apprunner_service#connection_arn ApprunnerService#connection_arn}.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

