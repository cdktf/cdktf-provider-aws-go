package mqbroker


type MqBrokerUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/mq_broker#password MqBroker#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/mq_broker#username MqBroker#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/mq_broker#console_access MqBroker#console_access}.
	ConsoleAccess interface{} `field:"optional" json:"consoleAccess" yaml:"consoleAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/mq_broker#groups MqBroker#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
}

