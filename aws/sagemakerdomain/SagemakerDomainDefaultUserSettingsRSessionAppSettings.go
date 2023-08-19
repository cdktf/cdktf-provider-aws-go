package sagemakerdomain


type SagemakerDomainDefaultUserSettingsRSessionAppSettings struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#custom_image SagemakerDomain#custom_image}
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

