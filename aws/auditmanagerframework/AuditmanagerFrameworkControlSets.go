package auditmanagerframework


type AuditmanagerFrameworkControlSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/auditmanager_framework#name AuditmanagerFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/auditmanager_framework#controls AuditmanagerFramework#controls}
	Controls interface{} `field:"optional" json:"controls" yaml:"controls"`
}

