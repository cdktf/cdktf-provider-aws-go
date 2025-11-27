// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeyvaluestorekeysexclusive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontkeyvaluestoreKeysExclusiveConfig struct {
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
	// The Amazon Resource Name (ARN) of the Key Value Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#key_value_store_arn CloudfrontkeyvaluestoreKeysExclusive#key_value_store_arn}
	KeyValueStoreArn *string `field:"required" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
	// Maximum resource key values pairs that you wills update in a single API request.
	//
	// AWS has a default quota of 50 keys or a 3 MB payload, whichever is reached first
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#max_batch_size CloudfrontkeyvaluestoreKeysExclusive#max_batch_size}
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// resource_key_value_pair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#resource_key_value_pair CloudfrontkeyvaluestoreKeysExclusive#resource_key_value_pair}
	ResourceKeyValuePair interface{} `field:"optional" json:"resourceKeyValuePair" yaml:"resourceKeyValuePair"`
}

