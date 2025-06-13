// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeyvaluestorekeysexclusive


type CloudfrontkeyvaluestoreKeysExclusiveResourceKeyValuePair struct {
	// The key to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#key CloudfrontkeyvaluestoreKeysExclusive#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/cloudfrontkeyvaluestore_keys_exclusive#value CloudfrontkeyvaluestoreKeysExclusive#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

