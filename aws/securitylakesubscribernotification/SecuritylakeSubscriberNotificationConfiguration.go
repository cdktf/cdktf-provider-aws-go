// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscribernotification


type SecuritylakeSubscriberNotificationConfiguration struct {
	// https_notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/securitylake_subscriber_notification#https_notification_configuration SecuritylakeSubscriberNotification#https_notification_configuration}
	HttpsNotificationConfiguration interface{} `field:"optional" json:"httpsNotificationConfiguration" yaml:"httpsNotificationConfiguration"`
	// sqs_notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/securitylake_subscriber_notification#sqs_notification_configuration SecuritylakeSubscriberNotification#sqs_notification_configuration}
	SqsNotificationConfiguration interface{} `field:"optional" json:"sqsNotificationConfiguration" yaml:"sqsNotificationConfiguration"`
}

