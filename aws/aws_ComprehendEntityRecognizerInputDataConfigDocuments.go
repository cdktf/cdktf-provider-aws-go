// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type ComprehendEntityRecognizerInputDataConfigDocuments struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#s3_uri ComprehendEntityRecognizer#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#input_format ComprehendEntityRecognizer#input_format}.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#test_s3_uri ComprehendEntityRecognizer#test_s3_uri}.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

