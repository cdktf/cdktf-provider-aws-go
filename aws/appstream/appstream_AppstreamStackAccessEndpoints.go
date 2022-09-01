package appstream


type AppstreamStackAccessEndpoints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#endpoint_type AppstreamStack#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#vpce_id AppstreamStack#vpce_id}.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

