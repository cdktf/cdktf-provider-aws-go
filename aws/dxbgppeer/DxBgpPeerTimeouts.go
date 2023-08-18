package dxbgppeer


type DxBgpPeerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/dx_bgp_peer#create DxBgpPeer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/dx_bgp_peer#delete DxBgpPeer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

