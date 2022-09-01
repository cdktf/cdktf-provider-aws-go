package glue


type GlueTriggerActionsNotificationProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#notify_delay_after GlueTrigger#notify_delay_after}.
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

