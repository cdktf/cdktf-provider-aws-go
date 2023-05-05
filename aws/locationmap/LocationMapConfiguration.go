package locationmap


type LocationMapConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/location_map#style LocationMap#style}.
	Style *string `field:"required" json:"style" yaml:"style"`
}

