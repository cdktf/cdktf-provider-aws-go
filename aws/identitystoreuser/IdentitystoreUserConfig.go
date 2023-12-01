// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentitystoreUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#display_name IdentitystoreUser#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#identity_store_id IdentitystoreUser#identity_store_id}.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#name IdentitystoreUser#name}
	Name *IdentitystoreUserName `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#user_name IdentitystoreUser#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#addresses IdentitystoreUser#addresses}
	Addresses *IdentitystoreUserAddresses `field:"optional" json:"addresses" yaml:"addresses"`
	// emails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#emails IdentitystoreUser#emails}
	Emails *IdentitystoreUserEmails `field:"optional" json:"emails" yaml:"emails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#id IdentitystoreUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#locale IdentitystoreUser#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#nickname IdentitystoreUser#nickname}.
	Nickname *string `field:"optional" json:"nickname" yaml:"nickname"`
	// phone_numbers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#phone_numbers IdentitystoreUser#phone_numbers}
	PhoneNumbers *IdentitystoreUserPhoneNumbers `field:"optional" json:"phoneNumbers" yaml:"phoneNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#preferred_language IdentitystoreUser#preferred_language}.
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#profile_url IdentitystoreUser#profile_url}.
	ProfileUrl *string `field:"optional" json:"profileUrl" yaml:"profileUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#timezone IdentitystoreUser#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#title IdentitystoreUser#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/identitystore_user#user_type IdentitystoreUser#user_type}.
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

