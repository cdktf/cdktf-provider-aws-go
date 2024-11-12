// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplicationversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticBeanstalkApplicationVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#application ElasticBeanstalkApplicationVersion#application}.
	Application *string `field:"required" json:"application" yaml:"application"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#bucket ElasticBeanstalkApplicationVersion#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#key ElasticBeanstalkApplicationVersion#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#name ElasticBeanstalkApplicationVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#description ElasticBeanstalkApplicationVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#force_delete ElasticBeanstalkApplicationVersion#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#id ElasticBeanstalkApplicationVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#process ElasticBeanstalkApplicationVersion#process}.
	Process interface{} `field:"optional" json:"process" yaml:"process"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#tags ElasticBeanstalkApplicationVersion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elastic_beanstalk_application_version#tags_all ElasticBeanstalkApplicationVersion#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

