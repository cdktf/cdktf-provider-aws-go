package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParametersBasic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/cloudwatch_event_connection#password CloudwatchEventConnection#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/cloudwatch_event_connection#username CloudwatchEventConnection#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

