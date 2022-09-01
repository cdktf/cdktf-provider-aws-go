package glue


type GlueTriggerEventBatchingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#batch_size GlueTrigger#batch_size}.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#batch_window GlueTrigger#batch_window}.
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}

