/*
The TencentCloud provider is used to interact with many resources supported by [TencentCloud](https://intl.cloud.tencent.com).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** From version 1.9.0 (June 18, 2019), the provider start to support Terraform 0.12.x.

Example Usage

```hcl
terraform {
  required_providers {
    tencentcloud = {
      source = "tencentcloudstack/tencentcloud"
    }
  }
}

# Configure the TencentCloud Provider
provider "tencentcloud" {
  secret_id  = var.secret_id
  secret_key = var.secret_key
  region     = var.region
}

#Configure the TencentCloud Provider with STS
provider "tencentcloud" {
  secret_id  = var.secret_id
  secret_key = var.secret_key
  region     = var.region
  assume_role {
    role_arn         = var.assume_role_arn
    session_name     = var.session_name
    session_duration = var.session_duration
    policy           = var.policy
  }
}
```

Resources List

Provider Data Sources
  tencentcloud_availability_regions
  tencentcloud_availability_zones_by_product
  tencentcloud_availability_zones

Project
  Data Source
    tencentcloud_projects

  Resource
    tencentcloud_project

Anti-DDoS
  Data Source
    tencentcloud_antiddos_basic_device_status
    tencentcloud_antiddos_bgp_biz_trend
    tencentcloud_antiddos_list_listener
    tencentcloud_antiddos_overview_attack_trend

Anti-DDoS(DayuV2)
  Data Source
    tencentcloud_dayu_eip
    tencentcloud_dayu_l4_rules_v2
    tencentcloud_dayu_l7_rules_v2
    tencentcloud_antiddos_pending_risk_info
    tencentcloud_antiddos_overview_index
    tencentcloud_antiddos_overview_ddos_trend
    tencentcloud_antiddos_overview_ddos_event_list
    tencentcloud_antiddos_overview_cc_trend

  Resource
    tencentcloud_dayu_eip
    tencentcloud_dayu_l4_rule
    tencentcloud_dayu_l7_rule_v2
    tencentcloud_dayu_ddos_policy_v2
    tencentcloud_dayu_cc_policy_v2
    tencentcloud_dayu_ddos_ip_attachment_v2
    tencentcloud_antiddos_ddos_black_white_ip
    tencentcloud_antiddos_ddos_geo_ip_block_config
    tencentcloud_antiddos_ddos_speed_limit_config
    tencentcloud_antiddos_default_alarm_threshold
    tencentcloud_antiddos_scheduling_domain_user_name
    tencentcloud_antiddos_ip_alarm_threshold_config

Anti-DDoS(Dayu)
  Data Source
    tencentcloud_dayu_cc_http_policies
    tencentcloud_dayu_cc_https_policies
    tencentcloud_dayu_ddos_policies
    tencentcloud_dayu_ddos_policy_attachments
    tencentcloud_dayu_ddos_policy_cases
    tencentcloud_dayu_l4_rules
    tencentcloud_dayu_l7_rules

  Resource
    tencentcloud_dayu_cc_http_policy
    tencentcloud_dayu_cc_https_policy
    tencentcloud_dayu_ddos_policy
    tencentcloud_dayu_ddos_policy_attachment
    tencentcloud_dayu_ddos_policy_case
    tencentcloud_dayu_l4_rule
    tencentcloud_dayu_l7_rule

API GateWay
  Data Source
    tencentcloud_api_gateway_apis
    tencentcloud_api_gateway_services
    tencentcloud_api_gateway_throttling_services
    tencentcloud_api_gateway_throttling_apis
    tencentcloud_api_gateway_usage_plans
    tencentcloud_api_gateway_ip_strategies
    tencentcloud_api_gateway_customer_domains
    tencentcloud_api_gateway_usage_plan_environments
    tencentcloud_api_gateway_api_keys
    tencentcloud_api_gateway_api_docs
    tencentcloud_api_gateway_api_apps
    tencentcloud_api_gateway_plugins
    tencentcloud_api_gateway_upstreams
    tencentcloud_api_gateway_api_usage_plans
    tencentcloud_api_gateway_api_app_service
    tencentcloud_api_gateway_bind_api_apps_status
    tencentcloud_api_gateway_api_app_api
    tencentcloud_api_gateway_api_plugins
    tencentcloud_api_gateway_service_release_versions
    tencentcloud_api_gateway_service_environment_list

  Resource
    tencentcloud_api_gateway_api
    tencentcloud_api_gateway_service
    tencentcloud_api_gateway_custom_domain
    tencentcloud_api_gateway_usage_plan
    tencentcloud_api_gateway_usage_plan_attachment
    tencentcloud_api_gateway_ip_strategy
    tencentcloud_api_gateway_strategy_attachment
    tencentcloud_api_gateway_api_key
    tencentcloud_api_gateway_api_key_attachment
    tencentcloud_api_gateway_service_release
    tencentcloud_api_gateway_plugin
    tencentcloud_api_gateway_plugin_attachment
    tencentcloud_api_gateway_api_doc
    tencentcloud_api_gateway_api_app
    tencentcloud_api_gateway_upstream
    tencentcloud_api_gateway_api_app_attachment
    tencentcloud_api_gateway_update_api_app_key
    tencentcloud_api_gateway_import_open_api

Cloud Audit(Audit)
  Data Source
    tencentcloud_audit_cos_regions
    tencentcloud_audit_key_alias
    tencentcloud_audits

  Resource
    tencentcloud_audit
    tencentcloud_audit_track

Auto Scaling(AS)
  Data Source
    tencentcloud_as_scaling_configs
    tencentcloud_as_scaling_groups
    tencentcloud_as_scaling_policies
    tencentcloud_as_instances
    tencentcloud_as_advices
    tencentcloud_as_limits
    tencentcloud_as_last_activity

  Resource
    tencentcloud_as_scaling_config
    tencentcloud_as_scaling_group
    tencentcloud_as_scaling_group_status
    tencentcloud_as_attachment
    tencentcloud_as_scaling_policy
    tencentcloud_as_schedule
    tencentcloud_as_lifecycle_hook
    tencentcloud_as_notification
    tencentcloud_as_remove_instances
    tencentcloud_as_protect_instances
    tencentcloud_as_start_instances
    tencentcloud_as_stop_instances
    tencentcloud_as_scale_in_instances
    tencentcloud_as_scale_out_instances
    tencentcloud_as_execute_scaling_policy
    tencentcloud_as_complete_lifecycle

Content Delivery Network(CDN)
  Data Source
    tencentcloud_cdn_domains
    tencentcloud_cdn_domain_verifier

  Resource
    tencentcloud_cdn_domain
    tencentcloud_cdn_url_push
    tencentcloud_cdn_url_purge

Cloud Kafka(ckafka)
  Data Source
    tencentcloud_ckafka_users
    tencentcloud_ckafka_acls
    tencentcloud_ckafka_topics
    tencentcloud_ckafka_instances
    tencentcloud_ckafka_connect_resource
    tencentcloud_ckafka_region
    tencentcloud_ckafka_datahub_topic
    tencentcloud_ckafka_datahub_group_offsets
    tencentcloud_ckafka_datahub_task
    tencentcloud_ckafka_group
    tencentcloud_ckafka_group_offsets
    tencentcloud_ckafka_group_info
    tencentcloud_ckafka_task_status
    tencentcloud_ckafka_topic_flow_ranking
    tencentcloud_ckafka_topic_produce_connection
    tencentcloud_ckafka_topic_subscribe_group
    tencentcloud_ckafka_topic_sync_replica
    tencentcloud_ckafka_zone

  Resource
    tencentcloud_ckafka_instance
    tencentcloud_ckafka_user
    tencentcloud_ckafka_acl
    tencentcloud_ckafka_topic
    tencentcloud_ckafka_datahub_topic
    tencentcloud_ckafka_connect_resource
    tencentcloud_ckafka_renew_instance
    tencentcloud_ckafka_acl_rule
    tencentcloud_ckafka_consumer_group
    tencentcloud_ckafka_consumer_group_modify_offset
    tencentcloud_ckafka_datahub_task
    tencentcloud_ckafka_route

Cloud Access Management(CAM)
  Data Source
    tencentcloud_cam_group_memberships
    tencentcloud_cam_group_policy_attachments
    tencentcloud_cam_groups
    tencentcloud_cam_policies
    tencentcloud_cam_role_policy_attachments
    tencentcloud_cam_roles
    tencentcloud_cam_saml_providers
    tencentcloud_cam_user_policy_attachments
    tencentcloud_cam_users
    tencentcloud_user_info
    tencentcloud_cam_list_entities_for_policy
    tencentcloud_cam_secret_last_used_time
    tencentcloud_cam_account_summary
    tencentcloud_cam_policy_granting_service_access
    tencentcloud_cam_oidc_config
    tencentcloud_cam_group_user_account

  Resource
    tencentcloud_cam_role
    tencentcloud_cam_role_by_name
    tencentcloud_cam_role_policy_attachment
    tencentcloud_cam_role_policy_attachment_by_name
    tencentcloud_cam_policy
    tencentcloud_cam_policy_by_name
    tencentcloud_cam_user
    tencentcloud_cam_user_policy_attachment
    tencentcloud_cam_group
    tencentcloud_cam_group_policy_attachment
    tencentcloud_cam_group_membership
    tencentcloud_cam_saml_provider
    tencentcloud_cam_oidc_sso
    tencentcloud_cam_role_sso
    tencentcloud_cam_service_linked_role
    tencentcloud_cam_mfa_flag
    tencentcloud_cam_access_key
    tencentcloud_cam_user_saml_config
    tencentcloud_cam_tag_role_attachment
    tencentcloud_cam_policy_version
    tencentcloud_cam_set_policy_version_config
    tencentcloud_cam_user_permission_boundary_attachment
    tencentcloud_cam_role_permission_boundary_attachment

Customer Identity and Access Management(CIAM)
  Resource
    tencentcloud_ciam_user_store
    tencentcloud_ciam_user_group

Cloud Block Storage(CBS)
  Data Source
    tencentcloud_cbs_snapshots
    tencentcloud_cbs_storages
    tencentcloud_cbs_storages_set
    tencentcloud_cbs_snapshot_policies

  Resource
    tencentcloud_cbs_storage
    tencentcloud_cbs_storage_set
    tencentcloud_cbs_storage_attachment
    tencentcloud_cbs_storage_set_attachment
    tencentcloud_cbs_snapshot
    tencentcloud_cbs_snapshot_policy
    tencentcloud_cbs_snapshot_policy_attachment
    tencentcloud_cbs_snapshot_share_permission
    tencentcloud_cbs_disk_backup
    tencentcloud_cbs_disk_backup_rollback_operation

Cloud Connect Network(CCN)
  Data Source
    tencentcloud_ccn_bandwidth_limits
    tencentcloud_ccn_instances
    tencentcloud_ccn_cross_border_compliance
    tencentcloud_ccn_tenant_instances
    tencentcloud_ccn_cross_border_flow_monitor
    tencentcloud_ccn_cross_border_region_bandwidth_limits

  Resource
    tencentcloud_ccn
    tencentcloud_ccn_attachment
    tencentcloud_ccn_bandwidth_limit
    tencentcloud_ccn_routes
    tencentcloud_ccn_instances_accept_attach
    tencentcloud_ccn_instances_reject_attach
    tencentcloud_ccn_instances_reset_attach

CVM Dedicated Host(CDH)
  Data Source
    tencentcloud_cdh_instances

  Resource
    tencentcloud_cdh_instance

Cloud File Storage(CFS)
  Data Source
    tencentcloud_cfs_access_groups
    tencentcloud_cfs_access_rules
    tencentcloud_cfs_file_systems
    tencentcloud_cfs_mount_targets
    tencentcloud_cfs_file_system_clients
    tencentcloud_cfs_available_zone

  Resource
    tencentcloud_cfs_file_system
    tencentcloud_cfs_access_group
    tencentcloud_cfs_access_rule
    tencentcloud_cfs_auto_snapshot_policy
    tencentcloud_cfs_auto_snapshot_policy_attachment
    tencentcloud_cfs_snapshot
    tencentcloud_cfs_sign_up_cfs_service

Container Cluster
  Data Source
    tencentcloud_container_cluster_instances
    tencentcloud_container_clusters

  Resource
    tencentcloud_container_cluster
    tencentcloud_container_cluster_instance

Cloud Load Balancer(CLB)
  Data Source
    tencentcloud_clb_attachments
    tencentcloud_clb_instances
    tencentcloud_clb_listener_rules
    tencentcloud_clb_listeners
    tencentcloud_clb_redirections
    tencentcloud_clb_target_groups
    tencentcloud_clb_cluster_resources
    tencentcloud_clb_cross_targets
    tencentcloud_clb_exclusive_clusters
    tencentcloud_clb_idle_instances
    tencentcloud_clb_listeners_by_targets
    tencentcloud_clb_instance_by_cert_id
    tencentcloud_clb_instance_traffic
    tencentcloud_clb_instance_detail
    tencentcloud_clb_resources
    tencentcloud_clb_target_group_list
    tencentcloud_clb_target_health

  Resource
    tencentcloud_clb_instance
    tencentcloud_clb_listener
    tencentcloud_clb_listener_rule
    tencentcloud_clb_attachment
    tencentcloud_clb_redirection
    tencentcloud_lb
    tencentcloud_alb_server_attachment
    tencentcloud_clb_target_group
    tencentcloud_clb_target_group_instance_attachment
    tencentcloud_clb_target_group_attachment
    tencentcloud_clb_log_set
    tencentcloud_clb_log_topic
    tencentcloud_clb_customized_config
    tencentcloud_clb_snat_ip
    tencentcloud_clb_function_targets_attachment
    tencentcloud_clb_instance_sla_config
    tencentcloud_clb_instance_mix_ip_target_config
    tencentcloud_clb_replace_cert_for_lbs
    tencentcloud_clb_security_group_attachment

Cloud Object Storage(COS)
  Data Source
    tencentcloud_cos_bucket_object
    tencentcloud_cos_buckets
    tencentcloud_cos_batchs
    tencentcloud_cos_bucket_inventorys
    tencentcloud_cos_bucket_multipart_uploads

  Resource
    tencentcloud_cos_bucket
    tencentcloud_cos_bucket_object
    tencentcloud_cos_bucket_policy
    tencentcloud_cos_bucket_referer
    tencentcloud_cos_bucket_version
    tencentcloud_cos_bucket_domain_certificate_attachment
    tencentcloud_cos_bucket_inventory
    tencentcloud_cos_batch
    tencentcloud_cos_object_abort_multipart_upload_operation
    tencentcloud_cos_object_copy_operation
    tencentcloud_cos_object_restore_operation
    tencentcloud_cos_bucket_generate_inventory_immediately_operation
    tencentcloud_cos_object_download_operation

Cloud Virtual Machine(CVM)
  Data Source
    tencentcloud_image
    tencentcloud_images
    tencentcloud_instance_types
    tencentcloud_instances
    tencentcloud_instances_set
    tencentcloud_key_pairs
    tencentcloud_eip
    tencentcloud_eips
    tencentcloud_eip_address_quota
    tencentcloud_eip_network_account_type
    tencentcloud_placement_groups
    tencentcloud_reserved_instance_configs
    tencentcloud_reserved_instances
    tencentcloud_cvm_instances_modification
    tencentcloud_cvm_instance_vnc_url
    tencentcloud_cvm_disaster_recover_group_quota
    tencentcloud_cvm_chc_hosts
    tencentcloud_cvm_chc_denied_actions
    tencentcloud_cvm_image_quota
    tencentcloud_cvm_image_share_permission
    tencentcloud_cvm_import_image_os

  Resource
    tencentcloud_instance
    tencentcloud_instance_set
    tencentcloud_eip
    tencentcloud_eip_association
    tencentcloud_eip_address_transform
    tencentcloud_eip_public_address_adjust
    tencentcloud_eip_normal_address_return
    tencentcloud_key_pair
    tencentcloud_placement_group
    tencentcloud_reserved_instance
    tencentcloud_image
    tencentcloud_cvm_hpc_cluster
    tencentcloud_cvm_launch_template
    tencentcloud_cvm_launch_template_version
    tencentcloud_cvm_launch_template_default_version
    tencentcloud_cvm_security_group_attachment
    tencentcloud_cvm_reboot_instance
    tencentcloud_cvm_chc_config
    tencentcloud_cvm_renew_instance
    tencentcloud_cvm_sync_image
    tencentcloud_cvm_export_images
    tencentcloud_cvm_image_share_permission

TDSQL-C MySQL(CynosDB)
  Data Source
    tencentcloud_cynosdb_clusters
    tencentcloud_cynosdb_instances
    tencentcloud_cynosdb_zone_config
    tencentcloud_cynosdb_accounts
    tencentcloud_cynosdb_cluster_instance_groups
    tencentcloud_cynosdb_cluster_params
    tencentcloud_cynosdb_param_templates
    tencentcloud_cynosdb_audit_logs
    tencentcloud_cynosdb_binlog_download_url
    tencentcloud_cynosdb_cluster_detail_databases
    tencentcloud_cynosdb_cluster_param_logs
    tencentcloud_cynosdb_cluster
    tencentcloud_cynosdb_describe_instance_slow_queries
    tencentcloud_cynosdb_describe_instance_error_logs
    tencentcloud_cynosdb_account_all_grant_privileges
    tencentcloud_cynosdb_resource_package_list
    tencentcloud_cynosdb_project_security_groups
    tencentcloud_cynosdb_resource_package_sale_specs
    tencentcloud_cynosdb_rollback_time_range
    tencentcloud_cynosdb_zone
    tencentcloud_cynosdb_instance_slow_queries
    tencentcloud_cynosdb_proxy_node
    tencentcloud_cynosdb_proxy_version

  Resource
    tencentcloud_cynosdb_cluster_resource_packages_attachment
    tencentcloud_cynosdb_cluster
    tencentcloud_cynosdb_readonly_instance
    tencentcloud_cynosdb_security_group
    tencentcloud_cynosdb_audit_log_file
    tencentcloud_cynosdb_cluster_password_complexity
    tencentcloud_cynosdb_export_instance_error_logs
    tencentcloud_cynosdb_export_instance_slow_queries
    tencentcloud_cynosdb_account_privileges
    tencentcloud_cynosdb_account
    tencentcloud_cynosdb_binlog_save_days
    tencentcloud_cynosdb_cluster_databases
    tencentcloud_cynosdb_instance_param
    tencentcloud_cynosdb_isolate_instance
    tencentcloud_cynosdb_param_template
    tencentcloud_cynosdb_restart_instance
    tencentcloud_cynosdb_roll_back_cluster
    tencentcloud_cynosdb_wan
    tencentcloud_cynosdb_proxy
    tencentcloud_cynosdb_reload_proxy_node
    tencentcloud_cynosdb_cluster_slave_zone
    tencentcloud_cynosdb_read_only_instance_exclusive_access
    tencentcloud_cynosdb_proxy_end_point
    tencentcloud_cynosdb_upgrade_proxy_version

Direct Connect(DC)
  Data Source
    tencentcloud_dc_instances
    tencentcloud_dc_access_points
    tencentcloud_dcx_instances
    tencentcloud_dc_internet_address_quota
    tencentcloud_dc_internet_address_statistics
    tencentcloud_dc_public_direct_connect_tunnel_routes

  Resource
    tencentcloud_dc_instance
    tencentcloud_dcx
    tencentcloud_dcx_extra_config
    tencentcloud_dc_share_dcx_config
    tencentcloud_dc_internet_address
    tencentcloud_dc_internet_address_config

Direct Connect Gateway(DCG)
  Data Source
    tencentcloud_dc_gateway_ccn_routes
    tencentcloud_dc_gateway_instances

  Resource
    tencentcloud_dc_gateway
    tencentcloud_dc_gateway_ccn_route
    tencentcloud_dc_gateway_attachment

Domain
  Data Source
    tencentcloud_domains

Elasticsearch Service(ES)
  Data Source
    tencentcloud_elasticsearch_instances
    tencentcloud_elasticsearch_instance_logs
    tencentcloud_elasticsearch_instance_operations
    tencentcloud_elasticsearch_logstash_instance_logs
    tencentcloud_elasticsearch_logstash_instance_operations
    tencentcloud_elasticsearch_views
    tencentcloud_elasticsearch_diagnose
    tencentcloud_elasticsearch_instance_plugin_list

  Resource
    tencentcloud_elasticsearch_instance
    tencentcloud_elasticsearch_security_group
    tencentcloud_elasticsearch_logstash
    tencentcloud_elasticsearch_logstash_pipeline
    tencentcloud_elasticsearch_restart_logstash_instance_operation
    tencentcloud_elasticsearch_start_logstash_pipeline_operation
    tencentcloud_elasticsearch_stop_logstash_pipeline_operation
    tencentcloud_elasticsearch_index
    tencentcloud_elasticsearch_restart_instance_operation
    tencentcloud_elasticsearch_restart_nodes_operation
    tencentcloud_elasticsearch_restart_kibana_operation
    tencentcloud_elasticsearch_diagnose
    tencentcloud_elasticsearch_diagnose_instance
    tencentcloud_elasticsearch_update_plugins_operation

Global Application Acceleration(GAAP)
  Data Source
    tencentcloud_gaap_certificates
    tencentcloud_gaap_http_domains
    tencentcloud_gaap_http_rules
    tencentcloud_gaap_layer4_listeners
    tencentcloud_gaap_layer7_listeners
    tencentcloud_gaap_proxies
    tencentcloud_gaap_realservers
    tencentcloud_gaap_security_policies
    tencentcloud_gaap_security_rules
    tencentcloud_gaap_domain_error_pages
    tencentcloud_gaap_access_regions
    tencentcloud_gaap_access_regions_by_dest_region
    tencentcloud_gaap_black_header
    tencentcloud_gaap_country_area_mapping
    tencentcloud_gaap_custom_header
    tencentcloud_gaap_dest_regions
    tencentcloud_gaap_proxy_detail
    tencentcloud_gaap_proxy_groups
    tencentcloud_gaap_proxy_statistics
    tencentcloud_gaap_proxy_group_statistics
    tencentcloud_gaap_real_servers_status
    tencentcloud_gaap_rule_real_servers
    tencentcloud_gaap_resources_by_tag
    tencentcloud_gaap_region_and_price
    tencentcloud_gaap_proxy_and_statistics_listeners
    tencentcloud_gaap_proxies_status
    tencentcloud_gaap_listener_statistics
    tencentcloud_gaap_listener_real_servers
    tencentcloud_gaap_group_and_statistics_proxy
    tencentcloud_gaap_domain_error_page_infos
    tencentcloud_gaap_check_proxy_create

  Resource
    tencentcloud_gaap_proxy
    tencentcloud_gaap_realserver
    tencentcloud_gaap_layer4_listener
    tencentcloud_gaap_layer7_listener
    tencentcloud_gaap_http_domain
    tencentcloud_gaap_http_rule
    tencentcloud_gaap_certificate
    tencentcloud_gaap_security_policy
    tencentcloud_gaap_security_rule
    tencentcloud_gaap_domain_error_page
    tencentcloud_gaap_global_domain_dns
    tencentcloud_gaap_global_domain

Key Management Service(KMS)
  Data Source
    tencentcloud_kms_keys
    tencentcloud_kms_public_key
    tencentcloud_kms_get_parameters_for_import
    tencentcloud_kms_describe_keys
    tencentcloud_kms_white_box_key_details
    tencentcloud_kms_list_keys
    tencentcloud_kms_white_box_decrypt_key
    tencentcloud_kms_white_box_device_fingerprints
    tencentcloud_kms_list_algorithms

  Resource
    tencentcloud_kms_key
    tencentcloud_kms_external_key
    tencentcloud_kms_white_box_key
    tencentcloud_kms_cloud_resource_attachment
    tencentcloud_kms_overwrite_white_box_device_fingerprints

Tencent Kubernetes Engine(TKE)
  Data Source
    tencentcloud_kubernetes_clusters
    tencentcloud_kubernetes_cluster_levels
    tencentcloud_kubernetes_charts
    tencentcloud_kubernetes_cluster_common_names
    tencentcloud_kubernetes_available_cluster_versions
    tencentcloud_kubernetes_cluster_authentication_options

  Resource
    tencentcloud_kubernetes_cluster
    tencentcloud_kubernetes_scale_worker
    tencentcloud_kubernetes_cluster_attachment
    tencentcloud_kubernetes_node_pool
    tencentcloud_kubernetes_serverless_node_pool
    tencentcloud_kubernetes_backup_storage_location
    tencentcloud_kubernetes_encryption_protection
    tencentcloud_kubernetes_auth_attachment
    tencentcloud_kubernetes_addon_attachment
    tencentcloud_kubernetes_cluster_endpoint

TDMQ for Pulsar(tpulsar)
  Data Source
    tencentcloud_tdmq_environment_attributes
    tencentcloud_tdmq_publisher_summary
    tencentcloud_tdmq_publishers
    tencentcloud_tdmq_pro_instances
    tencentcloud_tdmq_pro_instance_detail

  Resource
    tencentcloud_tdmq_instance
    tencentcloud_tdmq_namespace
    tencentcloud_tdmq_topic
    tencentcloud_tdmq_role
    tencentcloud_tdmq_namespace_role_attachment

TencentDB for MongoDB(mongodb)
  Data Source
    tencentcloud_mongodb_instances
    tencentcloud_mongodb_zone_config
    tencentcloud_mongodb_instance_backups
    tencentcloud_mongodb_instance_connections
    tencentcloud_mongodb_instance_current_op
    tencentcloud_mongodb_instance_params
    tencentcloud_mongodb_instance_slow_log

  Resource
    tencentcloud_mongodb_instance
    tencentcloud_mongodb_sharding_instance
    tencentcloud_mongodb_standby_instance
    tencentcloud_mongodb_instance_account
    tencentcloud_mongodb_instance_backup

TencentDB for MySQL(cdb)
  Data Source
    tencentcloud_mysql_backup_list
    tencentcloud_mysql_instance
    tencentcloud_mysql_parameter_list
    tencentcloud_mysql_default_params
    tencentcloud_mysql_zone_config
    tencentcloud_mysql_backup_overview
    tencentcloud_mysql_backup_summaries
    tencentcloud_mysql_bin_log
    tencentcloud_mysql_binlog_backup_overview
    tencentcloud_mysql_clone_list
    tencentcloud_mysql_data_backup_overview
    tencentcloud_mysql_db_features
    tencentcloud_mysql_inst_tables
    tencentcloud_mysql_instance_charset
    tencentcloud_mysql_instance_info
    tencentcloud_mysql_instance_param_record
    tencentcloud_mysql_instance_reboot_time
    tencentcloud_mysql_rollback_range_time
    tencentcloud_mysql_slow_log
    tencentcloud_mysql_slow_log_data
    tencentcloud_mysql_supported_privileges
    tencentcloud_mysql_switch_record
    tencentcloud_mysql_user_task
    tencentcloud_mysql_databases
    tencentcloud_mysql_error_log
    tencentcloud_mysql_project_security_group
    tencentcloud_mysql_ro_min_scale

  Resource
    tencentcloud_mysql_instance
    tencentcloud_mysql_database
    tencentcloud_mysql_readonly_instance
    tencentcloud_mysql_account
    tencentcloud_mysql_privilege
    tencentcloud_mysql_account_privilege
    tencentcloud_mysql_backup_policy
    tencentcloud_mysql_time_window
    tencentcloud_mysql_param_template
    tencentcloud_mysql_deploy_group
    tencentcloud_mysql_security_groups_attachment
    tencentcloud_mysql_local_binlog_config
    tencentcloud_mysql_audit_log_file
    tencentcloud_mysql_backup_download_restriction
    tencentcloud_mysql_renew_db_instance_operation
    tencentcloud_mysql_backup_encryption_status
    tencentcloud_mysql_dr_instance_to_mater
    tencentcloud_mysql_instance_encryption_operation
    tencentcloud_mysql_password_complexity
    tencentcloud_mysql_remote_backup_config
    tencentcloud_mysql_restart_db_instances_operation
    tencentcloud_mysql_switch_for_upgrade
    tencentcloud_mysql_rollback
    tencentcloud_mysql_rollback_stop
    tencentcloud_mysql_ro_group
    tencentcloud_mysql_ro_instance_ip
    tencentcloud_mysql_ro_group_load_operation
    tencentcloud_mysql_switch_master_slave_operation
    tencentcloud_mysql_proxy
    tencentcloud_mysql_reset_root_account
    tencentcloud_mysql_verify_root_account
    tencentcloud_mysql_reload_balance_proxy_node
    tencentcloud_mysql_ro_start_replication
    tencentcloud_mysql_ro_stop_replication
    tencentcloud_mysql_isolate_instance

Cloud Monitor(Monitor)
  Data Source
    tencentcloud_monitor_policy_conditions
    tencentcloud_monitor_data
    tencentcloud_monitor_product_event
    tencentcloud_monitor_binding_objects
    tencentcloud_monitor_policy_groups
    tencentcloud_monitor_product_namespace
    tencentcloud_monitor_alarm_notices
    tencentcloud_monitor_alarm_history
    tencentcloud_monitor_alarm_metric
    tencentcloud_monitor_alarm_policy
    tencentcloud_monitor_alarm_basic_alarms
    tencentcloud_monitor_alarm_basic_metric
    tencentcloud_monitor_alarm_conditions_template
    tencentcloud_monitor_alarm_notice_callbacks
    tencentcloud_monitor_alarm_all_namespaces
    tencentcloud_monitor_alarm_monitor_type

  Resource
    tencentcloud_monitor_policy_group
    tencentcloud_monitor_binding_object
    tencentcloud_monitor_policy_binding_object
    tencentcloud_monitor_binding_receiver
    tencentcloud_monitor_alarm_policy
    tencentcloud_monitor_alarm_notice
    tencentcloud_monitor_alarm_policy_set_default


Managed Service for Prometheus(TMP)
  Data Source
    tencentcloud_monitor_tmp_regions

  Resource
    tencentcloud_monitor_tmp_instance
    tencentcloud_monitor_tmp_alert_rule
    tencentcloud_monitor_tmp_exporter_integration
    tencentcloud_monitor_tmp_cvm_agent
    tencentcloud_monitor_tmp_scrape_job
    tencentcloud_monitor_tmp_recording_rule
    tencentcloud_monitor_tmp_manage_grafana_attachment
    tencentcloud_monitor_tmp_tke_template
    tencentcloud_monitor_tmp_tke_template_attachment
    tencentcloud_monitor_tmp_tke_alert_policy
    tencentcloud_monitor_tmp_tke_config
    tencentcloud_monitor_tmp_tke_record_rule_yaml
    tencentcloud_monitor_tmp_tke_global_notification
    tencentcloud_monitor_tmp_tke_cluster_agent
    tencentcloud_monitor_tmp_tke_basic_config

TencentCloud Managed Service for Grafana(TCMG)
  Data Source
    tencentcloud_monitor_grafana_plugin_overviews

  Resource
    tencentcloud_monitor_grafana_instance
    tencentcloud_monitor_grafana_integration
    tencentcloud_monitor_grafana_notification_channel
    tencentcloud_monitor_grafana_plugin
    tencentcloud_monitor_grafana_sso_account
    tencentcloud_monitor_tmp_grafana_config
    tencentcloud_monitor_grafana_dns_config
    tencentcloud_monitor_grafana_env_config
    tencentcloud_monitor_grafana_whitelist_config
    tencentcloud_monitor_grafana_sso_cam_config
    tencentcloud_monitor_grafana_sso_config
    tencentcloud_monitor_grafana_version_upgrade

TencentDB for PostgreSQL(PostgreSQL)
  Data Source
    tencentcloud_postgresql_instances
    tencentcloud_postgresql_specinfos
    tencentcloud_postgresql_xlogs
    tencentcloud_postgresql_parameter_templates
    tencentcloud_postgresql_readonly_groups
    tencentcloud_postgresql_base_backups
    tencentcloud_postgresql_log_backups
    tencentcloud_postgresql_backup_download_urls
    tencentcloud_postgresql_db_instance_classes
    tencentcloud_postgresql_default_parameters
    tencentcloud_postgresql_recovery_time
    tencentcloud_postgresql_regions
    tencentcloud_postgresql_db_instance_versions
    tencentcloud_postgresql_zones

  Resource
    tencentcloud_postgresql_instance
    tencentcloud_postgresql_readonly_instance
    tencentcloud_postgresql_readonly_group
    tencentcloud_postgresql_readonly_attachment
    tencentcloud_postgresql_parameter_template
    tencentcloud_postgresql_backup_plan_config
    tencentcloud_postgresql_security_group_config
    tencentcloud_postgresql_backup_download_restriction_config
    tencentcloud_postgresql_restart_db_instance_operation
    tencentcloud_postgresql_renew_db_instance_operation
    tencentcloud_postgresql_isolate_db_instance_operation
    tencentcloud_postgresql_disisolate_db_instance_operation
    tencentcloud_postgresql_rebalance_readonly_group_operation
    tencentcloud_postgresql_delete_log_backup_operation
    tencentcloud_postgresql_modify_account_remark_operation
    tencentcloud_postgresql_modify_switch_time_period_operation
    tencentcloud_postgresql_base_backup

TencentDB for Redis(crs)
  Data Source
    tencentcloud_redis_zone_config
    tencentcloud_redis_instances
    tencentcloud_redis_backup
    tencentcloud_redis_backup_download_info
    tencentcloud_redis_param_records
    tencentcloud_redis_instance_shards
    tencentcloud_redis_instance_zone_info
    tencentcloud_redis_instance_task_list
    tencentcloud_redis_instance_node_info

  Resource
    tencentcloud_redis_instance
    tencentcloud_redis_backup_config
    tencentcloud_redis_param_template
    tencentcloud_redis_account
    tencentcloud_redis_read_only
    tencentcloud_redis_ssl
    tencentcloud_redis_backup_download_restriction
    tencentcloud_redis_clear_instance_operation
    tencentcloud_redis_renew_instance_operation
    tencentcloud_redis_startup_instance_operation
    tencentcloud_redis_upgrade_cache_version_operation
    tencentcloud_redis_upgrade_multi_zone_operation
    tencentcloud_redis_upgrade_proxy_version_operation
    tencentcloud_redis_maintenance_window
    tencentcloud_redis_replica_readonly
    tencentcloud_redis_switch_master
    tencentcloud_redis_replicate_attachment
    tencentcloud_redis_backup_operation
    tencentcloud_redis_security_group_attachment
    tencentcloud_redis_connection_config

Serverless Cloud Function(SCF)
  Data Source
    tencentcloud_scf_functions
    tencentcloud_scf_logs
    tencentcloud_scf_namespaces
    tencentcloud_scf_account_info
    tencentcloud_scf_async_event_management
    tencentcloud_scf_triggers
    tencentcloud_scf_async_event_status
    tencentcloud_scf_function_address
    tencentcloud_scf_request_status
    tencentcloud_scf_function_aliases
    tencentcloud_scf_layer_versions
    tencentcloud_scf_layers
    tencentcloud_scf_function_versions

  Resource
    tencentcloud_scf_function
    tencentcloud_scf_function_version
    tencentcloud_scf_function_event_invoke_config
    tencentcloud_scf_reserved_concurrency_config
    tencentcloud_scf_provisioned_concurrency_config
    tencentcloud_scf_invoke_function
    tencentcloud_scf_sync_invoke_function
    tencentcloud_scf_terminate_async_event
    tencentcloud_scf_namespace
    tencentcloud_scf_layer
    tencentcloud_scf_function_alias
    tencentcloud_scf_trigger_config

SQLServer
  Data Source
    tencentcloud_sqlserver_zone_config
    tencentcloud_sqlserver_instances
    tencentcloud_sqlserver_dbs
    tencentcloud_sqlserver_accounts
    tencentcloud_sqlserver_account_db_attachments
      tencentcloud_sqlserver_readonly_groups
    tencentcloud_sqlserver_publish_subscribes
    tencentcloud_sqlserver_basic_instances
    tencentcloud_sqlserver_backup_commands
    tencentcloud_sqlserver_backup_by_flow_id
    tencentcloud_sqlserver_backup_upload_size
    tencentcloud_sqlserver_cross_region_zone
    tencentcloud_sqlserver_db_charsets
    tencentcloud_sqlserver_instance_param_records
    tencentcloud_sqlserver_project_security_groups
    tencentcloud_sqlserver_regions
    tencentcloud_sqlserver_rollback_time
    tencentcloud_sqlserver_slowlogs
    tencentcloud_sqlserver_upload_backup_info
    tencentcloud_sqlserver_upload_incremental_info
    tencentcloud_sqlserver_query_xevent
    tencentcloud_sqlserver_ins_attribute

  Resource
    tencentcloud_sqlserver_instance
    tencentcloud_sqlserver_readonly_instance
    tencentcloud_sqlserver_db
    tencentcloud_sqlserver_account
    tencentcloud_sqlserver_account_db_attachment
    tencentcloud_sqlserver_publish_subscribe
    tencentcloud_sqlserver_basic_instance
    tencentcloud_sqlserver_migration
    tencentcloud_sqlserver_config_backup_strategy
    tencentcloud_sqlserver_general_backup
    tencentcloud_sqlserver_general_clone
    tencentcloud_sqlserver_full_backup_migration
    tencentcloud_sqlserver_incre_backup_migration
    tencentcloud_sqlserver_business_intelligence_file
    tencentcloud_sqlserver_business_intelligence_instance
    tencentcloud_sqlserver_general_communication
    tencentcloud_sqlserver_general_cloud_instance
    tencentcloud_sqlserver_complete_expansion
    tencentcloud_sqlserver_config_database_cdc
    tencentcloud_sqlserver_config_database_ct
    tencentcloud_sqlserver_config_database_mdf
    tencentcloud_sqlserver_config_instance_param
    tencentcloud_sqlserver_config_instance_ro_group
    tencentcloud_sqlserver_renew_db_instance
    tencentcloud_sqlserver_renew_postpaid_db_instance
    tencentcloud_sqlserver_restart_db_instance
    tencentcloud_sqlserver_config_terminate_db_instance
    tencentcloud_sqlserver_restore_instance
    tencentcloud_sqlserver_rollback_instance
    tencentcloud_sqlserver_start_backup_full_migration
    tencentcloud_sqlserver_start_backup_incremental_migration
    tencentcloud_sqlserver_start_xevent
    tencentcloud_sqlserver_instance_tde
    tencentcloud_sqlserver_database_tde
    tencentcloud_sqlserver_general_cloud_ro_instance

SSL Certificates
  Data Source
    tencentcloud_ssl_certificates
    tencentcloud_ssl_describe_certificate
    tencentcloud_ssl_describe_companies
    tencentcloud_ssl_describe_host_api_gateway_instance_list
    tencentcloud_ssl_describe_host_cdn_instance_list
    tencentcloud_ssl_describe_host_clb_instance_list
    tencentcloud_ssl_describe_host_cos_instance_list
    tencentcloud_ssl_describe_host_ddos_instance_list
    tencentcloud_ssl_describe_host_lighthouse_instance_list
    tencentcloud_ssl_describe_host_live_instance_list
    tencentcloud_ssl_describe_host_teo_instance_list
    tencentcloud_ssl_describe_host_tke_instance_list
    tencentcloud_ssl_describe_host_vod_instance_list
    tencentcloud_ssl_describe_host_waf_instance_list
    tencentcloud_ssl_describe_host_deploy_record
    tencentcloud_ssl_describe_host_deploy_record_detail
    tencentcloud_ssl_describe_host_update_record
    tencentcloud_ssl_describe_host_update_record_detail
    tencentcloud_ssl_describe_managers
    tencentcloud_ssl_describe_manager_detail

  Resource
    tencentcloud_ssl_certificate
    tencentcloud_ssl_pay_certificate
    tencentcloud_ssl_free_certificate
    tencentcloud_ssl_replace_certificate_operation
    tencentcloud_ssl_revoke_certificate_operation
    tencentcloud_ssl_update_certificate_instance_operation
    tencentcloud_ssl_update_certificate_record_retry_operation
    tencentcloud_ssl_update_certificate_record_rollback_operation
    tencentcloud_ssl_upload_revoke_letter_operation
    tencentcloud_ssl_complete_certificate_operation
    tencentcloud_ssl_check_certificate_chain_operation
    tencentcloud_ssl_deploy_certificate_instance_operation
    tencentcloud_ssl_deploy_certificate_record_retry_operation
    tencentcloud_ssl_deploy_certificate_record_rollback_operation
    tencentcloud_ssl_download_certificate_operation

Secrets Manager(SSM)
  Data Source
    tencentcloud_ssm_products
    tencentcloud_ssm_secrets
    tencentcloud_ssm_secret_versions
    tencentcloud_ssm_rotation_detail
    tencentcloud_ssm_rotation_history
    tencentcloud_ssm_service_status
    tencentcloud_ssm_ssh_key_pair_value

  Resource
    tencentcloud_ssm_secret
    tencentcloud_ssm_secret_version
    tencentcloud_ssm_product_secret
    tencentcloud_ssm_ssh_key_pair_secret
    tencentcloud_ssm_rotate_product_secret

TcaplusDB
  Data Source
    tencentcloud_tcaplus_clusters
    tencentcloud_tcaplus_idls
    tencentcloud_tcaplus_tables
    tencentcloud_tcaplus_tablegroups

  Resource
    tencentcloud_tcaplus_cluster
    tencentcloud_tcaplus_tablegroup
    tencentcloud_tcaplus_idl
    tencentcloud_tcaplus_table

Tencent Container Registry(TCR)
  Data Source
    tencentcloud_tcr_instances
    tencentcloud_tcr_namespaces
    tencentcloud_tcr_repositories
    tencentcloud_tcr_tokens
    tencentcloud_tcr_vpc_attachments
    tencentcloud_tcr_webhook_trigger_logs
    tencentcloud_tcr_images
    tencentcloud_tcr_image_manifests
    tencentcloud_tcr_tag_retention_execution_tasks
    tencentcloud_tcr_tag_retention_executions
    tencentcloud_tcr_replication_instance_create_tasks
    tencentcloud_tcr_replication_instance_sync_status

  Resource
    tencentcloud_tcr_instance
    tencentcloud_tcr_namespace
    tencentcloud_tcr_repository
    tencentcloud_tcr_token
    tencentcloud_tcr_vpc_attachment
    tencentcloud_tcr_tag_retention_rule
    tencentcloud_tcr_webhook_trigger
    tencentcloud_tcr_manage_replication_operation
    tencentcloud_tcr_customized_domain
    tencentcloud_tcr_immutable_tag_rule
    tencentcloud_tcr_delete_image_operation
    tencentcloud_tcr_create_image_signature_operation
    tencentcloud_tcr_tag_retention_execution_config
    tencentcloud_tcr_service_account

Video on Demand(VOD)
  Data Source
    tencentcloud_vod_adaptive_dynamic_streaming_templates
    tencentcloud_vod_snapshot_by_time_offset_templates
    tencentcloud_vod_super_player_configs
    tencentcloud_vod_image_sprite_templates
    tencentcloud_vod_procedure_templates


  Resource
    tencentcloud_vod_adaptive_dynamic_streaming_template
    tencentcloud_vod_procedure_template
    tencentcloud_vod_snapshot_by_time_offset_template
    tencentcloud_vod_image_sprite_template
    tencentcloud_vod_super_player_config
    tencentcloud_vod_sub_application

Oceanus
  Data Source
    tencentcloud_oceanus_resource_related_job
    tencentcloud_oceanus_savepoint_list
    tencentcloud_oceanus_system_resource
    tencentcloud_oceanus_work_spaces
    tencentcloud_oceanus_clusters
    tencentcloud_oceanus_tree_jobs
    tencentcloud_oceanus_tree_resources
    tencentcloud_oceanus_job_submission_log
    tencentcloud_oceanus_check_savepoint

  Resource
    tencentcloud_oceanus_job
    tencentcloud_oceanus_job_config
    tencentcloud_oceanus_job_copy
    tencentcloud_oceanus_run_job
    tencentcloud_oceanus_stop_job
    tencentcloud_oceanus_trigger_job_savepoint
    tencentcloud_oceanus_resource
    tencentcloud_oceanus_resource_config
    tencentcloud_oceanus_work_space

Virtual Private Cloud(VPC)
  Data Source
    tencentcloud_route_table
    tencentcloud_security_group
    tencentcloud_security_groups
    tencentcloud_address_templates
    tencentcloud_address_template_groups
    tencentcloud_protocol_templates
    tencentcloud_protocol_template_groups
    tencentcloud_subnet
    tencentcloud_vpc
    tencentcloud_vpc_acls
    tencentcloud_vpc_account_attributes
    tencentcloud_vpc_classic_link_instances
    tencentcloud_vpc_gateway_flow_monitor_detail
    tencentcloud_vpc_gateway_flow_qos
    tencentcloud_vpc_cvm_instances
    tencentcloud_vpc_net_detect_states
    tencentcloud_vpc_net_detect_state_check
    tencentcloud_vpc_network_interface_limit
    tencentcloud_vpc_private_ip_addresses
    tencentcloud_vpc_product_quota
    tencentcloud_vpc_resource_dashboard
    tencentcloud_vpc_route_conflicts
    tencentcloud_vpc_security_group_limits
    tencentcloud_vpc_security_group_references
    tencentcloud_vpc_sg_snapshot_file_content
    tencentcloud_vpc_snapshot_files
    tencentcloud_vpc_subnet_resource_dashboard
    tencentcloud_vpc_template_limits
    tencentcloud_vpc_used_ip_address
    tencentcloud_vpc_limits
    tencentcloud_vpc_instances
    tencentcloud_vpc_route_tables
    tencentcloud_vpc_subnets
    tencentcloud_dnats
    tencentcloud_enis
    tencentcloud_ha_vip_eip_attachments
    tencentcloud_ha_vips
    tencentcloud_nat_gateways
    tencentcloud_nat_gateway_snats
    tencentcloud_nats
    tencentcloud_nat_dc_route
    tencentcloud_vpc_bandwidth_package_quota
    tencentcloud_vpc_bandwidth_package_bill_usage

  Resource
    tencentcloud_eni
    tencentcloud_eni_attachment
    tencentcloud_eni_sg_attachment
    tencentcloud_vpc
    tencentcloud_vpc_acl
    tencentcloud_vpc_acl_attachment
    tencentcloud_vpc_traffic_package
    tencentcloud_vpc_snapshot_policy
    tencentcloud_vpc_snapshot_policy_attachment
    tencentcloud_vpc_snapshot_policy_config
    tencentcloud_vpc_net_detect
    tencentcloud_vpc_dhcp_ip
    tencentcloud_vpc_ipv6_cidr_block
    tencentcloud_vpc_ipv6_subnet_cidr_block
    tencentcloud_vpc_ipv6_eni_address
    tencentcloud_vpc_local_gateway
    tencentcloud_vpc_resume_snapshot_instance
    tencentcloud_subnet
    tencentcloud_security_group
    tencentcloud_security_group_rule
    tencentcloud_security_group_rule_set
    tencentcloud_security_group_lite_rule
    tencentcloud_address_template
    tencentcloud_address_template_group
    tencentcloud_protocol_template
    tencentcloud_protocol_template_group
    tencentcloud_route_table
	tencentcloud_route_table_association
    tencentcloud_route_entry
    tencentcloud_route_table_entry
    tencentcloud_dnat
    tencentcloud_nat_gateway
    tencentcloud_nat_gateway_snat
    tencentcloud_nat_refresh_nat_dc_route
    tencentcloud_ha_vip
    tencentcloud_ha_vip_eip_attachment
    tencentcloud_vpc_bandwidth_package
    tencentcloud_vpc_bandwidth_package_attachment
    tencentcloud_ipv6_address_bandwidth

Private Link(PLS)
  Resource
    tencentcloud_vpc_end_point_service
    tencentcloud_vpc_end_point
    tencentcloud_vpc_enable_end_point_connect
    tencentcloud_vpc_end_point_service_white_list

Flow Logs(FL)
  Resource
     tencentcloud_vpc_flow_log
    tencentcloud_vpc_flow_log_config

VPN Connections(VPN)
  Data Source
    tencentcloud_vpn_connections
    tencentcloud_vpn_customer_gateways
    tencentcloud_vpn_gateways
    tencentcloud_vpn_gateway_routes
    tencentcloud_vpn_customer_gateway_vendors
    tencentcloud_vpn_default_health_check_ip

  Resource
    tencentcloud_vpn_customer_gateway
    tencentcloud_vpn_gateway
    tencentcloud_vpn_gateway_route
    tencentcloud_vpn_connection
    tencentcloud_vpn_ssl_server
    tencentcloud_vpn_ssl_client
    tencentcloud_vpn_connection_reset
    tencentcloud_vpn_customer_gateway_configuration_download
    tencentcloud_vpn_gateway_ssl_client_cert
    tencentcloud_vpn_gateway_ccn_routes

MapReduce(EMR)
  Data Source
    tencentcloud_emr
    tencentcloud_emr_nodes
    tencentcloud_emr_cvm_quota

  Resource
    tencentcloud_emr_cluster
    tencentcloud_emr_user_manager

DNSPOD
  Resource
    tencentcloud_dnspod_domain_instance
    tencentcloud_dnspod_domain_alias
    tencentcloud_dnspod_record
    tencentcloud_dnspod_record_group
    tencentcloud_dnspod_modify_record_group_operation
    tencentcloud_dnspod_modify_domain_owner_operation
    tencentcloud_dnspod_download_snapshot_operation
    tencentcloud_dnspod_custom_line
    tencentcloud_dnspod_snapshot_config
    tencentcloud_dnspod_domain_lock

  Data Source
    tencentcloud_dnspod_records
    tencentcloud_dnspod_domain_list
    tencentcloud_dnspod_domain_analytics
    tencentcloud_dnspod_domain_log_list
    tencentcloud_dnspod_record_analytics
    tencentcloud_dnspod_record_line_list
    tencentcloud_dnspod_record_list
    tencentcloud_dnspod_record_type

PrivateDNS
  Resource
    tencentcloud_private_dns_zone
    tencentcloud_private_dns_record
    tencentcloud_private_dns_zone_vpc_attachment
  Data Source
    tencentcloud_private_dns_records

Cloud Log Service(CLS)
  Resource
    tencentcloud_cls_logset
    tencentcloud_cls_topic
    tencentcloud_cls_config
    tencentcloud_cls_config_extra
    tencentcloud_cls_config_attachment
    tencentcloud_cls_machine_group
    tencentcloud_cls_cos_shipper
    tencentcloud_cls_index
    tencentcloud_cls_alarm
    tencentcloud_cls_alarm_notice
    tencentcloud_cls_ckafka_consumer
    tencentcloud_cls_kafka_recharge
    tencentcloud_cls_cos_recharge
    tencentcloud_cls_export
    tencentcloud_cls_scheduled_sql
    tencentcloud_cls_data_transform

  Data Source
    tencentcloud_cls_shipper_tasks
    tencentcloud_cls_machines
    tencentcloud_cls_machine_group_configs

TencentCloud Lighthouse(Lighthouse)
  Resource
    tencentcloud_lighthouse_instance
    tencentcloud_lighthouse_blueprint
    tencentcloud_lighthouse_firewall_rule
    tencentcloud_lighthouse_disk_backup
    tencentcloud_lighthouse_apply_disk_backup
    tencentcloud_lighthouse_disk_attachment
    tencentcloud_lighthouse_key_pair
    tencentcloud_lighthouse_snapshot
    tencentcloud_lighthouse_apply_instance_snapshot
    tencentcloud_lighthouse_start_instance
    tencentcloud_lighthouse_stop_instance
    tencentcloud_lighthouse_reboot_instance
    tencentcloud_lighthouse_key_pair_attachment
    tencentcloud_lighthouse_disk
    tencentcloud_lighthouse_renew_disk
    tencentcloud_lighthouse_renew_instance
    tencentcloud_lighthouse_firewall_template

  Data Source
    tencentcloud_lighthouse_firewall_rules_template
    tencentcloud_lighthouse_bundle
    tencentcloud_lighthouse_zone
    tencentcloud_lighthouse_scene
    tencentcloud_lighthouse_reset_instance_blueprint
    tencentcloud_lighthouse_region
    tencentcloud_lighthouse_instance_vnc_url
    tencentcloud_lighthouse_instance_traffic_package
    tencentcloud_lighthouse_instance_disk_num
    tencentcloud_lighthouse_instance_blueprint
    tencentcloud_lighthouse_disk_config
    tencentcloud_lighthouse_all_scene
    tencentcloud_lighthouse_modify_instance_bundle
    tencentcloud_lighthouse_disks

TencentCloud Elastic Microservice(TEM)
  Resource
    tencentcloud_tem_environment
    tencentcloud_tem_application
    tencentcloud_tem_workload
    tencentcloud_tem_app_config
    tencentcloud_tem_log_config
    tencentcloud_tem_scale_rule
    tencentcloud_tem_gateway
    tencentcloud_tem_application_service

TencentCloud EdgeOne(TEO)
  Data Source
    tencentcloud_teo_zone_available_plans
    tencentcloud_teo_rule_engine_settings

  Resource
    tencentcloud_teo_zone
    tencentcloud_teo_zone_setting
    tencentcloud_teo_origin_group
    tencentcloud_teo_rule_engine
    tencentcloud_teo_application_proxy_rule
    tencentcloud_teo_ownership_verify
    tencentcloud_teo_certificate_config
    tencentcloud_teo_acceleration_domain

TencentCloud ServiceMesh(TCM)
  Data Source
    tencentcloud_tcm_mesh
  Resource
    tencentcloud_tcm_mesh
    tencentcloud_tcm_cluster_attachment
    tencentcloud_tcm_prometheus_attachment
    tencentcloud_tcm_tracing_config
    tencentcloud_tcm_access_log_config

Simple Email Service(SES)
  Data Source
    tencentcloud_ses_receivers
    tencentcloud_ses_send_tasks
    tencentcloud_ses_email_identities
    tencentcloud_ses_black_email_address
    tencentcloud_ses_statistics_report
    tencentcloud_ses_send_email_status

  Resource
    tencentcloud_ses_domain
    tencentcloud_ses_template
    tencentcloud_ses_email_address
    tencentcloud_ses_receiver
    tencentcloud_ses_send_email
    tencentcloud_ses_batch_send_email
    tencentcloud_ses_verify_domain
    tencentcloud_ses_black_list_delete

Security Token Service(STS)
  Data Source
    tencentcloud_sts_caller_identity

TDSQL for MySQL(DCDB)
  Data Source
    tencentcloud_dcdb_instances
    tencentcloud_dcdb_accounts
    tencentcloud_dcdb_databases
    tencentcloud_dcdb_parameters
    tencentcloud_dcdb_shards
    tencentcloud_dcdb_security_groups
    tencentcloud_dcdb_database_objects
    tencentcloud_dcdb_database_tables
    tencentcloud_dcdb_file_download_url
    tencentcloud_dcdb_log_files
    tencentcloud_dcdb_instance_node_info
    tencentcloud_dcdb_orders
    tencentcloud_dcdb_price
    tencentcloud_dcdb_project_security_groups
    tencentcloud_dcdb_projects
    tencentcloud_dcdb_renewal_price
    tencentcloud_dcdb_sale_info
    tencentcloud_dcdb_shard_spec
    tencentcloud_dcdb_slow_logs
    tencentcloud_dcdb_upgrade_price

  Resource
    tencentcloud_dcdb_account
    tencentcloud_dcdb_hourdb_instance
    tencentcloud_dcdb_security_group_attachment
    tencentcloud_dcdb_account_privileges
    tencentcloud_dcdb_db_parameters
    tencentcloud_dcdb_db_sync_mode_config
    tencentcloud_dcdb_encrypt_attributes_config
    tencentcloud_dcdb_instance_config
    tencentcloud_dcdb_cancel_dcn_job_operation
    tencentcloud_dcdb_activate_hour_instance_operation
    tencentcloud_dcdb_isolate_hour_instance_operation
    tencentcloud_dcdb_flush_binlog_operation
    tencentcloud_dcdb_switch_db_instance_ha_operation

Short Message Service(SMS)
  Resource
    tencentcloud_sms_sign
    tencentcloud_sms_template

Cloud Automated Testing(CAT)
  Data Source
    tencentcloud_cat_probe_data
    tencentcloud_cat_node
    tencentcloud_cat_metric_data

  Resource
    tencentcloud_cat_task_set

TencentDB for MariaDB(MariaDB)
  Data Source
    tencentcloud_mariadb_db_instances
    tencentcloud_mariadb_accounts
    tencentcloud_mariadb_security_groups
    tencentcloud_mariadb_database_objects
    tencentcloud_mariadb_databases
    tencentcloud_mariadb_database_table
    tencentcloud_mariadb_dcn_detail
    tencentcloud_mariadb_file_download_url
    tencentcloud_mariadb_flow
    tencentcloud_mariadb_instance_specs
    tencentcloud_mariadb_log_files
    tencentcloud_mariadb_orders
    tencentcloud_mariadb_price
    tencentcloud_mariadb_project_security_groups
    tencentcloud_mariadb_renewal_price
    tencentcloud_mariadb_sale_info
    tencentcloud_mariadb_slow_logs
    tencentcloud_mariadb_upgrade_price

  Resource
    tencentcloud_mariadb_dedicatedcluster_db_instance
    tencentcloud_mariadb_instance
    tencentcloud_mariadb_hour_db_instance
    tencentcloud_mariadb_account
    tencentcloud_mariadb_parameters
    tencentcloud_mariadb_log_file_retention_period
    tencentcloud_mariadb_security_groups
    tencentcloud_mariadb_account_privileges
    tencentcloud_mariadb_operate_hour_db_instance
    tencentcloud_mariadb_backup_time
    tencentcloud_mariadb_cancel_dcn_job
    tencentcloud_mariadb_flush_binlog
    tencentcloud_mariadb_switch_ha
    tencentcloud_mariadb_restart_instance
    tencentcloud_mariadb_renew_instance
    tencentcloud_mariadb_instance_config

Real User Monitoring(RUM)
  Data Source
    tencentcloud_rum_project
    tencentcloud_rum_offline_log_config
    tencentcloud_rum_whitelist
    tencentcloud_rum_taw_instance
    tencentcloud_rum_custom_url
    tencentcloud_rum_event_url
    tencentcloud_rum_fetch_url_info
    tencentcloud_rum_fetch_url
    tencentcloud_rum_group_log
    tencentcloud_rum_log_url_statistics
    tencentcloud_rum_performance_page
    tencentcloud_rum_pv_url_info
    tencentcloud_rum_pv_url_statistics
    tencentcloud_rum_report_count
    tencentcloud_rum_scores
    tencentcloud_rum_set_url_statistics
    tencentcloud_rum_sign
    tencentcloud_rum_static_project
    tencentcloud_rum_static_resource
    tencentcloud_rum_static_url
    tencentcloud_rum_web_vitals_page
    tencentcloud_rum_log_export_list

  Resource
    tencentcloud_rum_project
    tencentcloud_rum_taw_instance
    tencentcloud_rum_whitelist
    tencentcloud_rum_offline_log_config_attachment
    tencentcloud_rum_instance_status_config
    tencentcloud_rum_project_status_config

Cloud Streaming Services(CSS)
  Resource
    tencentcloud_css_watermark
    tencentcloud_css_watermark_rule_attachment
    tencentcloud_css_pull_stream_task
    tencentcloud_css_live_transcode_template
    tencentcloud_css_live_transcode_rule_attachment
    tencentcloud_css_domain
    tencentcloud_css_authenticate_domain_owner_operation
    tencentcloud_css_play_domain_cert_attachment
    tencentcloud_css_play_auth_key_config
    tencentcloud_css_push_auth_key_config
    tencentcloud_css_backup_stream
    tencentcloud_css_callback_rule_attachment
    tencentcloud_css_callback_template
    tencentcloud_css_domain_referer
    tencentcloud_css_enable_optimal_switching
    tencentcloud_css_record_rule_attachment
    tencentcloud_css_snapshot_rule_attachment
    tencentcloud_css_snapshot_template
    tencentcloud_css_pad_template
    tencentcloud_css_pad_rule_attachment
    tencentcloud_css_timeshift_template
    tencentcloud_css_timeshift_rule_attachment
    tencentcloud_css_stream_monitor
    tencentcloud_css_start_stream_monitor
    tencentcloud_css_pull_stream_task_restart

  Data Source
    tencentcloud_css_domains
    tencentcloud_css_backup_stream
    tencentcloud_css_monitor_report
    tencentcloud_css_pad_templates
    tencentcloud_css_pull_stream_task_status
    tencentcloud_css_stream_monitor_list
    tencentcloud_css_time_shift_record_detail
    tencentcloud_css_time_shift_stream_list
    tencentcloud_css_watermarks
    tencentcloud_css_xp2p_detail_info_list

Performance Testing Service(PTS)
  Data Source
    tencentcloud_pts_scenario_with_jobs

  Resource
    tencentcloud_pts_project
    tencentcloud_pts_alert_channel
    tencentcloud_pts_scenario
    tencentcloud_pts_file
    tencentcloud_pts_job
    tencentcloud_pts_cron_job
    tencentcloud_pts_tmp_key_generate
    tencentcloud_pts_cron_job_restart
    tencentcloud_pts_job_abort
    tencentcloud_pts_cron_job_abort

TencentCloud Automation Tools(TAT)
  Data Source
    tencentcloud_tat_command
    tencentcloud_tat_invoker
    tencentcloud_tat_invoker_records
    tencentcloud_tat_agent
    tencentcloud_tat_invocation_task
  Resource
    tencentcloud_tat_command
    tencentcloud_tat_invoker
    tencentcloud_tat_invoker_config
    tencentcloud_tat_invocation_invoke_attachment
    tencentcloud_tat_invocation_command_attachment

Tencent Cloud Organization (TCO)
  Data Source
    tencentcloud_organization_org_auth_node
    tencentcloud_organization_org_financial_by_member
    tencentcloud_organization_org_financial_by_month
    tencentcloud_organization_org_financial_by_product
  Resource
    tencentcloud_organization_instance
    tencentcloud_organization_org_node
    tencentcloud_organization_org_member
    tencentcloud_organization_org_identity
    tencentcloud_organization_org_member_email
    tencentcloud_organization_org_member_auth_identity_attachment
    tencentcloud_organization_org_member_policy_attachment
    tencentcloud_organization_policy_sub_account_attachment
    tencentcloud_organization_quit_organization_operation

TDSQL-C for PostgreSQL(TDCPG)
  Data Source
    tencentcloud_tdcpg_clusters
    tencentcloud_tdcpg_instances
  Resource
    tencentcloud_tdcpg_cluster
    tencentcloud_tdcpg_instance

TencentDB for DBbrain(dbbrain)
  Data Source
    tencentcloud_dbbrain_sql_filters
    tencentcloud_dbbrain_security_audit_log_export_tasks
    tencentcloud_dbbrain_diag_event
    tencentcloud_dbbrain_diag_events
    tencentcloud_dbbrain_diag_history
    tencentcloud_dbbrain_security_audit_log_download_urls
    tencentcloud_dbbrain_slow_log_time_series_stats
    tencentcloud_dbbrain_slow_log_top_sqls
    tencentcloud_dbbrain_slow_log_user_host_stats
    tencentcloud_dbbrain_slow_log_user_sql_advice
    tencentcloud_dbbrain_slow_logs
    tencentcloud_dbbrain_health_scores
    tencentcloud_dbbrain_sql_templates
    tencentcloud_dbbrain_db_space_status
    tencentcloud_dbbrain_top_space_schemas
    tencentcloud_dbbrain_top_space_tables
    tencentcloud_dbbrain_top_space_schema_time_series
    tencentcloud_dbbrain_top_space_table_time_series
    tencentcloud_dbbrain_diag_db_instances
    tencentcloud_dbbrain_mysql_process_list
    tencentcloud_dbbrain_no_primary_key_tables
    tencentcloud_dbbrain_redis_top_big_keys
    tencentcloud_dbbrain_redis_top_key_prefix_list

  Resource
    tencentcloud_dbbrain_sql_filter
    tencentcloud_dbbrain_security_audit_log_export_task
    tencentcloud_dbbrain_db_diag_report_task
    tencentcloud_dbbrain_modify_diag_db_instance_operation
    tencentcloud_dbbrain_tdsql_audit_log

Data Transmission Service(DTS)
  Data Source
    tencentcloud_dts_sync_jobs
    tencentcloud_dts_migrate_jobs
    tencentcloud_dts_compare_tasks
    tencentcloud_dts_migrate_db_instances

  Resource
    tencentcloud_dts_sync_job
    tencentcloud_dts_sync_config
    tencentcloud_dts_sync_check_job_operation
    tencentcloud_dts_sync_job_resume_operation
    tencentcloud_dts_sync_job_start_operation
    tencentcloud_dts_sync_job_stop_operation
    tencentcloud_dts_sync_job_resize_operation
    tencentcloud_dts_sync_job_recover_operation
    tencentcloud_dts_sync_job_isolate_operation
    tencentcloud_dts_sync_job_continue_operation
    tencentcloud_dts_sync_job_pause_operation
    tencentcloud_dts_migrate_service
    tencentcloud_dts_migrate_job
    tencentcloud_dts_migrate_job_config
    tencentcloud_dts_migrate_job_start_operation
    tencentcloud_dts_migrate_job_resume_operation
    tencentcloud_dts_compare_task_stop_operation
    tencentcloud_dts_compare_task

TDMQ for RocketMQ(trocket)
  Data Source
    tencentcloud_tdmq_rocketmq_cluster
    tencentcloud_tdmq_rocketmq_namespace
    tencentcloud_tdmq_rocketmq_topic
    tencentcloud_tdmq_rocketmq_role
    tencentcloud_tdmq_rocketmq_group
    tencentcloud_tdmq_rocketmq_messages

  Resource
    tencentcloud_tdmq_rocketmq_cluster
    tencentcloud_tdmq_rocketmq_namespace
    tencentcloud_tdmq_rocketmq_role
    tencentcloud_tdmq_rocketmq_topic
    tencentcloud_tdmq_rocketmq_group
    tencentcloud_tdmq_rocketmq_environment_role
    tencentcloud_tdmq_send_rocketmq_message
    tencentcloud_tdmq_rocketmq_vip_instance
    tencentcloud_trocket_rocketmq_instance
    tencentcloud_trocket_rocketmq_topic
    tencentcloud_trocket_rocketmq_consumer_group
    tencentcloud_trocket_rocketmq_role

TDMQ for RabbitMQ(trabbit)
  Resource
    tencentcloud_tdmq_rabbitmq_user
    tencentcloud_tdmq_rabbitmq_virtual_host
    tencentcloud_tdmq_rabbitmq_vip_instance


Cloud Infinite(CI)
  Resource
    tencentcloud_ci_bucket_attachment
    tencentcloud_ci_bucket_pic_style
    tencentcloud_ci_hot_link
    tencentcloud_ci_media_snapshot_template
    tencentcloud_ci_media_transcode_template
    tencentcloud_ci_media_animation_template
    tencentcloud_ci_media_concat_template
    tencentcloud_ci_media_video_process_template
    tencentcloud_ci_media_video_montage_template
    tencentcloud_ci_media_voice_separate_template
    tencentcloud_ci_media_super_resolution_template
    tencentcloud_ci_media_pic_process_template
    tencentcloud_ci_media_watermark_template
    tencentcloud_ci_media_tts_template
    tencentcloud_ci_media_transcode_pro_template
    tencentcloud_ci_media_smart_cover_template
    tencentcloud_ci_media_speech_recognition_template
    tencentcloud_ci_guetzli
    tencentcloud_ci_original_image_protection

TDMQ for CMQ
  Data Source
    tencentcloud_tcmq_queue
    tencentcloud_tcmq_topic
    tencentcloud_tcmq_subscribe

  Resource
    tencentcloud_tcmq_queue
    tencentcloud_tcmq_topic
    tencentcloud_tcmq_subscribe

Tencent Service Framework(TSF)
  Data Source
    tencentcloud_tsf_application
    tencentcloud_tsf_application_config
    tencentcloud_tsf_application_file_config
    tencentcloud_tsf_application_public_config
    tencentcloud_tsf_cluster
    tencentcloud_tsf_microservice
    tencentcloud_tsf_unit_rules
    tencentcloud_tsf_config_summary
    tencentcloud_tsf_delivery_config_by_group_id
    tencentcloud_tsf_delivery_configs
    tencentcloud_tsf_public_config_summary
    tencentcloud_tsf_api_group
    tencentcloud_tsf_application_attribute
    tencentcloud_tsf_business_log_configs
    tencentcloud_tsf_api_detail
    tencentcloud_tsf_microservice_api_version
    tencentcloud_tsf_repository
    tencentcloud_tsf_pod_instances
    tencentcloud_tsf_gateway_all_group_apis
    tencentcloud_tsf_group_gateways
    tencentcloud_tsf_usable_unit_namespaces
    tencentcloud_tsf_group_instances
    tencentcloud_tsf_group_config_release
    tencentcloud_tsf_container_group
    tencentcloud_tsf_groups
    tencentcloud_tsf_ms_api_list

  Resource
    tencentcloud_tsf_cluster
    tencentcloud_tsf_microservice
    tencentcloud_tsf_application_config
    tencentcloud_tsf_api_group
    tencentcloud_tsf_namespace
    tencentcloud_tsf_path_rewrite
    tencentcloud_tsf_unit_rule
    tencentcloud_tsf_task
    tencentcloud_tsf_config_template
    tencentcloud_tsf_api_rate_limit_rule
    tencentcloud_tsf_application_release_config
    tencentcloud_tsf_lane
    tencentcloud_tsf_lane_rule
    tencentcloud_tsf_group
    tencentcloud_tsf_application
    tencentcloud_tsf_application_public_config_release
    tencentcloud_tsf_application_public_config
    tencentcloud_tsf_application_file_config_release
    tencentcloud_tsf_instances_attachment
    tencentcloud_tsf_bind_api_group
    tencentcloud_tsf_application_file_config
    tencentcloud_tsf_enable_unit_rule
    tencentcloud_tsf_deploy_container_group
    tencentcloud_tsf_deploy_vm_group
    tencentcloud_tsf_release_api_group
    tencentcloud_tsf_operate_container_group
    tencentcloud_tsf_operate_group
    tencentcloud_tsf_unit_namespace

Media Processing Service(MPS)
  Data Source
    tencentcloud_mps_schedules
    tencentcloud_mps_tasks
    tencentcloud_mps_parse_live_stream_process_notification
    tencentcloud_mps_parse_notification
    tencentcloud_mps_media_meta_data

  Resource
    tencentcloud_mps_workflow
    tencentcloud_mps_enable_workflow_config
    tencentcloud_mps_transcode_template
    tencentcloud_mps_watermark_template
    tencentcloud_mps_image_sprite_template
    tencentcloud_mps_snapshot_by_timeoffset_template
    tencentcloud_mps_sample_snapshot_template
    tencentcloud_mps_animated_graphics_template
    tencentcloud_mps_ai_recognition_template
    tencentcloud_mps_ai_analysis_template
    tencentcloud_mps_adaptive_dynamic_streaming_template
    tencentcloud_mps_person_sample
    tencentcloud_mps_withdraws_watermark_operation
    tencentcloud_mps_process_live_stream_operation
    tencentcloud_mps_edit_media_operation
    tencentcloud_mps_word_sample
    tencentcloud_mps_schedule
    tencentcloud_mps_enable_schedule_config
    tencentcloud_mps_flow
    tencentcloud_mps_input
    tencentcloud_mps_output
    tencentcloud_mps_content_review_template
    tencentcloud_mps_start_flow_operation
    tencentcloud_mps_event
    tencentcloud_mps_manage_task_operation
    tencentcloud_mps_execute_function_operation
    tencentcloud_mps_process_media_operation

Cloud HDFS(CHDFS)
  Data Source
    tencentcloud_chdfs_access_groups
    tencentcloud_chdfs_mount_points
    tencentcloud_chdfs_file_systems

  Resource
    tencentcloud_chdfs_access_group
    tencentcloud_chdfs_access_rule
    tencentcloud_chdfs_file_system
    tencentcloud_chdfs_life_cycle_rule
    tencentcloud_chdfs_mount_point
    tencentcloud_chdfs_mount_point_attachment

StreamLive(MDL)
  Resource
    tencentcloud_mdl_stream_live_input

Application Performance Management(APM)
  Resource
    tencentcloud_apm_instance

Tencent Cloud Service Engine(TSE)
  Data Source
    tencentcloud_tse_access_address
    tencentcloud_tse_nacos_replicas
    tencentcloud_tse_zookeeper_replicas
    tencentcloud_tse_zookeeper_server_interfaces
    tencentcloud_tse_nacos_server_interfaces
    tencentcloud_tse_groups
    tencentcloud_tse_gateways
    tencentcloud_tse_gateway_nodes
    tencentcloud_tse_gateway_routes
    tencentcloud_tse_gateway_canary_rules
    tencentcloud_tse_gateway_services
    tencentcloud_tse_gateway_certificates

  Resource
    tencentcloud_tse_instance
    tencentcloud_tse_cngw_service
    tencentcloud_tse_cngw_canary_rule
    tencentcloud_tse_cngw_gateway
    tencentcloud_tse_cngw_group
    tencentcloud_tse_cngw_service_rate_limit
    tencentcloud_tse_cngw_route
    tencentcloud_tse_cngw_route_rate_limit
    tencentcloud_tse_cngw_certificate

ClickHouse(CDWCH)
  Data Source
    tencentcloud_clickhouse_backup_jobs
    tencentcloud_clickhouse_backup_job_detail
    tencentcloud_clickhouse_backup_tables

  Resource
    tencentcloud_clickhouse_instance
    tencentcloud_clickhouse_backup
    tencentcloud_clickhouse_backup_strategy
    tencentcloud_clickhouse_recover_backup_job
    tencentcloud_clickhouse_delete_backup_data
    tencentcloud_clickhouse_account
    tencentcloud_clickhouse_account_permission

Tag
  Resource
    tencentcloud_tag
    tencentcloud_tag_attachment

EventBridge(EB)
  Data Source
    tencentcloud_eb_bus
    tencentcloud_eb_event_rules
    tencentcloud_eb_platform_event_names
    tencentcloud_eb_platform_event_patterns
    tencentcloud_eb_platform_products
    tencentcloud_eb_plateform_event_template

  Resource
    tencentcloud_eb_event_transform
    tencentcloud_eb_event_bus
    tencentcloud_eb_event_rule
    tencentcloud_eb_event_target
    tencentcloud_eb_put_events
    tencentcloud_eb_event_connector

Data Lake Compute(DLC)
  Data Source
    tencentcloud_dlc_describe_user_type
    tencentcloud_dlc_describe_user_info
    tencentcloud_dlc_describe_user_roles
    tencentcloud_dlc_describe_data_engine
    tencentcloud_dlc_describe_data_engine_image_versions
    tencentcloud_dlc_describe_data_engine_python_spark_images
    tencentcloud_dlc_describe_engine_usage_info
    tencentcloud_dlc_describe_work_group_info
    tencentcloud_dlc_check_data_engine_image_can_be_rollback
    tencentcloud_dlc_check_data_engine_image_can_be_upgrade
    tencentcloud_dlc_check_data_engine_config_pairs_validity
    tencentcloud_dlc_describe_updatable_data_engines

  Resource
    tencentcloud_dlc_work_group
    tencentcloud_dlc_user
    tencentcloud_dlc_data_engine
    tencentcloud_dlc_rollback_data_engine_image_operation
    tencentcloud_dlc_add_users_to_work_group_attachment
    tencentcloud_dlc_store_location_config
    tencentcloud_dlc_suspend_resume_data_engine
    tencentcloud_dlc_modify_data_engine_description_operation
    tencentcloud_dlc_modify_user_typ_operation
    tencentcloud_dlc_renew_data_engine_operation
    tencentcloud_dlc_restart_data_engine_operation
    tencentcloud_dlc_switch_data_engine_image_operation
    tencentcloud_dlc_update_data_engine_config_operation
    tencentcloud_dlc_upgrade_data_engine_image_operation
    tencentcloud_dlc_user_data_engine_config
    tencentcloud_dlc_update_row_filter_operation
    tencentcloud_dlc_bind_work_groups_to_user_attachment

Web Application Firewall(WAF)
  Data Source
    tencentcloud_waf_ciphers
    tencentcloud_waf_tls_versions
    tencentcloud_waf_domains
    tencentcloud_waf_find_domains
    tencentcloud_waf_ports
    tencentcloud_waf_user_domains
    tencentcloud_waf_attack_log_histogram
    tencentcloud_waf_attack_log_list
    tencentcloud_waf_attack_overview
    tencentcloud_waf_attack_total_count
    tencentcloud_waf_peak_points
    tencentcloud_waf_instance_qps_limit
    tencentcloud_waf_user_clb_regions

  Resource
    tencentcloud_waf_custom_rule
    tencentcloud_waf_custom_white_rule
    tencentcloud_waf_clb_domain
    tencentcloud_waf_saas_domain
    tencentcloud_waf_clb_instance
    tencentcloud_waf_saas_instance
    tencentcloud_waf_anti_fake
    tencentcloud_waf_anti_info_leak
    tencentcloud_waf_auto_deny_rules
    tencentcloud_waf_module_status
    tencentcloud_waf_protection_mode
    tencentcloud_waf_web_shell
    tencentcloud_waf_cc
    tencentcloud_waf_cc_auto_status
    tencentcloud_waf_cc_session
    tencentcloud_waf_ip_access_control
    tencentcloud_waf_modify_access_period

Wedata
  Data Source
    tencentcloud_wedata_rule_templates
    tencentcloud_wedata_data_source_list
    tencentcloud_wedata_data_source_without_info

  Resource
    tencentcloud_wedata_datasource
    tencentcloud_wedata_function
    tencentcloud_wedata_resource
    tencentcloud_wedata_script
    tencentcloud_wedata_dq_rule
    tencentcloud_wedata_rule_template
    tencentcloud_wedata_baseline
    tencentcloud_wedata_integration_offline_task
    tencentcloud_wedata_integration_realtime_task
    tencentcloud_wedata_integration_task_node

Cloud Firewall(CFW)
  Data Source
    tencentcloud_cfw_nat_fw_switches
    tencentcloud_cfw_vpc_fw_switches
    tencentcloud_cfw_edge_fw_switches

  Resource
    tencentcloud_cfw_address_template
    tencentcloud_cfw_block_ignore
    tencentcloud_cfw_edge_policy
    tencentcloud_cfw_nat_instance
    tencentcloud_cfw_nat_policy
    tencentcloud_cfw_vpc_instance
    tencentcloud_cfw_vpc_policy
    tencentcloud_cfw_sync_asset
    tencentcloud_cfw_sync_route
    tencentcloud_cfw_nat_firewall_switch
    tencentcloud_cfw_vpc_firewall_switch
    tencentcloud_cfw_edge_firewall_switch

Bastion Host(BH)
  Resource
    tencentcloud_dasb_acl
    tencentcloud_dasb_cmd_template
    tencentcloud_dasb_device_group
    tencentcloud_dasb_user
    tencentcloud_dasb_device_account
    tencentcloud_dasb_device_group_members
    tencentcloud_dasb_user_group_members
    tencentcloud_dasb_bind_device_resource
    tencentcloud_dasb_resource
    tencentcloud_dasb_device
    tencentcloud_dasb_user_group
    tencentcloud_dasb_reset_user
    tencentcloud_dasb_bind_device_account_private_key
    tencentcloud_dasb_bind_device_account_password

Cwp
  Data Source
    tencentcloud_cwp_machines_simple

  Resource
    tencentcloud_cwp_license_order
    tencentcloud_cwp_license_bind_attachment

Business Intelligence(BI)
  Data Source
    tencentcloud_bi_project
    tencentcloud_bi_user_project

  Resource
    tencentcloud_bi_project
    tencentcloud_bi_user_role
    tencentcloud_bi_project_user_role
    tencentcloud_bi_datasource
    tencentcloud_bi_datasource_cloud
    tencentcloud_bi_embed_token_apply
    tencentcloud_bi_embed_interval_apply

CDWPG
  Resource
    tencentcloud_cdwpg_instance
*/
package tencentcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/connectivity"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

const (
	PROVIDER_SECRET_ID                    = "TENCENTCLOUD_SECRET_ID"
	PROVIDER_SECRET_KEY                   = "TENCENTCLOUD_SECRET_KEY"
	PROVIDER_SECURITY_TOKEN               = "TENCENTCLOUD_SECURITY_TOKEN"
	PROVIDER_REGION                       = "TENCENTCLOUD_REGION"
	PROVIDER_PROTOCOL                     = "TENCENTCLOUD_PROTOCOL"
	PROVIDER_DOMAIN                       = "TENCENTCLOUD_DOMAIN"
	PROVIDER_ASSUME_ROLE_ARN              = "TENCENTCLOUD_ASSUME_ROLE_ARN"
	PROVIDER_ASSUME_ROLE_SESSION_NAME     = "TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME"
	PROVIDER_ASSUME_ROLE_SESSION_DURATION = "TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION"
	PROVIDER_SHARED_CREDENTIALS_DIR       = "TENCENTCLOUD_SHARED_CREDENTIALS_DIR"
	PROVIDER_PROFILE                      = "TENCENTCLOUD_PROFILE"
)

const (
	DEFAULT_REGION  = "ap-guangzhou"
	DEFAULT_PROFILE = "default"
)

type TencentCloudClient struct {
	apiV3Conn *connectivity.TencentCloudClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECURITY_TOKEN, nil),
				Description: "TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).",
				Sensitive:   true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description: "This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is ap-guangzhou.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "HTTPS"),
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS"}),
				Description:  "The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"assume_role": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_ARN, nil),
							Description: "The ARN of the role to assume. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN`.",
						},
						"session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_SESSION_NAME, nil),
							Description: "The session name to use when making the AssumeRole call. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`.",
						},
						"session_duration": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); v != "" {
									return strconv.Atoi(v)
								}
								return 7200, nil
							},
							ValidateFunc: validateIntegerInRange(0, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION`.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call. Its content must not contains `principal` elements. Notice: more syntax references, please refer to: [policies syntax logic](https://intl.cloud.tencent.com/document/product/598/10603).",
						},
					},
				},
			},
			"shared_credentials_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SHARED_CREDENTIALS_DIR, nil),
				Description: "The directory of the shared credentials. It can also be sourced from the `TENCENTCLOUD_SHARED_CREDENTIALS_DIR` environment variable. If not set this defaults to ~/.tccli.",
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_PROFILE, nil),
				Description: "The profile name as set in the shared credentials. It can also be sourced from the `TENCENTCLOUD_PROFILE` environment variable. If not set, the default profile created with `tccli configure` will be used.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tencentcloud_availability_regions":                         dataSourceTencentCloudAvailabilityRegions(),
			"tencentcloud_emr":                                          dataSourceTencentCloudEmr(),
			"tencentcloud_emr_nodes":                                    dataSourceTencentCloudEmrNodes(),
			"tencentcloud_emr_cvm_quota":                                dataSourceTencentCloudEmrCvmQuota(),
			"tencentcloud_availability_zones":                           dataSourceTencentCloudAvailabilityZones(),
			"tencentcloud_availability_zones_by_product":                dataSourceTencentCloudAvailabilityZonesByProduct(),
			"tencentcloud_projects":                                     dataSourceTencentCloudProjects(),
			"tencentcloud_instances":                                    dataSourceTencentCloudInstances(),
			"tencentcloud_instances_set":                                dataSourceTencentCloudInstancesSet(),
			"tencentcloud_reserved_instances":                           dataSourceTencentCloudReservedInstances(),
			"tencentcloud_placement_groups":                             dataSourceTencentCloudPlacementGroups(),
			"tencentcloud_key_pairs":                                    dataSourceTencentCloudKeyPairs(),
			"tencentcloud_image":                                        dataSourceTencentCloudImage(),
			"tencentcloud_images":                                       dataSourceTencentCloudImages(),
			"tencentcloud_instance_types":                               dataSourceInstanceTypes(),
			"tencentcloud_reserved_instance_configs":                    dataSourceTencentCloudReservedInstanceConfigs(),
			"tencentcloud_vpc_instances":                                dataSourceTencentCloudVpcInstances(),
			"tencentcloud_vpc_subnets":                                  dataSourceTencentCloudVpcSubnets(),
			"tencentcloud_vpc_route_tables":                             dataSourceTencentCloudVpcRouteTables(),
			"tencentcloud_vpc":                                          dataSourceTencentCloudVpc(),
			"tencentcloud_vpc_acls":                                     dataSourceTencentCloudVpcAcls(),
			"tencentcloud_vpc_bandwidth_package_quota":                  dataSourceTencentCloudVpcBandwidthPackageQuota(),
			"tencentcloud_vpc_bandwidth_package_bill_usage":             dataSourceTencentCloudVpcBandwidthPackageBillUsage(),
			"tencentcloud_vpc_account_attributes":                       dataSourceTencentCloudVpcAccountAttributes(),
			"tencentcloud_vpc_classic_link_instances":                   dataSourceTencentCloudVpcClassicLinkInstances(),
			"tencentcloud_vpc_gateway_flow_monitor_detail":              dataSourceTencentCloudVpcGatewayFlowMonitorDetail(),
			"tencentcloud_vpc_gateway_flow_qos":                         dataSourceTencentCloudVpcGatewayFlowQos(),
			"tencentcloud_vpc_cvm_instances":                            dataSourceTencentCloudVpcCvmInstances(),
			"tencentcloud_vpc_net_detect_states":                        dataSourceTencentCloudVpcNetDetectStates(),
			"tencentcloud_vpc_network_interface_limit":                  dataSourceTencentCloudVpcNetworkInterfaceLimit(),
			"tencentcloud_vpc_private_ip_addresses":                     dataSourceTencentCloudVpcPrivateIpAddresses(),
			"tencentcloud_vpc_product_quota":                            dataSourceTencentCloudVpcProductQuota(),
			"tencentcloud_vpc_resource_dashboard":                       dataSourceTencentCloudVpcResourceDashboard(),
			"tencentcloud_vpc_route_conflicts":                          dataSourceTencentCloudVpcRouteConflicts(),
			"tencentcloud_vpc_security_group_limits":                    dataSourceTencentCloudVpcSecurityGroupLimits(),
			"tencentcloud_vpc_security_group_references":                dataSourceTencentCloudVpcSecurityGroupReferences(),
			"tencentcloud_vpc_sg_snapshot_file_content":                 dataSourceTencentCloudVpcSgSnapshotFileContent(),
			"tencentcloud_vpc_snapshot_files":                           dataSourceTencentCloudVpcSnapshotFiles(),
			"tencentcloud_vpc_subnet_resource_dashboard":                dataSourceTencentCloudVpcSubnetResourceDashboard(),
			"tencentcloud_vpc_template_limits":                          dataSourceTencentCloudVpcTemplateLimits(),
			"tencentcloud_vpc_limits":                                   dataSourceTencentCloudVpcLimits(),
			"tencentcloud_vpc_used_ip_address":                          dataSourceTencentCloudVpcUsedIpAddress(),
			"tencentcloud_vpc_net_detect_state_check":                   dataSourceTencentCloudVpcNetDetectStateCheck(),
			"tencentcloud_subnet":                                       dataSourceTencentCloudSubnet(),
			"tencentcloud_route_table":                                  dataSourceTencentCloudRouteTable(),
			"tencentcloud_domains":                                      dataSourceTencentCloudDomains(),
			"tencentcloud_eip":                                          dataSourceTencentCloudEip(),
			"tencentcloud_eips":                                         dataSourceTencentCloudEips(),
			"tencentcloud_eip_address_quota":                            dataSourceTencentCloudEipAddressQuota(),
			"tencentcloud_eip_network_account_type":                     dataSourceTencentCloudEipNetworkAccountType(),
			"tencentcloud_enis":                                         dataSourceTencentCloudEnis(),
			"tencentcloud_nats":                                         dataSourceTencentCloudNats(),
			"tencentcloud_dnats":                                        dataSourceTencentCloudDnats(),
			"tencentcloud_nat_gateways":                                 dataSourceTencentCloudNatGateways(),
			"tencentcloud_nat_gateway_snats":                            dataSourceTencentCloudNatGatewaySnats(),
			"tencentcloud_nat_dc_route":                                 dataSourceTencentCloudNatDcRoute(),
			"tencentcloud_oceanus_resource_related_job":                 dataSourceTencentCloudOceanusResourceRelatedJob(),
			"tencentcloud_oceanus_savepoint_list":                       dataSourceTencentCloudOceanusSavepointList(),
			"tencentcloud_oceanus_system_resource":                      dataSourceTencentCloudOceanusSystemResource(),
			"tencentcloud_oceanus_work_spaces":                          dataSourceTencentCloudOceanusWorkSpaces(),
			"tencentcloud_oceanus_clusters":                             dataSourceTencentCloudOceanusClusters(),
			"tencentcloud_oceanus_tree_jobs":                            dataSourceTencentCloudOceanusTreeJobs(),
			"tencentcloud_oceanus_tree_resources":                       dataSourceTencentCloudOceanusTreeResources(),
			"tencentcloud_oceanus_job_submission_log":                   dataSourceTencentCloudOceanusJobSubmissionLog(),
			"tencentcloud_oceanus_check_savepoint":                      dataSourceTencentCloudOceanusCheckSavepoint(),
			"tencentcloud_vpn_customer_gateways":                        dataSourceTencentCloudVpnCustomerGateways(),
			"tencentcloud_vpn_gateways":                                 dataSourceTencentCloudVpnGateways(),
			"tencentcloud_vpn_gateway_routes":                           dataSourceTencentCloudVpnGatewayRoutes(),
			"tencentcloud_vpn_connections":                              dataSourceTencentCloudVpnConnections(),
			"tencentcloud_vpn_customer_gateway_vendors":                 dataSourceTencentCloudVpnCustomerGatewayVendors(),
			"tencentcloud_vpn_default_health_check_ip":                  dataSourceTencentCloudVpnDefaultHealthCheckIp(),
			"tencentcloud_ha_vips":                                      dataSourceTencentCloudHaVips(),
			"tencentcloud_ha_vip_eip_attachments":                       dataSourceTencentCloudHaVipEipAttachments(),
			"tencentcloud_ccn_instances":                                dataSourceTencentCloudCcnInstances(),
			"tencentcloud_ccn_bandwidth_limits":                         dataSourceTencentCloudCcnBandwidthLimits(),
			"tencentcloud_ccn_cross_border_compliance":                  dataSourceTencentCloudCcnCrossBorderCompliance(),
			"tencentcloud_ccn_tenant_instances":                         dataSourceTencentCloudCcnTenantInstance(),
			"tencentcloud_ccn_cross_border_flow_monitor":                dataSourceTencentCloudCcnCrossBorderFlowMonitor(),
			"tencentcloud_ccn_cross_border_region_bandwidth_limits":     dataSourceTencentCloudCcnCrossBorderRegionBandwidthLimits(),
			"tencentcloud_dc_instances":                                 dataSourceTencentCloudDcInstances(),
			"tencentcloud_dc_access_points":                             dataSourceTencentCloudDcAccessPoints(),
			"tencentcloud_dc_internet_address_quota":                    dataSourceTencentCloudDcInternetAddressQuota(),
			"tencentcloud_dc_internet_address_statistics":               dataSourceTencentCloudDcInternetAddressStatistics(),
			"tencentcloud_dc_public_direct_connect_tunnel_routes":       dataSourceTencentCloudDcPublicDirectConnectTunnelRoutes(),
			"tencentcloud_dcx_instances":                                dataSourceTencentCloudDcxInstances(),
			"tencentcloud_dc_gateway_instances":                         dataSourceTencentCloudDcGatewayInstances(),
			"tencentcloud_dc_gateway_ccn_routes":                        dataSourceTencentCloudDcGatewayCCNRoutes(),
			"tencentcloud_security_group":                               dataSourceTencentCloudSecurityGroup(),
			"tencentcloud_security_groups":                              dataSourceTencentCloudSecurityGroups(),
			"tencentcloud_kubernetes_clusters":                          dataSourceTencentCloudKubernetesClusters(),
			"tencentcloud_kubernetes_charts":                            dataSourceTencentCloudKubernetesCharts(),
			"tencentcloud_kubernetes_cluster_levels":                    datasourceTencentCloudKubernetesClusterLevels(),
			"tencentcloud_kubernetes_cluster_common_names":              datasourceTencentCloudKubernetesClusterCommonNames(),
			"tencentcloud_kubernetes_cluster_authentication_options":    dataSourceTencentCloudKubernetesClusterAuthenticationOptions(),
			"tencentcloud_kubernetes_available_cluster_versions":        dataSourceTencentCloudKubernetesAvailableClusterVersions(),
			"tencentcloud_eks_clusters":                                 dataSourceTencentCloudEKSClusters(),
			"tencentcloud_eks_cluster_credential":                       datasourceTencentCloudEksClusterCredential(),
			"tencentcloud_container_clusters":                           dataSourceTencentCloudContainerClusters(),
			"tencentcloud_container_cluster_instances":                  dataSourceTencentCloudContainerClusterInstances(),
			"tencentcloud_mysql_backup_list":                            dataSourceTencentMysqlBackupList(),
			"tencentcloud_mysql_zone_config":                            dataSourceTencentMysqlZoneConfig(),
			"tencentcloud_mysql_parameter_list":                         dataSourceTencentCloudMysqlParameterList(),
			"tencentcloud_mysql_default_params":                         datasourceTencentCloudMysqlDefaultParams(),
			"tencentcloud_mysql_instance":                               dataSourceTencentCloudMysqlInstance(),
			"tencentcloud_mysql_backup_overview":                        dataSourceTencentCloudMysqlBackupOverview(),
			"tencentcloud_mysql_backup_summaries":                       dataSourceTencentCloudMysqlBackupSummaries(),
			"tencentcloud_mysql_bin_log":                                dataSourceTencentCloudMysqlBinLog(),
			"tencentcloud_mysql_binlog_backup_overview":                 dataSourceTencentCloudMysqlBinlogBackupOverview(),
			"tencentcloud_mysql_clone_list":                             dataSourceTencentCloudMysqlCloneList(),
			"tencentcloud_mysql_data_backup_overview":                   dataSourceTencentCloudMysqlDataBackupOverview(),
			"tencentcloud_mysql_db_features":                            dataSourceTencentCloudMysqlDbFeatures(),
			"tencentcloud_mysql_inst_tables":                            dataSourceTencentCloudMysqlInstTables(),
			"tencentcloud_mysql_instance_charset":                       dataSourceTencentCloudMysqlInstanceCharset(),
			"tencentcloud_mysql_instance_info":                          dataSourceTencentCloudMysqlInstanceInfo(),
			"tencentcloud_mysql_instance_param_record":                  dataSourceTencentCloudMysqlInstanceParamRecord(),
			"tencentcloud_mysql_instance_reboot_time":                   dataSourceTencentCloudMysqlInstanceRebootTime(),
			"tencentcloud_mysql_proxy_custom":                           dataSourceTencentCloudMysqlProxyCustom(),
			"tencentcloud_mysql_rollback_range_time":                    dataSourceTencentCloudMysqlRollbackRangeTime(),
			"tencentcloud_mysql_slow_log":                               dataSourceTencentCloudMysqlSlowLog(),
			"tencentcloud_mysql_slow_log_data":                          dataSourceTencentCloudMysqlSlowLogData(),
			"tencentcloud_mysql_supported_privileges":                   dataSourceTencentCloudMysqlSupportedPrivileges(),
			"tencentcloud_mysql_switch_record":                          dataSourceTencentCloudMysqlSwitchRecord(),
			"tencentcloud_mysql_user_task":                              dataSourceTencentCloudMysqlUserTask(),
			"tencentcloud_mysql_databases":                              dataSourceTencentCloudMysqlDatabases(),
			"tencentcloud_mysql_error_log":                              dataSourceTencentCloudMysqlErrorLog(),
			"tencentcloud_mysql_project_security_group":                 dataSourceTencentCloudMysqlProjectSecurityGroup(),
			"tencentcloud_mysql_ro_min_scale":                           dataSourceTencentCloudMysqlRoMinScale(),
			"tencentcloud_cos_bucket_object":                            dataSourceTencentCloudCosBucketObject(),
			"tencentcloud_cos_buckets":                                  dataSourceTencentCloudCosBuckets(),
			"tencentcloud_cos_batchs":                                   dataSourceTencentCloudCosBatchs(),
			"tencentcloud_cos_bucket_inventorys":                        dataSourceTencentCloudCosBucketInventorys(),
			"tencentcloud_cos_bucket_multipart_uploads":                 dataSourceTencentCloudCosBucketMultipartUploads(),
			"tencentcloud_cfs_file_systems":                             dataSourceTencentCloudCfsFileSystems(),
			"tencentcloud_cfs_access_groups":                            dataSourceTencentCloudCfsAccessGroups(),
			"tencentcloud_cfs_access_rules":                             dataSourceTencentCloudCfsAccessRules(),
			"tencentcloud_cfs_mount_targets":                            dataSourceTencentCloudCfsMountTargets(),
			"tencentcloud_cfs_file_system_clients":                      dataSourceTencentCloudCfsFileSystemClients(),
			"tencentcloud_cfs_available_zone":                           dataSourceTencentCloudCfsAvailableZone(),
			"tencentcloud_redis_zone_config":                            dataSourceTencentRedisZoneConfig(),
			"tencentcloud_redis_instances":                              dataSourceTencentRedisInstances(),
			"tencentcloud_redis_backup":                                 dataSourceTencentCloudRedisBackup(),
			"tencentcloud_redis_backup_download_info":                   dataSourceTencentCloudRedisBackupDownloadInfo(),
			"tencentcloud_redis_param_records":                          dataSourceTencentCloudRedisRecordsParam(),
			"tencentcloud_redis_instance_shards":                        dataSourceTencentCloudRedisInstanceShards(),
			"tencentcloud_redis_instance_zone_info":                     dataSourceTencentCloudRedisInstanceZoneInfo(),
			"tencentcloud_redis_instance_task_list":                     dataSourceTencentCloudRedisInstanceTaskList(),
			"tencentcloud_redis_instance_node_info":                     dataSourceTencentCloudRedisInstanceNodeInfo(),
			"tencentcloud_as_scaling_configs":                           dataSourceTencentCloudAsScalingConfigs(),
			"tencentcloud_as_scaling_groups":                            dataSourceTencentCloudAsScalingGroups(),
			"tencentcloud_as_scaling_policies":                          dataSourceTencentCloudAsScalingPolicies(),
			"tencentcloud_cbs_storages":                                 dataSourceTencentCloudCbsStorages(),
			"tencentcloud_cbs_storages_set":                             dataSourceTencentCloudCbsStoragesSet(),
			"tencentcloud_cbs_snapshots":                                dataSourceTencentCloudCbsSnapshots(),
			"tencentcloud_cbs_snapshot_policies":                        dataSourceTencentCloudCbsSnapshotPolicies(),
			"tencentcloud_clb_instances":                                dataSourceTencentCloudClbInstances(),
			"tencentcloud_clb_listeners":                                dataSourceTencentCloudClbListeners(),
			"tencentcloud_clb_listener_rules":                           dataSourceTencentCloudClbListenerRules(),
			"tencentcloud_clb_attachments":                              dataSourceTencentCloudClbServerAttachments(),
			"tencentcloud_clb_redirections":                             dataSourceTencentCloudClbRedirections(),
			"tencentcloud_clb_target_groups":                            dataSourceTencentCloudClbTargetGroups(),
			"tencentcloud_clb_cluster_resources":                        dataSourceTencentCloudClbClusterResources(),
			"tencentcloud_clb_cross_targets":                            dataSourceTencentCloudClbCrossTargets(),
			"tencentcloud_clb_exclusive_clusters":                       dataSourceTencentCloudClbExclusiveClusters(),
			"tencentcloud_clb_idle_instances":                           dataSourceTencentCloudClbIdleInstances(),
			"tencentcloud_clb_listeners_by_targets":                     dataSourceTencentCloudClbListenersByTargets(),
			"tencentcloud_clb_instance_by_cert_id":                      dataSourceTencentCloudClbInstanceByCertId(),
			"tencentcloud_clb_instance_traffic":                         dataSourceTencentCloudClbInstanceTraffic(),
			"tencentcloud_clb_instance_detail":                          dataSourceTencentCloudClbInstanceDetail(),
			"tencentcloud_clb_resources":                                dataSourceTencentCloudClbResources(),
			"tencentcloud_clb_target_group_list":                        dataSourceTencentCloudClbTargetGroupList(),
			"tencentcloud_clb_target_health":                            dataSourceTencentCloudClbTargetHealth(),
			"tencentcloud_elasticsearch_instances":                      dataSourceTencentCloudElasticsearchInstances(),
			"tencentcloud_elasticsearch_instance_logs":                  dataSourceTencentCloudElasticsearchInstanceLogs(),
			"tencentcloud_elasticsearch_instance_operations":            dataSourceTencentCloudElasticsearchInstanceOperations(),
			"tencentcloud_elasticsearch_logstash_instance_logs":         dataSourceTencentCloudElasticsearchLogstashInstanceLogs(),
			"tencentcloud_elasticsearch_logstash_instance_operations":   dataSourceTencentCloudElasticsearchLogstashInstanceOperations(),
			"tencentcloud_elasticsearch_views":                          dataSourceTencentCloudElasticsearchViews(),
			"tencentcloud_elasticsearch_diagnose":                       dataSourceTencentCloudElasticsearchDiagnose(),
			"tencentcloud_elasticsearch_instance_plugin_list":           dataSourceTencentCloudElasticsearchInstancePluginList(),
			"tencentcloud_mongodb_zone_config":                          dataSourceTencentCloudMongodbZoneConfig(),
			"tencentcloud_mongodb_instances":                            dataSourceTencentCloudMongodbInstances(),
			"tencentcloud_mongodb_instance_backups":                     dataSourceTencentCloudMongodbInstanceBackups(),
			"tencentcloud_mongodb_instance_connections":                 dataSourceTencentCloudMongodbInstanceConnections(),
			"tencentcloud_mongodb_instance_current_op":                  dataSourceTencentCloudMongodbInstanceCurrentOp(),
			"tencentcloud_mongodb_instance_params":                      dataSourceTencentCloudMongodbInstanceParams(),
			"tencentcloud_mongodb_instance_slow_log":                    dataSourceTencentCloudMongodbInstanceSlowLog(),
			"tencentcloud_dayu_cc_https_policies":                       dataSourceTencentCloudDayuCCHttpsPolicies(),
			"tencentcloud_dayu_cc_http_policies":                        dataSourceTencentCloudDayuCCHttpPolicies(),
			"tencentcloud_dayu_ddos_policies":                           dataSourceTencentCloudDayuDdosPolicies(),
			"tencentcloud_dayu_ddos_policy_cases":                       dataSourceTencentCloudDayuDdosPolicyCases(),
			"tencentcloud_dayu_ddos_policy_attachments":                 dataSourceTencentCloudDayuDdosPolicyAttachments(),
			"tencentcloud_dayu_l4_rules":                                dataSourceTencentCloudDayuL4Rules(),
			"tencentcloud_dayu_l4_rules_v2":                             dataSourceTencentCloudDayuL4RulesV2(),
			"tencentcloud_dayu_l7_rules":                                dataSourceTencentCloudDayuL7Rules(),
			"tencentcloud_dayu_l7_rules_v2":                             dataSourceTencentCloudDayuL7RulesV2(),
			"tencentcloud_antiddos_pending_risk_info":                   dataSourceTencentCloudAntiddosPendingRiskInfo(),
			"tencentcloud_antiddos_overview_index":                      dataSourceTencentCloudAntiddosOverviewIndex(),
			"tencentcloud_antiddos_overview_ddos_trend":                 dataSourceTencentCloudAntiddosOverviewDdosTrend(),
			"tencentcloud_antiddos_overview_ddos_event_list":            dataSourceTencentCloudAntiddosOverviewDdosEventList(),
			"tencentcloud_antiddos_overview_cc_trend":                   dataSourceTencentCloudAntiddosOverviewCcTrend(),
			"tencentcloud_gaap_proxies":                                 dataSourceTencentCloudGaapProxies(),
			"tencentcloud_gaap_realservers":                             dataSourceTencentCloudGaapRealservers(),
			"tencentcloud_gaap_layer4_listeners":                        dataSourceTencentCloudGaapLayer4Listeners(),
			"tencentcloud_gaap_layer7_listeners":                        dataSourceTencentCloudGaapLayer7Listeners(),
			"tencentcloud_gaap_http_domains":                            dataSourceTencentCloudGaapHttpDomains(),
			"tencentcloud_gaap_http_rules":                              dataSourceTencentCloudGaapHttpRules(),
			"tencentcloud_gaap_security_policies":                       dataSourceTencentCloudGaapSecurityPolices(),
			"tencentcloud_gaap_security_rules":                          dataSourceTencentCloudGaapSecurityRules(),
			"tencentcloud_gaap_certificates":                            dataSourceTencentCloudGaapCertificates(),
			"tencentcloud_gaap_domain_error_pages":                      dataSourceTencentCloudGaapDomainErrorPageInfoList(),
			"tencentcloud_gaap_access_regions":                          dataSourceTencentCloudGaapAccessRegions(),
			"tencentcloud_gaap_access_regions_by_dest_region":           dataSourceTencentCloudGaapAccessRegionsByDestRegion(),
			"tencentcloud_gaap_black_header":                            dataSourceTencentCloudGaapBlackHeader(),
			"tencentcloud_gaap_country_area_mapping":                    dataSourceTencentCloudGaapCountryAreaMapping(),
			"tencentcloud_gaap_custom_header":                           dataSourceTencentCloudGaapCustomHeader(),
			"tencentcloud_gaap_dest_regions":                            dataSourceTencentCloudGaapDestRegions(),
			"tencentcloud_gaap_proxy_detail":                            dataSourceTencentCloudGaapProxyDetail(),
			"tencentcloud_gaap_proxy_groups":                            dataSourceTencentCloudGaapProxyGroups(),
			"tencentcloud_gaap_proxy_group_statistics":                  dataSourceTencentCloudGaapProxyGroupStatistics(),
			"tencentcloud_gaap_proxy_statistics":                        dataSourceTencentCloudGaapProxyStatistics(),
			"tencentcloud_gaap_real_servers_status":                     dataSourceTencentCloudGaapRealServersStatus(),
			"tencentcloud_gaap_rule_real_servers":                       dataSourceTencentCloudGaapRuleRealServers(),
			"tencentcloud_gaap_resources_by_tag":                        dataSourceTencentCloudGaapResourcesByTag(),
			"tencentcloud_gaap_region_and_price":                        dataSourceTencentCloudGaapRegionAndPrice(),
			"tencentcloud_gaap_proxy_and_statistics_listeners":          dataSourceTencentCloudGaapProxyAndStatisticsListeners(),
			"tencentcloud_gaap_proxies_status":                          dataSourceTencentCloudGaapProxiesStatus(),
			"tencentcloud_gaap_listener_statistics":                     dataSourceTencentCloudGaapListenerStatistics(),
			"tencentcloud_gaap_listener_real_servers":                   dataSourceTencentCloudGaapListenerRealServers(),
			"tencentcloud_gaap_group_and_statistics_proxy":              dataSourceTencentCloudGaapGroupAndStatisticsProxy(),
			"tencentcloud_gaap_domain_error_page_infos":                 dataSourceTencentCloudGaapDomainErrorPageInfos(),
			"tencentcloud_gaap_check_proxy_create":                      dataSourceTencentCloudGaapCheckProxyCreate(),
			"tencentcloud_ssl_certificates":                             dataSourceTencentCloudSslCertificates(),
			"tencentcloud_ssl_describe_certificate":                     dataSourceTencentCloudSslDescribeCertificate(),
			"tencentcloud_ssl_describe_companies":                       dataSourceTencentCloudSslDescribeCompanies(),
			"tencentcloud_ssl_describe_host_api_gateway_instance_list":  dataSourceTencentCloudSslDescribeHostApiGatewayInstanceList(),
			"tencentcloud_ssl_describe_host_cdn_instance_list":          dataSourceTencentCloudSslDescribeHostCdnInstanceList(),
			"tencentcloud_ssl_describe_host_clb_instance_list":          dataSourceTencentCloudSslDescribeHostClbInstanceList(),
			"tencentcloud_ssl_describe_host_cos_instance_list":          dataSourceTencentCloudSslDescribeHostCosInstanceList(),
			"tencentcloud_ssl_describe_host_ddos_instance_list":         dataSourceTencentCloudSslDescribeHostDdosInstanceList(),
			"tencentcloud_ssl_describe_host_lighthouse_instance_list":   dataSourceTencentCloudSslDescribeHostLighthouseInstanceList(),
			"tencentcloud_ssl_describe_host_live_instance_list":         dataSourceTencentCloudSslDescribeHostLiveInstanceList(),
			"tencentcloud_ssl_describe_host_teo_instance_list":          dataSourceTencentCloudSslDescribeHostTeoInstanceList(),
			"tencentcloud_ssl_describe_host_tke_instance_list":          dataSourceTencentCloudSslDescribeHostTkeInstanceList(),
			"tencentcloud_ssl_describe_host_update_record":              dataSourceTencentCloudSslDescribeHostUpdateRecord(),
			"tencentcloud_ssl_describe_host_update_record_detail":       dataSourceTencentCloudSslDescribeHostUpdateRecordDetail(),
			"tencentcloud_ssl_describe_host_vod_instance_list":          dataSourceTencentCloudSslDescribeHostVodInstanceList(),
			"tencentcloud_ssl_describe_host_waf_instance_list":          dataSourceTencentCloudSslDescribeHostWafInstanceList(),
			"tencentcloud_ssl_describe_manager_detail":                  dataSourceTencentCloudSslDescribeManagerDetail(),
			"tencentcloud_ssl_describe_managers":                        dataSourceTencentCloudSslDescribeManagers(),
			"tencentcloud_ssl_describe_host_deploy_record":              dataSourceTencentCloudSslDescribeHostDeployRecord(),
			"tencentcloud_ssl_describe_host_deploy_record_detail":       dataSourceTencentCloudSslDescribeHostDeployRecordDetail(),
			"tencentcloud_cam_roles":                                    dataSourceTencentCloudCamRoles(),
			"tencentcloud_cam_users":                                    dataSourceTencentCloudCamUsers(),
			"tencentcloud_cam_groups":                                   dataSourceTencentCloudCamGroups(),
			"tencentcloud_cam_group_memberships":                        dataSourceTencentCloudCamGroupMemberships(),
			"tencentcloud_cam_policies":                                 dataSourceTencentCloudCamPolicies(),
			"tencentcloud_cam_role_policy_attachments":                  dataSourceTencentCloudCamRolePolicyAttachments(),
			"tencentcloud_cam_user_policy_attachments":                  dataSourceTencentCloudCamUserPolicyAttachments(),
			"tencentcloud_cam_group_policy_attachments":                 dataSourceTencentCloudCamGroupPolicyAttachments(),
			"tencentcloud_cam_saml_providers":                           dataSourceTencentCloudCamSAMLProviders(),
			"tencentcloud_cam_list_entities_for_policy":                 dataSourceTencentCloudCamListEntitiesForPolicy(),
			"tencentcloud_cam_account_summary":                          dataSourceTencentCloudCamAccountSummary(),
			"tencentcloud_cam_oidc_config":                              dataSourceTencentCloudCamOidcConfig(),
			"tencentcloud_user_info":                                    datasourceTencentCloudUserInfo(),
			"tencentcloud_cdn_domains":                                  dataSourceTencentCloudCdnDomains(),
			"tencentcloud_cdn_domain_verifier":                          dataSourceTencentCloudCdnDomainVerifyRecord(),
			"tencentcloud_scf_functions":                                dataSourceTencentCloudScfFunctions(),
			"tencentcloud_scf_namespaces":                               dataSourceTencentCloudScfNamespaces(),
			"tencentcloud_scf_account_info":                             dataSourceTencentCloudScfAccountInfo(),
			"tencentcloud_scf_async_event_management":                   dataSourceTencentCloudScfAsyncEventManagement(),
			"tencentcloud_scf_triggers":                                 dataSourceTencentCloudScfTriggers(),
			"tencentcloud_scf_async_event_status":                       dataSourceTencentCloudScfAsyncEventStatus(),
			"tencentcloud_scf_function_address":                         dataSourceTencentCloudScfFunctionAddress(),
			"tencentcloud_scf_request_status":                           dataSourceTencentCloudScfRequestStatus(),
			"tencentcloud_scf_function_aliases":                         dataSourceTencentCloudScfFunctionAliases(),
			"tencentcloud_scf_layer_versions":                           dataSourceTencentCloudScfLayerVersions(),
			"tencentcloud_scf_layers":                                   dataSourceTencentCloudScfLayers(),
			"tencentcloud_scf_function_versions":                        dataSourceTencentCloudScfFunctionVersions(),
			"tencentcloud_scf_logs":                                     dataSourceTencentCloudScfLogs(),
			"tencentcloud_tcaplus_clusters":                             dataSourceTencentCloudTcaplusClusters(),
			"tencentcloud_tcaplus_tablegroups":                          dataSourceTencentCloudTcaplusTableGroups(),
			"tencentcloud_tcaplus_tables":                               dataSourceTencentCloudTcaplusTables(),
			"tencentcloud_tcaplus_idls":                                 dataSourceTencentCloudTcaplusIdls(),
			"tencentcloud_monitor_policy_conditions":                    dataSourceTencentMonitorPolicyConditions(),
			"tencentcloud_monitor_data":                                 dataSourceTencentMonitorData(),
			"tencentcloud_monitor_product_event":                        dataSourceTencentMonitorProductEvent(),
			"tencentcloud_monitor_binding_objects":                      dataSourceTencentMonitorBindingObjects(),
			"tencentcloud_monitor_policy_groups":                        dataSourceTencentMonitorPolicyGroups(),
			"tencentcloud_monitor_product_namespace":                    dataSourceTencentMonitorProductNamespace(),
			"tencentcloud_monitor_alarm_notices":                        dataSourceTencentMonitorAlarmNotices(),
			"tencentcloud_monitor_alarm_metric":                         dataSourceTencentCloudMonitorAlarmMetric(),
			"tencentcloud_monitor_alarm_policy":                         dataSourceTencentCloudMonitorAlarmPolicy(),
			"tencentcloud_monitor_alarm_history":                        dataSourceTencentCloudMonitorAlarmHistory(),
			"tencentcloud_monitor_alarm_basic_alarms":                   dataSourceTencentCloudMonitorAlarmBasicAlarms(),
			"tencentcloud_monitor_alarm_basic_metric":                   dataSourceTencentCloudMonitorAlarmBasicMetric(),
			"tencentcloud_monitor_alarm_conditions_template":            dataSourceTencentCloudMonitorAlarmConditionsTemplate(),
			"tencentcloud_monitor_grafana_plugin_overviews":             dataSourceTencentCloudMonitorGrafanaPluginOverviews(),
			"tencentcloud_monitor_alarm_notice_callbacks":               dataSourceTencentCloudMonitorAlarmNoticeCallbacks(),
			"tencentcloud_monitor_alarm_all_namespaces":                 dataSourceTencentCloudMonitorAlarmAllNamespaces(),
			"tencentcloud_monitor_alarm_monitor_type":                   dataSourceTencentCloudMonitorAlarmMonitorType(),
			"tencentcloud_monitor_statistic_data":                       dataSourceTencentCloudMonitorStatisticData(),
			"tencentcloud_monitor_tmp_regions":                          dataSourceTencentCloudMonitorTmpRegions(),
			"tencentcloud_postgresql_instances":                         dataSourceTencentCloudPostgresqlInstances(),
			"tencentcloud_postgresql_specinfos":                         dataSourceTencentCloudPostgresqlSpecinfos(),
			"tencentcloud_postgresql_xlogs":                             datasourceTencentCloudPostgresqlXlogs(),
			"tencentcloud_postgresql_parameter_templates":               dataSourceTencentCloudPostgresqlParameterTemplates(),
			"tencentcloud_postgresql_readonly_groups":                   dataSourceTencentCloudPostgresqlReadonlyGroups(),
			"tencentcloud_postgresql_base_backups":                      dataSourceTencentCloudPostgresqlBaseBackups(),
			"tencentcloud_postgresql_log_backups":                       dataSourceTencentCloudPostgresqlLogBackups(),
			"tencentcloud_postgresql_backup_download_urls":              dataSourceTencentCloudPostgresqlBackupDownloadUrls(),
			"tencentcloud_postgresql_db_instance_classes":               dataSourceTencentCloudPostgresqlDbInstanceClasses(),
			"tencentcloud_postgresql_default_parameters":                dataSourceTencentCloudPostgresqlDefaultParameters(),
			"tencentcloud_postgresql_recovery_time":                     dataSourceTencentCloudPostgresqlRecoveryTime(),
			"tencentcloud_postgresql_regions":                           dataSourceTencentCloudPostgresqlRegions(),
			"tencentcloud_postgresql_db_instance_versions":              dataSourceTencentCloudPostgresqlDbInstanceVersions(),
			"tencentcloud_postgresql_zones":                             dataSourceTencentCloudPostgresqlZones(),
			"tencentcloud_sqlserver_zone_config":                        dataSourceTencentSqlserverZoneConfig(),
			"tencentcloud_sqlserver_instances":                          dataSourceTencentCloudSqlserverInstances(),
			"tencentcloud_sqlserver_backups":                            dataSourceTencentCloudSqlserverBackups(),
			"tencentcloud_sqlserver_dbs":                                dataSourceTencentSqlserverDBs(),
			"tencentcloud_sqlserver_accounts":                           dataSourceTencentCloudSqlserverAccounts(),
			"tencentcloud_sqlserver_account_db_attachments":             dataSourceTencentCloudSqlserverAccountDBAttachments(),
			"tencentcloud_sqlserver_readonly_groups":                    dataSourceTencentCloudSqlserverReadonlyGroups(),
			"tencentcloud_sqlserver_backup_commands":                    dataSourceTencentCloudSqlserverBackupCommands(),
			"tencentcloud_sqlserver_backup_by_flow_id":                  dataSourceTencentCloudSqlserverBackupByFlowId(),
			"tencentcloud_sqlserver_backup_upload_size":                 dataSourceTencentCloudSqlserverBackupUploadSize(),
			"tencentcloud_sqlserver_cross_region_zone":                  dataSourceTencentCloudSqlserverCrossRegionZone(),
			"tencentcloud_sqlserver_db_charsets":                        dataSourceTencentCloudSqlserverDBCharsets(),
			"tencentcloud_ckafka_users":                                 dataSourceTencentCloudCkafkaUsers(),
			"tencentcloud_ckafka_acls":                                  dataSourceTencentCloudCkafkaAcls(),
			"tencentcloud_ckafka_topics":                                dataSourceTencentCloudCkafkaTopics(),
			"tencentcloud_ckafka_instances":                             dataSourceTencentCloudCkafkaInstances(),
			"tencentcloud_ckafka_connect_resource":                      dataSourceTencentCloudCkafkaConnectResource(),
			"tencentcloud_ckafka_region":                                dataSourceTencentCloudCkafkaRegion(),
			"tencentcloud_ckafka_datahub_topic":                         dataSourceTencentCloudCkafkaDatahubTopic(),
			"tencentcloud_ckafka_datahub_group_offsets":                 dataSourceTencentCloudCkafkaDatahubGroupOffsets(),
			"tencentcloud_ckafka_datahub_task":                          dataSourceTencentCloudCkafkaDatahubTask(),
			"tencentcloud_ckafka_group":                                 dataSourceTencentCloudCkafkaGroup(),
			"tencentcloud_ckafka_group_offsets":                         dataSourceTencentCloudCkafkaGroupOffsets(),
			"tencentcloud_ckafka_group_info":                            dataSourceTencentCloudCkafkaGroupInfo(),
			"tencentcloud_ckafka_task_status":                           dataSourceTencentCloudCkafkaTaskStatus(),
			"tencentcloud_ckafka_topic_flow_ranking":                    dataSourceTencentCloudCkafkaTopicFlowRanking(),
			"tencentcloud_ckafka_topic_produce_connection":              dataSourceTencentCloudCkafkaTopicProduceConnection(),
			"tencentcloud_ckafka_topic_subscribe_group":                 dataSourceTencentCloudCkafkaTopicSubscribeGroup(),
			"tencentcloud_ckafka_topic_sync_replica":                    dataSourceTencentCloudCkafkaTopicSyncReplica(),
			"tencentcloud_ckafka_zone":                                  dataSourceTencentCloudCkafkaZone(),
			"tencentcloud_audit_cos_regions":                            dataSourceTencentCloudAuditCosRegions(),
			"tencentcloud_audit_key_alias":                              dataSourceTencentCloudAuditKeyAlias(),
			"tencentcloud_audits":                                       dataSourceTencentCloudAudits(),
			"tencentcloud_cynosdb_clusters":                             dataSourceTencentCloudCynosdbClusters(),
			"tencentcloud_cynosdb_instances":                            dataSourceTencentCloudCynosdbInstances(),
			"tencentcloud_cynosdb_zone_config":                          dataSourceTencentCynosdbZoneConfig(),
			"tencentcloud_cynosdb_instance_slow_queries":                dataSourceTencentCloudCynosdbInstanceSlowQueries(),
			"tencentcloud_vod_adaptive_dynamic_streaming_templates":     dataSourceTencentCloudVodAdaptiveDynamicStreamingTemplates(),
			"tencentcloud_vod_image_sprite_templates":                   dataSourceTencentCloudVodImageSpriteTemplates(),
			"tencentcloud_vod_procedure_templates":                      dataSourceTencentCloudVodProcedureTemplates(),
			"tencentcloud_vod_snapshot_by_time_offset_templates":        dataSourceTencentCloudVodSnapshotByTimeOffsetTemplates(),
			"tencentcloud_vod_super_player_configs":                     dataSourceTencentCloudVodSuperPlayerConfigs(),
			"tencentcloud_sqlserver_publish_subscribes":                 dataSourceTencentSqlserverPublishSubscribes(),
			"tencentcloud_sqlserver_instance_param_records":             dataSourceTencentCloudSqlserverInstanceParamRecords(),
			"tencentcloud_sqlserver_project_security_groups":            dataSourceTencentCloudSqlserverProjectSecurityGroups(),
			"tencentcloud_sqlserver_regions":                            dataSourceTencentCloudSqlserverRegions(),
			"tencentcloud_sqlserver_rollback_time":                      dataSourceTencentCloudSqlserverRollbackTime(),
			"tencentcloud_sqlserver_slowlogs":                           dataSourceTencentCloudSqlserverSlowlogs(),
			"tencentcloud_sqlserver_upload_backup_info":                 dataSourceTencentCloudSqlserverUploadBackupInfo(),
			"tencentcloud_sqlserver_upload_incremental_info":            dataSourceTencentCloudSqlserverUploadIncrementalInfo(),
			"tencentcloud_api_gateway_usage_plans":                      dataSourceTencentCloudAPIGatewayUsagePlans(),
			"tencentcloud_api_gateway_ip_strategies":                    dataSourceTencentCloudAPIGatewayIpStrategy(),
			"tencentcloud_api_gateway_customer_domains":                 dataSourceTencentCloudAPIGatewayCustomerDomains(),
			"tencentcloud_api_gateway_usage_plan_environments":          dataSourceTencentCloudAPIGatewayUsagePlanEnvironments(),
			"tencentcloud_api_gateway_throttling_services":              dataSourceTencentCloudAPIGatewayThrottlingServices(),
			"tencentcloud_api_gateway_throttling_apis":                  dataSourceTencentCloudAPIGatewayThrottlingApis(),
			"tencentcloud_api_gateway_apis":                             dataSourceTencentCloudAPIGatewayAPIs(),
			"tencentcloud_api_gateway_services":                         dataSourceTencentCloudAPIGatewayServices(),
			"tencentcloud_api_gateway_api_keys":                         dataSourceTencentCloudAPIGatewayAPIKeys(),
			"tencentcloud_api_gateway_plugins":                          dataSourceTencentCloudAPIGatewayPlugins(),
			"tencentcloud_api_gateway_upstreams":                        dataSourceTencentCloudAPIGatewayUpstreams(),
			"tencentcloud_api_gateway_api_usage_plans":                  dataSourceTencentCloudAPIGatewayApiUsagePlans(),
			"tencentcloud_api_gateway_api_app_service":                  dataSourceTencentCloudAPIGatewayApiAppService(),
			"tencentcloud_api_gateway_bind_api_apps_status":             dataSourceTencentCloudApiGatewayBindApiAppsStatus(),
			"tencentcloud_api_gateway_api_app_api":                      dataSourceTencentCloudApiGatewayApiAppApi(),
			"tencentcloud_api_gateway_api_plugins":                      dataSourceTencentCloudApiGatewayApiPlugins(),
			"tencentcloud_api_gateway_service_release_versions":         dataSourceTencentCloudApiGatewayServiceReleaseVersions(),
			"tencentcloud_api_gateway_service_environment_list":         dataSourceTencentCloudApiGatewayServiceEnvironmentList(),
			"tencentcloud_sqlserver_basic_instances":                    dataSourceTencentCloudSqlserverBasicInstances(),
			"tencentcloud_sqlserver_query_xevent":                       dataSourceTencentCloudSqlserverQueryXevent(),
			"tencentcloud_sqlserver_ins_attribute":                      dataSourceTencentCloudSqlserverInsAttribute(),
			"tencentcloud_tcr_instances":                                dataSourceTencentCloudTCRInstances(),
			"tencentcloud_tcr_namespaces":                               dataSourceTencentCloudTCRNamespaces(),
			"tencentcloud_tcr_tokens":                                   dataSourceTencentCloudTCRTokens(),
			"tencentcloud_tcr_vpc_attachments":                          dataSourceTencentCloudTCRVPCAttachments(),
			"tencentcloud_tcr_repositories":                             dataSourceTencentCloudTCRRepositories(),
			"tencentcloud_tcr_webhook_trigger_logs":                     dataSourceTencentCloudTcrWebhookTriggerLogs(),
			"tencentcloud_tcr_images":                                   dataSourceTencentCloudTcrImages(),
			"tencentcloud_tcr_image_manifests":                          dataSourceTencentCloudTcrImageManifests(),
			"tencentcloud_tcr_tag_retention_execution_tasks":            dataSourceTencentCloudTcrTagRetentionExecutionTasks(),
			"tencentcloud_tcr_tag_retention_executions":                 dataSourceTencentCloudTcrTagRetentionExecutions(),
			"tencentcloud_tcr_replication_instance_create_tasks":        dataSourceTencentCloudTcrReplicationInstanceCreateTasks(),
			"tencentcloud_tcr_replication_instance_sync_status":         dataSourceTencentCloudTcrReplicationInstanceSyncStatus(),
			"tencentcloud_address_templates":                            dataSourceTencentCloudAddressTemplates(),
			"tencentcloud_address_template_groups":                      dataSourceTencentCloudAddressTemplateGroups(),
			"tencentcloud_protocol_templates":                           dataSourceTencentCloudProtocolTemplates(),
			"tencentcloud_protocol_template_groups":                     dataSourceTencentCloudProtocolTemplateGroups(),
			"tencentcloud_kms_keys":                                     dataSourceTencentCloudKmsKeys(),
			"tencentcloud_kms_public_key":                               dataSourceTencentCloudKmsPublicKey(),
			"tencentcloud_kms_get_parameters_for_import":                dataSourceTencentCloudKmsGetParametersForImport(),
			"tencentcloud_kms_describe_keys":                            dataSourceTencentCloudKmsDescribeKeys(),
			"tencentcloud_kms_white_box_key_details":                    dataSourceTencentCloudKmsWhiteBoxKeyDetails(),
			"tencentcloud_kms_list_keys":                                dataSourceTencentCloudKmsListKeys(),
			"tencentcloud_kms_white_box_decrypt_key":                    dataSourceTencentCloudKmsWhiteBoxDecryptKey(),
			"tencentcloud_kms_white_box_device_fingerprints":            dataSourceTencentCloudKmsWhiteBoxDeviceFingerprints(),
			"tencentcloud_kms_list_algorithms":                          dataSourceTencentCloudKmsListAlgorithms(),
			"tencentcloud_ssm_products":                                 dataSourceTencentCloudSsmProducts(),
			"tencentcloud_ssm_secrets":                                  dataSourceTencentCloudSsmSecrets(),
			"tencentcloud_ssm_secret_versions":                          dataSourceTencentCloudSsmSecretVersions(),
			"tencentcloud_ssm_rotation_detail":                          dataSourceTencentCloudSsmRotationDetail(),
			"tencentcloud_ssm_rotation_history":                         dataSourceTencentCloudSsmRotationHistory(),
			"tencentcloud_ssm_service_status":                           dataSourceTencentCloudSsmServiceStatus(),
			"tencentcloud_ssm_ssh_key_pair_value":                       dataSourceTencentCloudSsmSshKeyPairValue(),
			"tencentcloud_cdh_instances":                                dataSourceTencentCloudCdhInstances(),
			"tencentcloud_dayu_eip":                                     dataSourceTencentCloudDayuEip(),
			"tencentcloud_teo_zone_available_plans":                     dataSourceTencentCloudTeoZoneAvailablePlans(),
			"tencentcloud_teo_rule_engine_settings":                     dataSourceTencentCloudTeoRuleEngineSettings(),
			"tencentcloud_sts_caller_identity":                          dataSourceTencentCloudStsCallerIdentity(),
			"tencentcloud_dcdb_instances":                               dataSourceTencentCloudDcdbInstances(),
			"tencentcloud_dcdb_accounts":                                dataSourceTencentCloudDcdbAccounts(),
			"tencentcloud_dcdb_databases":                               dataSourceTencentCloudDcdbDatabases(),
			"tencentcloud_dcdb_parameters":                              dataSourceTencentCloudDcdbParameters(),
			"tencentcloud_dcdb_shards":                                  dataSourceTencentCloudDcdbShards(),
			"tencentcloud_dcdb_security_groups":                         dataSourceTencentCloudDcdbSecurityGroups(),
			"tencentcloud_dcdb_database_objects":                        dataSourceTencentCloudDcdbDatabaseObjects(),
			"tencentcloud_dcdb_database_tables":                         dataSourceTencentCloudDcdbDatabaseTables(),
			"tencentcloud_dcdb_file_download_url":                       dataSourceTencentCloudDcdbFileDownloadUrl(),
			"tencentcloud_dcdb_log_files":                               dataSourceTencentCloudDcdbLogFiles(),
			"tencentcloud_dcdb_instance_node_info":                      dataSourceTencentCloudDcdbInstanceNodeInfo(),
			"tencentcloud_dcdb_orders":                                  dataSourceTencentCloudDcdbOrders(),
			"tencentcloud_dcdb_price":                                   dataSourceTencentCloudDcdbPrice(),
			"tencentcloud_dcdb_project_security_groups":                 dataSourceTencentCloudDcdbProjectSecurityGroups(),
			"tencentcloud_dcdb_projects":                                dataSourceTencentCloudDcdbProjects(),
			"tencentcloud_dcdb_renewal_price":                           dataSourceTencentCloudDcdbRenewalPrice(),
			"tencentcloud_dcdb_sale_info":                               dataSourceTencentCloudDcdbSaleInfo(),
			"tencentcloud_dcdb_shard_spec":                              dataSourceTencentCloudDcdbShardSpec(),
			"tencentcloud_dcdb_slow_logs":                               dataSourceTencentCloudDcdbSlowLogs(),
			"tencentcloud_dcdb_upgrade_price":                           dataSourceTencentCloudDcdbUpgradePrice(),
			"tencentcloud_mariadb_db_instances":                         dataSourceTencentCloudMariadbDbInstances(),
			"tencentcloud_mariadb_accounts":                             dataSourceTencentCloudMariadbAccounts(),
			"tencentcloud_mariadb_security_groups":                      dataSourceTencentCloudMariadbSecurityGroups(),
			"tencentcloud_mariadb_database_objects":                     dataSourceTencentCloudMariadbDatabaseObjects(),
			"tencentcloud_mariadb_databases":                            dataSourceTencentCloudMariadbDatabases(),
			"tencentcloud_mariadb_database_table":                       dataSourceTencentCloudMariadbDatabaseTable(),
			"tencentcloud_mariadb_dcn_detail":                           dataSourceTencentCloudMariadbDcnDetail(),
			"tencentcloud_mariadb_file_download_url":                    dataSourceTencentCloudMariadbFileDownloadUrl(),
			"tencentcloud_mariadb_flow":                                 dataSourceTencentCloudMariadbFlow(),
			"tencentcloud_mariadb_instance_node_info":                   dataSourceTencentCloudMariadbInstanceNodeInfo(),
			"tencentcloud_mariadb_instance_specs":                       dataSourceTencentCloudMariadbInstanceSpecs(),
			"tencentcloud_mariadb_log_files":                            dataSourceTencentCloudMariadbLogFiles(),
			"tencentcloud_mariadb_orders":                               dataSourceTencentCloudMariadbOrders(),
			"tencentcloud_mariadb_price":                                dataSourceTencentCloudMariadbPrice(),
			"tencentcloud_mariadb_project_security_groups":              dataSourceTencentCloudMariadbProjectSecurityGroups(),
			"tencentcloud_mariadb_renewal_price":                        dataSourceTencentCloudMariadbRenewalPrice(),
			"tencentcloud_mariadb_sale_info":                            dataSourceTencentCloudMariadbSaleInfo(),
			"tencentcloud_mariadb_slow_logs":                            dataSourceTencentCloudMariadbSlowLogs(),
			"tencentcloud_mariadb_upgrade_price":                        dataSourceTencentCloudMariadbUpgradePrice(),
			"tencentcloud_mps_schedules":                                dataSourceTencentCloudMpsSchedules(),
			"tencentcloud_mps_tasks":                                    dataSourceTencentCloudMpsTasks(),
			"tencentcloud_mps_parse_live_stream_process_notification":   dataSourceTencentCloudMpsParseLiveStreamProcessNotification(),
			"tencentcloud_mps_parse_notification":                       dataSourceTencentCloudMpsParseNotification(),
			"tencentcloud_mps_media_meta_data":                          dataSourceTencentCloudMpsMediaMetaData(),
			"tencentcloud_tdcpg_clusters":                               dataSourceTencentCloudTdcpgClusters(),
			"tencentcloud_tdcpg_instances":                              dataSourceTencentCloudTdcpgInstances(),
			"tencentcloud_cat_probe_data":                               dataSourceTencentCloudCatProbeData(),
			"tencentcloud_cat_node":                                     dataSourceTencentCloudCatNode(),
			"tencentcloud_cat_metric_data":                              dataSourceTencentCloudCatMetricData(),
			"tencentcloud_rum_project":                                  dataSourceTencentCloudRumProject(),
			"tencentcloud_rum_offline_log_config":                       dataSourceTencentCloudRumOfflineLogConfig(),
			"tencentcloud_rum_whitelist":                                dataSourceTencentCloudRumWhitelist(),
			"tencentcloud_rum_taw_instance":                             dataSourceTencentCloudRumTawInstance(),
			"tencentcloud_rum_custom_url":                               dataSourceTencentCloudRumCustomUrl(),
			"tencentcloud_rum_event_url":                                dataSourceTencentCloudRumEventUrl(),
			"tencentcloud_rum_fetch_url_info":                           dataSourceTencentCloudRumFetchUrlInfo(),
			"tencentcloud_rum_fetch_url":                                dataSourceTencentCloudRumFetchUrl(),
			"tencentcloud_rum_group_log":                                dataSourceTencentCloudRumGroupLog(),
			"tencentcloud_rum_log_list":                                 dataSourceTencentCloudRumLogList(),
			"tencentcloud_rum_log_url_statistics":                       dataSourceTencentCloudRumLogUrlStatistics(),
			"tencentcloud_rum_performance_page":                         dataSourceTencentCloudRumPerformancePage(),
			"tencentcloud_rum_pv_url_info":                              dataSourceTencentCloudRumPvUrlInfo(),
			"tencentcloud_rum_pv_url_statistics":                        dataSourceTencentCloudRumPvUrlStatistics(),
			"tencentcloud_rum_report_count":                             dataSourceTencentCloudRumReportCount(),
			"tencentcloud_rum_log_stats_log_list":                       dataSourceTencentCloudRumLogStatsLogList(),
			"tencentcloud_rum_scores":                                   dataSourceTencentCloudRumScores(),
			"tencentcloud_rum_set_url_statistics":                       dataSourceTencentCloudRumSetUrlStatistics(),
			"tencentcloud_rum_sign":                                     dataSourceTencentCloudRumSign(),
			"tencentcloud_rum_static_project":                           dataSourceTencentCloudRumStaticProject(),
			"tencentcloud_rum_static_resource":                          dataSourceTencentCloudRumStaticResource(),
			"tencentcloud_rum_static_url":                               dataSourceTencentCloudRumStaticUrl(),
			"tencentcloud_rum_taw_area":                                 dataSourceTencentCloudRumTawArea(),
			"tencentcloud_rum_web_vitals_page":                          dataSourceTencentCloudRumWebVitalsPage(),
			"tencentcloud_rum_log_export":                               dataSourceTencentCloudRumLogExport(),
			"tencentcloud_rum_log_export_list":                          dataSourceTencentCloudRumLogExportList(),
			"tencentcloud_dnspod_records":                               dataSourceTencentCloudDnspodRecords(),
			"tencentcloud_dnspod_domain_list":                           dataSourceTencentCloudDnspodDomainList(),
			"tencentcloud_dnspod_domain_analytics":                      dataSourceTencentCloudDnspodDomainAnalytics(),
			"tencentcloud_dnspod_domain_log_list":                       dataSourceTencentCloudDnspodDomainLogList(),
			"tencentcloud_dnspod_record_analytics":                      dataSourceTencentCloudDnspodRecordAnalytics(),
			"tencentcloud_dnspod_record_line_list":                      dataSourceTencentCloudDnspodRecordLineList(),
			"tencentcloud_dnspod_record_list":                           dataSourceTencentCloudDnspodRecordList(),
			"tencentcloud_dnspod_record_type":                           dataSourceTencentCloudDnspodRecordType(),
			"tencentcloud_tat_command":                                  dataSourceTencentCloudTatCommand(),
			"tencentcloud_tat_invoker":                                  dataSourceTencentCloudTatInvoker(),
			"tencentcloud_tat_invoker_records":                          dataSourceTencentCloudTatInvokerRecords(),
			"tencentcloud_tat_agent":                                    dataSourceTencentCloudTatAgent(),
			"tencentcloud_tat_invocation_task":                          dataSourceTencentCloudTatInvocationTask(),
			"tencentcloud_dbbrain_sql_filters":                          dataSourceTencentCloudDbbrainSqlFilters(),
			"tencentcloud_dbbrain_security_audit_log_export_tasks":      dataSourceTencentCloudDbbrainSecurityAuditLogExportTasks(),
			"tencentcloud_dbbrain_diag_event":                           dataSourceTencentCloudDbbrainDiagEvent(),
			"tencentcloud_dbbrain_diag_events":                          dataSourceTencentCloudDbbrainDiagEvents(),
			"tencentcloud_dbbrain_diag_history":                         dataSourceTencentCloudDbbrainDiagHistory(),
			"tencentcloud_dbbrain_security_audit_log_download_urls":     dataSourceTencentCloudDbbrainSecurityAuditLogDownloadUrls(),
			"tencentcloud_dbbrain_slow_log_time_series_stats":           dataSourceTencentCloudDbbrainSlowLogTimeSeriesStats(),
			"tencentcloud_dbbrain_slow_log_top_sqls":                    dataSourceTencentCloudDbbrainSlowLogTopSqls(),
			"tencentcloud_dbbrain_slow_log_user_host_stats":             dataSourceTencentCloudDbbrainSlowLogUserHostStats(),
			"tencentcloud_dbbrain_slow_log_user_sql_advice":             dataSourceTencentCloudDbbrainSlowLogUserSqlAdvice(),
			"tencentcloud_dbbrain_slow_logs":                            dataSourceTencentCloudDbbrainSlowLogs(),
			"tencentcloud_dbbrain_health_scores":                        dataSourceTencentCloudDbbrainHealthScores(),
			"tencentcloud_dbbrain_sql_templates":                        dataSourceTencentCloudDbbrainSqlTemplates(),
			"tencentcloud_dbbrain_db_space_status":                      dataSourceTencentCloudDbbrainDbSpaceStatus(),
			"tencentcloud_dbbrain_top_space_schemas":                    dataSourceTencentCloudDbbrainTopSpaceSchemas(),
			"tencentcloud_dbbrain_top_space_tables":                     dataSourceTencentCloudDbbrainTopSpaceTables(),
			"tencentcloud_dbbrain_top_space_schema_time_series":         dataSourceTencentCloudDbbrainTopSpaceSchemaTimeSeries(),
			"tencentcloud_dbbrain_top_space_table_time_series":          dataSourceTencentCloudDbbrainTopSpaceTableTimeSeries(),
			"tencentcloud_dbbrain_diag_db_instances":                    dataSourceTencentCloudDbbrainDiagDbInstances(),
			"tencentcloud_dbbrain_mysql_process_list":                   dataSourceTencentCloudDbbrainMysqlProcessList(),
			"tencentcloud_dbbrain_no_primary_key_tables":                dataSourceTencentCloudDbbrainNoPrimaryKeyTables(),
			"tencentcloud_dbbrain_redis_top_big_keys":                   dataSourceTencentCloudDbbrainRedisTopBigKeys(),
			"tencentcloud_dbbrain_redis_top_key_prefix_list":            dataSourceTencentCloudDbbrainRedisTopKeyPrefixList(),
			"tencentcloud_dts_sync_jobs":                                dataSourceTencentCloudDtsSyncJobs(),
			"tencentcloud_dts_compare_tasks":                            dataSourceTencentCloudDtsCompareTasks(),
			"tencentcloud_dts_migrate_jobs":                             dataSourceTencentCloudDtsMigrateJobs(),
			"tencentcloud_dts_migrate_db_instances":                     dataSourceTencentCloudDtsMigrateDbInstances(),
			"tencentcloud_tdmq_rocketmq_cluster":                        dataSourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloud_tdmq_rocketmq_namespace":                      dataSourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloud_tdmq_rocketmq_topic":                          dataSourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloud_tdmq_rocketmq_role":                           dataSourceTencentCloudTdmqRocketmqRole(),
			"tencentcloud_tdmq_rocketmq_group":                          dataSourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloud_tdmq_environment_attributes":                  dataSourceTencentCloudTdmqEnvironmentAttributes(),
			"tencentcloud_tdmq_publisher_summary":                       dataSourceTencentCloudTdmqPublisherSummary(),
			"tencentcloud_tdmq_publishers":                              dataSourceTencentCloudTdmqPublishers(),
			"tencentcloud_tdmq_rabbitmq_node_list":                      dataSourceTencentCloudTdmqRabbitmqNodeList(),
			"tencentcloud_tdmq_rabbitmq_vip_instance":                   dataSourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloud_tdmq_vip_instance":                            dataSourceTencentCloudTdmqVipInstance(),
			"tencentcloud_tdmq_rocketmq_messages":                       dataSourceTencentCloudTdmqRocketmqMessages(),
			"tencentcloud_tdmq_pro_instances":                           dataSourceTencentCloudTdmqProInstances(),
			"tencentcloud_tdmq_pro_instance_detail":                     dataSourceTencentCloudTdmqProInstanceDetail(),
			"tencentcloud_tcmq_queue":                                   dataSourceTencentCloudTcmqQueue(),
			"tencentcloud_tcmq_topic":                                   dataSourceTencentCloudTcmqTopic(),
			"tencentcloud_tcmq_subscribe":                               dataSourceTencentCloudTcmqSubscribe(),
			"tencentcloud_as_instances":                                 dataSourceTencentCloudAsInstances(),
			"tencentcloud_as_advices":                                   dataSourceTencentCloudAsAdvices(),
			"tencentcloud_as_limits":                                    dataSourceTencentCloudAsLimits(),
			"tencentcloud_as_last_activity":                             dataSourceTencentCloudAsLastActivity(),
			"tencentcloud_cynosdb_accounts":                             dataSourceTencentCloudCynosdbAccounts(),
			"tencentcloud_cynosdb_cluster_instance_groups":              dataSourceTencentCloudCynosdbClusterInstanceGroups(),
			"tencentcloud_cynosdb_cluster_params":                       dataSourceTencentCloudCynosdbClusterParams(),
			"tencentcloud_cynosdb_param_templates":                      dataSourceTencentCloudCynosdbParamTemplates(),
			"tencentcloud_cynosdb_zone":                                 dataSourceTencentCloudCynosdbZone(),
			"tencentcloud_cvm_instances_modification":                   dataSourceTencentCloudCvmInstancesModification(),
			"tencentcloud_cynosdb_audit_logs":                           dataSourceTencentCloudCynosdbAuditLogs(),
			"tencentcloud_cynosdb_backup_download_url":                  dataSourceTencentCloudCynosdbBackupDownloadUrl(),
			"tencentcloud_cynosdb_binlog_download_url":                  dataSourceTencentCloudCynosdbBinlogDownloadUrl(),
			"tencentcloud_cynosdb_cluster_detail_databases":             dataSourceTencentCloudCynosdbClusterDetailDatabases(),
			"tencentcloud_cynosdb_cluster_param_logs":                   dataSourceTencentCloudCynosdbClusterParamLogs(),
			"tencentcloud_cynosdb_cluster":                              dataSourceTencentCloudCynosdbCluster(),
			"tencentcloud_cynosdb_describe_instance_slow_queries":       dataSourceTencentCloudCynosdbDescribeInstanceSlowQueries(),
			"tencentcloud_cynosdb_describe_instance_error_logs":         dataSourceTencentCloudCynosdbDescribeInstanceErrorLogs(),
			"tencentcloud_cynosdb_account_all_grant_privileges":         dataSourceTencentCloudCynosdbAccountAllGrantPrivileges(),
			"tencentcloud_cynosdb_resource_package_list":                dataSourceTencentCloudCynosdbResourcePackageList(),
			"tencentcloud_cynosdb_project_security_groups":              dataSourceTencentCloudCynosdbProjectSecurityGroups(),
			"tencentcloud_cynosdb_resource_package_sale_specs":          dataSourceTencentCloudCynosdbResourcePackageSaleSpecs(),
			"tencentcloud_cynosdb_rollback_time_range":                  dataSourceTencentCloudCynosdbRollbackTimeRange(),
			"tencentcloud_cynosdb_proxy_node":                           dataSourceTencentCloudCynosdbProxyNode(),
			"tencentcloud_cynosdb_proxy_version":                        dataSourceTencentCloudCynosdbProxyVersion(),
			"tencentcloud_css_domains":                                  dataSourceTencentCloudCssDomains(),
			"tencentcloud_css_backup_stream":                            dataSourceTencentCloudCssBackupStream(),
			"tencentcloud_css_deliver_log_down_list":                    dataSourceTencentCloudCssDeliverLogDownList(),
			"tencentcloud_css_monitor_report":                           dataSourceTencentCloudCssMonitorReport(),
			"tencentcloud_css_pad_templates":                            dataSourceTencentCloudCssPadTemplates(),
			"tencentcloud_css_pull_stream_task_status":                  dataSourceTencentCloudCssPullStreamTaskStatus(),
			"tencentcloud_css_stream_monitor_list":                      dataSourceTencentCloudCssStreamMonitorList(),
			"tencentcloud_css_time_shift_record_detail":                 dataSourceTencentCloudCssTimeShiftRecordDetail(),
			"tencentcloud_css_time_shift_stream_list":                   dataSourceTencentCloudCssTimeShiftStreamList(),
			"tencentcloud_css_watermarks":                               dataSourceTencentCloudCssWatermarks(),
			"tencentcloud_css_xp2p_detail_info_list":                    dataSourceTencentCloudCssXp2pDetailInfoList(),
			"tencentcloud_chdfs_access_groups":                          dataSourceTencentCloudChdfsAccessGroups(),
			"tencentcloud_chdfs_mount_points":                           dataSourceTencentCloudChdfsMountPoints(),
			"tencentcloud_chdfs_file_systems":                           dataSourceTencentCloudChdfsFileSystems(),
			"tencentcloud_tcm_mesh":                                     dataSourceTencentCloudTcmMesh(),
			"tencentcloud_lighthouse_firewall_rules_template":           dataSourceTencentCloudLighthouseFirewallRulesTemplate(),
			"tencentcloud_cvm_instance_vnc_url":                         dataSourceTencentCloudCvmInstanceVncUrl(),
			"tencentcloud_cvm_disaster_recover_group_quota":             dataSourceTencentCloudCvmDisasterRecoverGroupQuota(),
			"tencentcloud_cvm_chc_hosts":                                dataSourceTencentCloudCvmChcHosts(),
			"tencentcloud_cvm_chc_denied_actions":                       dataSourceTencentCloudCvmChcDeniedActions(),
			"tencentcloud_cvm_image_quota":                              dataSourceTencentCloudCvmImageQuota(),
			"tencentcloud_cvm_import_image_os":                          dataSourceTencentCloudCvmImportImageOs(),
			"tencentcloud_tsf_application":                              dataSourceTencentCloudTsfApplication(),
			"tencentcloud_tsf_application_config":                       dataSourceTencentCloudTsfApplicationConfig(),
			"tencentcloud_tsf_application_file_config":                  dataSourceTencentCloudTsfApplicationFileConfig(),
			"tencentcloud_tsf_application_public_config":                dataSourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloud_cvm_image_share_permission":                   dataSourceTencentCloudCvmImageSharePermission(),
			"tencentcloud_tsf_cluster":                                  dataSourceTencentCloudTsfCluster(),
			"tencentcloud_tsf_microservice":                             dataSourceTencentCloudTsfMicroservice(),
			"tencentcloud_tsf_unit_rules":                               dataSourceTencentCloudTsfUnitRules(),
			"tencentcloud_tsf_config_summary":                           dataSourceTencentCloudTsfConfigSummary(),
			"tencentcloud_tsf_delivery_config_by_group_id":              dataSourceTencentCloudTsfDeliveryConfigByGroupId(),
			"tencentcloud_tsf_delivery_configs":                         dataSourceTencentCloudTsfDeliveryConfigs(),
			"tencentcloud_tsf_public_config_summary":                    dataSourceTencentCloudTsfPublicConfigSummary(),
			"tencentcloud_tsf_api_group":                                dataSourceTencentCloudTsfApiGroup(),
			"tencentcloud_tsf_application_attribute":                    dataSourceTencentCloudTsfApplicationAttribute(),
			"tencentcloud_tsf_business_log_configs":                     dataSourceTencentCloudTsfBusinessLogConfigs(),
			"tencentcloud_tsf_api_detail":                               dataSourceTencentCloudTsfApiDetail(),
			"tencentcloud_tsf_microservice_api_version":                 dataSourceTencentCloudTsfMicroserviceApiVersion(),
			"tencentcloud_tsf_repository":                               dataSourceTencentCloudTsfRepository(),
			"tencentcloud_tsf_pod_instances":                            dataSourceTencentCloudTsfPodInstances(),
			"tencentcloud_tsf_gateway_all_group_apis":                   dataSourceTencentCloudTsfGatewayAllGroupApis(),
			"tencentcloud_tsf_group_gateways":                           dataSourceTencentCloudTsfGroupGateways(),
			"tencentcloud_tsf_usable_unit_namespaces":                   dataSourceTencentCloudTsfUsableUnitNamespaces(),
			"tencentcloud_tsf_group_instances":                          dataSourceTencentCloudTsfGroupInstances(),
			"tencentcloud_tsf_group_config_release":                     dataSourceTencentCloudTsfGroupConfigRelease(),
			"tencentcloud_tsf_container_group":                          dataSourceTencentCloudTsfContainerGroup(),
			"tencentcloud_tsf_groups":                                   dataSourceTencentCloudTsfGroups(),
			"tencentcloud_tsf_ms_api_list":                              dataSourceTencentCloudTsfMsApiList(),
			"tencentcloud_lighthouse_bundle":                            dataSourceTencentCloudLighthouseBundle(),
			"tencentcloud_api_gateway_api_docs":                         dataSourceTencentCloudAPIGatewayAPIDocs(),
			"tencentcloud_api_gateway_api_apps":                         dataSourceTencentCloudAPIGatewayAPIApps(),
			"tencentcloud_tse_access_address":                           dataSourceTencentCloudTseAccessAddress(),
			"tencentcloud_tse_nacos_replicas":                           dataSourceTencentCloudTseNacosReplicas(),
			"tencentcloud_tse_nacos_server_interfaces":                  dataSourceTencentCloudTseNacosServerInterfaces(),
			"tencentcloud_tse_zookeeper_replicas":                       dataSourceTencentCloudTseZookeeperReplicas(),
			"tencentcloud_tse_zookeeper_server_interfaces":              dataSourceTencentCloudTseZookeeperServerInterfaces(),
			"tencentcloud_tse_groups":                                   dataSourceTencentCloudTseGroups(),
			"tencentcloud_tse_gateways":                                 dataSourceTencentCloudTseGateways(),
			"tencentcloud_tse_gateway_nodes":                            dataSourceTencentCloudTseGatewayNodes(),
			"tencentcloud_tse_gateway_routes":                           dataSourceTencentCloudTseGatewayRoutes(),
			"tencentcloud_tse_gateway_canary_rules":                     dataSourceTencentCloudTseGatewayCanaryRules(),
			"tencentcloud_tse_gateway_services":                         dataSourceTencentCloudTseGatewayServices(),
			"tencentcloud_tse_gateway_certificates":                     dataSourceTencentCloudTseGatewayCertificates(),
			"tencentcloud_lighthouse_modify_instance_bundle":            dataSourceTencentCloudLighthouseModifyInstanceBundle(),
			"tencentcloud_lighthouse_zone":                              dataSourceTencentCloudLighthouseZone(),
			"tencentcloud_lighthouse_scene":                             dataSourceTencentCloudLighthouseScene(),
			"tencentcloud_lighthouse_all_scene":                         dataSourceTencentCloudLighthouseAllScene(),
			"tencentcloud_lighthouse_reset_instance_blueprint":          dataSourceTencentCloudLighthouseResetInstanceBlueprint(),
			"tencentcloud_lighthouse_region":                            dataSourceTencentCloudLighthouseRegion(),
			"tencentcloud_lighthouse_instance_vnc_url":                  dataSourceTencentCloudLighthouseInstanceVncUrl(),
			"tencentcloud_lighthouse_instance_traffic_package":          dataSourceTencentCloudLighthouseInstanceTrafficPackage(),
			"tencentcloud_lighthouse_instance_disk_num":                 dataSourceTencentCloudLighthouseInstanceDiskNum(),
			"tencentcloud_lighthouse_instance_blueprint":                dataSourceTencentCloudLighthouseInstanceBlueprint(),
			"tencentcloud_lighthouse_disk_config":                       dataSourceTencentCloudLighthouseDiskConfig(),
			"tencentcloud_lighthouse_disks":                             dataSourceTencentCloudLighthouseInstanceDisks(),
			"tencentcloud_clickhouse_backup_jobs":                       dataSourceTencentCloudClickhouseBackupJobs(),
			"tencentcloud_clickhouse_backup_job_detail":                 dataSourceTencentCloudClickhouseBackupJobDetail(),
			"tencentcloud_clickhouse_backup_tables":                     dataSourceTencentCloudClickhouseBackupTables(),
			"tencentcloud_cls_shipper_tasks":                            dataSourceTencentCloudClsShipperTasks(),
			"tencentcloud_cls_machines":                                 dataSourceTencentCloudClsMachines(),
			"tencentcloud_cls_machine_group_configs":                    dataSourceTencentCloudClsMachineGroupConfigs(),
			"tencentcloud_eb_search":                                    dataSourceTencentCloudEbSearch(),
			"tencentcloud_eb_bus":                                       dataSourceTencentCloudEbBus(),
			"tencentcloud_eb_event_rules":                               dataSourceTencentCloudEbEventRules(),
			"tencentcloud_eb_platform_event_names":                      dataSourceTencentCloudEbPlatformEventNames(),
			"tencentcloud_eb_platform_event_patterns":                   dataSourceTencentCloudEbPlatformEventPatterns(),
			"tencentcloud_eb_platform_products":                         dataSourceTencentCloudEbPlatformProducts(),
			"tencentcloud_eb_plateform_event_template":                  dataSourceTencentCloudEbPlateformEventTemplate(),
			"tencentcloud_wedata_rule_templates":                        dataSourceTencentCloudWedataRuleTemplates(),
			"tencentcloud_wedata_data_source_list":                      dataSourceTencentCloudWedataDataSourceList(),
			"tencentcloud_wedata_data_source_without_info":              dataSourceTencentCloudWedataDataSourceWithoutInfo(),
			"tencentcloud_private_dns_records":                          dataSourceTencentCloudPrivateDnsRecords(),
			"tencentcloud_waf_ciphers":                                  dataSourceTencentCloudWafCiphers(),
			"tencentcloud_waf_tls_versions":                             dataSourceTencentCloudWafTlsVersions(),
			"tencentcloud_waf_domains":                                  dataSourceTencentCloudWafDomains(),
			"tencentcloud_waf_find_domains":                             dataSourceTencentCloudWafFindDomains(),
			"tencentcloud_waf_waf_infos":                                dataSourceTencentCloudWafWafInfos(),
			"tencentcloud_waf_ports":                                    dataSourceTencentCloudWafPorts(),
			"tencentcloud_waf_user_domains":                             dataSourceTencentCloudWafUserDomains(),
			"tencentcloud_waf_attack_log_histogram":                     dataSourceTencentCloudWafAttackLogHistogram(),
			"tencentcloud_waf_attack_log_list":                          dataSourceTencentCloudWafAttackLogList(),
			"tencentcloud_waf_attack_overview":                          dataSourceTencentCloudWafAttackOverview(),
			"tencentcloud_waf_attack_total_count":                       dataSourceTencentCloudWafAttackTotalCount(),
			"tencentcloud_waf_peak_points":                              dataSourceTencentCloudWafPeakPoints(),
			"tencentcloud_waf_instance_qps_limit":                       dataSourceTencentCloudWafInstanceQpsLimit(),
			"tencentcloud_waf_user_clb_regions":                         dataSourceTencentCloudWafUserClbRegions(),
			"tencentcloud_cfw_nat_fw_switches":                          dataSourceTencentCloudCfwNatFwSwitches(),
			"tencentcloud_cfw_vpc_fw_switches":                          dataSourceTencentCloudCfwVpcFwSwitches(),
			"tencentcloud_cfw_edge_fw_switches":                         dataSourceTencentCloudCfwEdgeFwSwitches(),
			"tencentcloud_cwp_machines_simple":                          dataSourceTencentCloudCwpMachinesSimple(),
			"tencentcloud_ses_receivers":                                dataSourceTencentCloudSesReceivers(),
			"tencentcloud_ses_send_tasks":                               dataSourceTencentCloudSesSendTasks(),
			"tencentcloud_ses_email_identities":                         dataSourceTencentCloudSesEmailIdentities(),
			"tencentcloud_ses_black_email_address":                      dataSourceTencentCloudSesBlackEmailAddress(),
			"tencentcloud_ses_statistics_report":                        dataSourceTencentCloudSesStatisticsReport(),
			"tencentcloud_ses_send_email_status":                        dataSourceTencentCloudSesSendEmailStatus(),
			"tencentcloud_organization_org_financial_by_member":         dataSourceTencentCloudOrganizationOrgFinancialByMember(),
			"tencentcloud_organization_org_financial_by_month":          dataSourceTencentCloudOrganizationOrgFinancialByMonth(),
			"tencentcloud_organization_org_financial_by_product":        dataSourceTencentCloudOrganizationOrgFinancialByProduct(),
			"tencentcloud_organization_org_auth_node":                   dataSourceTencentCloudOrganizationOrgAuthNode(),
			"tencentcloud_pts_scenario_with_jobs":                       dataSourceTencentCloudPtsScenarioWithJobs(),
			"tencentcloud_cam_list_attached_user_policy":                dataSourceTencentCloudCamListAttachedUserPolicy(),
			"tencentcloud_cam_secret_last_used_time":                    dataSourceTencentCloudCamSecretLastUsedTime(),
			"tencentcloud_cam_policy_granting_service_access":           dataSourceTencentCloudCamPolicyGrantingServiceAccess(),
			"tencentcloud_cam_group_user_account":                       dataSourceTencentCloudCamGroupUserAccount(),
			"tencentcloud_dlc_check_data_engine_image_can_be_rollback":  dataSourceTencentCloudDlcCheckDataEngineImageCanBeRollback(),
			"tencentcloud_dlc_check_data_engine_image_can_be_upgrade":   dataSourceTencentCloudDlcCheckDataEngineImageCanBeUpgrade(),
			"tencentcloud_dlc_describe_user_type":                       dataSourceTencentCloudDlcDescribeUserType(),
			"tencentcloud_dlc_describe_user_info":                       dataSourceTencentCloudDlcDescribeUserInfo(),
			"tencentcloud_dlc_describe_user_roles":                      dataSourceTencentCloudDlcDescribeUserRoles(),
			"tencentcloud_dlc_describe_data_engine":                     dataSourceTencentCloudDlcDescribeDataEngine(),
			"tencentcloud_dlc_describe_data_engine_image_versions":      dataSourceTencentCloudDlcDescribeDataEngineImageVersions(),
			"tencentcloud_dlc_describe_data_engine_python_spark_images": dataSourceTencentCloudDlcDescribeDataEnginePythonSparkImages(),
			"tencentcloud_dlc_describe_engine_usage_info":               dataSourceTencentCloudDlcDescribeEngineUsageInfo(),
			"tencentcloud_dlc_describe_work_group_info":                 dataSourceTencentCloudDlcDescribeWorkGroupInfo(),
			"tencentcloud_dlc_check_data_engine_config_pairs_validity":  dataSourceTencentCloudDlcCheckDataEngineConfigPairsValidity(),
			"tencentcloud_dlc_describe_updatable_data_engines":          dataSourceTencentCloudDlcDescribeUpdatableDataEngines(),
			"tencentcloud_bi_project":                                   dataSourceTencentCloudBiProject(),
			"tencentcloud_bi_user_project":                              dataSourceTencentCloudBiUserProject(),
			"tencentcloud_antiddos_basic_device_status":                 dataSourceTencentCloudAntiddosBasicDeviceStatus(),
			"tencentcloud_antiddos_bgp_biz_trend":                       dataSourceTencentCloudAntiddosBgpBizTrend(),
			"tencentcloud_antiddos_list_listener":                       dataSourceTencentCloudAntiddosListListener(),
			"tencentcloud_antiddos_overview_attack_trend":               dataSourceTencentCloudAntiddosOverviewAttackTrend(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"tencentcloud_project":                                             resourceTencentCloudProject(),
			"tencentcloud_emr_cluster":                                         resourceTencentCloudEmrCluster(),
			"tencentcloud_emr_user_manager":                                    resourceTencentCloudEmrUserManager(),
			"tencentcloud_instance":                                            resourceTencentCloudInstance(),
			"tencentcloud_instance_set":                                        resourceTencentCloudInstanceSet(),
			"tencentcloud_reserved_instance":                                   resourceTencentCloudReservedInstance(),
			"tencentcloud_key_pair":                                            resourceTencentCloudKeyPair(),
			"tencentcloud_placement_group":                                     resourceTencentCloudPlacementGroup(),
			"tencentcloud_cbs_snapshot":                                        resourceTencentCloudCbsSnapshot(),
			"tencentcloud_cbs_snapshot_policy":                                 resourceTencentCloudCbsSnapshotPolicy(),
			"tencentcloud_cbs_storage":                                         resourceTencentCloudCbsStorage(),
			"tencentcloud_cbs_storage_set":                                     resourceTencentCloudCbsStorageSet(),
			"tencentcloud_cbs_storage_attachment":                              resourceTencentCloudCbsStorageAttachment(),
			"tencentcloud_cbs_storage_set_attachment":                          resourceTencentCloudCbsStorageSetAttachment(),
			"tencentcloud_cbs_snapshot_policy_attachment":                      resourceTencentCloudCbsSnapshotPolicyAttachment(),
			"tencentcloud_vpc":                                                 resourceTencentCloudVpcInstance(),
			"tencentcloud_vpc_acl":                                             resourceTencentCloudVpcACL(),
			"tencentcloud_vpc_acl_attachment":                                  resourceTencentCloudVpcAclAttachment(),
			"tencentcloud_vpc_network_acl_quintuple":                           resourceTencentCloudVpcNetworkAclQuintuple(),
			"tencentcloud_vpc_notify_routes":                                   resourceTencentCloudVpcNotifyRoutes(),
			"tencentcloud_vpc_bandwidth_package":                               resourceTencentCloudVpcBandwidthPackage(),
			"tencentcloud_vpc_bandwidth_package_attachment":                    resourceTencentCloudVpcBandwidthPackageAttachment(),
			"tencentcloud_vpc_traffic_package":                                 resourceTencentCloudVpcTrafficPackage(),
			"tencentcloud_vpc_snapshot_policy":                                 resourceTencentCloudVpcSnapshotPolicy(),
			"tencentcloud_vpc_snapshot_policy_attachment":                      resourceTencentCloudVpcSnapshotPolicyAttachment(),
			"tencentcloud_vpc_snapshot_policy_config":                          resourceTencentCloudVpcSnapshotPolicyConfig(),
			"tencentcloud_vpc_net_detect":                                      resourceTencentCloudVpcNetDetect(),
			"tencentcloud_vpc_flow_log_config":                                 resourceTencentCloudVpcFlowLogConfig(),
			"tencentcloud_vpc_classic_link_attachment":                         resourceTencentCloudVpcClassicLinkAttachment(),
			"tencentcloud_vpc_dhcp_ip":                                         resourceTencentCloudVpcDhcpIp(),
			"tencentcloud_vpc_ipv6_cidr_block":                                 resourceTencentCloudVpcIpv6CidrBlock(),
			"tencentcloud_vpc_ipv6_subnet_cidr_block":                          resourceTencentCloudVpcIpv6SubnetCidrBlock(),
			"tencentcloud_vpc_ipv6_eni_address":                                resourceTencentCloudVpcIpv6EniAddress(),
			"tencentcloud_vpc_dhcp_associate_address":                          resourceTencentCloudVpcDhcpAssociateAddress(),
			"tencentcloud_vpc_local_gateway":                                   resourceTencentCloudVpcLocalGateway(),
			"tencentcloud_vpc_resume_snapshot_instance":                        resourceTencentCloudVpcResumeSnapshotInstance(),
			"tencentcloud_ipv6_address_bandwidth":                              resourceTencentCloudIpv6AddressBandwidth(),
			"tencentcloud_subnet":                                              resourceTencentCloudVpcSubnet(),
			"tencentcloud_route_entry":                                         resourceTencentCloudRouteEntry(),
			"tencentcloud_route_table_entry":                                   resourceTencentCloudVpcRouteEntry(),
			"tencentcloud_route_table":                                         resourceTencentCloudVpcRouteTable(),
			"tencentcloud_route_table_association":                             resourceTencentCloudRouteTableAssociation(),
			"tencentcloud_dnat":                                                resourceTencentCloudDnat(),
			"tencentcloud_nat_gateway":                                         resourceTencentCloudNatGateway(),
			"tencentcloud_nat_gateway_snat":                                    resourceTencentCloudNatGatewaySnat(),
			"tencentcloud_nat_refresh_nat_dc_route":                            resourceTencentCloudNatRefreshNatDcRoute(),
			"tencentcloud_oceanus_job":                                         resourceTencentCloudOceanusJob(),
			"tencentcloud_oceanus_job_config":                                  resourceTencentCloudOceanusJobConfig(),
			"tencentcloud_oceanus_job_copy":                                    resourceTencentCloudOceanusJobCopy(),
			"tencentcloud_oceanus_run_job":                                     resourceTencentCloudOceanusRunJob(),
			"tencentcloud_oceanus_stop_job":                                    resourceTencentCloudOceanusStopJob(),
			"tencentcloud_oceanus_trigger_job_savepoint":                       resourceTencentCloudOceanusTriggerJobSavepoint(),
			"tencentcloud_oceanus_resource":                                    resourceTencentCloudOceanusResource(),
			"tencentcloud_oceanus_resource_config":                             resourceTencentCloudOceanusResourceConfig(),
			"tencentcloud_oceanus_work_space":                                  resourceTencentCloudOceanusWorkSpace(),
			"tencentcloud_tag":                                                 resourceTencentCloudTag(),
			"tencentcloud_tag_attachment":                                      resourceTencentCloudTagAttachment(),
			"tencentcloud_eip":                                                 resourceTencentCloudEip(),
			"tencentcloud_eip_association":                                     resourceTencentCloudEipAssociation(),
			"tencentcloud_eip_address_transform":                               resourceTencentCloudEipAddressTransform(),
			"tencentcloud_eip_public_address_adjust":                           resourceTencentCloudEipPublicAddressAdjust(),
			"tencentcloud_eip_normal_address_return":                           resourceTencentCloudEipNormalAddressReturn(),
			"tencentcloud_eni":                                                 resourceTencentCloudEni(),
			"tencentcloud_eni_attachment":                                      resourceTencentCloudEniAttachment(),
			"tencentcloud_eni_sg_attachment":                                   resourceTencentCloudEniSgAttachment(),
			"tencentcloud_ccn":                                                 resourceTencentCloudCcn(),
			"tencentcloud_ccn_attachment":                                      resourceTencentCloudCcnAttachment(),
			"tencentcloud_ccn_bandwidth_limit":                                 resourceTencentCloudCcnBandwidthLimit(),
			"tencentcloud_ccn_routes":                                          resourceTencentCloudCcnRoutes(),
			"tencentcloud_ccn_instances_accept_attach":                         resourceTencentCloudCcnInstancesAcceptAttach(),
			"tencentcloud_ccn_instances_reject_attach":                         resourceTencentCloudCcnInstancesRejectAttach(),
			"tencentcloud_ccn_instances_reset_attach":                          resourceTencentCloudCcnInstancesResetAttach(),
			"tencentcloud_dc_instance":                                         resourceTencentCloudDcInstance(),
			"tencentcloud_dcx":                                                 resourceTencentCloudDcxInstance(),
			"tencentcloud_dcx_extra_config":                                    resourceTencentCloudDcxExtraConfig(),
			"tencentcloud_dc_share_dcx_config":                                 resourceTencentCloudDcShareDcxConfig(),
			"tencentcloud_dc_internet_address":                                 resourceTencentCloudDcInternetAddress(),
			"tencentcloud_dc_internet_address_config":                          resourceTencentCloudDcInternetAddressConfig(),
			"tencentcloud_dc_gateway":                                          resourceTencentCloudDcGatewayInstance(),
			"tencentcloud_dc_gateway_ccn_route":                                resourceTencentCloudDcGatewayCcnRouteInstance(),
			"tencentcloud_dc_gateway_attachment":                               resourceTencentCloudDcGatewayAttachment(),
			"tencentcloud_vpn_customer_gateway":                                resourceTencentCloudVpnCustomerGateway(),
			"tencentcloud_vpn_gateway":                                         resourceTencentCloudVpnGateway(),
			"tencentcloud_vpn_gateway_route":                                   resourceTencentCloudVpnGatewayRoute(),
			"tencentcloud_vpn_connection":                                      resourceTencentCloudVpnConnection(),
			"tencentcloud_vpn_ssl_server":                                      resourceTencentCloudVpnSslServer(),
			"tencentcloud_vpn_ssl_client":                                      resourceTencentCloudVpnSslClient(),
			"tencentcloud_vpn_connection_reset":                                resourceTencentCloudVpnConnectionReset(),
			"tencentcloud_vpn_customer_gateway_configuration_download":         resourceTencentCloudVpnCustomerGatewayConfigurationDownload(),
			"tencentcloud_vpn_gateway_ssl_client_cert":                         resourceTencentCloudVpnGatewaySslClientCert(),
			"tencentcloud_vpn_gateway_ccn_routes":                              resourceTencentCloudVpnGatewayCcnRoutes(),
			"tencentcloud_ha_vip":                                              resourceTencentCloudHaVip(),
			"tencentcloud_ha_vip_eip_attachment":                               resourceTencentCloudHaVipEipAttachment(),
			"tencentcloud_security_group":                                      resourceTencentCloudSecurityGroup(),
			"tencentcloud_security_group_rule":                                 resourceTencentCloudSecurityGroupRule(),
			"tencentcloud_security_group_rule_set":                             resourceTencentCloudSecurityGroupRuleSet(),
			"tencentcloud_security_group_lite_rule":                            resourceTencentCloudSecurityGroupLiteRule(),
			"tencentcloud_lb":                                                  resourceTencentCloudLB(),
			"tencentcloud_alb_server_attachment":                               resourceTencentCloudAlbServerAttachment(),
			"tencentcloud_clb_instance":                                        resourceTencentCloudClbInstance(),
			"tencentcloud_clb_listener":                                        resourceTencentCloudClbListener(),
			"tencentcloud_clb_listener_rule":                                   resourceTencentCloudClbListenerRule(),
			"tencentcloud_clb_attachment":                                      resourceTencentCloudClbServerAttachment(),
			"tencentcloud_clb_redirection":                                     resourceTencentCloudClbRedirection(),
			"tencentcloud_clb_target_group":                                    resourceTencentCloudClbTargetGroup(),
			"tencentcloud_clb_target_group_instance_attachment":                resourceTencentCloudClbTGAttachmentInstance(),
			"tencentcloud_clb_target_group_attachment":                         resourceTencentCloudClbTargetGroupAttachment(),
			"tencentcloud_clb_log_set":                                         resourceTencentCloudClbLogSet(),
			"tencentcloud_clb_log_topic":                                       resourceTencentCloudClbLogTopic(),
			"tencentcloud_clb_customized_config":                               resourceTencentCloudClbCustomizedConfig(),
			"tencentcloud_clb_snat_ip":                                         resourceTencentCloudClbSnatIp(),
			"tencentcloud_clb_function_targets_attachment":                     resourceTencentCloudClbFunctionTargetsAttachment(),
			"tencentcloud_clb_instance_mix_ip_target_config":                   resourceTencentCloudClbInstanceMixIpTargetConfig(),
			"tencentcloud_clb_instance_sla_config":                             resourceTencentCloudClbInstanceSlaConfig(),
			"tencentcloud_clb_replace_cert_for_lbs":                            resourceTencentCloudClbReplaceCertForLbs(),
			"tencentcloud_clb_security_group_attachment":                       resourceTencentCloudClbSecurityGroupAttachment(),
			"tencentcloud_container_cluster":                                   resourceTencentCloudContainerCluster(),
			"tencentcloud_container_cluster_instance":                          resourceTencentCloudContainerClusterInstance(),
			"tencentcloud_kubernetes_cluster":                                  resourceTencentCloudTkeCluster(),
			"tencentcloud_kubernetes_cluster_endpoint":                         resourceTencentCloudTkeClusterEndpoint(),
			"tencentcloud_eks_cluster":                                         resourceTencentCloudEksCluster(),
			"tencentcloud_eks_container_instance":                              resourceTencentCloudEksContainerInstance(),
			"tencentcloud_kubernetes_addon_attachment":                         resourceTencentCloudTkeAddonAttachment(),
			"tencentcloud_kubernetes_auth_attachment":                          resourceTencentCloudTKEAuthAttachment(),
			"tencentcloud_kubernetes_as_scaling_group":                         resourceTencentCloudKubernetesAsScalingGroup(),
			"tencentcloud_kubernetes_scale_worker":                             resourceTencentCloudTkeScaleWorker(),
			"tencentcloud_kubernetes_cluster_attachment":                       resourceTencentCloudTkeClusterAttachment(),
			"tencentcloud_kubernetes_node_pool":                                resourceTencentCloudKubernetesNodePool(),
			"tencentcloud_kubernetes_serverless_node_pool":                     resourceTkeServerLessNodePool(),
			"tencentcloud_kubernetes_backup_storage_location":                  resourceTencentCloudTkeBackupStorageLocation(),
			"tencentcloud_kubernetes_encryption_protection":                    resourceTencentCloudKubernetesEncryptionProtection(),
			"tencentcloud_mysql_backup_policy":                                 resourceTencentCloudMysqlBackupPolicy(),
			"tencentcloud_mysql_account":                                       resourceTencentCloudMysqlAccount(),
			"tencentcloud_mysql_account_privilege":                             resourceTencentCloudMysqlAccountPrivilege(),
			"tencentcloud_mysql_privilege":                                     resourceTencentCloudMysqlPrivilege(),
			"tencentcloud_mysql_instance":                                      resourceTencentCloudMysqlInstance(),
			"tencentcloud_mysql_database":                                      resourceTencentCloudMysqlDatabase(),
			"tencentcloud_mysql_readonly_instance":                             resourceTencentCloudMysqlReadonlyInstance(),
			"tencentcloud_mysql_time_window":                                   resourceTencentCloudMysqlTimeWindow(),
			"tencentcloud_mysql_param_template":                                resourceTencentCloudMysqlParamTemplate(),
			"tencentcloud_mysql_security_groups_attachment":                    resourceTencentCloudMysqlSecurityGroupsAttachment(),
			"tencentcloud_mysql_deploy_group":                                  resourceTencentCloudMysqlDeployGroup(),
			"tencentcloud_mysql_local_binlog_config":                           resourceTencentCloudMysqlLocalBinlogConfig(),
			"tencentcloud_mysql_audit_log_file":                                resourceTencentCloudMysqlAuditLogFile(),
			"tencentcloud_mysql_backup_download_restriction":                   resourceTencentCloudMysqlBackupDownloadRestriction(),
			"tencentcloud_mysql_renew_db_instance_operation":                   resourceTencentCloudMysqlRenewDbInstanceOperation(),
			"tencentcloud_mysql_backup_encryption_status":                      resourceTencentCloudMysqlBackupEncryptionStatus(),
			"tencentcloud_mysql_db_import_job_operation":                       resourceTencentCloudMysqlDbImportJobOperation(),
			"tencentcloud_mysql_dr_instance_to_mater":                          resourceTencentCloudMysqlDrInstanceToMater(),
			"tencentcloud_mysql_instance_encryption_operation":                 resourceTencentCloudMysqlInstanceEncryptionOperation(),
			"tencentcloud_mysql_isolate_instance":                              resourceTencentCloudMysqlIsolateInstance(),
			"tencentcloud_mysql_password_complexity":                           resourceTencentCloudMysqlPasswordComplexity(),
			"tencentcloud_mysql_remote_backup_config":                          resourceTencentCloudMysqlRemoteBackupConfig(),
			"tencentcloud_mysql_restart_db_instances_operation":                resourceTencentCloudMysqlRestartDbInstancesOperation(),
			"tencentcloud_mysql_switch_for_upgrade":                            resourceTencentCloudMysqlSwitchForUpgrade(),
			"tencentcloud_mysql_rollback":                                      resourceTencentCloudMysqlRollback(),
			"tencentcloud_mysql_rollback_stop":                                 resourceTencentCloudMysqlRollbackStop(),
			"tencentcloud_mysql_ro_group":                                      resourceTencentCloudMysqlRoGroup(),
			"tencentcloud_mysql_ro_instance_ip":                                resourceTencentCloudMysqlRoInstanceIp(),
			"tencentcloud_mysql_ro_group_load_operation":                       resourceTencentCloudMysqlRoGroupLoadOperation(),
			"tencentcloud_mysql_switch_master_slave_operation":                 resourceTencentCloudMysqlSwitchMasterSlaveOperation(),
			"tencentcloud_mysql_proxy":                                         resourceTencentCloudMysqlProxy(),
			"tencentcloud_mysql_reset_root_account":                            resourceTencentCloudMysqlResetRootAccount(),
			"tencentcloud_mysql_verify_root_account":                           resourceTencentCloudMysqlVerifyRootAccount(),
			"tencentcloud_mysql_reload_balance_proxy_node":                     resourceTencentCloudMysqlReloadBalanceProxyNode(),
			"tencentcloud_mysql_ro_start_replication":                          resourceTencentCloudMysqlRoStartReplication(),
			"tencentcloud_mysql_ro_stop_replication":                           resourceTencentCloudMysqlRoStopReplication(),
			"tencentcloud_mysql_switch_proxy":                                  resourceTencentCloudMysqlSwitchProxy(),
			"tencentcloud_cos_bucket":                                          resourceTencentCloudCosBucket(),
			"tencentcloud_cos_bucket_object":                                   resourceTencentCloudCosBucketObject(),
			"tencentcloud_cos_bucket_referer":                                  resourceTencentCloudCosBucketReferer(),
			"tencentcloud_cos_bucket_version":                                  resourceTencentCloudCosBucketVersion(),
			"tencentcloud_cfs_file_system":                                     resourceTencentCloudCfsFileSystem(),
			"tencentcloud_cfs_access_group":                                    resourceTencentCloudCfsAccessGroup(),
			"tencentcloud_cfs_access_rule":                                     resourceTencentCloudCfsAccessRule(),
			"tencentcloud_cfs_auto_snapshot_policy":                            resourceTencentCloudCfsAutoSnapshotPolicy(),
			"tencentcloud_cfs_auto_snapshot_policy_attachment":                 resourceTencentCloudCfsAutoSnapshotPolicyAttachment(),
			"tencentcloud_cfs_snapshot":                                        resourceTencentCloudCfsSnapshot(),
			"tencentcloud_cfs_user_quota":                                      resourceTencentCloudCfsUserQuota(),
			"tencentcloud_cfs_sign_up_cfs_service":                             resourceTencentCloudCfsSignUpCfsService(),
			"tencentcloud_redis_instance":                                      resourceTencentCloudRedisInstance(),
			"tencentcloud_redis_backup_config":                                 resourceTencentCloudRedisBackupConfig(),
			"tencentcloud_redis_account":                                       resourceTencentCloudRedisAccount(),
			"tencentcloud_redis_param_template":                                resourceTencentCloudRedisParamTemplate(),
			"tencentcloud_redis_connection_config":                             resourceTencentCloudRedisConnectionConfig(),
			"tencentcloud_redis_param":                                         resourceTencentCloudRedisParam(),
			"tencentcloud_redis_read_only":                                     resourceTencentCloudRedisReadOnly(),
			"tencentcloud_redis_ssl":                                           resourceTencentCloudRedisSsl(),
			"tencentcloud_redis_backup_download_restriction":                   resourceTencentCloudRedisBackupDownloadRestriction(),
			"tencentcloud_redis_clear_instance_operation":                      resourceTencentCloudRedisClearInstanceOperation(),
			"tencentcloud_redis_renew_instance_operation":                      resourceTencentCloudRedisRenewInstanceOperation(),
			"tencentcloud_redis_startup_instance_operation":                    resourceTencentCloudRedisStartupInstanceOperation(),
			"tencentcloud_redis_upgrade_cache_version_operation":               resourceTencentCloudRedisUpgradeCacheVersionOperation(),
			"tencentcloud_redis_upgrade_multi_zone_operation":                  resourceTencentCloudRedisUpgradeMultiZoneOperation(),
			"tencentcloud_redis_upgrade_proxy_version_operation":               resourceTencentCloudRedisUpgradeProxyVersionOperation(),
			"tencentcloud_redis_maintenance_window":                            resourceTencentCloudRedisMaintenanceWindow(),
			"tencentcloud_redis_replica_readonly":                              resourceTencentCloudRedisReplicaReadonly(),
			"tencentcloud_redis_switch_master":                                 resourceTencentCloudRedisSwitchMaster(),
			"tencentcloud_redis_replicate_attachment":                          resourceTencentCloudRedisReplicateAttachment(),
			"tencentcloud_redis_backup_operation":                              resourceTencentCloudRedisBackupOperation(),
			"tencentcloud_redis_security_group_attachment":                     resourceTencentCloudRedisSecurityGroupAttachment(),
			"tencentcloud_as_load_balancer":                                    resourceTencentCloudAsLoadBalancer(),
			"tencentcloud_as_scaling_config":                                   resourceTencentCloudAsScalingConfig(),
			"tencentcloud_as_scaling_group":                                    resourceTencentCloudAsScalingGroup(),
			"tencentcloud_as_scaling_group_status":                             resourceTencentCloudAsScalingGroupStatus(),
			"tencentcloud_as_attachment":                                       resourceTencentCloudAsAttachment(),
			"tencentcloud_as_scaling_policy":                                   resourceTencentCloudAsScalingPolicy(),
			"tencentcloud_as_schedule":                                         resourceTencentCloudAsSchedule(),
			"tencentcloud_as_lifecycle_hook":                                   resourceTencentCloudAsLifecycleHook(),
			"tencentcloud_as_notification":                                     resourceTencentCloudAsNotification(),
			"tencentcloud_as_remove_instances":                                 resourceTencentCloudAsRemoveInstances(),
			"tencentcloud_as_protect_instances":                                resourceTencentCloudAsProtectInstances(),
			"tencentcloud_as_start_instances":                                  resourceTencentCloudAsStartInstances(),
			"tencentcloud_as_stop_instances":                                   resourceTencentCloudAsStopInstances(),
			"tencentcloud_as_scale_in_instances":                               resourceTencentCloudAsScaleInInstances(),
			"tencentcloud_as_scale_out_instances":                              resourceTencentCloudAsScaleOutInstances(),
			"tencentcloud_as_execute_scaling_policy":                           resourceTencentCloudAsExecuteScalingPolicy(),
			"tencentcloud_as_complete_lifecycle":                               resourceTencentCloudAsCompleteLifecycle(),
			"tencentcloud_mongodb_instance":                                    resourceTencentCloudMongodbInstance(),
			"tencentcloud_mongodb_sharding_instance":                           resourceTencentCloudMongodbShardingInstance(),
			"tencentcloud_mongodb_instance_account":                            resourceTencentCloudMongodbInstanceAccount(),
			"tencentcloud_mongodb_instance_backup":                             resourceTencentCloudMongodbInstanceBackup(),
			"tencentcloud_mongodb_instance_backup_download_task":               resourceTencentCloudMongodbInstanceBackupDownloadTask(),
			"tencentcloud_dayu_cc_http_policy":                                 resourceTencentCloudDayuCCHttpPolicy(),
			"tencentcloud_dayu_cc_https_policy":                                resourceTencentCloudDayuCCHttpsPolicy(),
			"tencentcloud_dayu_ddos_policy":                                    resourceTencentCloudDayuDdosPolicy(),
			"tencentcloud_dayu_cc_policy_v2":                                   resourceTencentCloudDayuCCPolicyV2(),
			"tencentcloud_dayu_ddos_policy_v2":                                 resourceTencentCloudDayuDdosPolicyV2(),
			"tencentcloud_dayu_ddos_policy_case":                               resourceTencentCloudDayuDdosPolicyCase(),
			"tencentcloud_dayu_ddos_policy_attachment":                         resourceTencentCloudDayuDdosPolicyAttachment(),
			"tencentcloud_dayu_l4_rule":                                        resourceTencentCloudDayuL4Rule(),
			"tencentcloud_dayu_l4_rule_v2":                                     resourceTencentCloudDayuL4RuleV2(),
			"tencentcloud_dayu_l7_rule":                                        resourceTencentCloudDayuL7Rule(),
			"tencentcloud_dayu_l7_rule_v2":                                     resourceTencentCloudDayuL7RuleV2(),
			"tencentcloud_dayu_eip":                                            resourceTencentCloudDayuEip(),
			"tencentcloud_gaap_proxy":                                          resourceTencentCloudGaapProxy(),
			"tencentcloud_gaap_realserver":                                     resourceTencentCloudGaapRealserver(),
			"tencentcloud_gaap_layer4_listener":                                resourceTencentCloudGaapLayer4Listener(),
			"tencentcloud_gaap_layer7_listener":                                resourceTencentCloudGaapLayer7Listener(),
			"tencentcloud_gaap_http_domain":                                    resourceTencentCloudGaapHttpDomain(),
			"tencentcloud_gaap_http_rule":                                      resourceTencentCloudGaapHttpRule(),
			"tencentcloud_gaap_certificate":                                    resourceTencentCloudGaapCertificate(),
			"tencentcloud_gaap_security_policy":                                resourceTencentCloudGaapSecurityPolicy(),
			"tencentcloud_gaap_security_rule":                                  resourceTencentCloudGaapSecurityRule(),
			"tencentcloud_gaap_domain_error_page":                              resourceTencentCloudGaapDomainErrorPageInfo(),
			"tencentcloud_gaap_global_domain_dns":                              resourceTencentCloudGaapGlobalDomainDns(),
			"tencentcloud_gaap_global_domain":                                  resourceTencentCloudGaapGlobalDomain(),
			"tencentcloud_ssl_certificate":                                     resourceTencentCloudSslCertificate(),
			"tencentcloud_ssl_pay_certificate":                                 resourceTencentCloudSSLInstance(),
			"tencentcloud_ssl_free_certificate":                                resourceTencentCloudSSLFreeCertificate(),
			"tencentcloud_cam_role":                                            resourceTencentCloudCamRole(),
			"tencentcloud_cam_role_by_name":                                    resourceTencentCloudCamRoleByName(),
			"tencentcloud_cam_user":                                            resourceTencentCloudCamUser(),
			"tencentcloud_cam_policy":                                          resourceTencentCloudCamPolicy(),
			"tencentcloud_cam_policy_by_name":                                  resourceTencentCloudCamPolicyByName(),
			"tencentcloud_cam_role_policy_attachment":                          resourceTencentCloudCamRolePolicyAttachment(),
			"tencentcloud_cam_role_policy_attachment_by_name":                  resourceTencentCloudCamRolePolicyAttachmentByName(),
			"tencentcloud_cam_user_policy_attachment":                          resourceTencentCloudCamUserPolicyAttachment(),
			"tencentcloud_cam_group_policy_attachment":                         resourceTencentCloudCamGroupPolicyAttachment(),
			"tencentcloud_cam_group":                                           resourceTencentCloudCamGroup(),
			"tencentcloud_cam_oidc_sso":                                        resourceTencentCloudCamOIDCSSO(),
			"tencentcloud_cam_role_sso":                                        resourceTencentCloudCamRoleSSO(),
			"tencentcloud_cam_group_membership":                                resourceTencentCloudCamGroupMembership(),
			"tencentcloud_cam_saml_provider":                                   resourceTencentCloudCamSAMLProvider(),
			"tencentcloud_cam_service_linked_role":                             resourceTencentCloudCamServiceLinkedRole(),
			"tencentcloud_cam_mfa_flag":                                        resourceTencentCloudCamMfaFlag(),
			"tencentcloud_cam_access_key":                                      resourceTencentCloudCamAccessKey(),
			"tencentcloud_cam_user_saml_config":                                resourceTencentCloudCamUserSamlConfig(),
			"tencentcloud_cam_tag_role_attachment":                             resourceTencentCloudCamTagRoleAttachment(),
			"tencentcloud_cam_policy_version":                                  resourceTencentCloudCamPolicyVersion(),
			"tencentcloud_cam_set_policy_version_config":                       resourceTencentCloudCamSetPolicyVersionConfig(),
			"tencentcloud_cam_user_permission_boundary_attachment":             resourceTencentCloudCamUserPermissionBoundaryAttachment(),
			"tencentcloud_cam_role_permission_boundary_attachment":             resourceTencentCloudCamRolePermissionBoundaryAttachment(),
			"tencentcloud_organization_quit_organization_operation":            resourceTencentCloudOrganizationQuitOrganizationOperation(),
			"tencentcloud_ciam_user_group":                                     resourceTencentCloudCiamUserGroup(),
			"tencentcloud_ciam_user_store":                                     resourceTencentCloudCiamUserStore(),
			"tencentcloud_scf_function":                                        resourceTencentCloudScfFunction(),
			"tencentcloud_scf_function_version":                                resourceTencentCloudScfFunctionVersion(),
			"tencentcloud_scf_function_event_invoke_config":                    resourceTencentCloudScfFunctionEventInvokeConfig(),
			"tencentcloud_scf_reserved_concurrency_config":                     resourceTencentCloudScfReservedConcurrencyConfig(),
			"tencentcloud_scf_provisioned_concurrency_config":                  resourceTencentCloudScfProvisionedConcurrencyConfig(),
			"tencentcloud_scf_invoke_function":                                 resourceTencentCloudScfInvokeFunction(),
			"tencentcloud_scf_sync_invoke_function":                            resourceTencentCloudScfSyncInvokeFunction(),
			"tencentcloud_scf_terminate_async_event":                           resourceTencentCloudScfTerminateAsyncEvent(),
			"tencentcloud_scf_namespace":                                       resourceTencentCloudScfNamespace(),
			"tencentcloud_scf_layer":                                           resourceTencentCloudScfLayer(),
			"tencentcloud_scf_function_alias":                                  resourceTencentCloudScfFunctionAlias(),
			"tencentcloud_scf_trigger_config":                                  resourceTencentCloudScfTriggerConfig(),
			"tencentcloud_tcaplus_cluster":                                     resourceTencentCloudTcaplusCluster(),
			"tencentcloud_tcaplus_tablegroup":                                  resourceTencentCloudTcaplusTableGroup(),
			"tencentcloud_tcaplus_idl":                                         resourceTencentCloudTcaplusIdl(),
			"tencentcloud_tcaplus_table":                                       resourceTencentCloudTcaplusTable(),
			"tencentcloud_cdn_domain":                                          resourceTencentCloudCdnDomain(),
			"tencentcloud_cdn_url_push":                                        resourceTencentCloudUrlPush(),
			"tencentcloud_cdn_url_purge":                                       resourceTencentCloudUrlPurge(),
			"tencentcloud_monitor_policy_group":                                resourceTencentCloudMonitorPolicyGroup(),
			"tencentcloud_monitor_binding_object":                              resourceTencentCloudMonitorBindingObject(),
			"tencentcloud_monitor_policy_binding_object":                       resourceTencentCloudMonitorPolicyBindingObject(),
			"tencentcloud_monitor_binding_receiver":                            resourceTencentCloudMonitorBindingAlarmReceiver(),
			"tencentcloud_monitor_alarm_policy":                                resourceTencentCloudMonitorAlarmPolicy(),
			"tencentcloud_monitor_alarm_notice":                                resourceTencentCloudMonitorAlarmNotice(),
			"tencentcloud_monitor_alarm_policy_set_default":                    resourceTencentCloudMonitorAlarmPolicySetDefault(),
			"tencentcloud_monitor_tmp_instance":                                resourceTencentCloudMonitorTmpInstance(),
			"tencentcloud_monitor_tmp_cvm_agent":                               resourceTencentCloudMonitorTmpCvmAgent(),
			"tencentcloud_monitor_tmp_scrape_job":                              resourceTencentCloudMonitorTmpScrapeJob(),
			"tencentcloud_monitor_tmp_exporter_integration":                    resourceTencentCloudMonitorTmpExporterIntegration(),
			"tencentcloud_monitor_tmp_alert_rule":                              resourceTencentCloudMonitorTmpAlertRule(),
			"tencentcloud_monitor_tmp_recording_rule":                          resourceTencentCloudMonitorTmpRecordingRule(),
			"tencentcloud_monitor_tmp_tke_template":                            resourceTencentCloudMonitorTmpTkeTemplate(),
			"tencentcloud_monitor_tmp_tke_template_attachment":                 resourceTencentCloudMonitorTmpTkeTemplateAttachment(),
			"tencentcloud_monitor_tmp_tke_alert_policy":                        resourceTencentCloudMonitorTmpTkeAlertPolicy(),
			"tencentcloud_monitor_tmp_tke_basic_config":                        resourceTencentCloudMonitorTmpTkeBasicConfig(),
			"tencentcloud_monitor_tmp_tke_cluster_agent":                       resourceTencentCloudMonitorTmpTkeClusterAgent(),
			"tencentcloud_monitor_tmp_tke_config":                              resourceTencentCloudMonitorTmpTkeConfig(),
			"tencentcloud_monitor_tmp_tke_record_rule_yaml":                    resourceTencentCloudMonitorTmpTkeRecordRuleYaml(),
			"tencentcloud_monitor_tmp_tke_global_notification":                 resourceTencentCloudMonitorTmpTkeGlobalNotification(),
			"tencentcloud_monitor_tmp_manage_grafana_attachment":               resourceTencentCloudMonitorTmpManageGrafanaAttachment(),
			"tencentcloud_monitor_grafana_instance":                            resourceTencentCloudMonitorGrafanaInstance(),
			"tencentcloud_monitor_grafana_integration":                         resourceTencentCloudMonitorGrafanaIntegration(),
			"tencentcloud_monitor_grafana_notification_channel":                resourceTencentCloudMonitorGrafanaNotificationChannel(),
			"tencentcloud_monitor_grafana_plugin":                              resourceTencentCloudMonitorGrafanaPlugin(),
			"tencentcloud_monitor_grafana_sso_account":                         resourceTencentCloudMonitorGrafanaSsoAccount(),
			"tencentcloud_monitor_tmp_grafana_config":                          resourceTencentCloudMonitorTmpGrafanaConfig(),
			"tencentcloud_monitor_grafana_dns_config":                          resourceTencentCloudMonitorGrafanaDnsConfig(),
			"tencentcloud_monitor_grafana_env_config":                          resourceTencentCloudMonitorGrafanaEnvConfig(),
			"tencentcloud_monitor_grafana_whitelist_config":                    resourceTencentCloudMonitorGrafanaWhitelistConfig(),
			"tencentcloud_monitor_grafana_sso_cam_config":                      resourceTencentCloudMonitorGrafanaSsoCamConfig(),
			"tencentcloud_monitor_grafana_sso_config":                          resourceTencentCloudMonitorGrafanaSsoConfig(),
			"tencentcloud_monitor_grafana_version_upgrade":                     resourceTencentCloudMonitorGrafanaVersionUpgrade(),
			"tencentcloud_mongodb_standby_instance":                            resourceTencentCloudMongodbStandbyInstance(),
			"tencentcloud_elasticsearch_instance":                              resourceTencentCloudElasticsearchInstance(),
			"tencentcloud_elasticsearch_security_group":                        resourceTencentCloudElasticsearchSecurityGroup(),
			"tencentcloud_elasticsearch_logstash":                              resourceTencentCloudElasticsearchLogstash(),
			"tencentcloud_elasticsearch_logstash_pipeline":                     resourceTencentCloudElasticsearchLogstashPipeline(),
			"tencentcloud_elasticsearch_restart_logstash_instance_operation":   resourceTencentCloudElasticsearchRestartLogstashInstanceOperation(),
			"tencentcloud_elasticsearch_start_logstash_pipeline_operation":     resourceTencentCloudElasticsearchStartLogstashPipelineOperation(),
			"tencentcloud_elasticsearch_stop_logstash_pipeline_operation":      resourceTencentCloudElasticsearchStopLogstashPipelineOperation(),
			"tencentcloud_elasticsearch_index":                                 resourceTencentCloudElasticsearchIndex(),
			"tencentcloud_elasticsearch_restart_instance_operation":            resourceTencentCloudElasticsearchRestartInstanceOperation(),
			"tencentcloud_elasticsearch_restart_kibana_operation":              resourceTencentCloudElasticsearchRestartKibanaOperation(),
			"tencentcloud_elasticsearch_restart_nodes_operation":               resourceTencentCloudElasticsearchRestartNodesOperation(),
			"tencentcloud_elasticsearch_diagnose":                              resourceTencentCloudElasticsearchDiagnose(),
			"tencentcloud_elasticsearch_diagnose_instance":                     resourceTencentCloudElasticsearchDiagnoseInstance(),
			"tencentcloud_elasticsearch_update_plugins_operation":              resourceTencentCloudElasticsearchUpdatePluginsOperation(),
			"tencentcloud_postgresql_instance":                                 resourceTencentCloudPostgresqlInstance(),
			"tencentcloud_postgresql_readonly_instance":                        resourceTencentCloudPostgresqlReadonlyInstance(),
			"tencentcloud_postgresql_readonly_group":                           resourceTencentCloudPostgresqlReadonlyGroup(),
			"tencentcloud_postgresql_readonly_attachment":                      resourceTencentCloudPostgresqlReadonlyAttachment(),
			"tencentcloud_postgresql_parameter_template":                       resourceTencentCloudPostgresqlParameterTemplate(),
			"tencentcloud_postgresql_base_backup":                              resourceTencentCloudPostgresqlBaseBackup(),
			"tencentcloud_postgresql_backup_plan_config":                       resourceTencentCloudPostgresqlBackupPlanConfig(),
			"tencentcloud_postgresql_security_group_config":                    resourceTencentCloudPostgresqlSecurityGroupConfig(),
			"tencentcloud_postgresql_backup_download_restriction_config":       resourceTencentCloudPostgresqlBackupDownloadRestrictionConfig(),
			"tencentcloud_postgresql_restart_db_instance_operation":            resourceTencentCloudPostgresqlRestartDbInstanceOperation(),
			"tencentcloud_postgresql_renew_db_instance_operation":              resourceTencentCloudPostgresqlRenewDbInstanceOperation(),
			"tencentcloud_postgresql_isolate_db_instance_operation":            resourceTencentCloudPostgresqlIsolateDbInstanceOperation(),
			"tencentcloud_postgresql_disisolate_db_instance_operation":         resourceTencentCloudPostgresqlDisisolateDbInstanceOperation(),
			"tencentcloud_postgresql_rebalance_readonly_group_operation":       resourceTencentCloudPostgresqlRebalanceReadonlyGroupOperation(),
			"tencentcloud_postgresql_delete_log_backup_operation":              resourceTencentCloudPostgresqlDeleteLogBackupOperation(),
			"tencentcloud_postgresql_modify_account_remark_operation":          resourceTencentCloudPostgresqlModifyAccountRemarkOperation(),
			"tencentcloud_postgresql_modify_switch_time_period_operation":      resourceTencentCloudPostgresqlModifySwitchTimePeriodOperation(),
			"tencentcloud_sqlserver_instance":                                  resourceTencentCloudSqlserverInstance(),
			"tencentcloud_sqlserver_db":                                        resourceTencentCloudSqlserverDB(),
			"tencentcloud_sqlserver_account":                                   resourceTencentCloudSqlserverAccount(),
			"tencentcloud_sqlserver_account_db_attachment":                     resourceTencentCloudSqlserverAccountDBAttachment(),
			"tencentcloud_sqlserver_readonly_instance":                         resourceTencentCloudSqlserverReadonlyInstance(),
			"tencentcloud_sqlserver_migration":                                 resourceTencentCloudSqlserverMigration(),
			"tencentcloud_sqlserver_config_backup_strategy":                    resourceTencentCloudSqlserverConfigBackupStrategy(),
			"tencentcloud_sqlserver_general_backup":                            resourceTencentCloudSqlserverGeneralBackup(),
			"tencentcloud_sqlserver_general_clone":                             resourceTencentCloudSqlserverGeneralClone(),
			"tencentcloud_sqlserver_full_backup_migration":                     resourceTencentCloudSqlserverFullBackupMigration(),
			"tencentcloud_sqlserver_incre_backup_migration":                    resourceTencentCloudSqlserverIncreBackupMigration(),
			"tencentcloud_sqlserver_business_intelligence_file":                resourceTencentCloudSqlserverBusinessIntelligenceFile(),
			"tencentcloud_sqlserver_business_intelligence_instance":            resourceTencentCloudSqlserverBusinessIntelligenceInstance(),
			"tencentcloud_sqlserver_general_communication":                     resourceTencentCloudSqlserverGeneralCommunication(),
			"tencentcloud_sqlserver_general_cloud_instance":                    resourceTencentCloudSqlserverGeneralCloudInstance(),
			"tencentcloud_sqlserver_complete_expansion":                        resourceTencentCloudSqlserverCompleteExpansion(),
			"tencentcloud_sqlserver_config_database_cdc":                       resourceTencentCloudSqlserverConfigDatabaseCDC(),
			"tencentcloud_sqlserver_config_database_ct":                        resourceTencentCloudSqlserverConfigDatabaseCT(),
			"tencentcloud_sqlserver_config_database_mdf":                       resourceTencentCloudSqlserverConfigDatabaseMdf(),
			"tencentcloud_sqlserver_config_instance_param":                     resourceTencentCloudSqlserverConfigInstanceParam(),
			"tencentcloud_sqlserver_config_instance_ro_group":                  resourceTencentCloudSqlserverConfigInstanceRoGroup(),
			"tencentcloud_sqlserver_config_instance_security_groups":           resourceTencentCloudSqlserverConfigInstanceSecurityGroups(),
			"tencentcloud_sqlserver_renew_db_instance":                         resourceTencentCloudSqlserverRenewDBInstance(),
			"tencentcloud_sqlserver_renew_postpaid_db_instance":                resourceTencentCloudSqlserverRenewPostpaidDBInstance(),
			"tencentcloud_sqlserver_restart_db_instance":                       resourceTencentCloudSqlserverRestartDBInstance(),
			"tencentcloud_sqlserver_config_terminate_db_instance":              resourceTencentCloudSqlserverConfigTerminateDBInstance(),
			"tencentcloud_sqlserver_restore_instance":                          resourceTencentCloudSqlserverRestoreInstance(),
			"tencentcloud_sqlserver_rollback_instance":                         resourceTencentCloudSqlserverRollbackInstance(),
			"tencentcloud_sqlserver_start_backup_full_migration":               resourceTencentCloudSqlserverStartBackupFullMigration(),
			"tencentcloud_sqlserver_start_backup_incremental_migration":        resourceTencentCloudSqlserverStartBackupIncrementalMigration(),
			"tencentcloud_sqlserver_start_xevent":                              resourceTencentCloudSqlserverStartXevent(),
			"tencentcloud_ckafka_instance":                                     resourceTencentCloudCkafkaInstance(),
			"tencentcloud_ckafka_user":                                         resourceTencentCloudCkafkaUser(),
			"tencentcloud_ckafka_acl":                                          resourceTencentCloudCkafkaAcl(),
			"tencentcloud_ckafka_topic":                                        resourceTencentCloudCkafkaTopic(),
			"tencentcloud_ckafka_datahub_topic":                                resourceTencentCloudCkafkaDatahubTopic(),
			"tencentcloud_ckafka_connect_resource":                             resourceTencentCloudCkafkaConnectResource(),
			"tencentcloud_ckafka_renew_instance":                               resourceTencentCloudCkafkaRenewInstance(),
			"tencentcloud_ckafka_acl_rule":                                     resourceTencentCloudCkafkaAclRule(),
			"tencentcloud_ckafka_consumer_group":                               resourceTencentCloudCkafkaConsumerGroup(),
			"tencentcloud_ckafka_consumer_group_modify_offset":                 resourceTencentCloudCkafkaConsumerGroupModifyOffset(),
			"tencentcloud_ckafka_datahub_task":                                 resourceTencentCloudCkafkaDatahubTask(),
			"tencentcloud_ckafka_route":                                        resourceTencentCloudCkafkaRoute(),
			"tencentcloud_audit":                                               resourceTencentCloudAudit(),
			"tencentcloud_audit_track":                                         resourceTencentCloudAuditTrack(),
			"tencentcloud_image":                                               resourceTencentCloudImage(),
			"tencentcloud_cynosdb_proxy":                                       resourceTencentCloudCynosdbProxy(),
			"tencentcloud_cynosdb_reload_proxy_node":                           resourceTencentCloudCynosdbReloadProxyNode(),
			"tencentcloud_cynosdb_cluster_resource_packages_attachment":        resourceTencentCloudCynosdbClusterResourcePackagesAttachment(),
			"tencentcloud_cynosdb_cluster":                                     resourceTencentCloudCynosdbCluster(),
			"tencentcloud_cynosdb_readonly_instance":                           resourceTencentCloudCynosdbReadonlyInstance(),
			"tencentcloud_cynosdb_cluster_password_complexity":                 resourceTencentCloudCynosdbClusterPasswordComplexity(),
			"tencentcloud_cynosdb_export_instance_error_logs":                  resourceTencentCloudCynosdbExportInstanceErrorLogs(),
			"tencentcloud_cynosdb_export_instance_slow_queries":                resourceTencentCloudCynosdbExportInstanceSlowQueries(),
			"tencentcloud_cynosdb_account_privileges":                          resourceTencentCloudCynosdbAccountPrivileges(),
			"tencentcloud_cynosdb_account":                                     resourceTencentCloudCynosdbAccount(),
			"tencentcloud_cynosdb_binlog_save_days":                            resourceTencentCloudCynosdbBinlogSaveDays(),
			"tencentcloud_cynosdb_cluster_databases":                           resourceTencentCloudCynosdbClusterDatabases(),
			"tencentcloud_cynosdb_instance_param":                              resourceTencentCloudCynosdbInstanceParam(),
			"tencentcloud_cynosdb_isolate_instance":                            resourceTencentCloudCynosdbIsolateInstance(),
			"tencentcloud_cynosdb_param_template":                              resourceTencentCloudCynosdbParamTemplate(),
			"tencentcloud_cynosdb_resource_package":                            resourceTencentCloudCynosdbResourcePackage(),
			"tencentcloud_cynosdb_restart_instance":                            resourceTencentCloudCynosdbRestartInstance(),
			"tencentcloud_cynosdb_roll_back_cluster":                           resourceTencentCloudCynosdbRollBackCluster(),
			"tencentcloud_cynosdb_wan":                                         resourceTencentCloudCynosdbWan(),
			"tencentcloud_cynosdb_cluster_slave_zone":                          resourceTencentCloudCynosdbClusterSlaveZone(),
			"tencentcloud_cynosdb_read_only_instance_exclusive_access":         resourceTencentCloudCynosdbReadOnlyInstanceExclusiveAccess(),
			"tencentcloud_cynosdb_proxy_end_point":                             resourceTencentCloudCynosdbProxyEndPoint(),
			"tencentcloud_cynosdb_upgrade_proxy_version":                       resourceTencentCloudCynosdbUpgradeProxyVersion(),
			"tencentcloud_vod_adaptive_dynamic_streaming_template":             resourceTencentCloudVodAdaptiveDynamicStreamingTemplate(),
			"tencentcloud_vod_image_sprite_template":                           resourceTencentCloudVodImageSpriteTemplate(),
			"tencentcloud_vod_procedure_template":                              resourceTencentCloudVodProcedureTemplate(),
			"tencentcloud_vod_snapshot_by_time_offset_template":                resourceTencentCloudVodSnapshotByTimeOffsetTemplate(),
			"tencentcloud_vod_super_player_config":                             resourceTencentCloudVodSuperPlayerConfig(),
			"tencentcloud_vod_sub_application":                                 resourceTencentCloudVodSubApplication(),
			"tencentcloud_sqlserver_publish_subscribe":                         resourceTencentCloudSqlserverPublishSubscribe(),
			"tencentcloud_api_gateway_usage_plan":                              resourceTencentCloudAPIGatewayUsagePlan(),
			"tencentcloud_api_gateway_usage_plan_attachment":                   resourceTencentCloudAPIGatewayUsagePlanAttachment(),
			"tencentcloud_api_gateway_api":                                     resourceTencentCloudAPIGatewayAPI(),
			"tencentcloud_api_gateway_service":                                 resourceTencentCloudAPIGatewayService(),
			"tencentcloud_api_gateway_custom_domain":                           resourceTencentCloudAPIGatewayCustomDomain(),
			"tencentcloud_api_gateway_ip_strategy":                             resourceTencentCloudAPIGatewayIPStrategy(),
			"tencentcloud_api_gateway_strategy_attachment":                     resourceTencentCloudAPIGatewayStrategyAttachment(),
			"tencentcloud_api_gateway_api_key":                                 resourceTencentCloudAPIGatewayAPIKey(),
			"tencentcloud_api_gateway_api_key_attachment":                      resourceTencentCloudAPIGatewayAPIKeyAttachment(),
			"tencentcloud_api_gateway_service_release":                         resourceTencentCloudAPIGatewayServiceRelease(),
			"tencentcloud_api_gateway_plugin":                                  resourceTencentCloudAPIGatewayPlugin(),
			"tencentcloud_api_gateway_plugin_attachment":                       resourceTencentCloudAPIGatewayPluginAttachment(),
			"tencentcloud_api_gateway_upstream":                                resourceTencentCloudAPIGatewayUpstream(),
			"tencentcloud_api_gateway_api_app_attachment":                      resourceTencentCloudAPIGatewayApiAppAttachment(),
			"tencentcloud_sqlserver_basic_instance":                            resourceTencentCloudSqlserverBasicInstance(),
			"tencentcloud_sqlserver_instance_tde":                              resourceTencentCloudSqlserverInstanceTDE(),
			"tencentcloud_sqlserver_database_tde":                              resourceTencentCloudSqlserverDatabaseTDE(),
			"tencentcloud_sqlserver_general_cloud_ro_instance":                 resourceTencentCloudSqlserverGeneralCloudRoInstance(),
			"tencentcloud_tcr_instance":                                        resourceTencentCloudTcrInstance(),
			"tencentcloud_tcr_namespace":                                       resourceTencentCloudTcrNamespace(),
			"tencentcloud_tcr_repository":                                      resourceTencentCloudTcrRepository(),
			"tencentcloud_tcr_token":                                           resourceTencentCloudTcrToken(),
			"tencentcloud_tcr_vpc_attachment":                                  resourceTencentCloudTcrVpcAttachment(),
			"tencentcloud_tcr_tag_retention_rule":                              resourceTencentCloudTcrTagRetentionRule(),
			"tencentcloud_tcr_webhook_trigger":                                 resourceTencentCloudTcrWebhookTrigger(),
			"tencentcloud_tcr_manage_replication_operation":                    resourceTencentCloudTcrManageReplicationOperation(),
			"tencentcloud_tcr_customized_domain":                               resourceTencentCloudTcrCustomizedDomain(),
			"tencentcloud_tcr_immutable_tag_rule":                              resourceTencentCloudTcrImmutableTagRule(),
			"tencentcloud_tcr_delete_image_operation":                          resourceTencentCloudTcrDeleteImageOperation(),
			"tencentcloud_tcr_create_image_signature_operation":                resourceTencentCloudTcrCreateImageSignatureOperation(),
			"tencentcloud_tcr_tag_retention_execution_config":                  resourceTencentCloudTcrTagRetentionExecutionConfig(),
			"tencentcloud_tcr_service_account":                                 resourceTencentCloudTcrServiceAccount(),
			"tencentcloud_tdmq_instance":                                       resourceTencentCloudTdmqInstance(),
			"tencentcloud_tdmq_namespace":                                      resourceTencentCloudTdmqNamespace(),
			"tencentcloud_tdmq_topic":                                          resourceTencentCloudTdmqTopic(),
			"tencentcloud_tdmq_role":                                           resourceTencentCloudTdmqRole(),
			"tencentcloud_tdmq_namespace_role_attachment":                      resourceTencentCloudTdmqNamespaceRoleAttachment(),
			"tencentcloud_tdmq_subscription_attachment":                        resourceTencentCloudTdmqSubscriptionAttachment(),
			"tencentcloud_tdmq_rabbitmq_user":                                  resourceTencentCloudTdmqRabbitmqUser(),
			"tencentcloud_tdmq_rabbitmq_virtual_host":                          resourceTencentCloudTdmqRabbitmqVirtualHost(),
			"tencentcloud_tdmq_rabbitmq_vip_instance":                          resourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloud_tdmq_send_rocketmq_message":                          resourceTencentCloudTdmqSendRocketmqMessage(),
			"tencentcloud_cos_bucket_policy":                                   resourceTencentCloudCosBucketPolicy(),
			"tencentcloud_cos_bucket_domain_certificate_attachment":            resourceTencentCloudCosBucketDomainCertificateAttachment(),
			"tencentcloud_cos_bucket_inventory":                                resourceTencentCloudCosBucketInventory(),
			"tencentcloud_cos_batch":                                           resourceTencentCloudCosBatch(),
			"tencentcloud_cos_object_abort_multipart_upload_operation":         resourceTencentCloudCosObjectAbortMultipartUploadOperation(),
			"tencentcloud_cos_object_copy_operation":                           resourceTencentCloudCosObjectCopyOperation(),
			"tencentcloud_cos_object_restore_operation":                        resourceTencentCloudCosObjectRestoreOperation(),
			"tencentcloud_cos_bucket_generate_inventory_immediately_operation": resourceTencentCloudCosBucketGenerateInventoryImmediatelyOperation(),
			"tencentcloud_cos_object_download_operation":                       resourceTencentCloudCosObjectDownloadOperation(),
			"tencentcloud_address_template":                                    resourceTencentCloudAddressTemplate(),
			"tencentcloud_address_template_group":                              resourceTencentCloudAddressTemplateGroup(),
			"tencentcloud_protocol_template":                                   resourceTencentCloudProtocolTemplate(),
			"tencentcloud_protocol_template_group":                             resourceTencentCloudProtocolTemplateGroup(),
			"tencentcloud_kms_key":                                             resourceTencentCloudKmsKey(),
			"tencentcloud_kms_external_key":                                    resourceTencentCloudKmsExternalKey(),
			"tencentcloud_kms_white_box_key":                                   resourceTencentCloudKmsWhiteBoxKey(),
			"tencentcloud_kms_cloud_resource_attachment":                       resourceTencentCloudKmsCloudResourceAttachment(),
			"tencentcloud_kms_overwrite_white_box_device_fingerprints":         resourceTencentCloudKmsOverwriteWhiteBoxDeviceFingerprints(),
			"tencentcloud_ssm_secret":                                          resourceTencentCloudSsmSecret(),
			"tencentcloud_ssm_ssh_key_pair_secret":                             resourceTencentCloudSsmSshKeyPairSecret(),
			"tencentcloud_ssm_product_secret":                                  resourceTencentCloudSsmProductSecret(),
			"tencentcloud_ssm_secret_version":                                  resourceTencentCloudSsmSecretVersion(),
			"tencentcloud_ssm_rotate_product_secret":                           resourceTencentCloudSsmRotateProductSecret(),
			"tencentcloud_cdh_instance":                                        resourceTencentCloudCdhInstance(),
			"tencentcloud_dnspod_domain_instance":                              resourceTencentCloudDnspodDomainInstance(),
			"tencentcloud_dnspod_domain_alias":                                 resourceTencentCloudDnspodDomainAlias(),
			"tencentcloud_dnspod_record":                                       resourceTencentCloudDnspodRecord(),
			"tencentcloud_dnspod_record_group":                                 resourceTencentCloudDnspodRecordGroup(),
			"tencentcloud_dnspod_modify_domain_owner_operation":                resourceTencentCloudDnspodModifyDomainOwnerOperation(),
			"tencentcloud_dnspod_modify_record_group_operation":                resourceTencentCloudDnspodModifyRecordGroupOperation(),
			"tencentcloud_dnspod_download_snapshot_operation":                  resourceTencentCloudDnspodDownloadSnapshotOperation(),
			"tencentcloud_dnspod_custom_line":                                  resourceTencentCloudDnspodCustomLine(),
			"tencentcloud_dnspod_snapshot_config":                              resourceTencentCloudDnspodSnapshotConfig(),
			"tencentcloud_dnspod_domain_lock":                                  resourceTencentCloudDnspodDomainLock(),
			"tencentcloud_private_dns_zone":                                    resourceTencentCloudPrivateDnsZone(),
			"tencentcloud_private_dns_record":                                  resourceTencentCloudPrivateDnsRecord(),
			"tencentcloud_private_dns_zone_vpc_attachment":                     resourceTencentCloudPrivateDnsZoneVpcAttachment(),
			"tencentcloud_cls_logset":                                          resourceTencentCloudClsLogset(),
			"tencentcloud_cls_topic":                                           resourceTencentCloudClsTopic(),
			"tencentcloud_cls_config":                                          resourceTencentCloudClsConfig(),
			"tencentcloud_cls_config_extra":                                    resourceTencentCloudClsConfigExtra(),
			"tencentcloud_cls_config_attachment":                               resourceTencentCloudClsConfigAttachment(),
			"tencentcloud_cls_machine_group":                                   resourceTencentCloudClsMachineGroup(),
			"tencentcloud_cls_cos_shipper":                                     resourceTencentCloudClsCosShipper(),
			"tencentcloud_cls_index":                                           resourceTencentCloudClsIndex(),
			"tencentcloud_cls_alarm":                                           resourceTencentCloudClsAlarm(),
			"tencentcloud_cls_alarm_notice":                                    resourceTencentCloudClsAlarmNotice(),
			"tencentcloud_cls_ckafka_consumer":                                 resourceTencentCloudClsCkafkaConsumer(),
			"tencentcloud_cls_cos_recharge":                                    resourceTencentCloudClsCosRecharge(),
			"tencentcloud_cls_export":                                          resourceTencentCloudClsExport(),
			"tencentcloud_cls_data_transform":                                  resourceTencentCloudClsDataTransform(),
			"tencentcloud_lighthouse_instance":                                 resourceTencentCloudLighthouseInstance(),
			"tencentcloud_lighthouse_firewall_template":                        resourceTencentCloudLighthouseFirewallTemplate(),
			"tencentcloud_tem_environment":                                     resourceTencentCloudTemEnvironment(),
			"tencentcloud_tem_application":                                     resourceTencentCloudTemApplication(),
			"tencentcloud_tem_workload":                                        resourceTencentCloudTemWorkload(),
			"tencentcloud_tem_app_config":                                      resourceTencentCloudTemAppConfig(),
			"tencentcloud_tem_log_config":                                      resourceTencentCloudTemLogConfig(),
			"tencentcloud_tem_scale_rule":                                      resourceTencentCloudTemScaleRule(),
			"tencentcloud_tem_gateway":                                         resourceTencentCloudTemGateway(),
			"tencentcloud_tem_application_service":                             resourceTencentCloudTemApplicationService(),
			"tencentcloud_teo_zone":                                            resourceTencentCloudTeoZone(),
			"tencentcloud_teo_zone_setting":                                    resourceTencentCloudTeoZoneSetting(),
			"tencentcloud_teo_origin_group":                                    resourceTencentCloudTeoOriginGroup(),
			"tencentcloud_teo_rule_engine":                                     resourceTencentCloudTeoRuleEngine(),
			"tencentcloud_teo_ownership_verify":                                resourceTencentCloudTeoOwnershipVerify(),
			"tencentcloud_teo_certificate_config":                              resourceTencentCloudTeoCertificateConfig(),
			"tencentcloud_teo_acceleration_domain":                             resourceTencentCloudTeoAccelerationDomain(),
			"tencentcloud_teo_application_proxy":                               resourceTencentCloudTeoApplicationProxy(),
			"tencentcloud_teo_application_proxy_rule":                          resourceTencentCloudTeoApplicationProxyRule(),
			"tencentcloud_tcm_mesh":                                            resourceTencentCloudTcmMesh(),
			"tencentcloud_tcm_cluster_attachment":                              resourceTencentCloudTcmClusterAttachment(),
			"tencentcloud_tcm_prometheus_attachment":                           resourceTencentCloudTcmPrometheusAttachment(),
			"tencentcloud_tcm_tracing_config":                                  resourceTencentCloudTcmTracingConfig(),
			"tencentcloud_tcm_access_log_config":                               resourceTencentCloudTcmAccessLogConfig(),
			"tencentcloud_ses_domain":                                          resourceTencentCloudSesDomain(),
			"tencentcloud_ses_template":                                        resourceTencentCloudSesTemplate(),
			"tencentcloud_ses_email_address":                                   resourceTencentCloudSesEmailAddress(),
			"tencentcloud_ses_receiver":                                        resourceTencentCloudSesReceiver(),
			"tencentcloud_ses_send_email":                                      resourceTencentCloudSesSendEmail(),
			"tencentcloud_ses_batch_send_email":                                resourceTencentCloudSesBatchSendEmail(),
			"tencentcloud_ses_verify_domain":                                   resourceTencentCloudSesVerifyDomain(),
			"tencentcloud_ses_black_list_delete":                               resourceTencentCloudSesBlackListDelete(),
			"tencentcloud_sms_sign":                                            resourceTencentCloudSmsSign(),
			"tencentcloud_sms_template":                                        resourceTencentCloudSmsTemplate(),
			"tencentcloud_dcdb_account":                                        resourceTencentCloudDcdbAccount(),
			"tencentcloud_dcdb_hourdb_instance":                                resourceTencentCloudDcdbHourdbInstance(),
			"tencentcloud_dcdb_security_group_attachment":                      resourceTencentCloudDcdbSecurityGroupAttachment(),
			"tencentcloud_dcdb_db_instance":                                    resourceTencentCloudDcdbDbInstance(),
			"tencentcloud_dcdb_account_privileges":                             resourceTencentCloudDcdbAccountPrivileges(),
			"tencentcloud_dcdb_db_parameters":                                  resourceTencentCloudDcdbDbParameters(),
			"tencentcloud_dcdb_encrypt_attributes_config":                      resourceTencentCloudDcdbEncryptAttributesConfig(),
			"tencentcloud_dcdb_db_sync_mode_config":                            resourceTencentCloudDcdbDbSyncModeConfig(),
			"tencentcloud_dcdb_instance_config":                                resourceTencentCloudDcdbInstanceConfig(),
			"tencentcloud_dcdb_activate_hour_instance_operation":               resourceTencentCloudDcdbActivateHourInstanceOperation(),
			"tencentcloud_dcdb_isolate_hour_instance_operation":                resourceTencentCloudDcdbIsolateHourInstanceOperation(),
			"tencentcloud_dcdb_cancel_dcn_job_operation":                       resourceTencentCloudDcdbCancelDcnJobOperation(),
			"tencentcloud_dcdb_flush_binlog_operation":                         resourceTencentCloudDcdbFlushBinlogOperation(),
			"tencentcloud_dcdb_switch_db_instance_ha_operation":                resourceTencentCloudDcdbSwitchDbInstanceHaOperation(),
			"tencentcloud_cat_task_set":                                        resourceTencentCloudCatTaskSet(),
			"tencentcloud_mariadb_dedicatedcluster_db_instance":                resourceTencentCloudMariadbDedicatedclusterDbInstance(),
			"tencentcloud_mariadb_instance":                                    resourceTencentCloudMariadbInstance(),
			"tencentcloud_mariadb_hour_db_instance":                            resourceTencentCloudMariadbHourDbInstance(),
			"tencentcloud_mariadb_account":                                     resourceTencentCloudMariadbAccount(),
			"tencentcloud_mariadb_parameters":                                  resourceTencentCloudMariadbParameters(),
			"tencentcloud_mariadb_log_file_retention_period":                   resourceTencentCloudMariadbLogFileRetentionPeriod(),
			"tencentcloud_mariadb_security_groups":                             resourceTencentCloudMariadbSecurityGroups(),
			"tencentcloud_mariadb_encrypt_attributes":                          resourceTencentCloudMariadbEncryptAttributes(),
			"tencentcloud_mariadb_account_privileges":                          resourceTencentCloudMariadbAccountPrivileges(),
			"tencentcloud_mariadb_operate_hour_db_instance":                    resourceTencentCloudMariadbOperateHourDbInstance(),
			"tencentcloud_mariadb_backup_time":                                 resourceTencentCloudMariadbBackupTime(),
			"tencentcloud_mariadb_cancel_dcn_job":                              resourceTencentCloudMariadbCancelDcnJob(),
			"tencentcloud_mariadb_flush_binlog":                                resourceTencentCloudMariadbFlushBinlog(),
			"tencentcloud_mariadb_switch_ha":                                   resourceTencentCloudMariadbSwitchHA(),
			"tencentcloud_mariadb_restart_instance":                            resourceTencentCloudMariadbRestartInstance(),
			"tencentcloud_mariadb_renew_instance":                              resourceTencentCloudMariadbRenewInstance(),
			"tencentcloud_mariadb_instance_config":                             resourceTencentCloudMariadbInstanceConfig(),
			"tencentcloud_tdcpg_cluster":                                       resourceTencentCloudTdcpgCluster(),
			"tencentcloud_tdcpg_instance":                                      resourceTencentCloudTdcpgInstance(),
			"tencentcloud_css_watermark":                                       resourceTencentCloudCssWatermark(),
			"tencentcloud_css_watermark_rule_attachment":                       resourceTencentCloudCssWatermarkRuleAttachment(),
			"tencentcloud_css_pull_stream_task":                                resourceTencentCloudCssPullStreamTask(),
			"tencentcloud_css_live_transcode_template":                         resourceTencentCloudCssLiveTranscodeTemplate(),
			"tencentcloud_css_live_transcode_rule_attachment":                  resourceTencentCloudCssLiveTranscodeRuleAttachment(),
			"tencentcloud_css_domain":                                          resourceTencentCloudCssDomain(),
			"tencentcloud_css_authenticate_domain_owner_operation":             resourceTencentCloudCssAuthenticateDomainOwnerOperation(),
			"tencentcloud_css_play_domain_cert_attachment":                     resourceTencentCloudCssPlayDomainCertAttachment(),
			"tencentcloud_css_play_auth_key_config":                            resourceTencentCloudCssPlayAuthKeyConfig(),
			"tencentcloud_css_push_auth_key_config":                            resourceTencentCloudCssPushAuthKeyConfig(),
			"tencentcloud_css_backup_stream":                                   resourceTencentCloudCssBackupStream(),
			"tencentcloud_css_callback_rule_attachment":                        resourceTencentCloudCssCallbackRuleAttachment(),
			"tencentcloud_css_callback_template":                               resourceTencentCloudCssCallbackTemplate(),
			"tencentcloud_css_domain_referer":                                  resourceTencentCloudCssDomainReferer(),
			"tencentcloud_css_enable_optimal_switching":                        resourceTencentCloudCssEnableOptimalSwitching(),
			"tencentcloud_css_record_rule_attachment":                          resourceTencentCloudCssRecordRuleAttachment(),
			"tencentcloud_css_record_template":                                 resourceTencentCloudCssRecordTemplate(),
			"tencentcloud_css_snapshot_rule_attachment":                        resourceTencentCloudCssSnapshotRuleAttachment(),
			"tencentcloud_css_snapshot_template":                               resourceTencentCloudCssSnapshotTemplate(),
			"tencentcloud_css_pad_template":                                    resourceTencentCloudCssPadTemplate(),
			"tencentcloud_css_pad_rule_attachment":                             resourceTencentCloudCssPadRuleAttachment(),
			"tencentcloud_css_timeshift_template":                              resourceTencentCloudCssTimeshiftTemplate(),
			"tencentcloud_css_timeshift_rule_attachment":                       resourceTencentCloudCssTimeshiftRuleAttachment(),
			"tencentcloud_css_stream_monitor":                                  resourceTencentCloudCssStreamMonitor(),
			"tencentcloud_css_start_stream_monitor":                            resourceTencentCloudCssStartStreamMonitor(),
			"tencentcloud_css_pull_stream_task_restart":                        resourceTencentCloudCssPullStreamTaskRestart(),
			"tencentcloud_pts_project":                                         resourceTencentCloudPtsProject(),
			"tencentcloud_pts_alert_channel":                                   resourceTencentCloudPtsAlertChannel(),
			"tencentcloud_pts_scenario":                                        resourceTencentCloudPtsScenario(),
			"tencentcloud_pts_file":                                            resourceTencentCloudPtsFile(),
			"tencentcloud_pts_job":                                             resourceTencentCloudPtsJob(),
			"tencentcloud_pts_cron_job":                                        resourceTencentCloudPtsCronJob(),
			"tencentcloud_pts_tmp_key_generate":                                resourceTencentCloudPtsTmpKeyGenerate(),
			"tencentcloud_pts_cron_job_restart":                                resourceTencentCloudPtsCronJobRestart(),
			"tencentcloud_pts_job_abort":                                       resourceTencentCloudPtsJobAbort(),
			"tencentcloud_pts_cron_job_abort":                                  resourceTencentCloudPtsCronJobAbort(),
			"tencentcloud_tat_command":                                         resourceTencentCloudTatCommand(),
			"tencentcloud_tat_invoker":                                         resourceTencentCloudTatInvoker(),
			"tencentcloud_tat_invoker_config":                                  resourceTencentCloudTatInvokerConfig(),
			"tencentcloud_tat_invocation_invoke_attachment":                    resourceTencentCloudTatInvocationInvokeAttachment(),
			"tencentcloud_tat_invocation_command_attachment":                   resourceTencentCloudTatInvocationCommandAttachment(),
			"tencentcloud_organization_org_node":                               resourceTencentCloudOrganizationOrgNode(),
			"tencentcloud_organization_org_member":                             resourceTencentCloudOrganizationOrgMember(),
			"tencentcloud_organization_org_identity":                           resourceTencentCloudOrganizationOrgIdentity(),
			"tencentcloud_organization_org_member_email":                       resourceTencentCloudOrganizationOrgMemberEmail(),
			"tencentcloud_organization_instance":                               resourceTencentCloudOrganizationOrganization(),
			"tencentcloud_organization_policy_sub_account_attachment":          resourceTencentCloudOrganizationPolicySubAccountAttachment(),
			"tencentcloud_organization_org_member_auth_identity_attachment":    resourceTencentCloudOrganizationOrgMemberAuthIdentityAttachment(),
			"tencentcloud_organization_org_member_policy_attachment":           resourceTencentCloudOrganizationOrgMemberPolicyAttachment(),
			"tencentcloud_dbbrain_sql_filter":                                  resourceTencentCloudDbbrainSqlFilter(),
			"tencentcloud_dbbrain_security_audit_log_export_task":              resourceTencentCloudDbbrainSecurityAuditLogExportTask(),
			"tencentcloud_dbbrain_db_diag_report_task":                         resourceTencentCloudDbbrainDbDiagReportTask(),
			"tencentcloud_dbbrain_modify_diag_db_instance_operation":           resourceTencentCloudDbbrainModifyDiagDbInstanceOperation(),
			"tencentcloud_dbbrain_tdsql_audit_log":                             resourceTencentCloudDbbrainTdsqlAuditLog(),
			"tencentcloud_rum_project":                                         resourceTencentCloudRumProject(),
			"tencentcloud_rum_taw_instance":                                    resourceTencentCloudRumTawInstance(),
			"tencentcloud_rum_whitelist":                                       resourceTencentCloudRumWhitelist(),
			"tencentcloud_rum_offline_log_config_attachment":                   resourceTencentCloudRumOfflineLogConfigAttachment(),
			"tencentcloud_rum_instance_status_config":                          resourceTencentCloudRumInstanceStatusConfig(),
			"tencentcloud_rum_project_status_config":                           resourceTencentCloudRumProjectStatusConfig(),
			"tencentcloud_rum_release_file":                                    resourceTencentCloudRumReleaseFile(),
			"tencentcloud_tdmq_rocketmq_cluster":                               resourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloud_tdmq_rocketmq_namespace":                             resourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloud_tdmq_rocketmq_role":                                  resourceTencentCloudTdmqRocketmqRole(),
			"tencentcloud_tdmq_rocketmq_topic":                                 resourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloud_tdmq_rocketmq_group":                                 resourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloud_tdmq_rocketmq_environment_role":                      resourceTencentCloudTdmqRocketmqEnvironmentRole(),
			"tencentcloud_tdmq_rocketmq_vip_instance":                          resourceTencentCloudTdmqRocketmqVipInstance(),
			"tencentcloud_trocket_rocketmq_instance":                           resourceTencentCloudTrocketRocketmqInstance(),
			"tencentcloud_trocket_rocketmq_topic":                              resourceTencentCloudTrocketRocketmqTopic(),
			"tencentcloud_trocket_rocketmq_consumer_group":                     resourceTencentCloudTrocketRocketmqConsumerGroup(),
			"tencentcloud_trocket_rocketmq_role":                               resourceTencentCloudTrocketRocketmqRole(),
			"tencentcloud_dts_sync_job":                                        resourceTencentCloudDtsSyncJob(),
			"tencentcloud_dts_sync_config":                                     resourceTencentCloudDtsSyncConfig(),
			"tencentcloud_dts_sync_check_job_operation":                        resourceTencentCloudDtsSyncCheckJobOperation(),
			"tencentcloud_dts_sync_job_resume_operation":                       resourceTencentCloudDtsSyncJobResumeOperation(),
			"tencentcloud_dts_sync_job_start_operation":                        resourceTencentCloudDtsSyncJobStartOperation(),
			"tencentcloud_dts_sync_job_stop_operation":                         resourceTencentCloudDtsSyncJobStopOperation(),
			"tencentcloud_dts_sync_job_resize_operation":                       resourceTencentCloudDtsSyncJobResizeOperation(),
			"tencentcloud_dts_sync_job_recover_operation":                      resourceTencentCloudDtsSyncJobRecoverOperation(),
			"tencentcloud_dts_sync_job_isolate_operation":                      resourceTencentCloudDtsSyncJobIsolateOperation(),
			"tencentcloud_dts_sync_job_continue_operation":                     resourceTencentCloudDtsSyncJobContinueOperation(),
			"tencentcloud_dts_sync_job_pause_operation":                        resourceTencentCloudDtsSyncJobPauseOperation(),
			"tencentcloud_dts_migrate_service":                                 resourceTencentCloudDtsMigrateService(),
			"tencentcloud_dts_migrate_job":                                     resourceTencentCloudDtsMigrateJob(),
			"tencentcloud_dts_migrate_job_config":                              resourceTencentCloudDtsMigrateJobConfig(),
			"tencentcloud_dts_migrate_job_start_operation":                     resourceTencentCloudDtsMigrateJobStartOperation(),
			"tencentcloud_dts_migrate_job_resume_operation":                    resourceTencentCloudDtsMigrateJobResumeOperation(),
			"tencentcloud_dts_compare_task_stop_operation":                     resourceTencentCloudDtsCompareTaskStopOperation(),
			"tencentcloud_dts_compare_task":                                    resourceTencentCloudDtsCompareTask(),
			"tencentcloud_cvm_hpc_cluster":                                     resourceTencentCloudCvmHpcCluster(),
			"tencentcloud_vpc_flow_log":                                        resourceTencentCloudVpcFlowLog(),
			"tencentcloud_vpc_end_point_service":                               resourceTencentCloudVpcEndPointService(),
			"tencentcloud_vpc_end_point":                                       resourceTencentCloudVpcEndPoint(),
			"tencentcloud_vpc_end_point_service_white_list":                    resourceTencentCloudVpcEndPointServiceWhiteList(),
			"tencentcloud_vpc_enable_end_point_connect":                        resourceTencentCloudVpcEnableEndPointConnect(),
			"tencentcloud_ci_bucket_attachment":                                resourceTencentCloudCiBucketAttachment(),
			"tencentcloud_tcmq_queue":                                          resourceTencentCloudTcmqQueue(),
			"tencentcloud_tcmq_topic":                                          resourceTencentCloudTcmqTopic(),
			"tencentcloud_tcmq_subscribe":                                      resourceTencentCloudTcmqSubscribe(),
			"tencentcloud_ci_bucket_pic_style":                                 resourceTencentCloudCiBucketPicStyle(),
			"tencentcloud_ci_hot_link":                                         resourceTencentCloudCiHotLink(),
			"tencentcloud_ci_media_snapshot_template":                          resourceTencentCloudCiMediaSnapshotTemplate(),
			"tencentcloud_ci_media_transcode_template":                         resourceTencentCloudCiMediaTranscodeTemplate(),
			"tencentcloud_ci_media_animation_template":                         resourceTencentCloudCiMediaAnimationTemplate(),
			"tencentcloud_ci_media_concat_template":                            resourceTencentCloudCiMediaConcatTemplate(),
			"tencentcloud_ci_media_video_process_template":                     resourceTencentCloudCiMediaVideoProcessTemplate(),
			"tencentcloud_ci_media_video_montage_template":                     resourceTencentCloudCiMediaVideoMontageTemplate(),
			"tencentcloud_ci_media_voice_separate_template":                    resourceTencentCloudCiMediaVoiceSeparateTemplate(),
			"tencentcloud_ci_media_super_resolution_template":                  resourceTencentCloudCiMediaSuperResolutionTemplate(),
			"tencentcloud_ci_media_pic_process_template":                       resourceTencentCloudCiMediaPicProcessTemplate(),
			"tencentcloud_ci_media_watermark_template":                         resourceTencentCloudCiMediaWatermarkTemplate(),
			"tencentcloud_ci_media_tts_template":                               resourceTencentCloudCiMediaTtsTemplate(),
			"tencentcloud_ci_media_transcode_pro_template":                     resourceTencentCloudCiMediaTranscodeProTemplate(),
			"tencentcloud_ci_media_smart_cover_template":                       resourceTencentCloudCiMediaSmartCoverTemplate(),
			"tencentcloud_ci_media_speech_recognition_template":                resourceTencentCloudCiMediaSpeechRecognitionTemplate(),
			"tencentcloud_ci_guetzli":                                          resourceTencentCloudCIGuetzli(),
			"tencentcloud_ci_original_image_protection":                        resourceTencentCloudCIOriginalImageProtection(),
			"tencentcloud_cynosdb_audit_log_file":                              resourceTencentCloudCynosdbAuditLogFile(),
			"tencentcloud_cynosdb_security_group":                              resourceTencentCloudCynosdbSecurityGroup(),
			"tencentcloud_dayu_ddos_ip_attachment_v2":                          resourceTencentCloudDayuDDosIpAttachmentV2(),
			"tencentcloud_antiddos_ddos_black_white_ip":                        resourceTencentCloudAntiddosDdosBlackWhiteIp(),
			"tencentcloud_antiddos_ddos_geo_ip_block_config":                   resourceTencentCloudAntiddosDdosGeoIpBlockConfig(),
			"tencentcloud_antiddos_ddos_speed_limit_config":                    resourceTencentCloudAntiddosDdosSpeedLimitConfig(),
			"tencentcloud_antiddos_default_alarm_threshold":                    resourceTencentCloudAntiddosDefaultAlarmThreshold(),
			"tencentcloud_antiddos_scheduling_domain_user_name":                resourceTencentCloudAntiddosSchedulingDomainUserName(),
			"tencentcloud_antiddos_ip_alarm_threshold_config":                  resourceTencentCloudAntiddosIpAlarmThresholdConfig(),
			"tencentcloud_tsf_microservice":                                    resourceTencentCloudTsfMicroservice(),
			"tencentcloud_tsf_application_config":                              resourceTencentCloudTsfApplicationConfig(),
			"tencentcloud_cvm_launch_template":                                 resourceTencentCloudCvmLaunchTemplate(),
			"tencentcloud_tsf_cluster":                                         resourceTencentCloudTsfCluster(),
			"tencentcloud_tsf_api_group":                                       resourceTencentCloudTsfApiGroup(),
			"tencentcloud_tsf_namespace":                                       resourceTencentCloudTsfNamespace(),
			"tencentcloud_tsf_path_rewrite":                                    resourceTencentCloudTsfPathRewrite(),
			"tencentcloud_tsf_unit_rule":                                       resourceTencentCloudTsfUnitRule(),
			"tencentcloud_tsf_task":                                            resourceTencentCloudTsfTask(),
			"tencentcloud_tsf_config_template":                                 resourceTencentCloudTsfConfigTemplate(),
			"tencentcloud_tsf_api_rate_limit_rule":                             resourceTencentCloudTsfApiRateLimitRule(),
			"tencentcloud_tsf_application_release_config":                      resourceTencentCloudTsfApplicationReleaseConfig(),
			"tencentcloud_tsf_lane":                                            resourceTencentCloudTsfLane(),
			"tencentcloud_tsf_lane_rule":                                       resourceTencentCloudTsfLaneRule(),
			"tencentcloud_tsf_group":                                           resourceTencentCloudTsfGroup(),
			"tencentcloud_tsf_repository":                                      resourceTencentCloudTsfRepository(),
			"tencentcloud_tsf_application":                                     resourceTencentCloudTsfApplication(),
			"tencentcloud_tsf_application_public_config_release":               resourceTencentCloudTsfApplicationPublicConfigRelease(),
			"tencentcloud_tsf_application_public_config":                       resourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloud_tsf_application_file_config_release":                 resourceTencentCloudTsfApplicationFileConfigRelease(),
			"tencentcloud_tsf_instances_attachment":                            resourceTencentCloudTsfInstancesAttachment(),
			"tencentcloud_tsf_bind_api_group":                                  resourceTencentCloudTsfBindApiGroup(),
			"tencentcloud_tsf_application_file_config":                         resourceTencentCloudTsfApplicationFileConfig(),
			"tencentcloud_tsf_enable_unit_rule":                                resourceTencentCloudTsfEnableUnitRule(),
			"tencentcloud_tsf_deploy_container_group":                          resourceTencentCloudTsfDeployContainerGroup(),
			"tencentcloud_tsf_deploy_vm_group":                                 resourceTencentCloudTsfDeployVmGroup(),
			"tencentcloud_tsf_release_api_group":                               resourceTencentCloudTsfReleaseApiGroup(),
			"tencentcloud_tsf_operate_container_group":                         resourceTencentCloudTsfOperateContainerGroup(),
			"tencentcloud_tsf_operate_group":                                   resourceTencentCloudTsfOperateGroup(),
			"tencentcloud_tsf_unit_namespace":                                  resourceTencentCloudTsfUnitNamespace(),
			"tencentcloud_mps_workflow":                                        resourceTencentCloudMpsWorkflow(),
			"tencentcloud_mps_enable_workflow_config":                          resourceTencentCloudMpsEnableWorkflowConfig(),
			"tencentcloud_mps_flow":                                            resourceTencentCloudMpsFlow(),
			"tencentcloud_mps_input":                                           resourceTencentCloudMpsInput(),
			"tencentcloud_mps_output":                                          resourceTencentCloudMpsOutput(),
			"tencentcloud_mps_content_review_template":                         resourceTencentCloudMpsContentReviewTemplate(),
			"tencentcloud_mps_start_flow_operation":                            resourceTencentCloudMpsStartFlowOperation(),
			"tencentcloud_mps_event":                                           resourceTencentCloudMpsEvent(),
			"tencentcloud_mps_execute_function_operation":                      resourceTencentCloudMpsExecuteFunctionOperation(),
			"tencentcloud_mps_manage_task_operation":                           resourceTencentCloudMpsManageTaskOperation(),
			"tencentcloud_mps_transcode_template":                              resourceTencentCloudMpsTranscodeTemplate(),
			"tencentcloud_mps_watermark_template":                              resourceTencentCloudMpsWatermarkTemplate(),
			"tencentcloud_mps_image_sprite_template":                           resourceTencentCloudMpsImageSpriteTemplate(),
			"tencentcloud_mps_snapshot_by_timeoffset_template":                 resourceTencentCloudMpsSnapshotByTimeoffsetTemplate(),
			"tencentcloud_mps_sample_snapshot_template":                        resourceTencentCloudMpsSampleSnapshotTemplate(),
			"tencentcloud_mps_animated_graphics_template":                      resourceTencentCloudMpsAnimatedGraphicsTemplate(),
			"tencentcloud_mps_ai_recognition_template":                         resourceTencentCloudMpsAiRecognitionTemplate(),
			"tencentcloud_mps_ai_analysis_template":                            resourceTencentCloudMpsAiAnalysisTemplate(),
			"tencentcloud_mps_adaptive_dynamic_streaming_template":             resourceTencentCloudMpsAdaptiveDynamicStreamingTemplate(),
			"tencentcloud_mps_person_sample":                                   resourceTencentCloudMpsPersonSample(),
			"tencentcloud_mps_withdraws_watermark_operation":                   resourceTencentCloudMpsWithdrawsWatermarkOperation(),
			"tencentcloud_mps_process_live_stream_operation":                   resourceTencentCloudMpsProcessLiveStreamOperation(),
			"tencentcloud_mps_edit_media_operation":                            resourceTencentCloudMpsEditMediaOperation(),
			"tencentcloud_mps_word_sample":                                     resourceTencentCloudMpsWordSample(),
			"tencentcloud_mps_schedule":                                        resourceTencentCloudMpsSchedule(),
			"tencentcloud_mps_enable_schedule_config":                          resourceTencentCloudMpsEnableScheduleConfig(),
			"tencentcloud_mps_process_media_operation":                         resourceTencentCloudMpsProcessMediaOperation(),
			"tencentcloud_cbs_disk_backup":                                     resourceTencentCloudCbsDiskBackup(),
			"tencentcloud_cbs_snapshot_share_permission":                       resourceTencentCloudCbsSnapshotSharePermission(),
			"tencentcloud_cbs_disk_backup_rollback_operation":                  resourceTencentCloudCbsDiskBackupRollbackOperation(),
			"tencentcloud_chdfs_access_group":                                  resourceTencentCloudChdfsAccessGroup(),
			"tencentcloud_chdfs_access_rule":                                   resourceTencentCloudChdfsAccessRule(),
			"tencentcloud_chdfs_file_system":                                   resourceTencentCloudChdfsFileSystem(),
			"tencentcloud_chdfs_life_cycle_rule":                               resourceTencentCloudChdfsLifeCycleRule(),
			"tencentcloud_chdfs_mount_point":                                   resourceTencentCloudChdfsMountPoint(),
			"tencentcloud_chdfs_mount_point_attachment":                        resourceTencentCloudChdfsMountPointAttachment(),
			"tencentcloud_mdl_stream_live_input":                               resourceTencentCloudMdlStreamLiveInput(),
			"tencentcloud_lighthouse_blueprint":                                resourceTencentCloudLighthouseBlueprint(),
			"tencentcloud_cvm_launch_template_version":                         resourceTencentCloudCvmLaunchTemplateVersion(),
			"tencentcloud_apm_instance":                                        resourceTencentCloudApmInstance(),
			"tencentcloud_cvm_launch_template_default_version":                 resourceTencentCloudCvmLaunchTemplateDefaultVersion(),
			"tencentcloud_lighthouse_firewall_rule":                            resourceTencentCloudLighthouseFirewallRule(),
			"tencentcloud_cvm_security_group_attachment":                       resourceTencentCloudCvmSecurityGroupAttachment(),
			"tencentcloud_cvm_reboot_instance":                                 resourceTencentCloudCvmRebootInstance(),
			"tencentcloud_cvm_chc_config":                                      resourceTencentCloudCvmChcConfig(),
			"tencentcloud_cvm_sync_image":                                      resourceTencentCloudCvmSyncImage(),
			"tencentcloud_cvm_renew_instance":                                  resourceTencentCloudCvmRenewInstance(),
			"tencentcloud_cvm_export_images":                                   resourceTencentCloudCvmExportImages(),
			"tencentcloud_cvm_image_share_permission":                          resourceTencentCloudCvmImageSharePermission(),
			"tencentcloud_cvm_import_image":                                    resourceTencentCloudCvmImportImage(),
			"tencentcloud_cvm_renew_host":                                      resourceTencentCloudCvmRenewHost(),
			"tencentcloud_cvm_program_fpga_image":                              resourceTencentCloudCvmProgramFpgaImage(),
			"tencentcloud_cvm_modify_instance_disk_type":                       resourceTencentCloudCvmModifyInstanceDiskType(),
			"tencentcloud_lighthouse_disk_backup":                              resourceTencentCloudLighthouseDiskBackup(),
			"tencentcloud_lighthouse_apply_disk_backup":                        resourceTencentCloudLighthouseApplyDiskBackup(),
			"tencentcloud_lighthouse_disk_attachment":                          resourceTencentCloudLighthouseDiskAttachment(),
			"tencentcloud_lighthouse_key_pair":                                 resourceTencentCloudLighthouseKeyPair(),
			"tencentcloud_lighthouse_snapshot":                                 resourceTencentCloudLighthouseSnapshot(),
			"tencentcloud_lighthouse_apply_instance_snapshot":                  resourceTencentCloudLighthouseApplyInstanceSnapshot(),
			"tencentcloud_lighthouse_start_instance":                           resourceTencentCloudLighthouseStartInstance(),
			"tencentcloud_lighthouse_stop_instance":                            resourceTencentCloudLighthouseStopInstance(),
			"tencentcloud_lighthouse_reboot_instance":                          resourceTencentCloudLighthouseRebootInstance(),
			"tencentcloud_lighthouse_key_pair_attachment":                      resourceTencentCloudLighthouseKeyPairAttachment(),
			"tencentcloud_lighthouse_disk":                                     resourceTencentCloudLighthouseDisk(),
			"tencentcloud_lighthouse_renew_disk":                               resourceTencentCloudLighthouseRenewDisk(),
			"tencentcloud_lighthouse_renew_instance":                           resourceTencentCloudLighthouseRenewInstance(),
			"tencentcloud_clickhouse_backup":                                   resourceTencentCloudClickhouseBackup(),
			"tencentcloud_clickhouse_backup_strategy":                          resourceTencentCloudClickhouseBackupStrategy(),
			"tencentcloud_clickhouse_recover_backup_job":                       resourceTencentCloudClickhouseRecoverBackupJob(),
			"tencentcloud_clickhouse_delete_backup_data":                       resourceTencentCloudClickhouseDeleteBackupData(),
			"tencentcloud_clickhouse_account":                                  resourceTencentCloudClickhouseAccount(),
			"tencentcloud_clickhouse_account_permission":                       resourceTencentCloudClickhouseAccountPermission(),
			"tencentcloud_api_gateway_api_doc":                                 resourceTencentCloudAPIGatewayAPIDoc(),
			"tencentcloud_api_gateway_api_app":                                 resourceTencentCloudAPIGatewayAPIApp(),
			"tencentcloud_api_gateway_update_api_app_key":                      resourceTencentCloudApiGatewayUpdateApiAppKey(),
			"tencentcloud_api_gateway_import_open_api":                         resourceTencentCloudApiGatewayImportOpenApi(),
			"tencentcloud_tse_instance":                                        resourceTencentCloudTseInstance(),
			"tencentcloud_tse_cngw_gateway":                                    resourceTencentCloudTseCngwGateway(),
			"tencentcloud_tse_cngw_group":                                      resourceTencentCloudTseCngwGroup(),
			"tencentcloud_tse_cngw_service":                                    resourceTencentCloudTseCngwService(),
			"tencentcloud_tse_cngw_service_rate_limit":                         resourceTencentCloudTseCngwServiceRateLimit(),
			"tencentcloud_tse_cngw_route":                                      resourceTencentCloudTseCngwRoute(),
			"tencentcloud_tse_cngw_route_rate_limit":                           resourceTencentCloudTseCngwRouteRateLimit(),
			"tencentcloud_tse_cngw_canary_rule":                                resourceTencentCloudTseCngwCanaryRule(),
			"tencentcloud_tse_cngw_certificate":                                resourceTencentCloudTseCngwCertificate(),
			"tencentcloud_clickhouse_instance":                                 resourceTencentCloudClickhouseInstance(),
			"tencentcloud_cls_kafka_recharge":                                  resourceTencentCloudClsKafkaRecharge(),
			"tencentcloud_cls_scheduled_sql":                                   resourceTencentCloudClsScheduledSql(),
			"tencentcloud_eb_event_transform":                                  resourceTencentCloudEbEventTransform(),
			"tencentcloud_eb_event_bus":                                        resourceTencentCloudEbEventBus(),
			"tencentcloud_eb_event_rule":                                       resourceTencentCloudEbEventRule(),
			"tencentcloud_eb_event_target":                                     resourceTencentCloudEbEventTarget(),
			"tencentcloud_eb_put_events":                                       resourceTencentCloudEbPutEvents(),
			"tencentcloud_eb_event_connector":                                  resourceTencentCloudEbEventConnector(),
			"tencentcloud_dlc_user":                                            resourceTencentCloudDlcUser(),
			"tencentcloud_dlc_work_group":                                      resourceTencentCloudDlcWorkGroup(),
			"tencentcloud_dlc_data_engine":                                     resourceTencentCloudDlcDataEngine(),
			"tencentcloud_dlc_suspend_resume_data_engine":                      resourceTencentCloudDlcSuspendResumeDataEngine(),
			"tencentcloud_dlc_rollback_data_engine_image_operation":            resourceTencentCloudDlcRollbackDataEngineImageOperation(),
			"tencentcloud_dlc_add_users_to_work_group_attachment":              resourceTencentCloudDlcAddUsersToWorkGroupAttachment(),
			"tencentcloud_dlc_store_location_config":                           resourceTencentCloudDlcStoreLocationConfig(),
			"tencentcloud_dlc_modify_data_engine_description_operation":        resourceTencentCloudDlcModifyDataEngineDescriptionOperation(),
			"tencentcloud_dlc_modify_user_typ_operation":                       resourceTencentCloudDlcModifyUserTypOperation(),
			"tencentcloud_dlc_renew_data_engine_operation":                     resourceTencentCloudDlcRenewDataEngineOperation(),
			"tencentcloud_dlc_restart_data_engine_operation":                   resourceTencentCloudDlcRestartDataEngineOperation(),
			"tencentcloud_dlc_attach_user_policy_operation":                    resourceTencentCloudDlcAttachUserPolicyOperation(),
			"tencentcloud_dlc_detach_user_policy_operation":                    resourceTencentCloudDlcDetachUserPolicyOperation(),
			"tencentcloud_dlc_attach_work_group_policy_operation":              resourceTencentCloudDlcAttachWorkGroupPolicyOperation(),
			"tencentcloud_dlc_detach_work_group_policy_operation":              resourceTencentCloudDlcDetachWorkGroupPolicyOperation(),
			"tencentcloud_dlc_switch_data_engine_image_operation":              resourceTencentCloudDlcSwitchDataEngineImageOperation(),
			"tencentcloud_dlc_update_data_engine_config_operation":             resourceTencentCloudDlcUpdateDataEngineConfigOperation(),
			"tencentcloud_dlc_upgrade_data_engine_image_operation":             resourceTencentCloudDlcUpgradeDataEngineImageOperation(),
			"tencentcloud_dlc_bind_work_groups_to_user_attachment":             resourceTencentCloudDlcBindWorkGroupsToUserAttachment(),
			"tencentcloud_dlc_update_row_filter_operation":                     resourceTencentCloudDlcUpdateRowFilterOperation(),
			"tencentcloud_dlc_user_data_engine_config":                         resourceTencentCloudDlcUserDataEngineConfig(),
			"tencentcloud_waf_custom_rule":                                     resourceTencentCloudWafCustomRule(),
			"tencentcloud_waf_custom_white_rule":                               resourceTencentCloudWafCustomWhiteRule(),
			"tencentcloud_waf_clb_domain":                                      resourceTencentCloudWafClbDomain(),
			"tencentcloud_waf_saas_domain":                                     resourceTencentCloudWafSaasDomain(),
			"tencentcloud_waf_clb_instance":                                    resourceTencentCloudWafClbInstance(),
			"tencentcloud_waf_saas_instance":                                   resourceTencentCloudWafSaasInstance(),
			"tencentcloud_waf_anti_fake":                                       resourceTencentCloudWafAntiFake(),
			"tencentcloud_waf_anti_info_leak":                                  resourceTencentCloudWafAntiInfoLeak(),
			"tencentcloud_waf_auto_deny_rules":                                 resourceTencentCloudWafAutoDenyRules(),
			"tencentcloud_waf_module_status":                                   resourceTencentCloudWafModuleStatus(),
			"tencentcloud_waf_protection_mode":                                 resourceTencentCloudWafProtectionMode(),
			"tencentcloud_waf_web_shell":                                       resourceTencentCloudWafWebShell(),
			"tencentcloud_waf_cc":                                              resourceTencentCloudWafCc(),
			"tencentcloud_waf_cc_auto_status":                                  resourceTencentCloudWafCcAutoStatus(),
			"tencentcloud_waf_cc_session":                                      resourceTencentCloudWafCcSession(),
			"tencentcloud_waf_ip_access_control":                               resourceTencentCloudWafIpAccessControl(),
			"tencentcloud_waf_modify_access_period":                            resourceTencentCloudWafModifyAccessPeriod(),
			"tencentcloud_wedata_rule_template":                                resourceTencentCloudWedataRuleTemplate(),
			"tencentcloud_wedata_datasource":                                   resourceTencentCloudWedataDatasource(),
			"tencentcloud_wedata_function":                                     resourceTencentCloudWedataFunction(),
			"tencentcloud_wedata_resource":                                     resourceTencentCloudWedataResource(),
			"tencentcloud_wedata_script":                                       resourceTencentCloudWedataScript(),
			"tencentcloud_wedata_dq_rule":                                      resourceTencentCloudWedataDqRule(),
			"tencentcloud_wedata_baseline":                                     resourceTencentCloudWedataBaseline(),
			"tencentcloud_wedata_integration_offline_task":                     resourceTencentCloudWedataIntegrationOfflineTask(),
			"tencentcloud_wedata_integration_realtime_task":                    resourceTencentCloudWedataIntegrationRealtimeTask(),
			"tencentcloud_wedata_integration_task_node":                        resourceTencentCloudWedataIntegrationTaskNode(),
			"tencentcloud_cfw_address_template":                                resourceTencentCloudCfwAddressTemplate(),
			"tencentcloud_cfw_block_ignore":                                    resourceTencentCloudCfwBlockIgnore(),
			"tencentcloud_cfw_edge_policy":                                     resourceTencentCloudCfwEdgePolicy(),
			"tencentcloud_cfw_nat_instance":                                    resourceTencentCloudCfwNatInstance(),
			"tencentcloud_cfw_nat_policy":                                      resourceTencentCloudCfwNatPolicy(),
			"tencentcloud_cfw_vpc_instance":                                    resourceTencentCloudCfwVpcInstance(),
			"tencentcloud_cfw_vpc_policy":                                      resourceTencentCloudCfwVpcPolicy(),
			"tencentcloud_cfw_sync_asset":                                      resourceTencentCloudCfwSyncAsset(),
			"tencentcloud_cfw_sync_route":                                      resourceTencentCloudCfwSyncRoute(),
			"tencentcloud_cfw_nat_firewall_switch":                             resourceTencentCloudCfwNatFirewallSwitch(),
			"tencentcloud_cfw_vpc_firewall_switch":                             resourceTencentCloudCfwVpcFirewallSwitch(),
			"tencentcloud_cfw_edge_firewall_switch":                            resourceTencentCloudCfwEdgeFirewallSwitch(),
			"tencentcloud_dasb_acl":                                            resourceTencentCloudDasbAcl(),
			"tencentcloud_dasb_cmd_template":                                   resourceTencentCloudDasbCmdTemplate(),
			"tencentcloud_dasb_device_group":                                   resourceTencentCloudDasbDeviceGroup(),
			"tencentcloud_dasb_user":                                           resourceTencentCloudDasbUser(),
			"tencentcloud_dasb_device_account":                                 resourceTencentCloudDasbDeviceAccount(),
			"tencentcloud_dasb_device_group_members":                           resourceTencentCloudDasbDeviceGroupMembers(),
			"tencentcloud_dasb_user_group_members":                             resourceTencentCloudDasbUserGroupMembers(),
			"tencentcloud_dasb_bind_device_resource":                           resourceTencentCloudDasbBindDeviceResource(),
			"tencentcloud_dasb_resource":                                       resourceTencentCloudDasbResource(),
			"tencentcloud_dasb_device":                                         resourceTencentCloudDasbDevice(),
			"tencentcloud_dasb_user_group":                                     resourceTencentCloudDasbUserGroup(),
			"tencentcloud_dasb_reset_user":                                     resourceTencentCloudDasbResetUser(),
			"tencentcloud_dasb_bind_device_account_private_key":                resourceTencentCloudDasbBindDeviceAccountPrivateKey(),
			"tencentcloud_dasb_bind_device_account_password":                   resourceTencentCloudDasbBindDeviceAccountPassword(),
			"tencentcloud_ssl_check_certificate_chain_operation":               resourceTencentCloudSslCheckCertificateChainOperation(),
			"tencentcloud_ssl_complete_certificate_operation":                  resourceTencentCloudSslCompleteCertificateOperation(),
			"tencentcloud_ssl_deploy_certificate_instance_operation":           resourceTencentCloudSslDeployCertificateInstanceOperation(),
			"tencentcloud_ssl_deploy_certificate_record_retry_operation":       resourceTencentCloudSslDeployCertificateRecordRetryOperation(),
			"tencentcloud_ssl_deploy_certificate_record_rollback_operation":    resourceTencentCloudSslDeployCertificateRecordRollbackOperation(),
			"tencentcloud_ssl_download_certificate_operation":                  resourceTencentCloudSslDownloadCertificateOperation(),
			"tencentcloud_cwp_license_order":                                   resourceTencentCloudCwpLicenseOrder(),
			"tencentcloud_cwp_license_bind_attachment":                         resourceTencentCloudCwpLicenseBindAttachment(),
			"tencentcloud_ssl_replace_certificate_operation":                   resourceTencentCloudSslReplaceCertificateOperation(),
			"tencentcloud_ssl_revoke_certificate_operation":                    resourceTencentCloudSslRevokeCertificateOperation(),
			"tencentcloud_ssl_update_certificate_instance_operation":           resourceTencentCloudSslUpdateCertificateInstanceOperation(),
			"tencentcloud_ssl_update_certificate_record_retry_operation":       resourceTencentCloudSslUpdateCertificateRecordRetryOperation(),
			"tencentcloud_ssl_update_certificate_record_rollback_operation":    resourceTencentCloudSslUpdateCertificateRecordRollbackOperation(),
			"tencentcloud_ssl_upload_revoke_letter_operation":                  resourceTencentCloudSslUploadRevokeLetterOperation(),
			"tencentcloud_bi_project":                                          resourceTencentCloudBiProject(),
			"tencentcloud_bi_user_role":                                        resourceTencentCloudBiUserRole(),
			"tencentcloud_bi_project_user_role":                                resourceTencentCloudBiProjectUserRole(),
			"tencentcloud_bi_datasource":                                       resourceTencentCloudBiDatasource(),
			"tencentcloud_bi_datasource_cloud":                                 resourceTencentCloudBiDatasourceCloud(),
			"tencentcloud_bi_embed_token_apply":                                resourceTencentCloudBiEmbedTokenApply(),
			"tencentcloud_bi_embed_interval_apply":                             resourceTencentCloudBiEmbedIntervalApply(),
			"tencentcloud_cdwpg_instance":                                      resourceTencentCloudCdwpgInstance(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//var getProviderConfig = func(str string, key string) string {
	//	if str == "" {
	//		value, err := getConfigFromProfile(d, key)
	//		if err == nil && value != nil {
	//			str = value.(string)
	//		}
	//	}
	//
	//	return str
	//}

	var getProviderConfig = func(key string) string {
		var str string
		value, err := getConfigFromProfile(d, key)
		if err == nil && value != nil {
			str = value.(string)
		}

		return str
	}

	var (
		secretId      string
		secretKey     string
		securityToken string
		region        string
		protocol      string
		domain        string
	)

	if v, ok := d.GetOk("secret_id"); ok {
		secretId = v.(string)
	} else {
		secretId = getProviderConfig("secretId")
	}

	if v, ok := d.GetOk("secret_key"); ok {
		secretKey = v.(string)
	} else {
		secretKey = getProviderConfig("secretKey")
	}

	if v, ok := d.GetOk("security_token"); ok {
		securityToken = v.(string)
	}

	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	} else {
		region = getProviderConfig("region")
		if region == "" {
			region = DEFAULT_REGION
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		protocol = v.(string)
	}

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	// standard client
	var tcClient TencentCloudClient
	tcClient.apiV3Conn = &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		Region:   region,
		Protocol: protocol,
		Domain:   domain,
	}

	var (
		assumeRoleArn             string
		assumeRoleSessionName     string
		assumeRoleSessionDuration int
		assumeRolePolicy          string
	)

	// get assume role from credential
	if providerConfig["role-arn"] != nil {
		assumeRoleArn = providerConfig["role-arn"].(string)
	}

	if providerConfig["role-session-name"] != nil {
		assumeRoleSessionName = providerConfig["role-session-name"].(string)
	}

	if assumeRoleArn != "" && assumeRoleSessionName != "" {
		assumeRoleSessionDuration = 7200
		assumeRolePolicy = ""

		_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
	}

	// get assume role from env
	envRoleArn := os.Getenv(PROVIDER_ASSUME_ROLE_ARN)
	envSessionName := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_NAME)
	if envRoleArn != "" && envSessionName != "" {
		if envSessionDuration := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); envSessionDuration != "" {
			var err error
			assumeRoleSessionDuration, err = strconv.Atoi(envSessionDuration)
			if err != nil {
				return nil, err
			}
		}

		if assumeRoleSessionDuration == 0 {
			assumeRoleSessionDuration = 7200
		}

		_ = genClientWithSTS(&tcClient, envRoleArn, envSessionName, assumeRoleSessionDuration, "")
	}

	// get assume role from tf
	if v, ok := d.GetOk("assume_role"); ok {
		assumeRoleList := v.(*schema.Set).List()
		if len(assumeRoleList) == 1 {
			assumeRole := assumeRoleList[0].(map[string]interface{})
			assumeRoleArn = assumeRole["role_arn"].(string)
			assumeRoleSessionName = assumeRole["session_name"].(string)
			assumeRoleSessionDuration = assumeRole["session_duration"].(int)
			assumeRolePolicy = assumeRole["policy"].(string)

			_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
		}
	}

	if secretId == "" || secretKey == "" {
		return nil, fmt.Errorf("Please set your `secret_id` and `secret_key`.")
	}

	return &tcClient, nil
}

func genClientWithSTS(tcClient *TencentCloudClient, assumeRoleArn, assumeRoleSessionName string, assumeRoleSessionDuration int, assumeRolePolicy string) error {
	// applying STS credentials
	request := sts.NewAssumeRoleRequest()
	request.RoleArn = helper.String(assumeRoleArn)
	request.RoleSessionName = helper.String(assumeRoleSessionName)
	request.DurationSeconds = helper.IntUint64(assumeRoleSessionDuration)
	if assumeRolePolicy != "" {
		request.Policy = helper.String(url.QueryEscape(assumeRolePolicy))
	}

	ratelimit.Check(request.GetAction())
	response, err := tcClient.apiV3Conn.UseStsClient().AssumeRole(request)
	if err != nil {
		return err
	}

	// using STS credentials
	tcClient.apiV3Conn.Credential = common.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)

	return nil
}

var providerConfig map[string]interface{}

func getConfigFromProfile(d *schema.ResourceData, ProfileKey string) (interface{}, error) {
	if providerConfig == nil {
		var (
			profile              string
			sharedCredentialsDir string
			credentialPath       string
			configurePath        string
		)

		if v, ok := d.GetOk("profile"); ok {
			profile = v.(string)
		} else {
			profile = DEFAULT_PROFILE
		}

		if v, ok := d.GetOk("shared_credentials_dir"); ok {
			sharedCredentialsDir = v.(string)
		}

		tmpSharedCredentialsDir, err := homedir.Expand(sharedCredentialsDir)
		if err != nil {
			return nil, err
		}

		if tmpSharedCredentialsDir == "" {
			credentialPath = fmt.Sprintf("%s/.tccli/%s.credential", os.Getenv("HOME"), profile)
			configurePath = fmt.Sprintf("%s/.tccli/%s.configure", os.Getenv("HOME"), profile)
			if runtime.GOOS == "windows" {
				credentialPath = fmt.Sprintf("%s/.tccli/%s.credential", os.Getenv("USERPROFILE"), profile)
				configurePath = fmt.Sprintf("%s/.tccli/%s.configure", os.Getenv("USERPROFILE"), profile)
			}
		} else {
			credentialPath = fmt.Sprintf("%s/%s.credential", tmpSharedCredentialsDir, profile)
			configurePath = fmt.Sprintf("%s/%s.configure", tmpSharedCredentialsDir, profile)
		}

		providerConfig = make(map[string]interface{})
		_, err = os.Stat(credentialPath)
		if !os.IsNotExist(err) {
			data, err := ioutil.ReadFile(credentialPath)
			if err != nil {
				return nil, err
			}

			config := map[string]interface{}{}
			err = json.Unmarshal(data, &config)
			if err != nil {
				return nil, err
			}

			for k, v := range config {
				providerConfig[k] = strings.TrimSpace(v.(string))
			}
		}

		_, err = os.Stat(configurePath)
		if !os.IsNotExist(err) {
			data, err := ioutil.ReadFile(configurePath)
			if err != nil {
				return nil, err
			}

			config := map[string]interface{}{}
			err = json.Unmarshal(data, &config)
			if err != nil {
				return nil, err
			}

		outerLoop:
			for k, v := range config {
				if k == "_sys_param" {
					tmpMap := v.(map[string]interface{})
					for tmpK, tmpV := range tmpMap {
						if tmpK == "region" {
							providerConfig[tmpK] = strings.TrimSpace(tmpV.(string))
							break outerLoop
						}
					}
				}
			}
		}
	}

	return providerConfig[ProfileKey], nil
}
