package slack

import (
	"encoding/json"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const unknownUserName = "Unknown User"

func MapUser(inputJson string) (string, error) {
	var input struct {
		ID       string `json:"id,omitempty"`
		Bot      bool   `json:"is_bot,omitempty"`
		App      bool   `json:"is_app_user,omitempty"`
		Deleted  bool   `json:"deleted,omitempty"`
		Admin    bool   `json:"is_admin,omitempty"`
		Timezone string `json:"tz,omitempty"`
		TeamId   string `json:"team_id,omitempty"`
		Profile  struct {
			Email     string `json:"email,omitempty"`
			Phone     string `json:"phone,omitempty"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Name      string `json:"real_name_normalized,omitempty"`
			Image192  string `json:"image_192,omitempty"`
		} `json:"profile"`
		OpenlineFields struct {
			OrganizationId string `json:"organization_id,omitempty"`
			TenantDomain   string `json:"tenant_domain,omitempty"`
			TenantTeamId   string `json:"tenant_team_id,omitempty"`
		} `json:"openline_fields,omitempty"`
	}

	if err := json.Unmarshal([]byte(inputJson), &input); err != nil {
		return "", err
	}

	if input.Bot || input.App {
		output := entity.UserData{
			BaseData: entity.BaseData{
				Skip:       true,
				SkipReason: "User is a bot or app",
			},
		}
		return utils.ToJson(output)
	}
	if input.Deleted {
		output := entity.UserData{
			BaseData: entity.BaseData{
				Skip:       true,
				SkipReason: "User is deleted",
			},
		}
		return utils.ToJson(output)
	}
	if !slackUserIsTenantUser(input.Profile.Email, input.TeamId, input.OpenlineFields.TenantDomain, input.OpenlineFields.TenantTeamId) {
		output := entity.UserData{
			BaseData: entity.BaseData{

				Skip:       true,
				SkipReason: "Slack user is not a tenant user",
			},
		}
		return utils.ToJson(output)
	}

	output := entity.UserData{
		BaseData: entity.BaseData{
			ExternalId: input.ID,
		},
		Email: input.Profile.Email,
		PhoneNumbers: []entity.PhoneNumber{{
			Number:  input.Profile.Phone,
			Primary: true}},
		FirstName: input.Profile.FirstName,
		LastName:  input.Profile.LastName,
		Name:      input.Profile.Name,
		Timezone:  input.Timezone,
	}
	if !strings.HasPrefix(input.Profile.Image192, "https://secure.gravatar.com") {
		output.ProfilePhotoUrl = input.Profile.Image192
	}

	return utils.ToJson(output)
}

func MapContact(inputJson string) (string, error) {
	var input struct {
		ID       string `json:"id,omitempty"`
		Bot      bool   `json:"is_bot,omitempty"`
		App      bool   `json:"is_app_user,omitempty"`
		Deleted  bool   `json:"deleted,omitempty"`
		Admin    bool   `json:"is_admin,omitempty"`
		Timezone string `json:"tz,omitempty"`
		TeamId   string `json:"team_id,omitempty"`
		Profile  struct {
			Email     string `json:"email,omitempty"`
			Phone     string `json:"phone,omitempty"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Name      string `json:"real_name_normalized,omitempty"`
			Image192  string `json:"image_192,omitempty"`
		} `json:"profile"`
		OpenlineFields struct {
			OrganizationId string `json:"organization_id,omitempty"`
			TenantDomain   string `json:"tenant_domain,omitempty"`
			TenantTeamId   string `json:"tenant_team_id,omitempty"`
		} `json:"openline_fields,omitempty"`
	}

	if err := json.Unmarshal([]byte(inputJson), &input); err != nil {
		return "", err
	}

	if input.Bot || input.App {
		output := entity.BaseData{
			Skip:       true,
			SkipReason: "User is a bot or app",
		}
		return utils.ToJson(output)
	}
	if input.Deleted {
		output := entity.BaseData{
			Skip:       true,
			SkipReason: "User is deleted",
		}
		return utils.ToJson(output)
	}
	if slackUserIsTenantUser(input.Profile.Email, input.TeamId, input.OpenlineFields.TenantDomain, input.OpenlineFields.TenantTeamId) {
		output := entity.BaseData{
			Skip:       true,
			SkipReason: "Slack user is not a contact",
		}
		return utils.ToJson(output)
	}

	output := entity.ContactData{
		BaseData: entity.BaseData{
			ExternalId: input.ID,
		},
		Email:        input.Profile.Email,
		PhoneNumbers: []entity.PhoneNumber{{Number: input.Profile.Phone, Primary: true}},
		FirstName:    input.Profile.FirstName,
		LastName:     input.Profile.LastName,
		Name:         input.Profile.Name,
		Timezone:     input.Timezone,
	}
	output.Organizations = append(output.Organizations, entity.ReferencedOrganization{
		Id: input.OpenlineFields.OrganizationId,
	})
	if !strings.HasPrefix(input.Profile.Image192, "https://secure.gravatar.com") {
		output.ProfilePhotoUrl = input.Profile.Image192
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(outputJson), nil
}

type OutputContent struct {
	Text   string `json:"text,omitempty"`
	Blocks []any  `json:"blocks,omitempty"`
}

func MapInteractionEvent(inputJson string) (string, error) {
	var input struct {
		Ts             string `json:"ts,omitempty"`
		Type           string `json:"type,omitempty"`
		SenderUser     string `json:"user,omitempty"`
		Text           string `json:"text,omitempty"`
		ThreadTs       string `json:"thread_ts,omitempty"`
		Blocks         []any  `json:"blocks,omitempty"`
		OpenlineFields struct {
			OrganizationId string            `json:"organization_id,omitempty"`
			UserIds        []string          `json:"channel_user_ids,omitempty"`
			UserNamesById  map[string]string `json:"channel_user_names,omitempty"`
			ChannelId      string            `json:"channel_id,omitempty"`
			ChannelName    string            `json:"channel_name,omitempty"`
			Permalink      string            `json:"permalink"`
		} `json:"openline_fields,omitempty"`
	}

	if err := json.Unmarshal([]byte(inputJson), &input); err != nil {
		return "", err
	}

	output := entity.InteractionEventData{
		BaseData: entity.BaseData{
			ExternalId:   input.OpenlineFields.ChannelId + "/" + input.Ts,
			CreatedAtStr: tsStrToRFC3339Nanos(input.Ts),
			ExternalUrl:  input.OpenlineFields.Permalink,
		},
		ContentType: "plain/text",
		Type:        "MESSAGE",
		Channel:     "CHAT",
	}

	// Do not use blocks in content for now.

	//outputContent := OutputContent{
	//	Text:   replaceUserMentionsInText(input.Text, input.OpenlineFields.UserNamesById),
	//	Blocks: addUserNameInBlocks(input.Blocks, input.OpenlineFields.UserNamesById),
	//}
	//outputContentJson, err := json.Marshal(outputContent)
	//if err != nil {
	//	return "", err
	//}
	//output.Content = string(outputContentJson)

	output.Content = replaceUserMentionsInText(input.Text, input.OpenlineFields.UserNamesById)
	if output.Content == "" {
		output.Skip = true
		output.SkipReason = "Empty text message"
		return utils.ToJson(output)
	}

	output.SentBy = entity.InteractionEventParticipant{
		ReferencedUser: entity.ReferencedUser{
			ExternalId: input.SenderUser,
		},
		ReferencedJobRole: entity.ReferencedJobRole{
			ReferencedOrganization: entity.ReferencedOrganization{
				Id: input.OpenlineFields.OrganizationId,
			},
			ReferencedContact: entity.ReferencedContact{
				ExternalId: input.SenderUser,
			},
		},
	}
	output.SessionDetails.Channel = "CHAT"
	output.SessionDetails.Type = "THREAD"
	output.SessionDetails.Status = "ACTIVE"
	output.SessionDetails.Name = input.OpenlineFields.ChannelName
	if input.ThreadTs != "" {
		if input.ThreadTs != input.Ts {
			output.Hide = true
		}
		output.SessionDetails.ExternalId = "session/" + input.OpenlineFields.ChannelId + "/" + input.ThreadTs
		output.SessionDetails.CreatedAtStr = tsStrToRFC3339Nanos(input.ThreadTs)
		output.SessionDetails.Identifier = input.OpenlineFields.ChannelId + "/" + input.ThreadTs
	} else {
		output.SessionDetails.ExternalId = "session/" + input.OpenlineFields.ChannelId + "/" + input.Ts
		output.SessionDetails.CreatedAtStr = tsStrToRFC3339Nanos(input.Ts)
		output.SessionDetails.Identifier = input.OpenlineFields.ChannelId + "/" + input.Ts
	}

	for _, user := range input.OpenlineFields.UserIds {
		if user != input.SenderUser {
			output.SentTo = append(output.SentTo,
				entity.InteractionEventParticipant{
					ReferencedUser: entity.ReferencedUser{
						ExternalId: user,
					},
					ReferencedJobRole: entity.ReferencedJobRole{
						ReferencedOrganization: entity.ReferencedOrganization{
							Id: input.OpenlineFields.OrganizationId,
						},
						ReferencedContact: entity.ReferencedContact{
							ExternalId: user,
						},
					},
				})
		}
	}

	output.SentTo = append(output.SentTo,
		entity.InteractionEventParticipant{
			ReferencedOrganization: entity.ReferencedOrganization{
				Id: input.OpenlineFields.OrganizationId,
			},
		})

	if input.Type != "message" {
		output.Skip = true
		output.SkipReason = "Not a message type. Type: " + input.Type
	}

	return utils.ToJson(output)
}

func tsStrToRFC3339Nanos(ts string) string {
	parts := strings.Split(ts, ".")
	secs, _ := strconv.ParseInt(parts[0], 10, 64)
	millis, _ := strconv.ParseInt(parts[1], 10, 64)
	t := time.Unix(secs, millis*1000).UTC()
	layout := "2006-01-02T15:04:05.000000Z"
	return t.Format(layout)
}

func replaceUserMentionsInText(text string, userNames map[string]string) string {
	re := regexp.MustCompile("<@(U[A-Z0-9]+)>")
	replaced := re.ReplaceAllStringFunc(text, func(mention string) string {
		id := mention[2 : len(mention)-1]
		name, ok := userNames[id]
		if !ok || name == "" {
			return unknownUserName
		}
		return markdownUserName(name)
	})
	return replaced
}

func markdownUserName(name string) string {
	// Replace spaces with underscores and add "@" at the beginning
	formattedName := "@" + strings.ReplaceAll(name, " ", "_")
	return formattedName
}

func addUserNameInBlocks(blocks []any, userNamesById map[string]string) []any {
	for _, block := range blocks {
		blockMap, ok := block.(map[string]any)
		if !ok {
			continue
		}

		if elements, exists := blockMap["elements"]; exists {
			elementsSlice, ok := elements.([]any)
			if !ok {
				continue
			}

			for _, element := range elementsSlice {
				elementMap, ok := element.(map[string]any)
				if !ok {
					continue
				}

				if innerElements, exists := elementMap["elements"]; exists {
					innerElementsSlice, ok := innerElements.([]any)
					if !ok {
						continue
					}
					for _, innerElement := range innerElementsSlice {
						innerElementMap, ok := innerElement.(map[string]any)
						if !ok {
							continue
						}
						if innerElementMap["type"] == "user" {
							userID := innerElementMap["user_id"].(string)
							if userName, exists := userNamesById[userID]; exists {
								innerElementMap["user_name"] = userName
							} else {
								innerElementMap["user_name"] = unknownUserName
							}
						}
					}
				}
			}
		}
	}
	return blocks
}

func slackUserIsTenantUser(email, userTeamId, tenantDomain, tenantTeamId string) bool {
	if tenantTeamId != "" && userTeamId == tenantTeamId {
		return true
	} else if tenantDomain != "" && strings.HasSuffix(email, "@"+tenantDomain) {
		return true
	}
	return false
}
