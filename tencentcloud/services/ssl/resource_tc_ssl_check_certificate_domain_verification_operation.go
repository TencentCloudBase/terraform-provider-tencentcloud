// Code generated by iacg; DO NOT EDIT.
package ssl

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sslv20191205 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudSslCheckCertificateDomainVerificationOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudSslCheckCertificateDomainVerificationOperationCreate,
		Read:   resourceTencentCloudSslCheckCertificateDomainVerificationOperationRead,
		Delete: resourceTencentCloudSslCheckCertificateDomainVerificationOperationDelete,
		Schema: map[string]*schema.Schema{
			"certificate_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The certificate ID.",
			},

			"verification_results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Domain name verification results.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name.",
						},
						"verify_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain Verify Type.",
						},
						"local_check": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Local inspection results.",
						},
						"ca_check": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CA inspection results.",
						},
						"local_check_fail_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Check the reason for the failure.",
						},
						"check_value": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Detected values.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"frequently": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether frequent requests.",
						},
						"issued": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether issued.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudSslCheckCertificateDomainVerificationOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_ssl_check_certificate_domain_verification_operation.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	var (
		certificateId string
	)
	var (
		request  = sslv20191205.NewCheckCertificateDomainVerificationRequest()
		response = sslv20191205.NewCheckCertificateDomainVerificationResponse()
	)

	if v, ok := d.GetOk("certificate_id"); ok {
		certificateId = v.(string)
	}

	if v, ok := d.GetOk("certificate_id"); ok {
		request.CertificateId = helper.String(v.(string))
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseSslV20191205Client().CheckCertificateDomainVerificationWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create ssl check certificate domain verification operation failed, reason:%+v", logId, err)
		return err
	}

	if err := resourceTencentCloudSslCheckCertificateDomainVerificationOperationCreatePreHandleResponse0(ctx, response); err != nil {
		return err
	}

	_ = response

	d.SetId(certificateId)

	return resourceTencentCloudSslCheckCertificateDomainVerificationOperationRead(d, meta)
}

func resourceTencentCloudSslCheckCertificateDomainVerificationOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_ssl_check_certificate_domain_verification_operation.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudSslCheckCertificateDomainVerificationOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_ssl_check_certificate_domain_verification_operation.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	return nil
}
