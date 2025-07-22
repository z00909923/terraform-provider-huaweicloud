// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product IdentityCenter
// ---------------------------------------------------------------

package identitycenter

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API IdentitySCIM POST /{tenant_id}/scim/v2/Users
// @API IdentitySCIM GET /{tenant_id}/scim/v2/Users/{user_id}
// @API IdentitySCIM PUT /{tenant_id}/scim/v2/Users/{user_id}
// @API IdentitySCIM DELETE /{tenant_id}/scim/v2/Users/{user_id}
//TODO Authorization header头在底层源码里会被覆盖，导致实际请求带的Authorization header为底层源码里的签名，Bearertoken无法认证 401
func ResourceIdentityCenterSCIMUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdentityCenterSCIMUserCreate,
		UpdateContext: resourceIdentityCenterSCIMUserUpdate,
		ReadContext:   resourceIdentityCenterSCIMUserRead,
		DeleteContext: resourceIdentityCenterSCIMUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceIdentityCenterSCIMUserImportState,
		},

		Description: "schema: Internal",
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the identity store`,
			},
			"header_authorization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the identity store`,
			},
			"schemas": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the identity store`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the username of the user.`,
			},
			"family_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the family name of the user.`,
			},
			"given_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the given name of the user.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the display name of the user.`,
			},
			"nick_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the family name of the user.`,
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the email of the user.`,
			},
			"external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the phone number of the user.`,
			},
			"phone_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the phone number of the user.`,
			},
			"profile_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the phone number of the user.`,
			},
			"user_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the type of the user.`,
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the title of the user.`,
			},
			"preferred_language": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the title of the user.`,
			},
			"locale": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the title of the user.`,
			},
			"timezone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the title of the user.`,
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the title of the user.`,
			},
			"addresses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Specifies the addresses information of the user.`,
				MaxItems:    1,
				Elem:        identityCenterSCIMUserAddressesSchema(),
			},
			"enterprise": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Specifies the enterprise information of the user.`,
				MaxItems:    1,
				Elem:        identityCenterSCIMUserEnterpriseSchema(),
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Specifies the manager of the enterprise.`,
			},
			"last_modified": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Specifies the manager of the enterprise.`,
			},
		},
	}
}

func identityCenterSCIMUserAddressesSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"country": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the country of the user.`,
			},
			"formatted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies a string containing a formatted version of the address to be displayed.`,
			},
			"locality": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the locality of the user.`,
			},
			"postal_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the postal code of the user.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the region of the user.`,
			},
			"street_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the street address of the user.`,
			},
		},
	}
	return &sc
}

func identityCenterSCIMUserEnterpriseSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"cost_center": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the cost center of the enterprise.`,
			},
			"department": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the department of the enterprise.`,
			},
			"division": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the division of the enterprise.`,
			},
			"employee_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the employee number of the enterprise.`,
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the organization of the enterprise.`,
			},
			"manager": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the manager of the enterprise.`,
			},
		},
	}
	return &sc
}

func resourceIdentityCenterSCIMUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createIdentityCenterUser: create IdentityCenter user
	var (
		createIdentityCenterUserHttpUrl = "{tenant_id}/scim/v2/Users"
		createIdentityCenterUserProduct = "identityscim"
	)
	createIdentityCenterUserClient, err := cfg.NewServiceClient(createIdentityCenterUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	createIdentityCenterUserPath := createIdentityCenterUserClient.Endpoint + createIdentityCenterUserHttpUrl
	createIdentityCenterUserPath = strings.ReplaceAll(createIdentityCenterUserPath, "{tenant_id}",
		fmt.Sprintf("%v", d.Get("tenant_id")))
	bearerToken := d.Get("header_authorization").(string)
	createIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Authorization": bearerToken,
		},
	}
	createIdentityCenterUserOpt.JSONBody = utils.RemoveNil(buildCreateIdentityCenterSCIMUserBodyParams(d))
	createIdentityCenterUserResp, err := createIdentityCenterUserClient.Request("POST",
		createIdentityCenterUserPath, &createIdentityCenterUserOpt)
	if err != nil {
		return diag.Errorf("error creating Identity Center User: %s", err)
	}

	createIdentityCenterUserRespBody, err := utils.FlattenResponse(createIdentityCenterUserResp)
	if err != nil {
		return diag.FromErr(err)
	}

	userId := utils.PathSearch("user_id", createIdentityCenterUserRespBody, "").(string)
	if userId == "" {
		return diag.Errorf("unable to find the Identity Center user ID from the API response")
	}
	d.SetId(userId)

	return resourceIdentityCenterUserRead(ctx, d, meta)
}

func buildCreateIdentityCenterSCIMUserBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"schemas":           d.Get("schemas"),
		"externalId":        utils.ValueIgnoreEmpty(d.Get("external_id")),
		"profileUrl":        utils.ValueIgnoreEmpty(d.Get("profile_url")),
		"nickName":          utils.ValueIgnoreEmpty(d.Get("nick_name")),
		"preferredLanguage": utils.ValueIgnoreEmpty(d.Get("preferred_language")),
		"locale":            utils.ValueIgnoreEmpty(d.Get("locale")),
		"timezone":          utils.ValueIgnoreEmpty(d.Get("timezone")),
		"active":            utils.ValueIgnoreEmpty(d.Get("active")),
		"userName":          utils.ValueIgnoreEmpty(d.Get("user_name")),
		"displayName":       utils.ValueIgnoreEmpty(d.Get("display_name")),
		"emails": []map[string]interface{}{{
			"primary": true,
			"type":    "Work",
			"value":   utils.ValueIgnoreEmpty(d.Get("email")),
		}},
		"name": map[string]interface{}{
			"familyName": utils.ValueIgnoreEmpty(d.Get("family_name")),
			"givenName":  utils.ValueIgnoreEmpty(d.Get("given_name")),
		},
		"title":    utils.ValueIgnoreEmpty(d.Get("title")),
		"userType": utils.ValueIgnoreEmpty(d.Get("user_type")),
		"phoneNumbers": []map[string]interface{}{{
			"primary": true,
			"type":    "Work",
			"value":   utils.ValueIgnoreEmpty(d.Get("phone_number")),
		}},
		"addresses": []map[string]interface{}{
			{
				"country":       utils.ValueIgnoreEmpty(d.Get("addresses.0.country")),
				"formatted":     utils.ValueIgnoreEmpty(d.Get("addresses.0.formatted")),
				"locality":      utils.ValueIgnoreEmpty(d.Get("addresses.0.locality")),
				"postalCode":    utils.ValueIgnoreEmpty(d.Get("addresses.0.postal_code")),
				"region":        utils.ValueIgnoreEmpty(d.Get("addresses.0.region")),
				"streetAddress": utils.ValueIgnoreEmpty(d.Get("addresses.0.street_address")),
			},
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": map[string]interface{}{
			"costCenter":     utils.ValueIgnoreEmpty(d.Get("enterprise.0.cost_center")),
			"department":     utils.ValueIgnoreEmpty(d.Get("enterprise.0.department")),
			"division":       utils.ValueIgnoreEmpty(d.Get("enterprise.0.division")),
			"employeeNumber": utils.ValueIgnoreEmpty(d.Get("enterprise.0.employee_number")),
			"manager": map[string]interface{}{
				"value": utils.ValueIgnoreEmpty(d.Get("enterprise.0.manager")),
			},
			"organization": utils.ValueIgnoreEmpty(d.Get("enterprise.0.organization")),
		},
	}
	return bodyParams
}

func resourceIdentityCenterSCIMUserRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getIdentityCenterUser: Query Identity Center user
	var (
		getIdentityCenterUserHttpUrl = "{tenant_id}/scim/v2/Users/{user_id}"
		getIdentityCenterUserProduct = "identityscim"
	)
	getIdentityCenterUserClient, err := cfg.NewServiceClient(getIdentityCenterUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	getIdentityCenterUserPath := getIdentityCenterUserClient.Endpoint + getIdentityCenterUserHttpUrl
	getIdentityCenterUserPath = strings.ReplaceAll(getIdentityCenterUserPath, "{tenant_id}",
		fmt.Sprintf("%v", d.Get("tenant_id")))
	getIdentityCenterUserPath = strings.ReplaceAll(getIdentityCenterUserPath, "{user_id}", d.Id())

	bearerToken := d.Get("header_authorization").(string)
	getIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Authorization": bearerToken,
		},
	}
	getIdentityCenterUserResp, err := getIdentityCenterUserClient.Request("GET", getIdentityCenterUserPath,
		&getIdentityCenterUserOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving Identity Center User")
	}

	getIdentityCenterUserRespBody, err := utils.FlattenResponse(getIdentityCenterUserResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("external_id", utils.PathSearch("externalId", getIdentityCenterUserRespBody, nil)),
		d.Set("schemas", utils.PathSearch("schemas", getIdentityCenterUserRespBody, nil)),
		d.Set("created", utils.PathSearch("meta.created", getIdentityCenterUserRespBody, nil)),
		d.Set("last_modified", utils.PathSearch("meta.lastModified", getIdentityCenterUserRespBody, nil)),
		d.Set("active", utils.PathSearch("active", getIdentityCenterUserRespBody, nil)),
		d.Set("preferred_language", utils.PathSearch("preferredLanguage", getIdentityCenterUserRespBody, nil)),
		d.Set("user_type", utils.PathSearch("userType", getIdentityCenterUserRespBody, nil)),
		d.Set("locale", utils.PathSearch("locale", getIdentityCenterUserRespBody, nil)),
		d.Set("timezone", utils.PathSearch("timezone", getIdentityCenterUserRespBody, nil)),
		d.Set("user_name", utils.PathSearch("userName", getIdentityCenterUserRespBody, nil)),
		d.Set("family_name", utils.PathSearch("name.familyName", getIdentityCenterUserRespBody, nil)),
		d.Set("given_name", utils.PathSearch("name.givenName", getIdentityCenterUserRespBody, nil)),
		d.Set("display_name", utils.PathSearch("displayName", getIdentityCenterUserRespBody, nil)),
		d.Set("email", utils.PathSearch("emails|[0].value", getIdentityCenterUserRespBody, nil)),
		d.Set("phone_number", utils.PathSearch("phoneNumbers|[0].value", getIdentityCenterUserRespBody, nil)),
		d.Set("title", utils.PathSearch("title", getIdentityCenterUserRespBody, nil)),
		d.Set("user_type", utils.PathSearch("userType", getIdentityCenterUserRespBody, nil)),
		d.Set("addresses", flattenIdentityCenterSCIMUserAddresses(utils.PathSearch("addresses|[0]", getIdentityCenterUserRespBody, nil))),
		d.Set("enterprise", flattenIdentityCenterSCIMUserEnterprise(utils.PathSearch("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", getIdentityCenterUserRespBody, nil))),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenIdentityCenterSCIMUserAddresses(address interface{}) []map[string]interface{} {
	if address == nil || len(address.(map[string]interface{})) == 0 {
		return nil
	}

	return []map[string]interface{}{
		{
			"country":        utils.PathSearch("country", address, nil),
			"formatted":      utils.PathSearch("formatted", address, nil),
			"locality":       utils.PathSearch("locality", address, nil),
			"postal_code":    utils.PathSearch("postalCode", address, nil),
			"region":         utils.PathSearch("region", address, nil),
			"street_address": utils.PathSearch("streetAddress", address, nil),
		},
	}
}

func flattenIdentityCenterSCIMUserEnterprise(enterprise interface{}) []map[string]interface{} {
	if enterprise == nil {
		return nil
	}

	// enterprise format: {"enterprise":{"manager":{}}}
	enterpriseMap := enterprise.(map[string]interface{})
	manager := utils.PathSearch("manager.value", enterprise, nil)
	if len(enterpriseMap) == 1 && manager == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"cost_center":     utils.PathSearch("costCenter", enterprise, nil),
			"department":      utils.PathSearch("department", enterprise, nil),
			"division":        utils.PathSearch("division", enterprise, nil),
			"employee_number": utils.PathSearch("employeeNumber", enterprise, nil),
			"organization":    utils.PathSearch("organization", enterprise, nil),
			"manager":         manager,
		},
	}
}

func resourceIdentityCenterSCIMUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// updateIdentityCenterUser: update Identity Center user
	var (
		updateIdentityCenterUserHttpUrl = "{tenant_id}/scim/v2/Users/{user_id}"
		updateIdentityCenterUserProduct = "identityscim"
	)
	updateIdentityCenterUserClient, err := cfg.NewServiceClient(updateIdentityCenterUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	updateIdentityCenterUserPath := updateIdentityCenterUserClient.Endpoint + updateIdentityCenterUserHttpUrl
	updateIdentityCenterUserPath = strings.ReplaceAll(updateIdentityCenterUserPath, "{tenant_id}",
		fmt.Sprintf("%v", d.Get("tenant_id")))
	updateIdentityCenterUserPath = strings.ReplaceAll(updateIdentityCenterUserPath, "{user_id}", d.Id())

	bearerToken := d.Get("header_authorization").(string)
	updateIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Authorization": bearerToken,
		},
	}
	updateIdentityCenterUserOpt.JSONBody = utils.RemoveNil(buildCreateIdentityCenterSCIMUserBodyParams(d))
	_, err = updateIdentityCenterUserClient.Request("PUT", updateIdentityCenterUserPath,
		&updateIdentityCenterUserOpt)
	if err != nil {
		return diag.Errorf("error updating Identity Center User: %s", err)
	}
	return resourceIdentityCenterUserRead(ctx, d, meta)
}

func resourceIdentityCenterSCIMUserDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// deleteIdentityCenterUser: Delete Identity Center user
	var (
		deleteIdentityCenterUserHttpUrl = "{tenant_id}/scim/v2/Users/{user_id}"
		deleteIdentityCenterUserProduct = "identityscim"
	)
	deleteIdentityCenterUserClient, err := cfg.NewServiceClient(deleteIdentityCenterUserProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	deleteIdentityCenterUserPath := deleteIdentityCenterUserClient.Endpoint + deleteIdentityCenterUserHttpUrl
	deleteIdentityCenterUserPath = strings.ReplaceAll(deleteIdentityCenterUserPath, "{tenant_id}",
		fmt.Sprintf("%v", d.Get("tenant_id")))
	deleteIdentityCenterUserPath = strings.ReplaceAll(deleteIdentityCenterUserPath, "{user_id}", d.Id())

	bearerToken := d.Get("header_authorization").(string)
	deleteIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Authorization": bearerToken,
		},
	}
	_, err = deleteIdentityCenterUserClient.Request("DELETE", deleteIdentityCenterUserPath,
		&deleteIdentityCenterUserOpt)
	if err != nil {
		return diag.Errorf("error deleting Identity Center User: %s", err)
	}

	return nil
}

func resourceIdentityCenterSCIMUserImportState(_ context.Context, d *schema.ResourceData,
	_ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid id format, must be <tenant_id>/<id>")
	}
	d.SetId(parts[1])
	mErr := multierror.Append(nil,
		d.Set("tenant_id", parts[0]),
	)
	if err := mErr.ErrorOrNil(); err != nil {
		return nil, fmt.Errorf("failed to set value to state when import, %s", err)
	}
	return []*schema.ResourceData{d}, nil
}
