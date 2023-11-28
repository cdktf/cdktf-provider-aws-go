// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorReplicationInfoListTopicReplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/msk_replicator#topics_to_replicate MskReplicator#topics_to_replicate}.
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/msk_replicator#copy_access_control_lists_for_topics MskReplicator#copy_access_control_lists_for_topics}.
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/msk_replicator#copy_topic_configurations MskReplicator#copy_topic_configurations}.
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/msk_replicator#detect_and_copy_new_topics MskReplicator#detect_and_copy_new_topics}.
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/msk_replicator#topics_to_exclude MskReplicator#topics_to_exclude}.
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

