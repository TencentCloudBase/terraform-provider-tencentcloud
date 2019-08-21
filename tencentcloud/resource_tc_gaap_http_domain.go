package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTencentCloudGaapHttpDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudGaapHttpDomainCreate,
		Read:   resourceTencentCloudGaapHttpDomainRead,
		Update: resourceTencentCloudGaapHttpDomainUpdate,
		Delete: resourceTencentCloudGaapHttpDomainDelete,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"client_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"basic_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"basic_auth_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realserver_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"realserver_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realserver_certificate_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gaap_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gaap_auth_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTencentCloudGaapHttpDomainCreate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_http_domain.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	listenerId := d.Get("listener_id").(string)
	domain := d.Get("domain").(string)

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	var (
		protocol        string
		forwardProtocol string
	)
	httpListeners, err := service.DescribeHTTPListeners(ctx, nil, &listenerId, nil, nil)
	if err != nil {
		return err
	}
	if len(httpListeners) > 0 {
		protocol = "HTTP"
	}

	httpsListeners, err := service.DescribeHTTPSListeners(ctx, nil, &listenerId, nil, nil)
	if err != nil {
		return err
	}
	if len(httpsListeners) > 0 {
		protocol = "HTTPS"
		if httpsListeners[0].ForwardProtocol == nil {
			return errors.New("https listener forward protocol is nil")
		}
		forwardProtocol = *httpsListeners[0].ForwardProtocol
	}

	switch protocol {
	case "HTTP":
		if err := service.CreateHTTPDomain(ctx, listenerId, domain); err != nil {
			return err
		}

		id := fmt.Sprintf("%s+%s+%s", listenerId, protocol, domain)
		d.SetId(id)

	case "HTTPS":
		var (
			certificateId               string
			clientCertificateId         string
			basicAuth                   bool
			basicAuthId                 *string
			realserverAuth              bool
			realserverCertificateId     *string
			realserverCertificateDomain *string
			gaapAuth                    bool
			gaapCertificateId           *string
		)

		certificateId = d.Get("certificate_id").(string)
		clientCertificateId = d.Get("client_certificate_id").(string)

		// basic auth
		basicAuth = d.Get("basic_auth").(bool)

		if raw, ok := d.GetOk("basic_auth_id"); ok {
			basicAuthId = stringToPointer(raw.(string))
		}
		if basicAuth && basicAuthId == nil {
			return errors.New("when use basic auth, basic auth id should be set")
		}

		// realserver certification
		realserverAuth = d.Get("realserver_auth").(bool)
		if forwardProtocol == "HTTP" && realserverAuth {
			return errors.New("when listener forward protocol is http, realserver_auth can't be true")
		}

		if raw, ok := d.GetOk("realserver_certificate_id"); ok {
			realserverCertificateId = stringToPointer(raw.(string))
		}
		if raw, ok := d.GetOk("realserver_certificate_domain"); ok {
			realserverCertificateDomain = stringToPointer(raw.(string))
		}

		if realserverAuth && (realserverCertificateId == nil || realserverCertificateDomain == nil) {
			return errors.New("when use realserver auth, realserver auth id and domain should be set")
		}

		// gaap certification
		gaapAuth = d.Get("gaap_auth").(bool)
		if forwardProtocol == "HTTP" && gaapAuth {
			return errors.New("when listener forward protocol is http, gaap_auth can't be set")
		}

		if raw, ok := d.GetOk("gaap_auth_id"); ok {
			gaapCertificateId = stringToPointer(raw.(string))
		}

		if gaapAuth && gaapCertificateId == nil {
			return errors.New("when use gaap auth, gaap auth id should be set")
		}

		if err := service.CreateHTTPSDomain(ctx, listenerId, domain, certificateId, clientCertificateId); err != nil {
			return err
		}

		id := fmt.Sprintf("%s+%s+%s", listenerId, protocol, domain)
		d.SetId(id)

		if err := service.SetAdvancedAuth(
			ctx,
			listenerId, domain,
			&realserverAuth, &basicAuth, &gaapAuth,
			realserverCertificateId, realserverCertificateDomain, basicAuthId, gaapCertificateId,
		); err != nil {
			return err
		}
	}

	return resourceTencentCloudGaapHttpDomainRead(d, m)
}

func resourceTencentCloudGaapHttpDomainRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_http_domain.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	var (
		listenerId string
		protocol   string
		domain     string
	)
	split := strings.Split(id, "+")
	listenerId, protocol, domain = split[0], split[1], split[2]

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	httpDomain, err := service.DescribeDomain(ctx, listenerId, domain)
	if err != nil {
		return err
	}

	if httpDomain == nil {
		d.SetId("")
		return nil
	}

	d.Set("domain", httpDomain.Domain)

	if protocol == "HTTP" {
		return nil
	}

	listeners, err := service.DescribeHTTPSListeners(ctx, nil, &listenerId, nil, nil)
	if err != nil {
		return err
	}

	// if listener doesn't exist, domain won't exist
	if len(listeners) == 0 {
		log.Printf("[DEBUG]%s domain listener doesn't exist", logId)
		d.SetId("")
		return nil
	}

	if listeners[0].ForwardProtocol == nil {
		return errors.New("listener forward protocol is nil")
	}
	forwardProtocol := *listeners[0].ForwardProtocol

	if httpDomain.CertificateId == nil {
		return errors.New("domain certificate id is nil")
	}
	d.Set("certificate_id", httpDomain.CertificateId)

	if httpDomain.ClientCertificateId == nil {
		return errors.New("domain client certificate id is nil")
	}
	d.Set("client_certificate_id", httpDomain.ClientCertificateId)

	if httpDomain.BasicAuth == nil {
		return errors.New("domain basic auth is nil")
	}
	d.Set("basic_auth", *httpDomain.BasicAuth == 1)

	if httpDomain.BasicAuthConfId != nil {
		d.Set("basic_auth_id", httpDomain.BasicAuthConfId)
	}

	if forwardProtocol == "HTTPS" {
		if httpDomain.RealServerAuth == nil {
			return errors.New("domain realserver auth is nil")
		}
		d.Set("realserver_auth", *httpDomain.RealServerAuth == 1)

		if httpDomain.RealServerCertificateId != nil {
			d.Set("realserver_certificate_id", httpDomain.RealServerCertificateId)
		}

		if httpDomain.RealServerCertificateDomain != nil {
			d.Set("realserver_certificate_domain", httpDomain.RealServerCertificateDomain)
		}

		if httpDomain.GaapAuth == nil {
			return errors.New("domain gaap auth is nil")
		}
		d.Set("gaap_auth", *httpDomain.GaapAuth == 1)

		if httpDomain.GaapCertificateId != nil {
			d.Set("gaap_auth_id", httpDomain.GaapCertificateId)
		}
	}

	return nil
}

func resourceTencentCloudGaapHttpDomainUpdate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_http_domain.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	var (
		listenerId string
		protocol   string
		domain     string
	)
	split := strings.Split(id, "+")
	listenerId, protocol, domain = split[0], split[1], split[2]

	// when protocol is http, nothing can be updated
	if protocol == "HTTP" {
		return errors.New("http listener can't set auth")
	}

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	listeners, err := service.DescribeHTTPSListeners(ctx, nil, &listenerId, nil, nil)
	if err != nil {
		return err
	}

	// if listener doesn't exist, domain won't exist
	if len(listeners) == 0 {
		log.Printf("[DEBUG]%s domain listener doesn't exist", logId)
		d.SetId("")
		return nil
	}

	if listeners[0].ForwardProtocol == nil {
		return errors.New("listener forward protocol is nil")
	}
	forwardProtocol := *listeners[0].ForwardProtocol

	d.Partial(true)

	if d.HasChange("certificate_id") || d.HasChange("client_certificate_id") {
		var (
			certificateId       *string
			clientCertificateId *string
		)

		if d.HasChange("certificate_id") {
			certificateId = stringToPointer(d.Get("certificate_id").(string))
		}

		if d.HasChange("client_certificate_id") {
			clientCertificateId = stringToPointer(d.Get("client_certificate_id").(string))
		}

		if err := service.ModifyDomainCertificate(ctx, listenerId, domain, certificateId, clientCertificateId); err != nil {
			return err
		}

		if d.HasChange("certificate_id") {
			d.SetPartial("certificate_id")
		}
		if d.HasChange("client_certificate_id") {
			d.SetPartial("client_certificate_id")
		}
	}

	var (
		realserverAuth              *bool
		realserverCertificateId     *string
		realserverCertificateDomain *string
		basicAuth                   *bool
		basicAuthId                 *string
		gaapAuth                    *bool
		gaapCertificateId           *string
		updateAdvancedAttr          []string
	)

	if d.HasChange("basic_auth") {
		updateAdvancedAttr = append(updateAdvancedAttr, "basic_auth")
		basicAuth = boolToPointer(d.Get("basic_auth").(bool))
	}
	if d.HasChange("basic_auth_id") {
		updateAdvancedAttr = append(updateAdvancedAttr, "basic_auth_id")
		basicAuthId = stringToPointer(d.Get("basic_auth_id").(string))
	}

	if forwardProtocol == "HTTP" {
		if d.HasChange("realserver_auth") ||
			d.HasChange("realserver_certificate_id") ||
			d.HasChange("realserver_certificate_domain") ||
			d.HasChange("gaap_auth") ||
			d.HasChange("gaap_auth_id") {
			return errors.New("when listener forward protocol is https, only can set basic auth")
		}
	}

	if d.HasChange("realserver_auth") {
		updateAdvancedAttr = append(updateAdvancedAttr, "realserver_auth")
		realserverAuth = boolToPointer(d.Get("realserver_auth").(bool))
	}
	if d.HasChange("realserver_certificate_id") {
		updateAdvancedAttr = append(updateAdvancedAttr, "realserver_certificate_id")
		realserverCertificateId = stringToPointer(d.Get("realserver_certificate_id").(string))
	}
	if d.HasChange("realserver_certificate_domain") {
		updateAdvancedAttr = append(updateAdvancedAttr, "realserver_certificate_domain")
		realserverCertificateDomain = stringToPointer(d.Get("realserver_certificate_domain").(string))
	}

	if d.HasChange("gaap_auth") {
		updateAdvancedAttr = append(updateAdvancedAttr, "gaap_auth")
		gaapAuth = boolToPointer(d.Get("gaap_auth").(bool))
	}
	if d.HasChange("gaap_auth_id") {
		updateAdvancedAttr = append(updateAdvancedAttr, "gaap_auth_id")
		gaapCertificateId = stringToPointer(d.Get("gaap_auth_id").(string))
	}

	if len(updateAdvancedAttr) > 0 {
		if err := service.SetAdvancedAuth(
			ctx,
			listenerId, domain, realserverAuth, basicAuth, gaapAuth,
			realserverCertificateId, realserverCertificateDomain, basicAuthId, gaapCertificateId,
		); err != nil {
			return err
		}

		for _, attr := range updateAdvancedAttr {
			d.SetPartial(attr)
		}
	}

	d.Partial(false)

	return resourceTencentCloudGaapHttpDomainRead(d, m)
}

func resourceTencentCloudGaapHttpDomainDelete(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_http_domain.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	var (
		listenerId string
		domain     string
	)
	split := strings.Split(id, "+")
	listenerId, domain = split[0], split[2]

	service := GaapService{client: m.(*TencentCloudClient).apiV3Conn}

	return service.DeleteDomain(ctx, listenerId, domain)
}
