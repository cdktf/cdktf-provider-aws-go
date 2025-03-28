// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrserverlessApplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#name EmrserverlessApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#release_label EmrserverlessApplication#release_label}.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#type EmrserverlessApplication#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#architecture EmrserverlessApplication#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// auto_start_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#auto_start_configuration EmrserverlessApplication#auto_start_configuration}
	AutoStartConfiguration *EmrserverlessApplicationAutoStartConfiguration `field:"optional" json:"autoStartConfiguration" yaml:"autoStartConfiguration"`
	// auto_stop_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#auto_stop_configuration EmrserverlessApplication#auto_stop_configuration}
	AutoStopConfiguration *EmrserverlessApplicationAutoStopConfiguration `field:"optional" json:"autoStopConfiguration" yaml:"autoStopConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#id EmrserverlessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#image_configuration EmrserverlessApplication#image_configuration}
	ImageConfiguration *EmrserverlessApplicationImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// initial_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#initial_capacity EmrserverlessApplication#initial_capacity}
	InitialCapacity interface{} `field:"optional" json:"initialCapacity" yaml:"initialCapacity"`
	// interactive_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#interactive_configuration EmrserverlessApplication#interactive_configuration}
	InteractiveConfiguration *EmrserverlessApplicationInteractiveConfiguration `field:"optional" json:"interactiveConfiguration" yaml:"interactiveConfiguration"`
	// maximum_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#maximum_capacity EmrserverlessApplication#maximum_capacity}
	MaximumCapacity *EmrserverlessApplicationMaximumCapacity `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#network_configuration EmrserverlessApplication#network_configuration}
	NetworkConfiguration *EmrserverlessApplicationNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#tags EmrserverlessApplication#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/emrserverless_application#tags_all EmrserverlessApplication#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

