package kinesisstream


type KinesisStreamStreamModeDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_stream#stream_mode KinesisStream#stream_mode}.
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

