package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParametersOauthOauthHttpParametersBody struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_connection#is_value_secret CloudwatchEventConnection#is_value_secret}.
	IsValueSecret interface{} `field:"optional" json:"isValueSecret" yaml:"isValueSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_connection#key CloudwatchEventConnection#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_connection#value CloudwatchEventConnection#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

