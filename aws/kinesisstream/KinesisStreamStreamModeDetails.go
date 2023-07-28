package kinesisstream


type KinesisStreamStreamModeDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/kinesis_stream#stream_mode KinesisStream#stream_mode}.
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

