package mqbroker


type MqBrokerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#create MqBroker#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#delete MqBroker#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#update MqBroker#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

