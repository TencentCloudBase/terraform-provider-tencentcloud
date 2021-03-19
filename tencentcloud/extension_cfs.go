package tencentcloud

const (
	CFS_PROTOCOL_NFS  = "NFS"
	CFS_PROTOCOL_CIFS = "CIFS"

	CFS_STORAGETYPE_SD = "SD"
	CFS_STORAGETYPE_HP = "HP"

	CFS_FILE_SYSTEM_STATUS_CREATING = "creating"
	CFS_FILE_SYSTEM_STATUS_SUCCESS  = "available"
	CFS_FILE_SYSTEM_STATUS_FAILED   = "create_failed"

	CFS_RW_PERMISSION_RO = "RO"
	CFS_RW_PERMISSION_RW = "RW"

	CFS_USER_PERMISSION_ALL_SQUASH     = "all_squash"
	CFS_USER_PERMISSION_NO_ALL_SQUASH  = "no_all_squash"
	CFS_USER_PERMISSION_ROOT_SQUASH    = "root_squash"
	CFS_USER_PERMISSION_NO_ROOT_SQUASH = "no_root_squash"
)

var CFS_STORAGETYPE = []string{
	CFS_STORAGETYPE_SD,
	CFS_STORAGETYPE_HP,
}

var CFS_PROTOCOL = []string{
	CFS_PROTOCOL_NFS,
	CFS_PROTOCOL_CIFS,
}

var CFS_RW_PERMISSION = []string{
	CFS_RW_PERMISSION_RO,
	CFS_RW_PERMISSION_RW,
}

var CFS_USER_PERMISSION = []string{
	CFS_USER_PERMISSION_ALL_SQUASH,
	CFS_USER_PERMISSION_NO_ALL_SQUASH,
	CFS_USER_PERMISSION_ROOT_SQUASH,
	CFS_USER_PERMISSION_NO_ROOT_SQUASH,
}

const CfsInvalidPgroup = "InvalidParameterValue.InvalidPgroup"
