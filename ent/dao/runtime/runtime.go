// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/adminrole"
	"aisecurity/ent/dao/department"
	"aisecurity/ent/dao/employee"
	"aisecurity/ent/dao/risk"
	"aisecurity/ent/dao/riskcategory"
	"aisecurity/ent/dao/risklocation"
	"aisecurity/ent/schema"
	"aisecurity/properties"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinHooks0 := adminMixin[0].Hooks()
	admin.Hooks[0] = adminMixinHooks0[0]
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescCreatedBy is the schema descriptor for created_by field.
	adminDescCreatedBy := adminMixinFields0[1].Descriptor()
	// admin.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	admin.CreatedByValidator = adminDescCreatedBy.Validators[0].(func(int) error)
	// adminDescUpdatedBy is the schema descriptor for updated_by field.
	adminDescUpdatedBy := adminMixinFields0[3].Descriptor()
	// admin.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	admin.UpdatedByValidator = adminDescUpdatedBy.Validators[0].(func(int) error)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[4].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescUsername is the schema descriptor for username field.
	adminDescUsername := adminFields[0].Descriptor()
	// admin.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	admin.UsernameValidator = adminDescUsername.Validators[0].(func(string) error)
	// adminDescPassword is the schema descriptor for password field.
	adminDescPassword := adminFields[1].Descriptor()
	// admin.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	admin.PasswordValidator = adminDescPassword.Validators[0].(func(string) error)
	// adminDescName is the schema descriptor for name field.
	adminDescName := adminFields[2].Descriptor()
	// admin.NameValidator is a validator for the "name" field. It is called by the builders before save.
	admin.NameValidator = adminDescName.Validators[0].(func(string) error)
	adminroleFields := schema.AdminRole{}.Fields()
	_ = adminroleFields
	// adminroleDescCreatedAt is the schema descriptor for created_at field.
	adminroleDescCreatedAt := adminroleFields[0].Descriptor()
	// adminrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	adminrole.DefaultCreatedAt = adminroleDescCreatedAt.Default.(func() time.Time)
	// adminroleDescCreatedBy is the schema descriptor for created_by field.
	adminroleDescCreatedBy := adminroleFields[1].Descriptor()
	// adminrole.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	adminrole.CreatedByValidator = adminroleDescCreatedBy.Validators[0].(func(int) error)
	// adminroleDescName is the schema descriptor for name field.
	adminroleDescName := adminroleFields[2].Descriptor()
	// adminrole.NameValidator is a validator for the "name" field. It is called by the builders before save.
	adminrole.NameValidator = adminroleDescName.Validators[0].(func(string) error)
	// adminroleDescUpdatedBy is the schema descriptor for updated_by field.
	adminroleDescUpdatedBy := adminroleFields[4].Descriptor()
	// adminrole.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	adminrole.UpdatedByValidator = adminroleDescUpdatedBy.Validators[0].(func(int) error)
	// adminroleDescUpdatedAt is the schema descriptor for updated_at field.
	adminroleDescUpdatedAt := adminroleFields[5].Descriptor()
	// adminrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	adminrole.DefaultUpdatedAt = adminroleDescUpdatedAt.Default.(func() time.Time)
	// adminrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	adminrole.UpdateDefaultUpdatedAt = adminroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentFields[0].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(func() time.Time)
	// departmentDescCreatedBy is the schema descriptor for created_by field.
	departmentDescCreatedBy := departmentFields[1].Descriptor()
	// department.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	department.CreatedByValidator = departmentDescCreatedBy.Validators[0].(func(int) error)
	// departmentDescTitle is the schema descriptor for title field.
	departmentDescTitle := departmentFields[2].Descriptor()
	// department.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	department.TitleValidator = departmentDescTitle.Validators[0].(func(string) error)
	// departmentDescParentID is the schema descriptor for parent_id field.
	departmentDescParentID := departmentFields[3].Descriptor()
	// department.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	department.ParentIDValidator = departmentDescParentID.Validators[0].(func(int) error)
	// departmentDescUpdatedBy is the schema descriptor for updated_by field.
	departmentDescUpdatedBy := departmentFields[5].Descriptor()
	// department.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	department.UpdatedByValidator = departmentDescUpdatedBy.Validators[0].(func(int) error)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentFields[6].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(func() time.Time)
	// department.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	department.UpdateDefaultUpdatedAt = departmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescCreatedAt is the schema descriptor for created_at field.
	employeeDescCreatedAt := employeeFields[0].Descriptor()
	// employee.DefaultCreatedAt holds the default value on creation for the created_at field.
	employee.DefaultCreatedAt = employeeDescCreatedAt.Default.(func() time.Time)
	// employeeDescCreatedBy is the schema descriptor for created_by field.
	employeeDescCreatedBy := employeeFields[1].Descriptor()
	// employee.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	employee.CreatedByValidator = employeeDescCreatedBy.Validators[0].(func(int) error)
	// employeeDescAdminID is the schema descriptor for admin_id field.
	employeeDescAdminID := employeeFields[2].Descriptor()
	// employee.AdminIDValidator is a validator for the "admin_id" field. It is called by the builders before save.
	employee.AdminIDValidator = employeeDescAdminID.Validators[0].(func(int) error)
	// employeeDescDepartmentID is the schema descriptor for department_id field.
	employeeDescDepartmentID := employeeFields[3].Descriptor()
	// employee.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	employee.DepartmentIDValidator = employeeDescDepartmentID.Validators[0].(func(int) error)
	// employeeDescUpdatedBy is the schema descriptor for updated_by field.
	employeeDescUpdatedBy := employeeFields[5].Descriptor()
	// employee.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	employee.UpdatedByValidator = employeeDescUpdatedBy.Validators[0].(func(int) error)
	// employeeDescUpdatedAt is the schema descriptor for updated_at field.
	employeeDescUpdatedAt := employeeFields[6].Descriptor()
	// employee.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	employee.DefaultUpdatedAt = employeeDescUpdatedAt.Default.(func() time.Time)
	// employee.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	employee.UpdateDefaultUpdatedAt = employeeDescUpdatedAt.UpdateDefault.(func() time.Time)
	riskMixin := schema.Risk{}.Mixin()
	riskMixinHooks0 := riskMixin[0].Hooks()
	risk.Hooks[0] = riskMixinHooks0[0]
	riskMixinFields0 := riskMixin[0].Fields()
	_ = riskMixinFields0
	riskFields := schema.Risk{}.Fields()
	_ = riskFields
	// riskDescCreatedAt is the schema descriptor for created_at field.
	riskDescCreatedAt := riskMixinFields0[0].Descriptor()
	// risk.DefaultCreatedAt holds the default value on creation for the created_at field.
	risk.DefaultCreatedAt = riskDescCreatedAt.Default.(func() time.Time)
	// riskDescCreatedBy is the schema descriptor for created_by field.
	riskDescCreatedBy := riskMixinFields0[1].Descriptor()
	// risk.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	risk.CreatedByValidator = riskDescCreatedBy.Validators[0].(func(int) error)
	// riskDescUpdatedBy is the schema descriptor for updated_by field.
	riskDescUpdatedBy := riskMixinFields0[3].Descriptor()
	// risk.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	risk.UpdatedByValidator = riskDescUpdatedBy.Validators[0].(func(int) error)
	// riskDescUpdatedAt is the schema descriptor for updated_at field.
	riskDescUpdatedAt := riskMixinFields0[4].Descriptor()
	// risk.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	risk.DefaultUpdatedAt = riskDescUpdatedAt.Default.(func() time.Time)
	// risk.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	risk.UpdateDefaultUpdatedAt = riskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// riskDescTitle is the schema descriptor for title field.
	riskDescTitle := riskFields[0].Descriptor()
	// risk.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	risk.TitleValidator = riskDescTitle.Validators[0].(func(string) error)
	// riskDescRiskCategoryID is the schema descriptor for risk_category_id field.
	riskDescRiskCategoryID := riskFields[3].Descriptor()
	// risk.RiskCategoryIDValidator is a validator for the "risk_category_id" field. It is called by the builders before save.
	risk.RiskCategoryIDValidator = riskDescRiskCategoryID.Validators[0].(func(int) error)
	// riskDescRiskLocationID is the schema descriptor for risk_location_id field.
	riskDescRiskLocationID := riskFields[4].Descriptor()
	// risk.RiskLocationIDValidator is a validator for the "risk_location_id" field. It is called by the builders before save.
	risk.RiskLocationIDValidator = riskDescRiskLocationID.Validators[0].(func(int) error)
	// riskDescMaintainer is the schema descriptor for maintainer field.
	riskDescMaintainer := riskFields[5].Descriptor()
	// risk.MaintainerValidator is a validator for the "maintainer" field. It is called by the builders before save.
	risk.MaintainerValidator = riskDescMaintainer.Validators[0].(func(int) error)
	// riskDescMaintainStatus is the schema descriptor for maintain_status field.
	riskDescMaintainStatus := riskFields[7].Descriptor()
	// risk.DefaultMaintainStatus holds the default value on creation for the maintain_status field.
	risk.DefaultMaintainStatus = properties.MaintainStatus(riskDescMaintainStatus.Default.(int))
	// risk.MaintainStatusValidator is a validator for the "maintain_status" field. It is called by the builders before save.
	risk.MaintainStatusValidator = riskDescMaintainStatus.Validators[0].(func(int) error)
	riskcategoryMixin := schema.RiskCategory{}.Mixin()
	riskcategoryMixinHooks0 := riskcategoryMixin[0].Hooks()
	riskcategory.Hooks[0] = riskcategoryMixinHooks0[0]
	riskcategoryMixinFields0 := riskcategoryMixin[0].Fields()
	_ = riskcategoryMixinFields0
	riskcategoryFields := schema.RiskCategory{}.Fields()
	_ = riskcategoryFields
	// riskcategoryDescCreatedAt is the schema descriptor for created_at field.
	riskcategoryDescCreatedAt := riskcategoryMixinFields0[0].Descriptor()
	// riskcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	riskcategory.DefaultCreatedAt = riskcategoryDescCreatedAt.Default.(func() time.Time)
	// riskcategoryDescCreatedBy is the schema descriptor for created_by field.
	riskcategoryDescCreatedBy := riskcategoryMixinFields0[1].Descriptor()
	// riskcategory.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	riskcategory.CreatedByValidator = riskcategoryDescCreatedBy.Validators[0].(func(int) error)
	// riskcategoryDescUpdatedBy is the schema descriptor for updated_by field.
	riskcategoryDescUpdatedBy := riskcategoryMixinFields0[3].Descriptor()
	// riskcategory.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	riskcategory.UpdatedByValidator = riskcategoryDescUpdatedBy.Validators[0].(func(int) error)
	// riskcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	riskcategoryDescUpdatedAt := riskcategoryMixinFields0[4].Descriptor()
	// riskcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	riskcategory.DefaultUpdatedAt = riskcategoryDescUpdatedAt.Default.(func() time.Time)
	// riskcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	riskcategory.UpdateDefaultUpdatedAt = riskcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// riskcategoryDescTitle is the schema descriptor for title field.
	riskcategoryDescTitle := riskcategoryFields[0].Descriptor()
	// riskcategory.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	riskcategory.TitleValidator = riskcategoryDescTitle.Validators[0].(func(string) error)
	risklocationMixin := schema.RiskLocation{}.Mixin()
	risklocationMixinHooks0 := risklocationMixin[0].Hooks()
	risklocation.Hooks[0] = risklocationMixinHooks0[0]
	risklocationMixinFields0 := risklocationMixin[0].Fields()
	_ = risklocationMixinFields0
	risklocationFields := schema.RiskLocation{}.Fields()
	_ = risklocationFields
	// risklocationDescCreatedAt is the schema descriptor for created_at field.
	risklocationDescCreatedAt := risklocationMixinFields0[0].Descriptor()
	// risklocation.DefaultCreatedAt holds the default value on creation for the created_at field.
	risklocation.DefaultCreatedAt = risklocationDescCreatedAt.Default.(func() time.Time)
	// risklocationDescCreatedBy is the schema descriptor for created_by field.
	risklocationDescCreatedBy := risklocationMixinFields0[1].Descriptor()
	// risklocation.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	risklocation.CreatedByValidator = risklocationDescCreatedBy.Validators[0].(func(int) error)
	// risklocationDescUpdatedBy is the schema descriptor for updated_by field.
	risklocationDescUpdatedBy := risklocationMixinFields0[3].Descriptor()
	// risklocation.UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	risklocation.UpdatedByValidator = risklocationDescUpdatedBy.Validators[0].(func(int) error)
	// risklocationDescUpdatedAt is the schema descriptor for updated_at field.
	risklocationDescUpdatedAt := risklocationMixinFields0[4].Descriptor()
	// risklocation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	risklocation.DefaultUpdatedAt = risklocationDescUpdatedAt.Default.(func() time.Time)
	// risklocation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	risklocation.UpdateDefaultUpdatedAt = risklocationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// risklocationDescTitle is the schema descriptor for title field.
	risklocationDescTitle := risklocationFields[0].Descriptor()
	// risklocation.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	risklocation.TitleValidator = risklocationDescTitle.Validators[0].(func(string) error)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)