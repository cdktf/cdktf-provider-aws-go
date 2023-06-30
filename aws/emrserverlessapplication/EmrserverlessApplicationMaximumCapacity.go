package emrserverlessapplication


type EmrserverlessApplicationMaximumCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/emrserverless_application#cpu EmrserverlessApplication#cpu}.
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/emrserverless_application#memory EmrserverlessApplication#memory}.
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/emrserverless_application#disk EmrserverlessApplication#disk}.
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

