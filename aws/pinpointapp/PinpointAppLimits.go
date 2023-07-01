package pinpointapp


type PinpointAppLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/pinpoint_app#daily PinpointApp#daily}.
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/pinpoint_app#maximum_duration PinpointApp#maximum_duration}.
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/pinpoint_app#messages_per_second PinpointApp#messages_per_second}.
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/pinpoint_app#total PinpointApp#total}.
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

