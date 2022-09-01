package directconnect


type DxGatewayAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dx_gateway_association#create DxGatewayAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dx_gateway_association#delete DxGatewayAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dx_gateway_association#update DxGatewayAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

