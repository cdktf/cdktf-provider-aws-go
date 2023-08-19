package elasticsearchdomain


type ElasticsearchDomainAutoTuneOptionsMaintenanceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticsearch_domain#cron_expression_for_recurrence ElasticsearchDomain#cron_expression_for_recurrence}.
	CronExpressionForRecurrence *string `field:"required" json:"cronExpressionForRecurrence" yaml:"cronExpressionForRecurrence"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticsearch_domain#duration ElasticsearchDomain#duration}
	Duration *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticsearch_domain#start_at ElasticsearchDomain#start_at}.
	StartAt *string `field:"required" json:"startAt" yaml:"startAt"`
}

