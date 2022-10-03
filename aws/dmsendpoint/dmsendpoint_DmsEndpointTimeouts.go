package dmsendpoint


type DmsEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint#create DmsEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint#delete DmsEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

