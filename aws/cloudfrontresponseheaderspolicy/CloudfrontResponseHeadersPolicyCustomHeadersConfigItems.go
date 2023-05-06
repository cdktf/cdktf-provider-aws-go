package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyCustomHeadersConfigItems struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudfront_response_headers_policy#header CloudfrontResponseHeadersPolicy#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudfront_response_headers_policy#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/cloudfront_response_headers_policy#value CloudfrontResponseHeadersPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

