package kinesis


type Kinesisanalyticsv2ApplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#create Kinesisanalyticsv2Application#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#delete Kinesisanalyticsv2Application#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#update Kinesisanalyticsv2Application#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

