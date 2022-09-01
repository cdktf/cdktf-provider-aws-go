package codedeploy


type CodedeployDeploymentGroupAlarmConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#alarms CodedeployDeploymentGroup#alarms}.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#enabled CodedeployDeploymentGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#ignore_poll_alarm_failure CodedeployDeploymentGroup#ignore_poll_alarm_failure}.
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

