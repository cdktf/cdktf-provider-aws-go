package ssoadmincustomermanagedpolicyattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoadminCustomerManagedPolicyAttachmentConfig struct {
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
	// customer_managed_policy_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ssoadmin_customer_managed_policy_attachment#customer_managed_policy_reference SsoadminCustomerManagedPolicyAttachment#customer_managed_policy_reference}
	CustomerManagedPolicyReference *SsoadminCustomerManagedPolicyAttachmentCustomerManagedPolicyReference `field:"required" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ssoadmin_customer_managed_policy_attachment#instance_arn SsoadminCustomerManagedPolicyAttachment#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ssoadmin_customer_managed_policy_attachment#permission_set_arn SsoadminCustomerManagedPolicyAttachment#permission_set_arn}.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ssoadmin_customer_managed_policy_attachment#id SsoadminCustomerManagedPolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

