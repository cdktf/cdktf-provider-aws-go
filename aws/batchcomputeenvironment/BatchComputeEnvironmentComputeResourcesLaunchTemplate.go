package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/batch_compute_environment#launch_template_id BatchComputeEnvironment#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/batch_compute_environment#launch_template_name BatchComputeEnvironment#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/batch_compute_environment#version BatchComputeEnvironment#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

