package elasticsearch


type ElasticsearchDomainAutoTuneOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#desired_state ElasticsearchDomain#desired_state}.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#maintenance_schedule ElasticsearchDomain#maintenance_schedule}
	MaintenanceSchedule interface{} `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#rollback_on_disable ElasticsearchDomain#rollback_on_disable}.
	RollbackOnDisable *string `field:"optional" json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
}

