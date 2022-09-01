package mq


type MqBrokerLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#audit MqBroker#audit}.
	Audit *string `field:"optional" json:"audit" yaml:"audit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#general MqBroker#general}.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

