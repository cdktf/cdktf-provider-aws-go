// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#access_key_id AppflowConnectorProfile#access_key_id}.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#datakey AppflowConnectorProfile#datakey}.
	Datakey *string `field:"required" json:"datakey" yaml:"datakey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#secret_access_key AppflowConnectorProfile#secret_access_key}.
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#user_id AppflowConnectorProfile#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

