package gluejob


type GlueJobNotificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/glue_job#notify_delay_after GlueJob#notify_delay_after}.
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

