package launchtemplate


type LaunchTemplateCpuOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/launch_template#amd_sev_snp LaunchTemplate#amd_sev_snp}.
	AmdSevSnp *string `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/launch_template#core_count LaunchTemplate#core_count}.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/launch_template#threads_per_core LaunchTemplate#threads_per_core}.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

