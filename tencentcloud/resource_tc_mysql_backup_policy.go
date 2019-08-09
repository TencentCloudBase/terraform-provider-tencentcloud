/*
Provides a mysql policy resource to create a backup policy.

Example Usage

```hcl
resource "tencentcloud_mysql_backup_policy" "default" {
  mysql_id = "cdb-dnqksd9f"
  retention_period = 7
  backup_model = "logical"
  backup_time ="02:00–06:00"
}
```
*/
package tencentcloud

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTencentCloudMysqlBackupPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMysqlBackupPolicyCreate,
		Read:   resourceTencentCloudMysqlBackupPolicyRead,
		Update: resourceTencentCloudMysqlBackupPolicyUpdate,
		Delete: resourceTencentCloudMysqlBackupPolicyDelete,

		Schema: map[string]*schema.Schema{
			"mysql_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Instance ID to which policies will be applied.",
			},
			"retention_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(7, 732),
				Optional:     true,
				Default:      7,
				Description:  "Instance backup retention days. Valid values: [7-730]. And default value is 7.",
			},
			"backup_model": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      MYSQL_ALLOW_BACKUP_MODEL[0],
				ValidateFunc: validateAllowedStringValue(MYSQL_ALLOW_BACKUP_MODEL),
				Description:  "Backup method. Supported values include: physical - physical backup, and logical - logical backup.",
			},
			"backup_time": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      MYSQL_ALLOW_BACKUP_TIME[0],
				ValidateFunc: validateAllowedStringValue(MYSQL_ALLOW_BACKUP_TIME),
				Description:  `Instance backup time, in the format of "HH:mm-HH:mm". Time setting interval is four hours. Default to "02:00-06:00". The following value can be supported: 02:00\-06:00, 06:00\-10:00, 10:00\-14:00, 14:00\-18:00, 18:00\-22:00, and 22:00\-02:00.`,
			},

			// Computed values
			"binlog_period": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Retention period for binlog in days.",
			},
		},
	}
}

func resourceTencentCloudMysqlBackupPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_backup_policy.create")()

	d.SetId(d.Get("mysql_id").(string))

	return resourceTencentCloudMysqlBackupPolicyUpdate(d, meta)
}

func resourceTencentCloudMysqlBackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("tencentcloud_mysql_backup_policy.read")()

	logId := getLogId(nil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}
	desResponse, err := mysqlService.DescribeBackupConfigByMysqlId(ctx, d.Id())

	if err != nil {
		if mysqlService.NotFoundMysqlInstance(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[API]Describe mysql backup policy fail,reason:%s", err.Error())
	}

	d.Set("mysql_id", d.Id())
	d.Set("retention_period", int(*desResponse.Response.BackupExpireDays))
	d.Set("backup_model", *desResponse.Response.BackupMethod)

	var buf bytes.Buffer

	if *desResponse.Response.StartTimeMin < 10 {
		buf.WriteString("0")
	}
	buf.WriteString(fmt.Sprintf("%d:00-", *desResponse.Response.StartTimeMin))

	if *desResponse.Response.StartTimeMax < 10 {
		buf.WriteString("0")
	}
	buf.WriteString(fmt.Sprintf("%d:00", *desResponse.Response.StartTimeMax))

	d.Set("backup_time", buf.String())
	d.Set("binlog_period", int(*desResponse.Response.BinlogExpireDays))

	return nil
}

func resourceTencentCloudMysqlBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_backup_policy.update")()

	logId := getLogId(nil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		isUpdate = false

		mysqlId         = d.Get("mysql_id").(string)
		retentionPeriod = int64(d.Get("retention_period").(int))
		backupModel     = d.Get("backup_model").(string)
		backupTime      = d.Get("backup_time").(string)
	)

	if d.HasChange("retention_period") || d.HasChange("backup_model") || d.HasChange("backup_time") {
		isUpdate = true
	}

	if isUpdate {
		err := mysqlService.ModifyBackupConfigByMysqlId(ctx, mysqlId, retentionPeriod, backupModel, backupTime)
		if err != nil {
			return err
		}
	}
	return resourceTencentCloudMysqlBackupPolicyRead(d, meta)
}

//set all config to default
func resourceTencentCloudMysqlBackupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_backup_policy.delete")()

	logId := getLogId(nil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		retentionPeriod int64 = 7
		backupModel           = MYSQL_ALLOW_BACKUP_MODEL[0]
		backupTime            = MYSQL_ALLOW_BACKUP_TIME[0]
	)
	err := mysqlService.ModifyBackupConfigByMysqlId(ctx, d.Id(), retentionPeriod, backupModel, backupTime)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
