package internetgatewayattachment


type InternetGatewayAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/internet_gateway_attachment#create InternetGatewayAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/internet_gateway_attachment#delete InternetGatewayAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

