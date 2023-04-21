package dataawsdmsendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDmsEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#endpoint_id DataAwsDmsEndpoint#endpoint_id}.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// elasticsearch_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#elasticsearch_settings DataAwsDmsEndpoint#elasticsearch_settings}
	ElasticsearchSettings interface{} `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#extra_connection_attributes DataAwsDmsEndpoint#extra_connection_attributes}.
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#id DataAwsDmsEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kafka_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#kafka_settings DataAwsDmsEndpoint#kafka_settings}
	KafkaSettings interface{} `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// mongodb_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#mongodb_settings DataAwsDmsEndpoint#mongodb_settings}
	MongodbSettings interface{} `field:"optional" json:"mongodbSettings" yaml:"mongodbSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/dms_endpoint#tags DataAwsDmsEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

