package appmesh


type AppmeshVirtualNodeSpecServiceDiscovery struct {
	// aws_cloud_map block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#aws_cloud_map AppmeshVirtualNode#aws_cloud_map}
	AwsCloudMap *AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#dns AppmeshVirtualNode#dns}
	Dns *AppmeshVirtualNodeSpecServiceDiscoveryDns `field:"optional" json:"dns" yaml:"dns"`
}

