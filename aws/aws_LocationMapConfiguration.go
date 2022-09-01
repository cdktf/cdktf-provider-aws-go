// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type LocationMapConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_map#style LocationMap#style}.
	Style *string `field:"required" json:"style" yaml:"style"`
}

