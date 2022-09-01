package apprunner


type ApprunnerServiceInstanceConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#cpu ApprunnerService#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#instance_role_arn ApprunnerService#instance_role_arn}.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#memory ApprunnerService#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

