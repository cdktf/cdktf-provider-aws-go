package ecs


type EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_task_definition#credentials_parameter EcsTaskDefinition#credentials_parameter}.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_task_definition#domain EcsTaskDefinition#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

