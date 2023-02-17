package auditmanagerframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerFrameworkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#name AuditmanagerFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#compliance_type AuditmanagerFramework#compliance_type}.
	ComplianceType *string `field:"optional" json:"complianceType" yaml:"complianceType"`
	// control_sets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#control_sets AuditmanagerFramework#control_sets}
	ControlSets interface{} `field:"optional" json:"controlSets" yaml:"controlSets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#description AuditmanagerFramework#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#tags AuditmanagerFramework#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

