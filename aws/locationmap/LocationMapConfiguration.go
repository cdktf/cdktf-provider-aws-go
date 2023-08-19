package locationmap


type LocationMapConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/location_map#style LocationMap#style}.
	Style *string `field:"required" json:"style" yaml:"style"`
}

