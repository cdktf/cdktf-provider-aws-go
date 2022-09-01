package cloudfront


type CloudfrontMonitoringSubscriptionMonitoringSubscription struct {
	// realtime_metrics_subscription_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription#realtime_metrics_subscription_config CloudfrontMonitoringSubscription#realtime_metrics_subscription_config}
	RealtimeMetricsSubscriptionConfig *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig `field:"required" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

