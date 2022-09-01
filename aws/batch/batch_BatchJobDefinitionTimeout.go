package batch


type BatchJobDefinitionTimeout struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#attempt_duration_seconds BatchJobDefinition#attempt_duration_seconds}.
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

