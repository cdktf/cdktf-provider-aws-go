package vpc


type InternetGatewayAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/internet_gateway_attachment#create InternetGatewayAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/internet_gateway_attachment#delete InternetGatewayAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

