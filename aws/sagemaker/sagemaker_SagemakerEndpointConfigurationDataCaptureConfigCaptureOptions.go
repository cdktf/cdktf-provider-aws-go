package sagemaker


type SagemakerEndpointConfigurationDataCaptureConfigCaptureOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#capture_mode SagemakerEndpointConfiguration#capture_mode}.
	CaptureMode *string `field:"required" json:"captureMode" yaml:"captureMode"`
}

