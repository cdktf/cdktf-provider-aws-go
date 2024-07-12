// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/medialive_channel#id MedialiveChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// media_package_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/medialive_channel#media_package_settings MedialiveChannel#media_package_settings}
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// multiplex_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/medialive_channel#multiplex_settings MedialiveChannel#multiplex_settings}
	MultiplexSettings *MedialiveChannelDestinationsMultiplexSettings `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/medialive_channel#settings MedialiveChannel#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

