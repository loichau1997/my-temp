package common

const (
	HeaderXRequestID = "x-request-id"
	HeaderUserID     = "x-user-id"
	HeaderUserMeta   = "x-user-type"
	HeaderTenantID   = "x-tenant-id"
	HeaderClientKey  = "x-client-key"
)

const (
	CODE_SUCCESS             = 100
	CODE_CREATE_SUCCESSFULLY = 101
	CODE_UPDATE_SUCCESSFULLY = 102
	CODE_DELETE_SUCCESSFULLY = 103
	CODE_BAD_REQUEST         = 400
	CODE_NOT_FOUND           = 404
	CODE_SERVER_ERROR        = 500
	CODE_UNAUTHORIZED        = 401
	CODE_FORBIDDEN           = 403
)

const (
	MAPPING_IN  = "in"
	MAPPING_OUT = "out"
)

const (
	mappingJsonPrefix  = "[json]"
	mappingArrayPrefix = "[]"
)

const (
	IS_DELETE_TRASH      = "trash"
	IS_DELETE_HARD_TRASH = "hard_trash"
	REVERT               = "revert"

	SQL_DELETED_AT      = "deleted_at"
	SQL_HARD_DELETED_AT = "hard_deleted_at"
	SQL_UPDATER_ID      = "updater_id"
)

const (
	KeySeparator = "|"
)

const (
	GinParamObject = "object"
)

const (
	GinObjectSeller = "seller"
)

var GinObject = []string{GinObjectSeller}
