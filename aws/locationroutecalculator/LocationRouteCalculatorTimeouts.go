// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationroutecalculator


type LocationRouteCalculatorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/location_route_calculator#create LocationRouteCalculator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/location_route_calculator#delete LocationRouteCalculator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/location_route_calculator#update LocationRouteCalculator#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

