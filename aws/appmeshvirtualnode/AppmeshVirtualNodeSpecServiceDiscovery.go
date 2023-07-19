package appmeshvirtualnode


type AppmeshVirtualNodeSpecServiceDiscovery struct {
	// aws_cloud_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appmesh_virtual_node#aws_cloud_map AppmeshVirtualNode#aws_cloud_map}
	AwsCloudMap *AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appmesh_virtual_node#dns AppmeshVirtualNode#dns}
	Dns *AppmeshVirtualNodeSpecServiceDiscoveryDns `field:"optional" json:"dns" yaml:"dns"`
}

