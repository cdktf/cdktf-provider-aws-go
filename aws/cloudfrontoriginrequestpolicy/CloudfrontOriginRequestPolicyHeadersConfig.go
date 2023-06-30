package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/cloudfront_origin_request_policy#header_behavior CloudfrontOriginRequestPolicy#header_behavior}.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/cloudfront_origin_request_policy#headers CloudfrontOriginRequestPolicy#headers}
	Headers *CloudfrontOriginRequestPolicyHeadersConfigHeaders `field:"optional" json:"headers" yaml:"headers"`
}

