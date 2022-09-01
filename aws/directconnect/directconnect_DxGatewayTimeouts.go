package directconnect


type DxGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dx_gateway#create DxGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dx_gateway#delete DxGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

