package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/emrcontainers_job_template#entry_point EmrcontainersJobTemplate#entry_point}.
	EntryPoint *string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/emrcontainers_job_template#entry_point_arguments EmrcontainersJobTemplate#entry_point_arguments}.
	EntryPointArguments *[]*string `field:"optional" json:"entryPointArguments" yaml:"entryPointArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/emrcontainers_job_template#spark_submit_parameters EmrcontainersJobTemplate#spark_submit_parameters}.
	SparkSubmitParameters *string `field:"optional" json:"sparkSubmitParameters" yaml:"sparkSubmitParameters"`
}

