// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkenvironment


type ElasticBeanstalkEnvironmentSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/elastic_beanstalk_environment#name ElasticBeanstalkEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/elastic_beanstalk_environment#namespace ElasticBeanstalkEnvironment#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/elastic_beanstalk_environment#value ElasticBeanstalkEnvironment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/elastic_beanstalk_environment#resource ElasticBeanstalkEnvironment#resource}.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

