package xray


type XrayGroupInsightsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/xray_group#insights_enabled XrayGroup#insights_enabled}.
	InsightsEnabled interface{} `field:"required" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/xray_group#notifications_enabled XrayGroup#notifications_enabled}.
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

