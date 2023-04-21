package ecstaskdefinition


type EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration struct {
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ecs_task_definition#authorization_config EcsTaskDefinition#authorization_config}
	AuthorizationConfig *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ecs_task_definition#file_system_id EcsTaskDefinition#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ecs_task_definition#root_directory EcsTaskDefinition#root_directory}.
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

