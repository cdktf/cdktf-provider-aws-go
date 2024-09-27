// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mqbroker


type MqBrokerLdapServerMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#hosts MqBroker#hosts}.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#role_base MqBroker#role_base}.
	RoleBase *string `field:"optional" json:"roleBase" yaml:"roleBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#role_name MqBroker#role_name}.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#role_search_matching MqBroker#role_search_matching}.
	RoleSearchMatching *string `field:"optional" json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#role_search_subtree MqBroker#role_search_subtree}.
	RoleSearchSubtree interface{} `field:"optional" json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#service_account_password MqBroker#service_account_password}.
	ServiceAccountPassword *string `field:"optional" json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#service_account_username MqBroker#service_account_username}.
	ServiceAccountUsername *string `field:"optional" json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#user_base MqBroker#user_base}.
	UserBase *string `field:"optional" json:"userBase" yaml:"userBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#user_role_name MqBroker#user_role_name}.
	UserRoleName *string `field:"optional" json:"userRoleName" yaml:"userRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#user_search_matching MqBroker#user_search_matching}.
	UserSearchMatching *string `field:"optional" json:"userSearchMatching" yaml:"userSearchMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/mq_broker#user_search_subtree MqBroker#user_search_subtree}.
	UserSearchSubtree interface{} `field:"optional" json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

