package appstream


type AppstreamImageBuilderAccessEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#endpoint_type AppstreamImageBuilder#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#vpce_id AppstreamImageBuilder#vpce_id}.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

