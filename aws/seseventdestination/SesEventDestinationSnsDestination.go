package seseventdestination


type SesEventDestinationSnsDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ses_event_destination#topic_arn SesEventDestination#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

