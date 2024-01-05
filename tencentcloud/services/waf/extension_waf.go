package waf

const (
	CUSTOM_RULE_ACTION_TYPE_1 = "1"
	CUSTOM_RULE_ACTION_TYPE_2 = "2"
	CUSTOM_RULE_ACTION_TYPE_3 = "3"
	CUSTOM_RULE_ACTION_TYPE_4 = "4"
)

var CUSTOM_RULE_ACTION_TYPE = []string{
	CUSTOM_RULE_ACTION_TYPE_1,
	CUSTOM_RULE_ACTION_TYPE_2,
	CUSTOM_RULE_ACTION_TYPE_3,
	CUSTOM_RULE_ACTION_TYPE_4,
}

const (
	CUSTOM_RULE_STATUS_0     = "0"
	CUSTOM_RULE_STATUS_1     = "1"
	CUSTOM_RULE_STATUS_0_INT = 0
	CUSTOM_RULE_STATUS_1_INT = 1
)

var CUSTOM_RULE_STATUS = []string{
	CUSTOM_RULE_STATUS_0,
	CUSTOM_RULE_STATUS_1,
}

const (
	CUSTOM_WHITE_RULE_STATUS_0     = "0"
	CUSTOM_WHITE_RULE_STATUS_1     = "1"
	CUSTOM_WHITE_RULE_STATUS_0_INT = 0
	CUSTOM_WHITE_RULE_STATUS_1_INT = 1
)

var CUSTOM_WHITE_RULE_STATUS = []string{
	CUSTOM_WHITE_RULE_STATUS_0,
	CUSTOM_WHITE_RULE_STATUS_1,
}

const (
	CLB_DOMAIN_STATUS_0 = 0
	CLB_DOMAIN_STATUS_1 = 1
)

var CLB_DOMAIN_STATUS = []int{
	CLB_DOMAIN_STATUS_0,
	CLB_DOMAIN_STATUS_1,
}

const (
	CLB_DOMAIN_ENGINE_10 = 10
	CLB_DOMAIN_ENGINE_11 = 11
	CLB_DOMAIN_ENGINE_12 = 12
	CLB_DOMAIN_ENGINE_20 = 20
	CLB_DOMAIN_ENGINE_21 = 21
	CLB_DOMAIN_ENGINE_22 = 22
)

var CLB_DOMAIN_ENGINE = []int{
	CLB_DOMAIN_ENGINE_10,
	CLB_DOMAIN_ENGINE_11,
	CLB_DOMAIN_ENGINE_12,
	CLB_DOMAIN_ENGINE_20,
	CLB_DOMAIN_ENGINE_21,
	CLB_DOMAIN_ENGINE_22,
}

const (
	ISCDN_0 = 0
	ISCDN_1 = 1
	ISCDN_2 = 2
	ISCDN_3 = 3
)

var ISCDN_STSTUS = []int{
	ISCDN_0,
	ISCDN_1,
	ISCDN_2,
	ISCDN_3,
}

const (
	FLOW_MODE_0 = 0
	FLOW_MODE_1 = 1
)

var FLOW_MODE_STATUS = []int{
	FLOW_MODE_0,
	FLOW_MODE_1,
}

const (
	CLS_STATUS_0 = 0
	CLS_STATUS_1 = 1
)

var CLS_STATUS = []int{
	CLS_STATUS_0,
	CLS_STATUS_1,
}

const (
	BOT_STATUS_0 = 0
	BOT_STATUS_1 = 1
	BOT_STATUS_2 = 2
	BOT_STATUS_3 = 3
)

var BOT_STATUS = []int{
	BOT_STATUS_0,
	BOT_STATUS_1,
}

const (
	API_SAFE_STATUS_0 = 0
	API_SAFE_STATUS_1 = 1
)

var API_SAFE_STATUS = []int{
	API_SAFE_STATUS_0,
	API_SAFE_STATUS_1,
}

const (
	PROTECTION_STATUS_0 = 0
	PROTECTION_STATUS_1 = 1
)

var PROTECTION_STATUS = []int{
	PROTECTION_STATUS_0,
	PROTECTION_STATUS_1,
}

const (
	IPV6_ON  = 1
	IPV6_OFF = 2
)

const (
	ALB_TYPE_CLB    = "clb"
	ALB_TYPE_APISIX = "apisix"
	ALB_TYPE_TSEGW  = "tsegw"
)

var ALB_TYPES = []string{
	ALB_TYPE_CLB,
	ALB_TYPE_APISIX,
	ALB_TYPE_TSEGW,
}

const (
	DescribeHostLimitSuccess = "Success"
)

const (
	CERT_TYPE_0 = 0
	CERT_TYPE_1 = 1
	CERT_TYPE_2 = 2
)

var CERT_TYPES = []int{
	CERT_TYPE_0,
	CERT_TYPE_1,
	CERT_TYPE_2,
}

const (
	UPSTREAM_SCHEME_HTTP  = "http"
	UPSTREAM_SCHEME_HTTPS = "https"
)

var UPSTREAM_SCHEMES = []string{
	UPSTREAM_SCHEME_HTTP,
	UPSTREAM_SCHEME_HTTPS,
}

const (
	IS_GRAY_0 = 0
	IS_GRAY_1 = 1
)

var IS_GRAY_STATUS = []int{
	IS_GRAY_0,
	IS_GRAY_1,
}

const (
	UP_STREAM_TYPE_0 = 0
	UP_STREAM_TYPE_1 = 1
)

var UP_STREAM_TYPES = []int{
	UP_STREAM_TYPE_0,
	UP_STREAM_TYPE_1,
}

const (
	IS_HTTP2_0 = 0
	IS_HTTP2_1 = 1
)

var IS_HTTP2_STATUS = []int{
	IS_HTTP2_0,
	IS_HTTP2_1,
}

const (
	IS_WEBSOCKET_0 = 0
	IS_WEBSOCKET_1 = 1
)

var IS_WEBSOCKET_STATUS = []int{
	IS_WEBSOCKET_0,
	IS_WEBSOCKET_1,
}

const (
	LOAD_BALANCE_0 = "0"
	LOAD_BALANCE_1 = "1"
	LOAD_BALANCE_2 = "2"
)

var LOAD_BALANCE_STATUS = []string{
	LOAD_BALANCE_0,
	LOAD_BALANCE_1,
	LOAD_BALANCE_2,
}

const (
	HTTPS_REWRITE_0 = 0
	HTTPS_REWRITE_1 = 1
)

var HTTPS_REWRITE_STATUS = []int{
	HTTPS_REWRITE_0,
	HTTPS_REWRITE_1,
}

const (
	ANYCAST_0 = 0
	ANYCAST_1 = 1
)

var ANYCAST_STATUS = []int{
	ANYCAST_0,
	ANYCAST_1,
}

const (
	ACTIVE_CHECK_0 = 0
	ACTIVE_CHECK_1 = 1
)

var ACTIVE_CHECK_STATUS = []int{
	ACTIVE_CHECK_0,
	ACTIVE_CHECK_1,
}

const (
	CIPHER_TEMPLATE_1 = 1
	CIPHER_TEMPLATE_2 = 2
	CIPHER_TEMPLATE_3 = 3
)

var CIPHER_TEMPLATES = []int{
	CIPHER_TEMPLATE_1,
	CIPHER_TEMPLATE_2,
	CIPHER_TEMPLATE_3,
}

const (
	PROXY_READ_TIMEOUT = 300
	PROXY_SEND_TIMEOUT = 300
)

const (
	SNI_TYPE_0 = 0
	SNI_TYPE_1 = 1
	SNI_TYPE_2 = 2
	SNI_TYPE_3 = 3
)

var SNI_TYPES = []int{
	SNI_TYPE_0,
	SNI_TYPE_1,
	SNI_TYPE_2,
	SNI_TYPE_3,
}

const (
	XFF_RESET_0 = 0
	XFF_RESET_1 = 1
)

var XFF_RESET_STATUS = []int{
	XFF_RESET_0,
	XFF_RESET_1,
}

const (
	DescribeDomainVerifyResultSUCCESS = 0
)

const (
	IPV6_STATUS_0 = 0
	IPV6_STATUS_1 = 1
)

var IPV6_STATUS = []int{
	IPV6_STATUS_0,
	IPV6_STATUS_1,
}

const (
	SAAS_DOMAIN_STATUS_0 = 0
	SAAS_DOMAIN_STATUS_1 = 1
)

var SAAS_DOMAIN_STATUS = []int{
	SAAS_DOMAIN_STATUS_0,
	SAAS_DOMAIN_STATUS_1,
}

const (
	TLS_VERSION_STATUS_3 = 3
)

const (
	IS_KEEP_ALIVE_0 = "0"
	IS_KEEP_ALIVE_1 = "1"
)

var IS_KEEP_ALIVE_STATUS = []string{
	IS_KEEP_ALIVE_0,
	IS_KEEP_ALIVE_1,
}

const (
	ORDER_ASC  = "asc"
	ORDER_DESC = "desc"
)

const (
	WAF_PREMIUM_SAAS    = "premium_saas"
	WAF_ENTERPRISE_SAAS = "enterprise_saas"
	WAF_ULTIMATE_SAAS   = "ultimate_saas"
)

var WAF_CATEGORY_SAAS = []string{
	WAF_PREMIUM_SAAS,
	WAF_ENTERPRISE_SAAS,
	WAF_ULTIMATE_SAAS,
}

const (
	WAF_PREMIUM_CLB    = "premium_clb"
	WAF_ENTERPRISE_CLB = "enterprise_clb"
	WAF_ULTIMATE_CLB   = "ultimate_clb"
)

var WAF_CATEGORY_CLB = []string{
	WAF_PREMIUM_CLB,
	WAF_ENTERPRISE_CLB,
	WAF_ULTIMATE_CLB,
}

const (
	TIME_UINT_D = "d"
	TIME_UINT_M = "m"
	TIME_UINT_Y = "y"
)

var TIME_UNIT = []string{
	TIME_UINT_D,
	TIME_UINT_M,
	TIME_UINT_Y,
}

const (
	REGION_ID_MAINLAND     = 1
	REGION_ID_NON_MAINLAND = 9
)

var WAF_CATEGORY_ID_SAAS = map[string]int{
	WAF_PREMIUM_SAAS:    102375,
	WAF_ENTERPRISE_SAAS: 102378,
	WAF_ULTIMATE_SAAS:   102369,
}

var SUB_PRODUCT_CODE_SAAS = map[string]string{
	WAF_PREMIUM_SAAS:    "sp_wsm_waf_premium",
	WAF_ENTERPRISE_SAAS: "sp_wsm_waf_enterprise",
	WAF_ULTIMATE_SAAS:   "sp_wsm_waf_ultimate",
}

var PID_SAAS = map[string]int{
	WAF_PREMIUM_SAAS:    1000827,
	WAF_ENTERPRISE_SAAS: 1000830,
	WAF_ULTIMATE_SAAS:   1000832,
}

var PKG_SAAS = map[string]int{
	"DOMAIN": 1000834,
	"QPS":    1000481,
}

var LABEL_TYPES_SAAS = map[string]string{
	WAF_PREMIUM_SAAS:    "sv_wsm_waf_package_premium",
	WAF_ENTERPRISE_SAAS: "sv_wsm_waf_package_enterprise",
	WAF_ULTIMATE_SAAS:   "sv_wsm_waf_package_ultimate",
}

var LABEL_TYPES_CLB = map[string]string{
	WAF_PREMIUM_CLB:    "sv_wsm_waf_package_premium_clb",
	WAF_ENTERPRISE_CLB: "sv_wsm_waf_package_enterprise_clb",
	WAF_ULTIMATE_CLB:   "sv_wsm_waf_package_ultimate_clb",
}

var BOT_MANAGEMENT_LABEL_TYPES_CLB = map[int]string{
	REGION_ID_MAINLAND:     "sv_wsm_waf_scene_bot_protection_clb",
	REGION_ID_NON_MAINLAND: "sv_wsm_waf_scene_bot_protection_clb_intl",
}

var BOT_MANAGEMENT_LABEL_TYPES_SAAS = map[int]string{
	REGION_ID_MAINLAND:     "sv_wsm_waf_scene_bot_protection",
	REGION_ID_NON_MAINLAND: "sv_wsm_waf_scene_bot_protection_intl",
}

var API_SECURITY_LABEL_TYPES_CLB_REGION1 = map[string]string{
	WAF_PREMIUM_CLB:    "sv_wsm_waf_scene_cpre",
	WAF_ENTERPRISE_CLB: "sv_wsm_waf_scene_cent",
	WAF_ULTIMATE_CLB:   "sv_wsm_waf_scene_cult",
}

var API_SECURITY_LABEL_TYPES_CLB_REGION9 = map[string]string{
	WAF_PREMIUM_CLB:    "sv_wsm_waf_scene_cipre",
	WAF_ENTERPRISE_CLB: "sv_wsm_waf_scene_cient",
	WAF_ULTIMATE_CLB:   "sv_wsm_waf_scene_ciult",
}

var API_SECURITY_LABEL_TYPES_SAAS_REGION1 = map[string]string{
	WAF_PREMIUM_SAAS:    "sv_wsm_waf_scene_pre",
	WAF_ENTERPRISE_SAAS: "sv_wsm_waf_scene_ent",
	WAF_ULTIMATE_SAAS:   "sv_wsm_waf_scene_ult",
}

var API_SECURITY_LABEL_TYPES_SAAS_REGION9 = map[string]string{
	WAF_PREMIUM_SAAS:    "sv_wsm_waf_scene_ipre",
	WAF_ENTERPRISE_SAAS: "sv_wsm_waf_scene_ient",
	WAF_ULTIMATE_SAAS:   "sv_wsm_waf_scene_iult",
}

var WAF_CATEGORY_ID_CLB = map[string]int{
	WAF_PREMIUM_CLB:    101198,
	WAF_ENTERPRISE_CLB: 101204,
	WAF_ULTIMATE_CLB:   101201,
}

var SUB_PRODUCT_CODE_CLB = map[string]string{
	WAF_PREMIUM_CLB:    "sp_wsm_waf_premium_clb",
	WAF_ENTERPRISE_CLB: "sp_wsm_waf_enterprise_clb",
	WAF_ULTIMATE_CLB:   "sp_wsm_waf_ultimate_clb",
}

var PID_CLB = map[string]int{
	WAF_PREMIUM_CLB:    1001150,
	WAF_ENTERPRISE_CLB: 1001152,
	WAF_ULTIMATE_CLB:   1001154,
}

const (
	AUTO_RENEW_FLAG_0 = 0
	AUTO_RENEW_FLAG_1 = 1
)

var AUTO_RENEW_FLAG = []int{
	AUTO_RENEW_FLAG_0,
	AUTO_RENEW_FLAG_1,
}

const (
	ELASTIC_MODE_0 = 0
	ELASTIC_MODE_1 = 1
)

var ELASTIC_MODE = []int{
	ELASTIC_MODE_0,
	ELASTIC_MODE_1,
}

const (
	MAINLAND_0 = 0
	MAINLAND_1 = 1
)

var MAINLAND = []int{
	MAINLAND_0,
	MAINLAND_1,
}

const (
	DOMIAN_CATEGORY_ID_SAAS      = 102372
	DOMAIN_SUB_PRODUCT_CODE_SAAS = "sp_wsm_waf_domain"
	DOMAIN_PID_SAAS              = 1000834
	DOMAIN_LABEL_TYPE_SAAS       = "sv_wsm_waf_domain"
)

const (
	QPS_CATEGORY_ID_SAAS      = 101040
	QPS_SUB_PRODUCT_CODE_SAAS = "sp_wsm_waf_qpsep"
	QPS_PID_SAAS              = 1000481
	QPS_LABEL_TYPE_SAAS       = "sv_wsm_waf_qps_ep"
)

const (
	DOMIAN_CATEGORY_ID_CLB               = 101207
	DOMAIN_SUB_PRODUCT_CODE_CLB          = "sp_wsm_waf_domain_clb"
	DOMAIN_PID_CLB                       = 1001156
	DOMAIN_LABEL_TYPE_CLB                = "sv_wsm_waf_domain_clb"
	BOT_MANAGEMENT_CATEGORY_ID_CLB       = 1025567
	BOT_MANAGEMENT_SUB_PRODUCT_CODE_CLB  = "sp_wsm_waf_bot_protection_clb"
	BOT_MANAGEMENT_PID_CLB               = 1017001
	API_SECURITY_CATEGORY_ID_CLB         = 1027183
	API_SECURITY_SUB_PRODUCT_CODE_CLB    = "sp_wsm_waf_apiclb"
	API_SECURITY_PID_CLB                 = 1028166
	BOT_MANAGEMENT_CATEGORY_ID_SAAS      = 1025564
	BOT_MANAGEMENT_SUB_PRODUCT_CODE_SAAS = "sp_wsm_waf_bot_protection"
	BOT_MANAGEMENT_PID_SAAS              = 1016997
	API_SECURITY_CATEGORY_ID_SAAS        = 1027180
	API_SECURITY_SUB_PRODUCT_CODE_SAAS   = "sp_wsm_waf_api"
	API_SECURITY_PID_SAAS                = 1028161
)

const (
	QPS_CATEGORY_ID_CLB      = 101210
	QPS_SUB_PRODUCT_CODE_CLB = "sp_wsm_waf_qpsep_clb"
	QPS_PID_CLB              = 1001160
	QPS_LABEL_TYPE_CLB       = "sv_wsm_waf_qps_ep_clb"
)

const (
	SAAS_REAL_REGION_MAINLAND_GZ = "gz"
	SAAS_REAL_REGION_MAINLAND_SH = "sh"
	SAAS_REAL_REGION_MAINLAND_BJ = "bj"
	SAAS_REAL_REGION_MAINLAND_CD = "cd"
)

var SAAS_REAL_REGION_MAINLAND = []string{
	SAAS_REAL_REGION_MAINLAND_GZ,
	SAAS_REAL_REGION_MAINLAND_SH,
	SAAS_REAL_REGION_MAINLAND_BJ,
	SAAS_REAL_REGION_MAINLAND_CD,
}

var SAAS_REAL_REGION_MAINLAND_ID_MAP = map[string]int{
	SAAS_REAL_REGION_MAINLAND_GZ: 1,
	SAAS_REAL_REGION_MAINLAND_SH: 4,
	SAAS_REAL_REGION_MAINLAND_BJ: 8,
	SAAS_REAL_REGION_MAINLAND_CD: 16,
}

const (
	SAAS_REAL_REGION_NON_MAINLAND_HK  = "hk"
	SAAS_REAL_REGION_NON_MAINLAND_SG  = "sg"
	SAAS_REAL_REGION_NON_MAINLAND_TH  = "th"
	SAAS_REAL_REGION_NON_MAINLAND_KR  = "kr"
	SAAS_REAL_REGION_NON_MAINLAND_IN  = "in"
	SAAS_REAL_REGION_NON_MAINLAND_DE  = "de"
	SAAS_REAL_REGION_NON_MAINLAND_CA  = "ca"
	SAAS_REAL_REGION_NON_MAINLAND_USE = "use"
	SAAS_REAL_REGION_NON_MAINLAND_SAO = "sao"
	SAAS_REAL_REGION_NON_MAINLAND_USW = "usw"
	SAAS_REAL_REGION_NON_MAINLAND_JKT = "jkt"
)

var SAAS_REAL_REGION_NON_MAINLAND = []string{
	SAAS_REAL_REGION_NON_MAINLAND_HK,
	SAAS_REAL_REGION_NON_MAINLAND_SG,
	SAAS_REAL_REGION_NON_MAINLAND_TH,
	SAAS_REAL_REGION_NON_MAINLAND_KR,
	SAAS_REAL_REGION_NON_MAINLAND_IN,
	SAAS_REAL_REGION_NON_MAINLAND_DE,
	SAAS_REAL_REGION_NON_MAINLAND_CA,
	SAAS_REAL_REGION_NON_MAINLAND_USE,
	SAAS_REAL_REGION_NON_MAINLAND_SAO,
	SAAS_REAL_REGION_NON_MAINLAND_USW,
	SAAS_REAL_REGION_NON_MAINLAND_JKT,
}

var SAAS_REAL_REGION_NON_MAINLAND_ID_MAP = map[string]int{
	SAAS_REAL_REGION_NON_MAINLAND_HK:  5,
	SAAS_REAL_REGION_NON_MAINLAND_SG:  9,
	SAAS_REAL_REGION_NON_MAINLAND_TH:  23,
	SAAS_REAL_REGION_NON_MAINLAND_KR:  18,
	SAAS_REAL_REGION_NON_MAINLAND_IN:  21,
	SAAS_REAL_REGION_NON_MAINLAND_DE:  17,
	SAAS_REAL_REGION_NON_MAINLAND_CA:  6,
	SAAS_REAL_REGION_NON_MAINLAND_USE: 22,
	SAAS_REAL_REGION_NON_MAINLAND_SAO: 43,
	SAAS_REAL_REGION_NON_MAINLAND_USW: 51,
	SAAS_REAL_REGION_NON_MAINLAND_JKT: 72,
}

var SAAS_REAL_REGIONS = append(SAAS_REAL_REGION_MAINLAND, SAAS_REAL_REGION_NON_MAINLAND...)

const (
	EDITION_SAAS = "sparta-waf"
	EDITION_CLB  = "clb-waf"
)

var EDITION_TYPE = []string{
	EDITION_SAAS,
	EDITION_CLB,
}

const (
	REGION_GZ = "ap-guangzhou"
	REGION_KR = "ap-seoul"
)

const (
	METRIC_NAME_ACCESS              = "access"
	METRIC_NAME_BOTACCESS           = "botAccess"
	METRIC_NAME_DOWN                = "down"
	METRIC_NAME_UP                  = "up"
	METRIC_NAME_ATTACK              = "attack"
	METRIC_NAME_CC                  = "cc"
	METRIC_NAME_STATUSSERVERERROR   = "StatusServerError"
	METRIC_NAME_STATUSCLIENTERROR   = "StatusClientError"
	METRIC_NAME_STATUSREDIRECT      = "StatusRedirect"
	METRIC_NAME_STATUSOK            = "StatusOk"
	METRIC_NAME_UPSTREAMSERVERERROR = "UpstreamServerError"
	METRIC_NAME_UPSTREAMCLIENTERROR = "UpstreamClientError"
	METRIC_NAME_UPSTREAMREDIRECT    = "UpstreamRedirect"
)

var MetricNameList = []string{
	METRIC_NAME_ACCESS,
	METRIC_NAME_BOTACCESS,
	METRIC_NAME_DOWN,
	METRIC_NAME_UP,
	METRIC_NAME_ATTACK,
	METRIC_NAME_CC,
	METRIC_NAME_STATUSSERVERERROR,
	METRIC_NAME_STATUSCLIENTERROR,
	METRIC_NAME_STATUSREDIRECT,
	METRIC_NAME_STATUSOK,
	METRIC_NAME_UPSTREAMSERVERERROR,
	METRIC_NAME_UPSTREAMCLIENTERROR,
	METRIC_NAME_UPSTREAMREDIRECT,
}

var (
	STATE_0 = 0
)

const (
	ANTI_FAKE_URL_STATUS_0 = 0
	ANTI_FAKE_URL_STATUS_1 = 1
	ANTI_FAKE_URL_STATUS_2 = 2
	ANTI_FAKE_URL_STATUS_3 = 3
)

var ANTI_FAKE_URL_STATUS = []int{
	ANTI_FAKE_URL_STATUS_0,
	ANTI_FAKE_URL_STATUS_1,
	ANTI_FAKE_URL_STATUS_2,
	ANTI_FAKE_URL_STATUS_3,
}

const (
	ANTI_INFO_LEAK_ACTION_TYPE_0 = 0
	ANTI_INFO_LEAK_ACTION_TYPE_1 = 1
	ANTI_INFO_LEAK_ACTION_TYPE_2 = 2
	ANTI_INFO_LEAK_ACTION_TYPE_3 = 3
	ANTI_INFO_LEAK_ACTION_TYPE_4 = 4
)

var ANTI_INFO_LEAK_ACTION_TYPE = []int{
	ANTI_INFO_LEAK_ACTION_TYPE_0,
	ANTI_INFO_LEAK_ACTION_TYPE_1,
	ANTI_INFO_LEAK_ACTION_TYPE_2,
	ANTI_INFO_LEAK_ACTION_TYPE_3,
	ANTI_INFO_LEAK_ACTION_TYPE_4,
}

const (
	STRATEGIES_FIELD_RETURNCODE  = "returncode"
	STRATEGIES_FIELD_KEYWORDS    = "keywords"
	STRATEGIES_FIELD_INFORMATION = "information"
)

var STRATEGIES_FIELD = []string{
	STRATEGIES_FIELD_RETURNCODE,
	STRATEGIES_FIELD_KEYWORDS,
	STRATEGIES_FIELD_INFORMATION,
}

const (
	ANTI_INFO_LEAK_RULE_STATUS_0 = 0
	ANTI_INFO_LEAK_RULE_STATUS_1 = 1
)

var ANTI_INFO_LEAK_RULE_STATUS = []int{
	ANTI_INFO_LEAK_RULE_STATUS_0,
	ANTI_INFO_LEAK_RULE_STATUS_1,
}

const (
	BOT_MANAGEMENT_STATUS_0 = 0
	BOT_MANAGEMENT_STATUS_1 = 1
)

var BOT_MANAGEMENT_STATUS = []int{
	BOT_MANAGEMENT_STATUS_0,
	BOT_MANAGEMENT_STATUS_1,
}

const (
	API_SECURITY_STATUS_0 = 0
	API_SECURITY_STATUS_1 = 1
)

var API_SECURITY_STATUS = []int{
	API_SECURITY_STATUS_0,
	API_SECURITY_STATUS_1,
}
