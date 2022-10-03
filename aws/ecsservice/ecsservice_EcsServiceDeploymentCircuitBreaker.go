package ecsservice


type EcsServiceDeploymentCircuitBreaker struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#enable EcsService#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#rollback EcsService#rollback}.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

