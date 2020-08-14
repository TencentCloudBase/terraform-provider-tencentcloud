package tencentcloud

const (
	SQLSERVER_CHARGE_TYPE_PREPAID  = "PREPAID"
	SQLSERVER_CHARGE_TYPE_POSTPAID = "POSTPAID_BY_HOUR"
)

var SQLSERVER_CHARGE_TYPE_NAME = map[string]string{
	"PRE":  SQLSERVER_CHARGE_TYPE_PREPAID,
	"POST": SQLSERVER_CHARGE_TYPE_POSTPAID,
	"ALL":  "ALL",
}

const (
	SQLSERVER_TASK_SUCCESS = 0
	SQLSERVER_TASK_FAIL    = 1
	SQLSERVER_TASK_RUNNING = 2
)

var SQLSERVER_CHARSET_LIST = []string{
	"Chinese_PRC_CI_AS",
	"Chinese_PRC_CS_AS",
	"Chinese_PRC_BIN",
	"Chinese_Taiwan_Stroke_CI_AS",
	"SQL_Latin1_General_CP1_CI_AS",
	"SQL_Latin1_General_CP1_CS_AS",
}

const (
	SQLSERVER_ACCOUNT_RW = "ReadWrite"
	SQLSERVER_ACCOUNT_RO = "ReadOnly"
)

var SQLSERVER_ACCOUNT_PRIVILEGE = []string{
	SQLSERVER_ACCOUNT_RW,
	SQLSERVER_ACCOUNT_RO,
}

const (
	SQLSERVER_DEFAULT_LIMIT  = 20
	SQLSERVER_DEFAULT_OFFSET = 0
)

const (
	SQLSERVER_DB_CREATING  = 1
	SQLSERVER_DB_RUNNING   = 2
	SQLSERVER_DB_MODIFYING = 3
	SQLSERVER_DB_DELETING  = -1
)

var SQLSERVER_DB_STATUS = map[int64]string{
	SQLSERVER_DB_CREATING:  "creating",
	SQLSERVER_DB_RUNNING:   "running",
	SQLSERVER_DB_MODIFYING: "modifying",
	SQLSERVER_DB_DELETING:  "deleting",
}

const (
	SQLSERVER_HA_FLAG_SINGLE  = "SINGLE"
	SQLSERVER_HA_FLAG_DAUL    = "MIRROR"
	SQLSERVER_HA_FLAG_CLUSTER = "ALWAYSON"
)

var SQLSERVER_HA_TYPE_FLAGS = map[string]string{
	SQLSERVER_HA_FLAG_SINGLE:  "SINGLE",
	SQLSERVER_HA_FLAG_DAUL:    "DUAL",
	SQLSERVER_HA_FLAG_CLUSTER: "CLUSTER",
}

const (
	BASIC_NETWORK = 0
	VPC_NEWORK    = 1
)
