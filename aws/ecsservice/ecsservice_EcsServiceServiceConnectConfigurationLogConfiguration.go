package ecsservice


type EcsServiceServiceConnectConfigurationLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#log_driver EcsService#log_driver}.
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#options EcsService#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#secret_option EcsService#secret_option}
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

