package ecsservice


type EcsServiceAlarms struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#alarm_names EcsService#alarm_names}.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#enable EcsService#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#rollback EcsService#rollback}.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

