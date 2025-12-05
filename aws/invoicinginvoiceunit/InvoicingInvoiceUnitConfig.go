// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package invoicinginvoiceunit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InvoicingInvoiceUnitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#invoice_receiver InvoicingInvoiceUnit#invoice_receiver}.
	InvoiceReceiver *string `field:"required" json:"invoiceReceiver" yaml:"invoiceReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#name InvoicingInvoiceUnit#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#description InvoicingInvoiceUnit#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#region InvoicingInvoiceUnit#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#rule InvoicingInvoiceUnit#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#tags InvoicingInvoiceUnit#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#tax_inheritance_disabled InvoicingInvoiceUnit#tax_inheritance_disabled}.
	TaxInheritanceDisabled interface{} `field:"optional" json:"taxInheritanceDisabled" yaml:"taxInheritanceDisabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/invoicing_invoice_unit#timeouts InvoicingInvoiceUnit#timeouts}
	Timeouts *InvoicingInvoiceUnitTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

