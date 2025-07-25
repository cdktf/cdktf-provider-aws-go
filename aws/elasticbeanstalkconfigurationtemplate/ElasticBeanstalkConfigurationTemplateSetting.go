// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkconfigurationtemplate


type ElasticBeanstalkConfigurationTemplateSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/elastic_beanstalk_configuration_template#name ElasticBeanstalkConfigurationTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/elastic_beanstalk_configuration_template#namespace ElasticBeanstalkConfigurationTemplate#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/elastic_beanstalk_configuration_template#value ElasticBeanstalkConfigurationTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/elastic_beanstalk_configuration_template#resource ElasticBeanstalkConfigurationTemplate#resource}.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

