package tencentcloud

const (
	SYSTEM_DISK_TYPE_LOCAL_BASIC   = "LOCAL_BASIC"
	SYSTEM_DISK_TYPE_LOCAL_SSD     = "LOCAL_SSD"
	SYSTEM_DISK_TYPE_CLOUD_BASIC   = "CLOUD_BASIC"
	SYSTEM_DISK_TYPE_CLOUD_PREMIUM = "CLOUD_PREMIUM"
	SYSTEM_DISK_TYPE_CLOUD_SSD     = "CLOUD_SSD"
)

var SYSTEM_DISK_ALLOW_TYPE = []string{
	SYSTEM_DISK_TYPE_CLOUD_PREMIUM,
	SYSTEM_DISK_TYPE_CLOUD_SSD,
}

const (
	INTERNET_CHARGE_TYPE_BANDWIDTH_PREPAID          = "BANDWIDTH_PREPAID"
	INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID_BY_HOUR   = "TRAFFIC_POSTPAID_BY_HOUR"
	INTERNET_CHARGE_TYPE_BANDWIDTH_POSTPAID_BY_HOUR = "BANDWIDTH_POSTPAID_BY_HOUR"
	INTERNET_CHARGE_TYPE_BANDWIDTH_PACKAGE          = "BANDWIDTH_PACKAGE"
)

var INTERNET_CHARGE_ALLOW_TYPE = []string{
	INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID_BY_HOUR,
}

const (
	INSTANCE_CHARGE_TYPE_POSTPAID = "POSTPAID_BY_HOUR"
	INSTANCE_CHARGE_TYPE_SPOTPAID = "SPOTPAID"
)

const (
	SCALING_GROUP_TERMINATION_POLICY_NEWEST_INSTANCE = "NEWEST_INSTANCE"
	SCALING_GROUP_TERMINATION_POLICY_OLDEST_INSTANCE = "OLDEST_INSTANCE"
)

const (
	SCALING_GROUP_RETRY_POLICY_IMMEDIATE_RETRY       = "IMMEDIATE_RETRY"
	SCALING_GROUP_RETRY_POLICY_INCREMENTAL_INTERVALS = "INCREMENTAL_INTERVALS"
)

const (
	SCALING_GROUP_ADJUSTMENT_TYPE_CHANGE_IN_CAPACITY         = "CHANGE_IN_CAPACITY"
	SCALING_GROUP_ADJUSTMENT_TYPE_EXACT_CAPACITY             = "EXACT_CAPACITY"
	SCALING_GROUP_ADJUSTMENT_TYPE_PERCENT_CHANGE_IN_CAPACITY = "PERCENT_CHANGE_IN_CAPACITY"
)

var SCALING_GROUP_ADJUSTMENT_TYPE = []string{
	SCALING_GROUP_ADJUSTMENT_TYPE_CHANGE_IN_CAPACITY,
	SCALING_GROUP_ADJUSTMENT_TYPE_EXACT_CAPACITY,
	SCALING_GROUP_ADJUSTMENT_TYPE_PERCENT_CHANGE_IN_CAPACITY,
}

const (
	SCALING_GROUP_COMPARISON_OPERATOR_GREATER       = "GREATER_THAN"
	SCALING_GROUP_COMPARISON_OPERATOR_GREATER_EQUAL = "GREATER_THAN_OR_EQUAL_TO"
	SCALING_GROUP_COMPARISON_OPERATOR_LESS          = "LESS_THAN"
	SCALING_GROUP_COMPARISON_OPERATOR_LESS_EQUAL    = "LESS_THAN_OR_EQUAL_TO"
	SCALING_GROUP_COMPARISON_OPERATOR_EQUAL         = "EQUAL_TO"
	SCALING_GROUP_COMPARISON_OPERATOR_NOT_EQUAL     = "NOT_EQUAL_TO"
)

var SCALING_GROUP_COMPARISON_OPERATOR = []string{
	SCALING_GROUP_COMPARISON_OPERATOR_GREATER,
	SCALING_GROUP_COMPARISON_OPERATOR_GREATER_EQUAL,
	SCALING_GROUP_COMPARISON_OPERATOR_LESS,
	SCALING_GROUP_COMPARISON_OPERATOR_LESS_EQUAL,
	SCALING_GROUP_COMPARISON_OPERATOR_EQUAL,
	SCALING_GROUP_COMPARISON_OPERATOR_NOT_EQUAL,
}

const (
	SCALING_GROUP_METRIC_NAME_CPU_UTILIZATION = "CPU_UTILIZATION"
	SCALING_GROUP_METRIC_NAME_MEM_UTILIZATION = "MEM_UTILIZATION"
	SCALING_GROUP_METRIC_NAME_LAN_TRAFFIC_OUT = "LAN_TRAFFIC_OUT"
	SCALING_GROUP_METRIC_NAME_LAN_TRAFFIC_IN  = "LAN_TRAFFIC_IN"
	SCALING_GROUP_METRIC_NAME_WAN_TRAFFIC_OUT = "WAN_TRAFFIC_OUT"
	SCALING_GROUP_METRIC_NAME_WAN_TRAFFIC_IN  = "WAN_TRAFFIC_IN"
)

var SCALING_GROUP_METRIC_NAME = []string{
	SCALING_GROUP_METRIC_NAME_CPU_UTILIZATION,
	SCALING_GROUP_METRIC_NAME_MEM_UTILIZATION,
	SCALING_GROUP_METRIC_NAME_LAN_TRAFFIC_OUT,
	SCALING_GROUP_METRIC_NAME_LAN_TRAFFIC_IN,
	SCALING_GROUP_METRIC_NAME_WAN_TRAFFIC_OUT,
	SCALING_GROUP_METRIC_NAME_WAN_TRAFFIC_IN,
}

const (
	SCALING_GROUP_STATISTIC_AVERAGE = "AVERAGE"
	SCALING_GROUP_STATISTIC_MAXIMUM = "MAXIMUM"
	SCALING_GROUP_STATISTIC_MINIMUM = "MINIMUM"
)

var SCALING_GROUP_STATISTIC = []string{
	SCALING_GROUP_STATISTIC_AVERAGE,
	SCALING_GROUP_STATISTIC_MAXIMUM,
	SCALING_GROUP_STATISTIC_MINIMUM,
}

const (
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_OUT_SUCCESS = "SCALE_OUT_SUCCESSFUL"
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_OUT_FAILED  = "SCALE_OUT_FAILED"
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_IN_SUCCESS  = "SCALE_IN_SUCCESSFUL"
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_IN_FAILED   = "SCALE_IN_FAILED"
	SCALING_GROUP_NOTIFICATION_TYPE_REPLACE_SUCCESS   = "REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL"
	SCALING_GROUP_NOTIFICATION_TYPE_REPLACE_FAILED    = "REPLACE_UNHEALTHY_INSTANCE_FAILED"
)

var SCALING_GROUP_NOTIFICATION_TYPE = []string{
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_OUT_SUCCESS,
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_OUT_FAILED,
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_IN_SUCCESS,
	SCALING_GROUP_NOTIFICATION_TYPE_SCALE_IN_FAILED,
	SCALING_GROUP_NOTIFICATION_TYPE_REPLACE_SUCCESS,
	SCALING_GROUP_NOTIFICATION_TYPE_REPLACE_FAILED,
}

const (
	SCALING_GROUP_ACTIVITY_STATUS_INIT                 = "INIT"
	SCALING_GROUP_ACTIVITY_STATUS_RUNNING              = "RUNNING"
	SCALING_GROUP_ACTIVITY_STATUS_SUCCESSFUL           = "SUCCESSFUL"
	SCALING_GROUP_ACTIVITY_STATUS_PARTIALLY_SUCCESSFUL = "PARTIALLY_SUCCESSFUL"
	SCALING_GROUP_ACTIVITY_STATUS_FAILED               = "FAILED"
	SCALING_GROUP_ACTIVITY_STATUS_CANCELLED            = "CANCELLED"
)
