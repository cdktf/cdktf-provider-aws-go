package cloudwatcheventtarget


type CloudwatchEventTargetKinesisTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_event_target#partition_key_path CloudwatchEventTarget#partition_key_path}.
	PartitionKeyPath *string `field:"optional" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

