package apprunner


type ApprunnerServiceSourceConfigurationAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#access_role_arn ApprunnerService#access_role_arn}.
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#connection_arn ApprunnerService#connection_arn}.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

