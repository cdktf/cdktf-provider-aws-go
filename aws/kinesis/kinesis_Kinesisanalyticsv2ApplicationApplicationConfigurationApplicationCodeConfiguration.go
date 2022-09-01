package kinesis


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#code_content_type Kinesisanalyticsv2Application#code_content_type}.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
	// code_content block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#code_content Kinesisanalyticsv2Application#code_content}
	CodeContent *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent `field:"optional" json:"codeContent" yaml:"codeContent"`
}

