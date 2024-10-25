// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogprovisioningartifact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogProvisioningArtifactConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#product_id ServicecatalogProvisioningArtifact#product_id}.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#accept_language ServicecatalogProvisioningArtifact#accept_language}.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#active ServicecatalogProvisioningArtifact#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#description ServicecatalogProvisioningArtifact#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#disable_template_validation ServicecatalogProvisioningArtifact#disable_template_validation}.
	DisableTemplateValidation interface{} `field:"optional" json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#guidance ServicecatalogProvisioningArtifact#guidance}.
	Guidance *string `field:"optional" json:"guidance" yaml:"guidance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#id ServicecatalogProvisioningArtifact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#name ServicecatalogProvisioningArtifact#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#template_physical_id ServicecatalogProvisioningArtifact#template_physical_id}.
	TemplatePhysicalId *string `field:"optional" json:"templatePhysicalId" yaml:"templatePhysicalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#template_url ServicecatalogProvisioningArtifact#template_url}.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#timeouts ServicecatalogProvisioningArtifact#timeouts}
	Timeouts *ServicecatalogProvisioningArtifactTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/servicecatalog_provisioning_artifact#type ServicecatalogProvisioningArtifact#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

