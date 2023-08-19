package sagemakerdomain


type SagemakerDomainDomainSettingsRStudioServerProDomainSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#domain_execution_role_arn SagemakerDomain#domain_execution_role_arn}.
	DomainExecutionRoleArn *string `field:"required" json:"domainExecutionRoleArn" yaml:"domainExecutionRoleArn"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#r_studio_connect_url SagemakerDomain#r_studio_connect_url}.
	RStudioConnectUrl *string `field:"optional" json:"rStudioConnectUrl" yaml:"rStudioConnectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#r_studio_package_manager_url SagemakerDomain#r_studio_package_manager_url}.
	RStudioPackageManagerUrl *string `field:"optional" json:"rStudioPackageManagerUrl" yaml:"rStudioPackageManagerUrl"`
}

