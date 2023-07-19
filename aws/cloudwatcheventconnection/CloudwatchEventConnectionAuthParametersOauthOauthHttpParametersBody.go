package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParametersOauthOauthHttpParametersBody struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_connection#is_value_secret CloudwatchEventConnection#is_value_secret}.
	IsValueSecret interface{} `field:"optional" json:"isValueSecret" yaml:"isValueSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_connection#key CloudwatchEventConnection#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_connection#value CloudwatchEventConnection#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

