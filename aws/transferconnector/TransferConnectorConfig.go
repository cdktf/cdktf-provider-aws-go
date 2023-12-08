// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferConnectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#access_role TransferConnector#access_role}.
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#url TransferConnector#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// as2_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#as2_config TransferConnector#as2_config}
	As2Config *TransferConnectorAs2Config `field:"optional" json:"as2Config" yaml:"as2Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#id TransferConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#logging_role TransferConnector#logging_role}.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// sftp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#sftp_config TransferConnector#sftp_config}
	SftpConfig *TransferConnectorSftpConfig `field:"optional" json:"sftpConfig" yaml:"sftpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#tags TransferConnector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/transfer_connector#tags_all TransferConnector#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

