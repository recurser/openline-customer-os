package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-auth/repository/postgres/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/model"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/utils"
	tokenOauth "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/oauth2/v2"
	"log"
	"net/http"
	"strings"
)

const APP_SOURCE = "user-admin-api"

func addRegistrationRoutes(rg *gin.RouterGroup, config *config.Config, services *service.Services) {
	rg.POST("/signin", func(ginContext *gin.Context) {
		log.Printf("Sign in User")
		apiKey := ginContext.GetHeader("X-Openline-Api-Key")
		if apiKey != config.Service.ApiKey {
			ginContext.JSON(http.StatusUnauthorized, gin.H{
				"result": fmt.Sprintf("invalid api key"),
			})
			return
		}
		log.Printf("api key is valid")
		var signInRequest model.SignInRequest
		if err := ginContext.BindJSON(&signInRequest); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			ginContext.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}
		log.Printf("parsed json: %v", signInRequest)

		//get wokrspace for user

		var userInfo *oauth2.Userinfo

		if signInRequest.Provider == "google" {

			conf := &tokenOauth.Config{
				ClientID:     config.GoogleOAuth.ClientId,
				ClientSecret: config.GoogleOAuth.ClientSecret,
				Endpoint:     google.Endpoint,
			}

			token := tokenOauth.Token{
				AccessToken:  signInRequest.OAuthToken.AccessToken,
				RefreshToken: signInRequest.OAuthToken.RefreshToken,
				Expiry:       signInRequest.OAuthToken.ExpiresAt,
				TokenType:    "Bearer",
			}

			client := conf.Client(ginContext, &token)

			oauth2Service, err := oauth2.New(client)

			if err != nil {
				log.Printf("unable to create oauth2 service: %v", err.Error())
				ginContext.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to create oauth2 service: %v", err.Error()),
				})
				return
			}
			userInfoService := oauth2.NewUserinfoV2MeService(oauth2Service)

			userInfo, err = userInfoService.Get().Do()

			if err != nil {
				log.Printf("unable to get user info: %v", err.Error())
				ginContext.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to get user info: %v", err.Error()),
				})
				return
			}

		}

		var tenantName *string
		if userInfo.Hd != "" {
			tenant, err := services.CustomerOsClient.GetTenantByWorkspace(&model.WorkspaceInput{
				Name:     userInfo.Hd,
				Provider: signInRequest.Provider,
			})
			if err != nil {
				log.Printf("unable to get workspace: %v", err.Error())
				ginContext.JSON(http.StatusInternalServerError, gin.H{
					"result": fmt.Sprintf("unable to get workspace: %v", err.Error()),
				})
				return
			}
			if tenant != nil {
				log.Printf("tenant found %s", *tenant)
				var appSource = APP_SOURCE

				playerExists := false
				userExists := false

				player, err := services.CustomerOsClient.GetPlayer(signInRequest.Email, signInRequest.Provider)
				if err != nil {
					log.Printf("unable to check if player exists: %v", err.Error())
					ginContext.JSON(http.StatusInternalServerError, gin.H{
						"result": fmt.Sprintf("unable to check if player exists: %v", err.Error()),
					})
					return
				}
				if player != nil && player.Id != "" {
					playerExists = true
				}

				userByEmail, err := services.CustomerOsClient.GetUserByEmail(*tenant, signInRequest.Email)
				if err != nil {
					log.Printf("unable to get user: %v", err.Error())
					ginContext.JSON(http.StatusInternalServerError, gin.H{
						"result": fmt.Sprintf("unable to get user: %v", err.Error()),
					})
					return
				}
				if userByEmail != nil && userByEmail.ID != "" {
					userExists = true
				}

				if !playerExists && !userExists {
					userByEmail, err = services.CustomerOsClient.CreateUser(&model.UserInput{
						FirstName: userInfo.GivenName,
						LastName:  userInfo.FamilyName,
						Email: model.EmailInput{
							Email:     signInRequest.Email,
							Primary:   true,
							AppSource: &appSource,
						},
						Player: model.PlayerInput{
							IdentityId: signInRequest.OAuthToken.ProviderAccountId,
							AuthId:     signInRequest.Email,
							Provider:   signInRequest.Provider,
							AppSource:  &appSource,
						},
						AppSource: &appSource,
					}, *tenant, []model.Role{model.RoleUser, model.RoleOwner})
					if err != nil {
						log.Printf("unable to create user: %v", err.Error())
						ginContext.JSON(http.StatusInternalServerError, gin.H{
							"result": fmt.Sprintf("unable to create user: %v", err.Error()),
						})
						return
					}
				} else {
					if !playerExists {
						err = services.CustomerOsClient.CreatePlayer(*tenant, userByEmail.ID, signInRequest.OAuthToken.ProviderAccountId, signInRequest.Email, signInRequest.Provider)
						if err != nil {
							log.Printf("unable to create player: %v", err.Error())
							ginContext.JSON(http.StatusInternalServerError, gin.H{
								"result": fmt.Sprintf("unable to create player: %v", err.Error()),
							})
							return
						}
					}
				}

				addDefaultMissingRoles(services, userByEmail, tenant, ginContext)

				tenantName = tenant
			} else {
				var appSource = APP_SOURCE
				tenantStr := utils.Sanitize(userInfo.Hd)
				log.Printf("tenant not found for workspace, creating new tenant %s", tenantStr)
				// Workspace is not mapped to a tenant create a new tenant and map it to the workspace
				id, failed := makeTenantAndUser(ginContext, services.CustomerOsClient, tenantStr, appSource, signInRequest, userInfo)
				if failed {
					return
				}
				log.Printf("user created: %s", id)
				tenantName = &tenantStr
			}
		} else {
			// no workspace for this e-mail
			// check tenant exists for this e-mail
			if userInfo.Email != "" {
				var err error
				tenantName, err = services.CustomerOsClient.GetTenantByUserEmail(userInfo.Email)
				if err != nil {
					log.Printf("unable to get tenant: %v", err.Error())
					ginContext.JSON(http.StatusInternalServerError, gin.H{
						"result": fmt.Sprintf("unable to get tenant: %v", err.Error()),
					})
					return
				}
			}
			// no tenant for this e-mail, invent a tenant name
			if tenantName == nil {
				var appSource = APP_SOURCE
				tenantStr := utils.GenerateName()
				log.Printf("user has no workspace, inventing tenant %s", tenantStr)

				id, failed := makeTenantAndUser(ginContext, services.CustomerOsClient, tenantStr, appSource, signInRequest, userInfo)
				if failed {
					return
				}
				log.Printf("user created: %s", id)
				tenantName = &tenantStr
			}
		}

		if isRequestEnablingOAuthSync(signInRequest) {
			//TODO Move this logic to a service
			var oauthToken, _ = services.AuthServices.OAuthTokenService.GetByPlayerIdAndProvider(signInRequest.OAuthToken.ProviderAccountId, signInRequest.Provider)
			if oauthToken == nil {
				oauthToken = &entity.OAuthTokenEntity{}
			}
			oauthToken.Provider = signInRequest.Provider
			oauthToken.TenantName = *tenantName
			oauthToken.PlayerIdentityId = signInRequest.OAuthToken.ProviderAccountId
			oauthToken.EmailAddress = signInRequest.Email
			oauthToken.AccessToken = signInRequest.OAuthToken.AccessToken
			oauthToken.RefreshToken = signInRequest.OAuthToken.RefreshToken
			oauthToken.IdToken = signInRequest.OAuthToken.IdToken
			oauthToken.ExpiresAt = signInRequest.OAuthToken.ExpiresAt
			oauthToken.Scope = signInRequest.OAuthToken.Scope
			oauthToken.NeedsManualRefresh = false
			if isRequestEnablingGmailSync(signInRequest) {
				oauthToken.GmailSyncEnabled = true
			}
			if isRequestEnablingGoogleCalendarSync(signInRequest) {
				oauthToken.GoogleCalendarSyncEnabled = true
			}
			services.AuthServices.OAuthTokenService.Save(*oauthToken)
		}
		ginContext.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	rg.POST("/google/revoke", func(ginContext *gin.Context) {
		log.Printf("revoke oauth token")

		apiKey := ginContext.GetHeader("X-Openline-Api-Key")
		if apiKey != config.Service.ApiKey {
			ginContext.JSON(http.StatusUnauthorized, gin.H{
				"result": fmt.Sprintf("invalid api key"),
			})
			return
		}

		var revokeRequest model.RevokeRequest
		if err := ginContext.BindJSON(&revokeRequest); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			ginContext.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}
		log.Printf("parsed json: %v", revokeRequest)

		var oauthToken, _ = services.AuthServices.OAuthTokenService.GetByPlayerIdAndProvider(revokeRequest.ProviderAccountId, "google")

		var resp *http.Response
		var err error

		if oauthToken.RefreshToken != "" {
			url := fmt.Sprintf("https://accounts.google.com/o/oauth2/revoke?token=%s", oauthToken.RefreshToken)
			resp, err = http.Get(url)
			if err != nil {
				ginContext.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		if resp == nil || resp.StatusCode == 200 {
			err := services.AuthServices.OAuthTokenService.DeleteByPlayerIdAndProvider(revokeRequest.ProviderAccountId, "google")
			if err != nil {
				ginContext.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		} else {
			if resp != nil && resp.StatusCode != 200 {
				ginContext.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		ginContext.JSON(http.StatusOK, gin.H{})
	})
}

func isRequestEnablingGmailSync(signInRequest model.SignInRequest) bool {
	if strings.Contains(signInRequest.OAuthToken.Scope, "gmail") {
		return true
	}
	return false
}

func isRequestEnablingGoogleCalendarSync(signInRequest model.SignInRequest) bool {
	if strings.Contains(signInRequest.OAuthToken.Scope, "calendar") {
		return true
	}
	return false
}

func isRequestEnablingOAuthSync(signInRequest model.SignInRequest) bool {
	if isRequestEnablingGmailSync(signInRequest) || isRequestEnablingGoogleCalendarSync(signInRequest) {
		return true
	}
	return false
}

func makeTenantAndUser(c *gin.Context, cosClient service.CustomerOsClient, tenantStr string, appSource string, req model.SignInRequest, userInfo *oauth2.Userinfo) (string, bool) {
	newTenantStr, err := cosClient.MergeTenant(&model.TenantInput{
		Name:      tenantStr,
		AppSource: &appSource})
	if err != nil {
		log.Printf("unable to create tenant: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("unable to create tenant: %v", err.Error()),
		})
		return "", true
	}

	if userInfo.Hd != "" {
		mergeWorkspaceRes, err := cosClient.MergeTenantToWorkspace(&model.WorkspaceInput{
			Name:      userInfo.Hd,
			Provider:  req.Provider,
			AppSource: &appSource,
		}, newTenantStr)

		if err != nil {
			log.Printf("unable to merge workspace: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to merge workspace: %v", err.Error()),
			})
			return "", true
		}
		if !mergeWorkspaceRes {
			log.Printf("unable to merge workspace")
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to merge workspace"),
			})
			return "", true
		}
	}

	user, err := cosClient.CreateUser(&model.UserInput{
		FirstName: userInfo.GivenName,
		LastName:  userInfo.FamilyName,
		Email: model.EmailInput{
			Email:     req.Email,
			Primary:   true,
			AppSource: &appSource,
		},
		Player: model.PlayerInput{
			IdentityId: req.OAuthToken.ProviderAccountId,
			AuthId:     req.Email,
			Provider:   req.Provider,
			AppSource:  &appSource,
		},
		AppSource: &appSource,
	}, newTenantStr, []model.Role{model.RoleUser, model.RoleOwner})
	if err != nil {
		log.Printf("unable to create user: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("unable to create user: %v", err.Error()),
		})
		return "", true
	}
	return user.ID, false
}

func addDefaultMissingRoles(services *service.Services, user *model.UserResponse, tenant *string, ginContext *gin.Context) {
	var rolesToAdd []model.Role

	if user.Roles == nil || len(*user.Roles) == 0 {
		rolesToAdd = []model.Role{model.RoleUser, model.RoleOwner}
	} else {
		userRoleFound := false
		ownerRoleFound := false
		for _, role := range *user.Roles {
			if role == model.RoleUser {
				userRoleFound = true
			}
			if role == model.RoleOwner {
				ownerRoleFound = true
			}
		}
		if !userRoleFound {
			rolesToAdd = append(rolesToAdd, model.RoleUser)
		}
		if !ownerRoleFound {
			rolesToAdd = append(rolesToAdd, model.RoleOwner)
		}
	}

	if len(rolesToAdd) > 0 {
		_, err := services.CustomerOsClient.AddUserRoles(*tenant, user.ID, rolesToAdd)
		if err != nil {
			log.Printf("unable to add role: %v", err.Error())
			ginContext.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to add role: %v", err.Error()),
			})
			return
		}
	}
}
