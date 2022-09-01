package elastictranscoder


type ElastictranscoderPipelineNotifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#completed ElastictranscoderPipeline#completed}.
	Completed *string `field:"optional" json:"completed" yaml:"completed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#error ElastictranscoderPipeline#error}.
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#progressing ElastictranscoderPipeline#progressing}.
	Progressing *string `field:"optional" json:"progressing" yaml:"progressing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#warning ElastictranscoderPipeline#warning}.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
}

