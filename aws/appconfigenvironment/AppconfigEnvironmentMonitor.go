package appconfigenvironment


type AppconfigEnvironmentMonitor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appconfig_environment#alarm_arn AppconfigEnvironment#alarm_arn}.
	AlarmArn *string `field:"required" json:"alarmArn" yaml:"alarmArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appconfig_environment#alarm_role_arn AppconfigEnvironment#alarm_role_arn}.
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

