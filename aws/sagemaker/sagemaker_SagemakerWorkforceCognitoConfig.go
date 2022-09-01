package sagemaker


type SagemakerWorkforceCognitoConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workforce#client_id SagemakerWorkforce#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workforce#user_pool SagemakerWorkforce#user_pool}.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

