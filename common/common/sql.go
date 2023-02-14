package common

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type DeleteRequest struct {
	ID     []uuid.UUID `json:"id,omitempty" valid:"Required"`
	Action string      `json:"action" valid:"Required"`
}

func GenRandomCode(tx *gorm.DB, model interface{}, field string, length int, upper bool) (string, error) {
	for {
		var count int64 = 0
		genCode := RandStringBytes(length, upper)

		if err := tx.Model(model).Where(field+" = ?", genCode).Count(&count).Error; err != nil {
			return "", err
		}

		if count > 0 {
			continue
		} else {
			return genCode, nil
		}
	}
}

func DeleteAction(tx *gorm.DB, req DeleteRequest, model interface{}, userId uuid.UUID) error {
	// get invalid id from request lists
	if err := checkMissingRecord(tx, model, req); err != nil {
		return err
	}

	switch req.Action {
	case IS_DELETE_TRASH:
		return deleteActionRule(tx, req.Action).Model(model).Where("id IN ?", req.ID).
			Updates(map[string]interface{}{
				SQL_DELETED_AT: gorm.Expr("CURRENT_TIMESTAMP"),
				SQL_UPDATER_ID: userId,
			}).Error
	case IS_DELETE_HARD_TRASH:
		return deleteActionRule(tx, req.Action).Model(model).Where("id IN ?", req.ID).
			Updates(map[string]interface{}{
				SQL_DELETED_AT:      gorm.Expr("CURRENT_TIMESTAMP"),
				SQL_HARD_DELETED_AT: gorm.Expr("CURRENT_TIMESTAMP"),
				SQL_UPDATER_ID:      userId,
			}).Error
	case REVERT:
		return deleteActionRule(tx, req.Action).Model(model).Where("id IN ?", req.ID).
			Updates(map[string]interface{}{
				SQL_DELETED_AT:      gorm.Expr("NULL"),
				SQL_HARD_DELETED_AT: gorm.Expr("NULL"),
				SQL_UPDATER_ID:      userId,
			}).Error
	default:
		return fmt.Errorf("invalid delete action (switch): %v", req.Action)
	}
}

func FilterDelete(tableName string, iSDelete *string, tx *gorm.DB) *gorm.DB {
	if iSDelete != nil && *iSDelete != "" {
		if tableName != "" {
			tableName = "\"" + tableName + "\"."
		}

		tx = tx.Unscoped()
		switch *iSDelete {
		case IS_DELETE_TRASH:
			tx = tx.Where(fmt.Sprintf("%sdeleted_at IS NOT NULL AND %shard_deleted_at IS NULL", tableName, tableName))
		case IS_DELETE_HARD_TRASH:
			tx = tx.Where(fmt.Sprintf("%sdeleted_at IS NOT NULL AND %shard_deleted_at IS NOT NULL", tableName, tableName))
		}
	}

	return tx
}

func CheckExist(tx *gorm.DB, model interface{}, field string, id interface{}) bool {
	var count int64 = 0
	err := tx.Model(model).Where(field+" = ?", id).Count(&count).Error
	if err != nil {
		return false
	}
	return count == 0
}

func FilterIfNotNil(destPtr interface{}, tx *gorm.DB, op func(query interface{}, args ...interface{}) (tx *gorm.DB), query interface{}, args ...interface{}) *gorm.DB {
	args = append(args, destPtr)
	switch destPtr.(type) {
	case *string:
		if destPtr != nil && destPtr.(*string) != nil {
			return op(query, args)
		}
	case *int:
		if destPtr != nil && destPtr.(*int) != nil {
			return op(query, args)
		}
	case *int8:
		if destPtr != nil && destPtr.(*int8) != nil {
			return op(query, args)
		}
	case *int16:
		if destPtr != nil && destPtr.(*int16) != nil {
			return op(query, args)
		}
	case *int32:
		if destPtr != nil && destPtr.(*int32) != nil {
			return op(query, args)
		}
	case *int64:
		if destPtr != nil && destPtr.(*int64) != nil {
			return op(query, args)
		}
	case *uint:
		if destPtr != nil && destPtr.(*uint) != nil {
			return op(query, args)
		}
	case *uint8:
		if destPtr != nil && destPtr.(*uint8) != nil {
			return op(query, args)
		}
	case *uint16:
		if destPtr != nil && destPtr.(*uint16) != nil {
			return op(query, args)
		}
	case *uint32:
		if destPtr != nil && destPtr.(*uint32) != nil {
			return op(query, args)
		}
	case *uint64:
		if destPtr != nil && destPtr.(*uint64) != nil {
			return op(query, args)
		}
	case *float32:
		if destPtr != nil && destPtr.(*float32) != nil {
			return op(query, args)
		}
	case *float64:
		if destPtr != nil && destPtr.(*float64) != nil {
			return op(query, args)
		}
	case *bool:
		if destPtr != nil && destPtr.(*bool) != nil {
			return op(query, args)
		}
	case *[]byte:
		if destPtr != nil && destPtr.(*[]byte) != nil {
			return op(query, args)
		}
	}
	return tx
}

func checkMissingRecord(tx *gorm.DB, model interface{}, req DeleteRequest) error {
	var missingIdList []string
	var subQuery1Tmp []string
	for _, v := range req.ID {
		subQuery1Tmp = append(subQuery1Tmp, fmt.Sprintf("SELECT '%v' AS id", v))
	}

	subQuery1 := tx.Raw(strings.Join(subQuery1Tmp, " UNION ALL "))
	subQuery2 := deleteActionRule(tx, req.Action).Select("id").Model(model).Where("id IN ?", req.ID)

	if err := tx.Table("(?) as id_list", subQuery1).Where("CAST(id AS uuid) NOT IN (?)", subQuery2).Pluck("id", &missingIdList).Error; err != nil {
		return fmt.Errorf("error getting missing id from delete request: %v", err)
	}

	if len(missingIdList) > 0 {
		return fmt.Errorf("record not found: id %v not exist", strings.Join(missingIdList, ", "))
	}

	return nil
}

func deleteActionRule(tx *gorm.DB, action string) *gorm.DB {
	switch action {
	case IS_DELETE_TRASH:
		return tx.Unscoped().Where("? IS NULL", gorm.Expr(SQL_DELETED_AT))
	case IS_DELETE_HARD_TRASH:
		return tx.Unscoped().Where("? IS NULL OR ? IS NULL", gorm.Expr(SQL_DELETED_AT), gorm.Expr(SQL_HARD_DELETED_AT))
	case REVERT:
		return tx.Unscoped().Where("? IS NOT NULL OR ? IS NOT NULL", gorm.Expr(SQL_DELETED_AT), gorm.Expr(SQL_HARD_DELETED_AT))
	}
	return tx
}
