package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides struct {
	// application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/emrcontainers_job_template#application_configuration EmrcontainersJobTemplate#application_configuration}
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/emrcontainers_job_template#monitoring_configuration EmrcontainersJobTemplate#monitoring_configuration}
	MonitoringConfiguration *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

