/*
Provides a resource to create a mysql instance_encryption_operation

Example Usage

```hcl
data "tencentcloud_availability_zones_by_product" "zones" {
  product = "cdb"
}

resource "tencentcloud_vpc" "vpc" {
  name       = "vpc-mysql"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloud_subnet" "subnet" {
  availability_zone = data.tencentcloud_availability_zones_by_product.zones.zones.0.name
  name              = "subnet-mysql"
  vpc_id            = tencentcloud_vpc.vpc.id
  cidr_block        = "10.0.0.0/16"
  is_multicast      = false
}

resource "tencentcloud_security_group" "security_group" {
  name        = "sg-mysql"
  description = "mysql test"
}

resource "tencentcloud_mysql_instance" "example" {
  internet_service  = 1
  engine_version    = "5.7"
  charge_type       = "POSTPAID"
  root_password     = "PassWord123"
  slave_deploy_mode = 0
  availability_zone = data.tencentcloud_availability_zones_by_product.zones.zones.0.name
  slave_sync_mode   = 1
  instance_name     = "tf-example-mysql"
  mem_size          = 4000
  volume_size       = 200
  vpc_id            = tencentcloud_vpc.vpc.id
  subnet_id         = tencentcloud_subnet.subnet.id
  intranet_port     = 3306
  security_groups   = [tencentcloud_security_group.security_group.id]

  tags = {
    name = "test"
  }

  parameters = {
    character_set_server = "utf8"
    max_connections      = "1000"
  }
}

resource "tencentcloud_mysql_instance_encryption_operation" "example" {
  instance_id = tencentcloud_mysql_instance.example.id
  key_id      = "KMS-CDB"
  key_region  = "ap-guangzhou"
}
```

*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	mysql "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudMysqlInstanceEncryptionOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMysqlInstanceEncryptionOperationCreate,
		Read:   resourceTencentCloudMysqlInstanceEncryptionOperationRead,
		Delete: resourceTencentCloudMysqlInstanceEncryptionOperationDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "TencentDB instance ID.",
			},

			"key_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.",
			},

			"key_region": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.",
			},
		},
	}
}

func resourceTencentCloudMysqlInstanceEncryptionOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_instance_encryption_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request    = mysql.NewOpenDBInstanceEncryptionRequest()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("key_id"); ok {
		request.KeyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("key_region"); ok {
		request.KeyRegion = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMysqlClient().OpenDBInstanceEncryption(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate mysql instanceEncryptionOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	service := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instanceInfo, err := service.DescribeMysqlInstanceInfoById(ctx, instanceId)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if *instanceInfo.Encryption == "YES" {
			return nil
		}
		if *instanceInfo.Encryption == "NO" {
			return resource.RetryableError(fmt.Errorf("%s instanceEncryption status is %s", instanceId, *instanceInfo.Encryption))
		}
		err = fmt.Errorf("%s operate mysql instanceEncryption status is %s,we won't wait for it finish", instanceId, *instanceInfo.Encryption)
		return resource.NonRetryableError(err)
	})

	if err != nil {
		log.Printf("[CRITAL]%s instanceEncryption fail, reason:%s\n ", logId, err.Error())
		return err
	}

	return resourceTencentCloudMysqlInstanceEncryptionOperationRead(d, meta)
}

func resourceTencentCloudMysqlInstanceEncryptionOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_instance_encryption_operation.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MysqlService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	result, err := service.DescribeMysqlInstanceInfoById(ctx, instanceId)
	if err != nil {
		return err
	}

	if result == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `MysqlInstanceInfo` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)

	if result.KeyId != nil {
		_ = d.Set("key_id", result.KeyId)
	}

	if result.KeyRegion != nil {
		_ = d.Set("key_region", result.KeyRegion)
	}

	return nil
}

func resourceTencentCloudMysqlInstanceEncryptionOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_instance_encryption_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
