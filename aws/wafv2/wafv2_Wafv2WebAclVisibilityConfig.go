package wafv2


type Wafv2WebAclVisibilityConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#cloudwatch_metrics_enabled Wafv2WebAcl#cloudwatch_metrics_enabled}.
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#metric_name Wafv2WebAcl#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#sampled_requests_enabled Wafv2WebAcl#sampled_requests_enabled}.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}
