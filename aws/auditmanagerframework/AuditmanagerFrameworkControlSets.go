package auditmanagerframework


type AuditmanagerFrameworkControlSets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#name AuditmanagerFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// controls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/auditmanager_framework#controls AuditmanagerFramework#controls}
	Controls interface{} `field:"optional" json:"controls" yaml:"controls"`
}

