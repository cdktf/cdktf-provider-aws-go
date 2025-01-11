// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationAzureBlobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#agent_arns DatasyncLocationAzureBlob#agent_arns}.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#authentication_type DatasyncLocationAzureBlob#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#container_url DatasyncLocationAzureBlob#container_url}.
	ContainerUrl *string `field:"required" json:"containerUrl" yaml:"containerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#access_tier DatasyncLocationAzureBlob#access_tier}.
	AccessTier *string `field:"optional" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#blob_type DatasyncLocationAzureBlob#blob_type}.
	BlobType *string `field:"optional" json:"blobType" yaml:"blobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#id DatasyncLocationAzureBlob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// sas_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#sas_configuration DatasyncLocationAzureBlob#sas_configuration}
	SasConfiguration *DatasyncLocationAzureBlobSasConfiguration `field:"optional" json:"sasConfiguration" yaml:"sasConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#subdirectory DatasyncLocationAzureBlob#subdirectory}.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#tags DatasyncLocationAzureBlob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/datasync_location_azure_blob#tags_all DatasyncLocationAzureBlob#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

