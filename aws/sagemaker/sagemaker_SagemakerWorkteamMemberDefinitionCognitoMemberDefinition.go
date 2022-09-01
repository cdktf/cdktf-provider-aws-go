package sagemaker


type SagemakerWorkteamMemberDefinitionCognitoMemberDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workteam#client_id SagemakerWorkteam#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workteam#user_group SagemakerWorkteam#user_group}.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workteam#user_pool SagemakerWorkteam#user_pool}.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

