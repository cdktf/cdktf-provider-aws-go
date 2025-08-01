// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticBeanstalkApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#name ElasticBeanstalkApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// appversion_lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#appversion_lifecycle ElasticBeanstalkApplication#appversion_lifecycle}
	AppversionLifecycle *ElasticBeanstalkApplicationAppversionLifecycle `field:"optional" json:"appversionLifecycle" yaml:"appversionLifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#description ElasticBeanstalkApplication#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#id ElasticBeanstalkApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#region ElasticBeanstalkApplication#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#tags ElasticBeanstalkApplication#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/elastic_beanstalk_application#tags_all ElasticBeanstalkApplication#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

