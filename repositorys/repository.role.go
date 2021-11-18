package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/golang-pos/models"
	"github.com/restuwahyu13/golang-pos/schemas"
)

type repositoryRole struct {
	db *gorm.DB
}

func NewRepositoryRole(db *gorm.DB) *repositoryRole {
	return &repositoryRole{db: db}
}

/**
* ==========================================
* Repository Create New Role Teritory
*===========================================
 */

func (r *repositoryRole) EntityCreate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role models.ModelRole
	role.RoleName = models.RoleAllowed(input.RoleName)
	role.RoleAccess = input.RoleAccess

	db := r.db.Model(&role)

	err := make(chan schemas.SchemaDatabaseError, 1)

	checkroleName := db.Debug().First(&role, "role_name = ?", input.RoleName)

	if checkroleName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &role, <-err
	}

	addRole := db.Debug().Create(&role).Commit()

	if addRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}

/**
* ==========================================
* Repository Results All Role Teritory
*===========================================
 */

func (r *repositoryRole) EntityResults() (*[]models.ModelRole, schemas.SchemaDatabaseError) {
	var role []models.ModelRole

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleName := db.Debug().Order("created_at DESC").Find(&role)

	if checkRoleName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryRole) EntityResult(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role models.ModelRole
	role.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleId := db.Debug().First(&role)

	if checkRoleId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryRole) EntityDelete(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role models.ModelRole
	role.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleId := db.Debug().First(&role)

	if checkRoleId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &role, <-err
	}

	deleterole := db.Debug().Delete(&role)

	if deleterole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryRole) EntityUpdate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError) {
	var role models.ModelRole
	role.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleId := db.Debug().First(&role)

	if checkRoleId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &role, <-err
	}

	role.RoleName = models.RoleAllowed(input.RoleName)
	role.RoleAccess = input.RoleAccess

	updateRole := db.Debug().Updates(&role)

	if updateRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}
