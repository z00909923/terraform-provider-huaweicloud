// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product IdentityCenter
// ---------------------------------------------------------------

package identitycenter

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API IdentityStore POST /v1/tokens
func ResourceIdentityCenterDeviceToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdentityCenterDeviceTokenCreate,
		ReadContext:   resourceIdentityCenterDeviceTokenRead,
		DeleteContext: resourceIdentityCenterDeviceTokenDelete,
		Description:   "schema: Internal",
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"grant_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"redirect_uri": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"refresh_token": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"access_token": {
				Type:     schema.TypeList,
				Computed: true,
			},
			"expires_in": {
				Type:     schema.TypeList,
				Computed: true,
			},
		},
	}
}

func resourceIdentityCenterDeviceTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createIdentityCenterClient: create IdentityCenter client
	var (
		createIdentityCenterDeviceTokenHttpUrl = "/v1/tokens"
		createIdentityCenterDeviceTokenProduct = "identityoidc"
	)
	createIdentityCenterDeviceTokenClient, err := cfg.NewServiceClient(createIdentityCenterDeviceTokenProduct, region)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	createIdentityCenterDeviceTokenPath := createIdentityCenterDeviceTokenClient.Endpoint + createIdentityCenterDeviceTokenHttpUrl
	createIdentityCenterDeviceTokenOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	log.Println(createIdentityCenterDeviceTokenOpt)
	createIdentityCenterDeviceTokenOpt.JSONBody = utils.RemoveNil(buildCreateIdentityCenterDeviceTokenBodyParams(d))
	createIdentityCenterClientResp, err := createIdentityCenterDeviceTokenClient.Request("POST",
		createIdentityCenterDeviceTokenPath, &createIdentityCenterDeviceTokenOpt)
	if err != nil {
		return diag.Errorf("error creating Identity Center Client: %s", err)
	}

	createIdentityCenterDeviceTokenRespBody, err := utils.FlattenResponse(createIdentityCenterClientResp)
	if err != nil {
		return diag.FromErr(err)
	}

	token := utils.PathSearch("access_token", createIdentityCenterDeviceTokenRespBody, "").(string)
	if token == "" {
		return diag.Errorf("unable to find the Identity Center access_token from the API response")
	}
	d.SetId(token)

	return nil
}

func buildCreateIdentityCenterDeviceTokenBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"client_id":     utils.ValueIgnoreEmpty(d.Get("client_id")),
		"client_secret": utils.ValueIgnoreEmpty(d.Get("client_secret")),
		"code":          utils.ValueIgnoreEmpty(d.Get("code")),
		"device_code":   utils.ValueIgnoreEmpty(d.Get("device_code")),
		"grant_type":    utils.ValueIgnoreEmpty(d.Get("grant_type")),
		"redirect_uri":  utils.ValueIgnoreEmpty(d.Get("redirect_uri")),
		"refresh_token": utils.ValueIgnoreEmpty(d.Get("refresh_token")),
		"scopes":        d.Get("scopes"),
	}
	return bodyParams
}

func resourceIdentityCenterDeviceTokenRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIdentityCenterDeviceTokenDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting token is not supported."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}
