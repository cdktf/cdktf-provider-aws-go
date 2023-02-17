package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityInitialCapacityConfigWorkerConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#cpu EmrserverlessApplication#cpu}.
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#memory EmrserverlessApplication#memory}.
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#disk EmrserverlessApplication#disk}.
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

