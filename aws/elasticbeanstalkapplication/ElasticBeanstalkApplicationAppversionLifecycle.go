// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication


type ElasticBeanstalkApplicationAppversionLifecycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/elastic_beanstalk_application#service_role ElasticBeanstalkApplication#service_role}.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/elastic_beanstalk_application#delete_source_from_s3 ElasticBeanstalkApplication#delete_source_from_s3}.
	DeleteSourceFromS3 interface{} `field:"optional" json:"deleteSourceFromS3" yaml:"deleteSourceFromS3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/elastic_beanstalk_application#max_age_in_days ElasticBeanstalkApplication#max_age_in_days}.
	MaxAgeInDays *float64 `field:"optional" json:"maxAgeInDays" yaml:"maxAgeInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/elastic_beanstalk_application#max_count ElasticBeanstalkApplication#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
}

