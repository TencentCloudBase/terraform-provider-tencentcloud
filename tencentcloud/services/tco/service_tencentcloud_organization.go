package tco

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"

	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"

	organization "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization/v20210331"

	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/connectivity"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type OrganizationService struct {
	client *connectivity.TencentCloudClient
}

func (me *OrganizationService) DescribeOrganizationOrgNode(ctx context.Context, nodeId string) (orgNode *organization.OrgNode, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationNodesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	var offset int64 = 0
	var pageSize int64 = 50
	instances := make([]*organization.OrgNode, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().DescribeOrganizationNodes(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		instances = append(instances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	if len(instances) < 1 {
		return
	}

	for _, instance := range instances {
		if helper.Int64ToStr(*instance.NodeId) == nodeId {
			orgNode = instance
		}
	}

	return
}

func (me *OrganizationService) DeleteOrganizationOrgNodeById(ctx context.Context, nodeId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationNodesRequest()

	request.NodeId = []*int64{helper.Int64(helper.StrToInt64(nodeId))}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().DeleteOrganizationNodes(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMember(ctx context.Context, uin string) (orgMember *organization.OrgMember, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationMembersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var offset uint64 = 0
	var pageSize uint64 = 50
	instances := make([]*organization.OrgMember, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMembers(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		instances = append(instances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	if len(instances) < 1 {
		return
	}

	for _, instance := range instances {
		if helper.Int64ToStr(*instance.MemberUin) == uin {
			orgMember = instance
		}
	}

	return

}

func (me *OrganizationService) DeleteOrganizationOrgMemberById(ctx context.Context, uin string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationMembersRequest()

	request.MemberUin = []*int64{helper.Int64(helper.StrToInt64(uin))}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().DeleteOrganizationMembers(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationPolicySubAccountAttachment(ctx context.Context, policyId, memberUin string) (policySubAccountAttachment *organization.OrgMemberAuthAccount, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationMemberAuthAccountsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()
	request.PolicyId = helper.StrToInt64Point(policyId)
	request.MemberUin = helper.StrToInt64Point(memberUin)
	request.Limit = helper.IntInt64(50)
	request.Offset = helper.IntInt64(0)

	response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberAuthAccounts(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	if len(response.Response.Items) < 1 {
		return
	}
	policySubAccountAttachment = response.Response.Items[0]
	return
}

func (me *OrganizationService) DeleteOrganizationPolicySubAccountAttachmentById(ctx context.Context, policyId, memberUin, orgSubAccountUin string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewCancelOrganizationMemberAuthAccountRequest()

	request.PolicyId = helper.StrToInt64Point(policyId)
	request.MemberUin = helper.StrToInt64Point(memberUin)
	request.OrgSubAccountUin = helper.StrToInt64Point(orgSubAccountUin)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseOrganizationClient().CancelOrganizationMemberAuthAccount(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMemberAuthIdentityById(ctx context.Context, memberUin int64) (identityIds []int64, errRet error) {
	logId := tccommon.GetLogId(ctx)
	request := organization.NewDescribeOrganizationMemberAuthIdentitiesRequest()
	request.MemberUin = &memberUin

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	var (
		offset int64 = 0
		limit  int64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberAuthIdentities(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if len(response.Response.Items) < 1 {
			return
		}

		for _, v := range response.Response.Items {
			if v.MemberUin != nil && *v.MemberUin == memberUin {
				if *v.IdentityId == 1 {
					continue
				}
				identityIds = append(identityIds, *v.IdentityId)
			}
		}
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *OrganizationService) DeleteOrganizationOrgMemberAuthIdentityById(ctx context.Context, memberUin string, identityId []string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationMemberAuthIdentityRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	for _, id := range identityId {
		log.Printf("[DEBUG]%s api[%s] delete identity, uin [%s], identityId [%s]\n", logId, request.GetAction(), memberUin, id)

		request.MemberUin = helper.StrToUint64Point(memberUin)
		request.IdentityId = helper.StrToUint64Point(id)

		response, err := me.client.UseOrganizationClient().DeleteOrganizationMemberAuthIdentity(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] delete identity success, request memberUin [%s], response body [%s]\n", logId, request.GetAction(), memberUin, response.ToJsonString())
	}

	return
}

func (me *OrganizationService) DescribeOrganizationOrgAuthNodeByFilter(ctx context.Context, param map[string]interface{}) (orgAuthNode []*organization.AuthNode, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationAuthNodeRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "AuthName" {
			request.AuthName = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationAuthNode(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		orgAuthNode = append(orgAuthNode, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *OrganizationService) DescribeOrganizationOrganizationById(ctx context.Context) (result *organization.DescribeOrganizationResponseParams, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDescribeOrganizationRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeOrganization(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	result = response.Response
	return
}

func (me *OrganizationService) DeleteOrganizationOrganizationById(ctx context.Context) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganization(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgMemberEmailById(ctx context.Context, memberUin int64, bindId uint64) (orgMemberEmail *organization.DescribeOrganizationMemberEmailBindResponseParams, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDescribeOrganizationMemberEmailBindRequest()
	request.MemberUin = &memberUin

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeOrganizationMemberEmailBind(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil {
		return
	}
	if *response.Response.BindId != bindId {
		return
	}
	orgMemberEmail = response.Response
	return
}

func (me *OrganizationService) DescribeOrganizationOrgFinancialByMemberByFilter(ctx context.Context, param map[string]interface{}) (orgFinancialByMember *organization.DescribeOrganizationFinancialByMemberResponseParams, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationFinancialByMemberRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Month" {
			request.Month = v.(*string)
		}
		if k == "EndMonth" {
			request.EndMonth = v.(*string)
		}
		if k == "MemberUins" {
			request.MemberUins = v.([]*int64)
		}
		if k == "ProductCodes" {
			request.ProductCodes = v.([]*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset    int64 = 0
		limit     int64 = 20
		items     []*organization.OrgMemberFinancial
		totalCost float64 = 0
		total     int64   = 0
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationFinancialByMember(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		items = append(items, response.Response.Items...)
		if response.Response != nil && response.Response.TotalCost != nil && totalCost == 0 && total == 0 {
			totalCost = *response.Response.TotalCost
			total = *response.Response.Total
		}
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}
	orgFinancialByMember = &organization.DescribeOrganizationFinancialByMemberResponseParams{
		TotalCost: &totalCost,
		Items:     items,
		Total:     &total,
	}
	return
}

func (me *OrganizationService) DescribeOrganizationOrgFinancialByMonthByFilter(ctx context.Context, param map[string]interface{}) (orgFinancialByMonth []*organization.OrgFinancialByMonth, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationFinancialByMonthRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "EndMonth" {
			request.EndMonth = v.(*string)
		}
		if k == "MemberUins" {
			request.MemberUins = v.([]*int64)
		}
		if k == "ProductCodes" {
			request.ProductCodes = v.([]*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeOrganizationFinancialByMonth(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.Items) < 1 {
		return
	}
	orgFinancialByMonth = append(orgFinancialByMonth, response.Response.Items...)
	return
}
func (me *OrganizationService) DescribeOrganizationOrgFinancialByProductByFilter(ctx context.Context, param map[string]interface{}) (orgFinancialByProduct *organization.DescribeOrganizationFinancialByProductResponseParams, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationFinancialByProductRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Month" {
			request.Month = v.(*string)
		}
		if k == "EndMonth" {
			request.EndMonth = v.(*string)
		}
		if k == "MemberUins" {
			request.MemberUins = v.([]*int64)
		}
		if k == "ProductCodes" {
			request.ProductCodes = v.([]*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset    int64 = 0
		limit     int64 = 20
		items     []*organization.OrgProductFinancial
		totalCost float64 = 0
		total     int64   = 0
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationFinancialByProduct(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		items = append(items, response.Response.Items...)
		if response.Response != nil && response.Response.TotalCost != nil && totalCost == 0 && total == 0 {
			totalCost = *response.Response.TotalCost
			total = *response.Response.Total
		}
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}
	orgFinancialByProduct = &organization.DescribeOrganizationFinancialByProductResponseParams{
		TotalCost: &totalCost,
		Items:     items,
		Total:     &total,
	}
	return
}

func (me *OrganizationService) DescribeOrganizationOrgIdentityById(ctx context.Context, identityId string) (orgIdentity *organization.OrgIdentity, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewListOrganizationIdentityRequest()
	request.IdentityId = helper.StrToUint64Point(identityId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	var tmp []*organization.OrgIdentity
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().ListOrganizationIdentity(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		tmp = append(tmp, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}
	}
	for _, item := range tmp {
		if *item.IdentityId == helper.StrToInt64(identityId) {
			orgIdentity = item
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgIdentityById(ctx context.Context, identityId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationIdentityRequest()
	request.IdentityId = helper.StrToUint64Point(identityId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganizationIdentity(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DeleteOrganizationOrgMemberPolicyAttachmentById(ctx context.Context, policyId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrganizationMembersPolicyRequest()
	request.PolicyId = helper.StrToUint64Point(policyId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteOrganizationMembersPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationMembersByFilter(ctx context.Context, param map[string]interface{}) (members []*organization.OrgMember, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeOrganizationMembersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Lang" {
			request.Lang = v.(*string)
		}
		if k == "SearchKey" {
			request.SearchKey = v.(*string)
		}
		if k == "AuthName" {
			request.AuthName = v.(*string)
		}
		if k == "Product" {
			request.Product = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeOrganizationMembers(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		members = append(members, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

func (me *OrganizationService) DescribeOrganizationOrgShareUnitById(ctx context.Context, area, unitId string) (orgShareUnit *organization.ManagerShareUnit, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDescribeShareUnitsRequest()
	request.SearchKey = &unitId
	request.Area = &area
	request.Limit = helper.IntUint64(20)
	request.Offset = helper.IntUint64(0)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeShareUnits(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || len(response.Response.Items) < 1 {
		return
	}

	orgShareUnit = response.Response.Items[0]
	return
}

func (me *OrganizationService) DeleteOrganizationOrgShareUnitById(ctx context.Context, unitId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteShareUnitRequest()
	request.UnitId = &unitId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteShareUnit(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgShareUnitMemberById(ctx context.Context, unitId, area, shareMemberUins string) (orgShareUnitMembers []*organization.ShareUnitMember, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDescribeShareUnitMembersRequest()
	request.UnitId = &unitId
	request.Area = &area

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	ratelimit.Check(request.GetAction())
	var (
		offset uint64 = 0
		limit  uint64 = 10
	)
	var tmp []*organization.ShareUnitMember
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().DescribeShareUnitMembers(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}
		tmp = append(tmp, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	memberIdMap := make(map[string]struct{})

	for _, item := range tmp {
		memberIdMap[helper.Int64ToStr(*item.ShareMemberUin)] = struct{}{}
	}

	split := strings.Split(shareMemberUins, tccommon.COMMA_SP)
	if len(split) < 1 {
		errRet = fmt.Errorf("shareMemberUins is broken, %s", shareMemberUins)
		return
	}
	for _, v := range split {
		if _, ok := memberIdMap[v]; !ok {
			return
		}
	}
	orgShareUnitMembers = tmp
	return
}

func (me *OrganizationService) DeleteOrganizationOrgShareUnitMemberById(ctx context.Context, unitId, area, shareMemberUins string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteShareUnitMembersRequest()
	request.UnitId = &unitId
	request.Area = &area
	split := strings.Split(shareMemberUins, tccommon.COMMA_SP)
	if len(split) < 1 {
		errRet = fmt.Errorf("shareMemberUins is broken, %s", shareMemberUins)
		return
	}
	for _, v := range split {
		request.Members = append(request.Members, &organization.ShareMember{ShareMemberUin: helper.StrToInt64Point(v)})
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeleteShareUnitMembers(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgShareAreaByFilter(ctx context.Context, param map[string]interface{}) (orgShareArea []*organization.ShareArea, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewDescribeShareAreasRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Lang" {
			request.Lang = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribeShareAreas(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.Items) < 1 {
		return
	}

	orgShareArea = response.Response.Items
	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyConfigById(ctx context.Context, organizationId string, policyType string) (OrgManagePolicyConfig *organization.DescribePolicyConfigResponseParams, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDescribePolicyConfigRequest()
	request.OrganizationId = helper.StrToUint64Point(organizationId)
	request.Type = helper.IntUint64(ServiceControlPolicyCode)

	if policyType == TagPolicyType {
		request.Type = helper.IntUint64(TagPolicyCode)
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DescribePolicyConfig(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if *response.Response.Status == 1 {
		OrgManagePolicyConfig = response.Response
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyConfigById(ctx context.Context, organizationId string, policyType string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDisablePolicyTypeRequest()
	request.OrganizationId = helper.StrToUint64Point(organizationId)
	request.PolicyType = &policyType

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DisablePolicyType(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyById(ctx context.Context, policyId, policyType string) (OrgManagePolicy *organization.DescribePolicyResponseParams, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewListPoliciesRequest()
	request.PolicyType = helper.String(policyType)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	pageStart := uint64(1)
	rp := uint64(PAGE_ITEM) //to save in extension
	result := make([]*organization.ListPolicyNode, 0)
	for {
		request.Page = &pageStart
		request.Rp = &rp
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().ListPolicies(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.List) < 1 {
			break
		}
		result = append(result, response.Response.List...)
		if len(response.Response.List) < PAGE_ITEM {
			break
		}
		pageStart += 1
	}

	for _, item := range result {
		if helper.UInt64ToStr(*item.PolicyId) == policyId {
			requestDescribe := organization.NewDescribePolicyRequest()
			requestDescribe.PolicyId = item.PolicyId
			requestDescribe.PolicyType = helper.String(policyType)
			responseDescribe, err := me.client.UseOrganizationClient().DescribePolicy(requestDescribe)
			if err != nil {
				errRet = err
				return
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), responseDescribe.ToJsonString())

			if responseDescribe == nil || responseDescribe.Response == nil {
				break
			}
			OrgManagePolicy = responseDescribe.Response
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyById(ctx context.Context, policyId, policyType string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeletePolicyRequest()
	request.PolicyId = helper.StrToUint64Point(policyId)
	request.Type = helper.String(policyType)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DeletePolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationOrgManagePolicyTargetById(ctx context.Context, policyType string, policyId string, targetType string, targetId string) (OrgManagePolicyTarget *organization.ListTargetsForPolicyNode, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewListTargetsForPolicyRequest()
	request.PolicyType = &policyType
	request.PolicyId = helper.StrToUint64Point(policyId)
	switch targetType {
	case TargetTypeNode:
		request.TargetType = helper.String(DescribeTargetTypeNode)
	case TargetTypeMember:
		request.TargetType = helper.String(DescribeTargetTypeMember)
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().ListTargetsForPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	for _, item := range response.Response.List {
		if item.Uin != nil && helper.UInt64ToStr(*item.Uin) == targetId {
			OrgManagePolicyTarget = item
		}
	}
	return
}

func (me *OrganizationService) DeleteOrganizationOrgManagePolicyTargetById(ctx context.Context, policyType string, policyId string, targetType string, targetId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDetachPolicyRequest()
	request.Type = &policyType
	request.PolicyId = helper.StrToUint64Point(policyId)
	request.TargetType = &targetType
	request.TargetId = helper.StrToUint64Point(targetId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseOrganizationClient().DetachPolicy(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *OrganizationService) DescribeOrganizationServiceAssignMemberById(ctx context.Context, serviceId string) (items []*organization.OrganizationServiceAssignMember, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewListOrgServiceAssignMemberRequest()
	serviceIdInt, _ := strconv.ParseUint(serviceId, 10, 64)
	request.ServiceId = &serviceIdInt

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var offset uint64 = 0
	var pageSize uint64 = 10
	items = make([]*organization.OrganizationServiceAssignMember, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseOrganizationClient().ListOrgServiceAssignMember(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}

		items = append(items, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}

		offset += pageSize
	}

	return
}

func (me *OrganizationService) DeleteOrganizationServiceAssignMemberById(ctx context.Context, serviceId string, memberUinList []*int64) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := organization.NewDeleteOrgServiceAssignRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	serviceIdInt, _ := strconv.ParseUint(serviceId, 10, 64)
	for _, memberUin := range memberUinList {
		ratelimit.Check(request.GetAction())
		request.ServiceId = &serviceIdInt
		request.MemberUin = memberUin
		response, err := me.client.UseOrganizationClient().DeleteOrgServiceAssign(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	}

	return
}

func (me *OrganizationService) DescribeOrganizationServicesByFilter(ctx context.Context, param map[string]interface{}) (members []*organization.OrganizationServiceAssign, errRet error) {
	var (
		logId   = tccommon.GetLogId(ctx)
		request = organization.NewListOrganizationServiceRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "SearchKey" {
			request.SearchKey = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset uint64 = 0
		limit  uint64 = 10
	)

	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseOrganizationClient().ListOrganizationService(request)
		if err != nil {
			errRet = err
			return
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Items) < 1 {
			break
		}

		members = append(members, response.Response.Items...)
		if len(response.Response.Items) < int(limit) {
			break
		}

		offset += limit
	}

	return
}
