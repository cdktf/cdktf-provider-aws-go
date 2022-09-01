package sagemaker


type SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#csv_content_types SagemakerEndpointConfiguration#csv_content_types}.
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#json_content_types SagemakerEndpointConfiguration#json_content_types}.
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

