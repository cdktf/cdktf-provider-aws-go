// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package invoicinginvoiceunit


type InvoicingInvoiceUnitRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/invoicing_invoice_unit#linked_accounts InvoicingInvoiceUnit#linked_accounts}.
	LinkedAccounts *[]*string `field:"required" json:"linkedAccounts" yaml:"linkedAccounts"`
}

