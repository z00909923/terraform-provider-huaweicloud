// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product IdentityCenter
// ---------------------------------------------------------------

package identitycenter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API IdentityStore POST /{tenant_id}/scim/v2/Users
// @API IdentityStore GET /{tenant_id}/scim/v2/Users/{user_id}
// @API IdentityStore PUT /{tenant_id}/scim/v2/Users/{user_id}
// @API IdentityStore DELETE /{tenant_id}/scim/v2/Users/{user_id}
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
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"schema": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the identity store`,
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
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the email of the user.`,
			},
			"phone_number": {
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
	bearerToken := d.Get("header_authorization")
	log.Println("ttttttttttttttttttttttttttt" + fmt.Sprintf("%v", bearerToken))
	createIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Authorization": "Bearer eyJraWQiOiJmZUV4U0oxRVppelJZM21FY2E1cHJYaWJRYXJ4cjEzMzlkOGVSWFNwRWhyQyIsImFsZyI6IlJTMjU2IiwidHlwIjoiSldUIn0.eyJhdWQiOiJjbi1ub3J0aC00Iiwic2NvcGUiOlsib3BlbmlkIl0sImlzcyI6Imh0dHBzOi8vb2lkYy5jbi1ub3J0aC00Lmh1YXdlaWNsb3VkLmNvbSIsImV4cCI6MTc4NDAzMjU2MiwiaWF0IjoxNzUyNDk2NTYyLCJqdGkiOiJDN2YyODhhNjgtMjIwYy00MWJmLWI0Y2MtNWNlNGQwMmZhNzJiIn0.OtF0Yf9EIx8HvPhXWQDMbC6Aj_G4PA-9mK6j2vjPRXLe7TQFN5GVnCzIJOdD5jTIdLQBuJpAmzQijBmI50U5HSemSyk8RBky4FRaxkkWzYYDeafmQQyXGah5LQEXzePohzFuUqO8U-NsSTH2D0mAVEHk8N7tA4bG4V5bZJADWEfmMun54qXbfUiK-qptDFc2o1GbucIx7zt6KxVirjMGgpxkOIm5EFRROwAyKpftHVL5jbHmMm6gjheD7A0BcvMy12KLt8LM1-slySQv69xoGeTz8GtCdl5ZzmYaZd1YkSQw68DtXmbcRR0rA9H_wfycf-xMggYUeGgan8znR0yQDuOAHWmTpjKc8Xfy4eXDLJkzjxKEkzzpJFTlz01tCZs_V-G4fG08-4zMnuBGu7ylxi1cROHga2s5fhvZuWt_mwn4qzajjZ0pGjlqPWJ9a5ZfLQyljED_LBQwxkn5IorIdFuLFgO2czPXsyHBHk6RGSUJK93ybx0NaIvMqYtuNd15Z9CRwYPLNmxRvrN8v4vloWsME918UudS3EcuQKTCkhBkKQtXcrkHtFiFW-gezWehGTMih_r6mMaDqDGIr9erdw6c8_0D6azcAgHqy-LadRnatg7TwHttl3tUtvLcM-5oXkmy4zzWTHIz06M0jx_wQ5eI9iLM95C2MtOOOQ8UuVQ",
		},
	}
	log.Println(createIdentityCenterUserOpt)
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
		"schemas":     []string{fmt.Sprintf("%v", d.Get("schema"))},
		"userName":    utils.ValueIgnoreEmpty(d.Get("user_name")),
		"displayName": utils.ValueIgnoreEmpty(d.Get("display_name")),
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
		"enterprise": map[string]interface{}{
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

	getIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
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
		d.Set("region", region),
		d.Set("userName", utils.PathSearch("userName", getIdentityCenterUserRespBody, nil)),
		d.Set("familyName", utils.PathSearch("name.familyName", getIdentityCenterUserRespBody, nil)),
		d.Set("givenName", utils.PathSearch("name.givenName", getIdentityCenterUserRespBody, nil)),
		d.Set("displayName", utils.PathSearch("displayName", getIdentityCenterUserRespBody, nil)),
		d.Set("email", utils.PathSearch("emails|[0].value", getIdentityCenterUserRespBody, nil)),
		d.Set("phoneNumber", utils.PathSearch("phoneNumbers|[0].value", getIdentityCenterUserRespBody, nil)),
		d.Set("title", utils.PathSearch("title", getIdentityCenterUserRespBody, nil)),
		d.Set("userType", utils.PathSearch("userType", getIdentityCenterUserRespBody, nil)),
		d.Set("addresses", flattenIdentityCenterSCIMUserAddresses(utils.PathSearch("addresses|[0]", getIdentityCenterUserRespBody, nil))),
		d.Set("enterprise", flattenIdentityCenterSCIMUserEnterprise(utils.PathSearch("enterprise", getIdentityCenterUserRespBody, nil))),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenIdentityCenterSCIMUserAddresses(address interface{}) []map[string]interface{} {
	if address == nil || len(address.(map[string]interface{})) == 0 {
		return nil
	}

	return []map[string]interface{}{
		{
			"country":       utils.PathSearch("country", address, nil),
			"formatted":     utils.PathSearch("formatted", address, nil),
			"locality":      utils.PathSearch("locality", address, nil),
			"postalCode":    utils.PathSearch("postalCode", address, nil),
			"region":        utils.PathSearch("region", address, nil),
			"streetAddress": utils.PathSearch("streetAddress", address, nil),
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
			"costCenter":     utils.PathSearch("costCenter", enterprise, nil),
			"department":     utils.PathSearch("department", enterprise, nil),
			"division":       utils.PathSearch("division", enterprise, nil),
			"employeeNumber": utils.PathSearch("employeeNumber", enterprise, nil),
			"organization":   utils.PathSearch("organization", enterprise, nil),
			"manager":        manager,
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

	updateIdentityCenterUserOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	updateIdentityCenterUserOpt.JSONBody = utils.RemoveNil(buildUpdateIdentityCenterSCIMUserBodyParams(d))
	_, err = updateIdentityCenterUserClient.Request("PUT", updateIdentityCenterUserPath,
		&updateIdentityCenterUserOpt)
	if err != nil {
		return diag.Errorf("error updating Identity Center User: %s", err)
	}
	return resourceIdentityCenterUserRead(ctx, d, meta)
}

func buildUpdateIdentityCenterSCIMUserBodyParams(d *schema.ResourceData) map[string]interface{} {
	operations := make([]map[string]interface{}, 0)
	if d.HasChanges("family_name", "given_name") {
		updateValue := map[string]interface{}{
			"family_name": utils.ValueIgnoreEmpty(d.Get("family_name")),
			"given_name":  utils.ValueIgnoreEmpty(d.Get("given_name")),
		}
		updateValueJson, _ := json.Marshal(updateValue)
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "name",
			"attribute_value": string(updateValueJson),
		})
	}
	if d.HasChange("display_name") {
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "display_name",
			"attribute_value": d.Get("display_name"),
		})
	}
	if d.HasChange("email") {
		updateValue := []map[string]interface{}{{
			"primary": true,
			"type":    "Work",
			"value":   utils.ValueIgnoreEmpty(d.Get("email")),
		}}
		updateValueJson, _ := json.Marshal(updateValue)
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "emails",
			"attribute_value": string(updateValueJson),
		})
	}

	if d.HasChange("phone_number") {
		updateValue := []map[string]interface{}{{
			"primary": true,
			"type":    "Work",
			"value":   utils.ValueIgnoreEmpty(d.Get("phone_number")),
		}}
		updateValueJson, _ := json.Marshal(updateValue)
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "phone_numbers",
			"attribute_value": string(updateValueJson),
		})
	}

	if d.HasChange("addresses") {
		updateValue := []map[string]interface{}{{
			"country":        utils.ValueIgnoreEmpty(d.Get("addresses.0.country")),
			"region":         utils.ValueIgnoreEmpty(d.Get("addresses.0.region")),
			"locality":       utils.ValueIgnoreEmpty(d.Get("addresses.0.locality")),
			"postal_code":    utils.ValueIgnoreEmpty(d.Get("addresses.0.postal_code")),
			"street_address": utils.ValueIgnoreEmpty(d.Get("addresses.0.street_address")),
			"formatted":      utils.ValueIgnoreEmpty(d.Get("addresses.0.formatted")),
		}}
		updateValueJson, _ := json.Marshal(updateValue)
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "addresses",
			"attribute_value": string(updateValueJson),
		})
	}

	if d.HasChange("user_type") {
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "user_type",
			"attribute_value": utils.ValueIgnoreEmpty(d.Get("user_type")),
		})
	}

	if d.HasChange("title") {
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "title",
			"attribute_value": utils.ValueIgnoreEmpty(d.Get("title")),
		})
	}

	if d.HasChange("enterprise") {
		updateValue := map[string]interface{}{
			"cost_center":     utils.ValueIgnoreEmpty(d.Get("enterprise.0.cost_center")),
			"department":      utils.ValueIgnoreEmpty(d.Get("enterprise.0.department")),
			"division":        utils.ValueIgnoreEmpty(d.Get("enterprise.0.division")),
			"employee_number": utils.ValueIgnoreEmpty(d.Get("enterprise.0.employee_number")),
			"organization":    utils.ValueIgnoreEmpty(d.Get("enterprise.0.organization")),
			"manager": map[string]interface{}{
				"value": utils.ValueIgnoreEmpty(d.Get("enterprise.0.manager")),
			},
		}
		updateValueJson, _ := json.Marshal(updateValue)
		operations = append(operations, map[string]interface{}{
			"attribute_path":  "enterprise",
			"attribute_value": string(updateValueJson),
		})
	}
	return map[string]interface{}{"operations": operations}
}

func resourceIdentityCenterSCIMUserDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIdentityCenterSCIMUserImportState(_ context.Context, d *schema.ResourceData,
	_ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid id format, must be <identity_store_id>/<id>")
	}
	d.SetId(parts[1])
	mErr := multierror.Append(nil,
		d.Set("identity_store_id", parts[0]),
	)
	if err := mErr.ErrorOrNil(); err != nil {
		return nil, fmt.Errorf("failed to set value to state when import, %s", err)
	}
	return []*schema.ResourceData{d}, nil
}
