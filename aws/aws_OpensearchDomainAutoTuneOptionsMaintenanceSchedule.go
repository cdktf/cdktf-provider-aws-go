// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainAutoTuneOptionsMaintenanceSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#cron_expression_for_recurrence OpensearchDomain#cron_expression_for_recurrence}.
	CronExpressionForRecurrence *string `field:"required" json:"cronExpressionForRecurrence" yaml:"cronExpressionForRecurrence"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#duration OpensearchDomain#duration}
	Duration *OpensearchDomainAutoTuneOptionsMaintenanceScheduleDuration `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain#start_at OpensearchDomain#start_at}.
	StartAt *string `field:"required" json:"startAt" yaml:"startAt"`
}

