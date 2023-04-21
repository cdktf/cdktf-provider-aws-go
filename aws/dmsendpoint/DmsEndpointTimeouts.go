package dmsendpoint


type DmsEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/dms_endpoint#create DmsEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/dms_endpoint#delete DmsEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

