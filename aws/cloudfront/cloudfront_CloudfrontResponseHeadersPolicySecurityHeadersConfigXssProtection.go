package cloudfront


type CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy#protection CloudfrontResponseHeadersPolicy#protection}.
	Protection interface{} `field:"required" json:"protection" yaml:"protection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy#mode_block CloudfrontResponseHeadersPolicy#mode_block}.
	ModeBlock interface{} `field:"optional" json:"modeBlock" yaml:"modeBlock"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy#report_uri CloudfrontResponseHeadersPolicy#report_uri}.
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
}

