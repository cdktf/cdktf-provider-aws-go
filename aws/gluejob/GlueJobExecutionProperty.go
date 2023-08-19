package gluejob


type GlueJobExecutionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_job#max_concurrent_runs GlueJob#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
}

