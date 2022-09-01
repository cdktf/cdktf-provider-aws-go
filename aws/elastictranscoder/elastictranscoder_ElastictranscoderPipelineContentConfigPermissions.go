package elastictranscoder


type ElastictranscoderPipelineContentConfigPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#access ElastictranscoderPipeline#access}.
	Access *[]*string `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#grantee ElastictranscoderPipeline#grantee}.
	Grantee *string `field:"optional" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline#grantee_type ElastictranscoderPipeline#grantee_type}.
	GranteeType *string `field:"optional" json:"granteeType" yaml:"granteeType"`
}

