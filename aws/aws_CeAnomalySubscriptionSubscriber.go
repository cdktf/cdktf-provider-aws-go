// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeAnomalySubscriptionSubscriber struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_anomaly_subscription#address CeAnomalySubscription#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_anomaly_subscription#type CeAnomalySubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

