package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryResolver_Organizations_FilterByNameLike(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	neo4jt.CreateOrganization(driver, tenantName, "A closed organization")
	neo4jt.CreateOrganization(driver, tenantName, "OPENLINE")
	neo4jt.CreateOrganization(driver, tenantName, "the openline")
	neo4jt.CreateOrganization(driver, tenantName, "some other open organization")
	neo4jt.CreateOrganization(driver, tenantName, "OpEnLiNe")

	require.Equal(t, 5, neo4jt.GetCountOfNodes(driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("get_organizations"),
		client.Var("page", 1),
		client.Var("limit", 3),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizations struct {
		Organizations model.OrganizationPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizations)
	require.Nil(t, err)
	require.NotNil(t, organizations)
	pagedOrganizations := organizations.Organizations
	require.Equal(t, 2, pagedOrganizations.TotalPages)
	require.Equal(t, int64(4), pagedOrganizations.TotalElements)
	require.Equal(t, "OPENLINE", pagedOrganizations.Content[0].Name)
	require.Equal(t, "OpEnLiNe", pagedOrganizations.Content[1].Name)
	require.Equal(t, "some other open organization", pagedOrganizations.Content[2].Name)
}

func TestQueryResolver_Organization(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organizationInput := entity.OrganizationEntity{
		Name:        "Organization name",
		Description: "Organization description",
		Domain:      "Organization domain",
		Website:     "Organization_website.com",
		Industry:    "tech",
		IsPublic:    true,
	}
	organizationId1 := neo4jt.CreateFullOrganization(driver, tenantName, organizationInput)
	neo4jt.CreateOrganization(driver, tenantName, "otherOrganization")

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("get_organization_by_id"),
		client.Var("organizationId", organizationId1),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organization struct {
		Organization model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)
	require.Equal(t, organizationId1, organization.Organization.ID)
	require.Equal(t, organizationInput.Name, organization.Organization.Name)
	require.Equal(t, organizationInput.Description, *organization.Organization.Description)
	require.Equal(t, organizationInput.Domain, *organization.Organization.Domain)
	require.Equal(t, organizationInput.Website, *organization.Organization.Website)
	require.Equal(t, organizationInput.IsPublic, *organization.Organization.IsPublic)
	require.Equal(t, organizationInput.Industry, *organization.Organization.Industry)
	require.NotNil(t, organization.Organization.CreatedAt)
}

func TestQueryResolver_Organizations_WithAddresses(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organization1 := neo4jt.CreateOrganization(driver, tenantName, "OPENLINE")
	organization2 := neo4jt.CreateOrganization(driver, tenantName, "some other organization")
	addressInput := entity.PlaceEntity{
		Source:   entity.DataSourceOpenline,
		Country:  "testCountry",
		State:    "testState",
		City:     "testCity",
		Address:  "testAddress",
		Address2: "testAddress2",
		Zip:      "testZip",
		Phone:    "testPhone",
		Fax:      "testFax",
	}
	address1 := neo4jt.CreateAddress(driver, addressInput)
	address2 := neo4jt.CreateAddress(driver, entity.PlaceEntity{
		Source: "manual",
	})
	neo4jt.OrganizationHasAddress(driver, organization1, address1)
	neo4jt.OrganizationHasAddress(driver, organization2, address2)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Address"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "LOCATED_AT"))

	rawResponse, err := c.RawPost(getQuery("get_organizations_with_addresses"),
		client.Var("page", 1),
		client.Var("limit", 3),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizations struct {
		Organizations model.OrganizationPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizations)
	require.Nil(t, err)
	require.NotNil(t, organizations)
	pagedOrganizations := organizations.Organizations
	require.Equal(t, int64(1), pagedOrganizations.TotalElements)
	require.Equal(t, 1, len(organizations.Organizations.Content[0].Addresses))
	address := organizations.Organizations.Content[0].Addresses[0]
	require.Equal(t, address1, address.ID)
	require.Equal(t, model.DataSourceOpenline, *address.Source)
	require.Equal(t, addressInput.Country, *address.Country)
	require.Equal(t, addressInput.City, *address.City)
	require.Equal(t, addressInput.State, *address.State)
	require.Equal(t, addressInput.Address, *address.Address)
	require.Equal(t, addressInput.Address2, *address.Address2)
	require.Equal(t, addressInput.Fax, *address.Fax)
	require.Equal(t, addressInput.Phone, *address.Phone)
	require.Equal(t, addressInput.Zip, *address.Zip)
}

func TestQueryResolver_Organization_WithNotes_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organizationId := neo4jt.CreateOrganization(driver, tenantName, "test org")
	userId := neo4jt.CreateDefaultUserWithId(driver, tenantName, testUserId)
	noteId1 := neo4jt.CreateNoteForOrganization(driver, tenantName, organizationId, "note1")
	noteId2 := neo4jt.CreateNoteForOrganization(driver, tenantName, organizationId, "note2")
	neo4jt.NoteCreatedByUser(driver, noteId1, userId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Note"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "NOTED"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "CREATED"))

	rawResponse, err := c.RawPost(getQuery("get_organization_with_notes_by_id"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)

	notes := searchedOrganization.Organization.Notes.Content
	require.Equal(t, 2, len(notes))
	var noteWithUser, noteWithoutUser *model.Note
	if noteId1 == notes[0].ID {
		noteWithUser = notes[0]
		noteWithoutUser = notes[1]
	} else {
		noteWithUser = notes[1]
		noteWithoutUser = notes[0]
	}
	require.Equal(t, noteId1, noteWithUser.ID)
	require.Equal(t, "note1", noteWithUser.HTML)
	require.NotNil(t, noteWithUser.CreatedAt)
	require.NotNil(t, noteWithUser.CreatedBy)
	require.Equal(t, userId, noteWithUser.CreatedBy.ID)
	require.Equal(t, "first", noteWithUser.CreatedBy.FirstName)
	require.Equal(t, "last", noteWithUser.CreatedBy.LastName)

	require.Equal(t, noteId2, noteWithoutUser.ID)
	require.Equal(t, "note2", noteWithoutUser.HTML)
	require.NotNil(t, noteWithoutUser.CreatedAt)
	require.Nil(t, noteWithoutUser.CreatedBy)
}

func TestMutationResolver_OrganizationCreate(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organizationTypeId := neo4jt.CreateOrganizationType(driver, tenantName, "COMPANY")

	// Ensure that the tenant and organization type nodes were created in the database.
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "OrganizationType"))
	require.Equal(t, 2, neo4jt.GetTotalCountOfNodes(driver))

	// Call the "create_organization" mutation.
	rawResponse, err := c.RawPost(getQuery("create_organization"),
		client.Var("organizationTypeId", organizationTypeId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the "organization" struct.
	var organization struct {
		Organization_Create model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)

	// Assign the organization to a shorter variable for easier reference.
	org := organization.Organization_Create

	// Ensure that the organization was created correctly.
	require.NotNil(t, org.ID)
	require.NotNil(t, org.CreatedAt)
	require.NotNil(t, org.UpdatedAt)
	require.Equal(t, "organization name", org.Name)
	require.Equal(t, "organization description", *org.Description)
	require.Equal(t, "organization domain", *org.Domain)
	require.Equal(t, "organization website", *org.Website)
	require.Equal(t, "organization industry", *org.Industry)
	require.Equal(t, true, *org.IsPublic)
	require.Equal(t, organizationTypeId, org.OrganizationType.ID)
	require.Equal(t, "COMPANY", org.OrganizationType.Name)
	require.Equal(t, model.DataSourceOpenline, org.Source)
	require.Equal(t, model.DataSourceOpenline, org.SourceOfTruth)
	require.Equal(t, "test", org.AppSource)

	// Check the number of nodes and relationships in the Neo4j database
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization_"+tenantName))

	// Check the labels on the nodes in the Neo4j database
	assertNeo4jLabels(t, driver, []string{"Tenant", "OrganizationType", "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationUpdate(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organizationId := neo4jt.CreateOrganization(driver, tenantName, "some organization")
	organizationTypeIdOrig := neo4jt.CreateOrganizationType(driver, tenantName, "ORIG")
	organizationTypeIdUpdate := neo4jt.CreateOrganizationType(driver, tenantName, "UPDATED")
	neo4jt.SetContactTypeForContact(driver, organizationTypeIdOrig, organizationTypeIdOrig)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "OrganizationType"))

	rawResponse, err := c.RawPost(getQuery("update_organization"),
		client.Var("organizationId", organizationId),
		client.Var("organizationTypeId", organizationTypeIdUpdate))
	assertRawResponseSuccess(t, rawResponse, err)

	var organization struct {
		Organization_Update model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)
	require.Equal(t, organizationId, organization.Organization_Update.ID)
	require.NotNil(t, organization.Organization_Update.CreatedAt)
	require.NotNil(t, organization.Organization_Update.UpdatedAt)
	require.Equal(t, "updated name", organization.Organization_Update.Name)
	require.Equal(t, "updated description", *organization.Organization_Update.Description)
	require.Equal(t, "updated domain", *organization.Organization_Update.Domain)
	require.Equal(t, "updated website", *organization.Organization_Update.Website)
	require.Equal(t, "updated industry", *organization.Organization_Update.Industry)
	require.Equal(t, true, *organization.Organization_Update.IsPublic)
	require.Equal(t, organizationTypeIdUpdate, organization.Organization_Update.OrganizationType.ID)
	require.Equal(t, "UPDATED", organization.Organization_Update.OrganizationType.Name)
	require.Equal(t, model.DataSourceOpenline, organization.Organization_Update.SourceOfTruth)

	// Check still single organization node exists after update, no new node created
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))
}

func TestMutationResolver_OrganizationDelete(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	organizationId := neo4jt.CreateOrganization(driver, tenantName, "LLC LLC")
	addressId := neo4jt.CreateAddress(driver, entity.PlaceEntity{
		Source: "manual",
	})
	neo4jt.OrganizationHasAddress(driver, organizationId, addressId)

	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "LOCATED_AT"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Address"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("delete_organization"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var result struct {
		Organization_Delete model.Result
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &result)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, true, result.Organization_Delete.Result)

	require.Equal(t, 0, neo4jt.GetCountOfRelationships(driver, "LOCATED_AT"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "Address"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "Organization"))

	assertNeo4jLabels(t, driver, []string{"Tenant"})
}

func TestQueryResolver_Organization_WithRoles_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId1 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, tenantName)
	organizationId := neo4jt.CreateOrganization(driver, tenantName, "some organization")
	role1 := neo4jt.ContactWorksForOrganization(driver, contactId1, organizationId, "CTO", false)
	role2 := neo4jt.ContactWorksForOrganization(driver, contactId2, organizationId, "CEO", true)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "JobRole"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "ROLE_IN"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "WORKS_AS"))

	rawResponse, err := c.RawPost(getQuery("get_organization_with_job_roles_by_id"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)

	roles := searchedOrganization.Organization.JobRoles
	require.Equal(t, 2, len(roles))
	var cto, ceo *model.JobRole
	ceo = roles[0]
	cto = roles[1]
	require.Equal(t, role1, cto.ID)
	require.Equal(t, "CTO", *cto.JobTitle)
	require.Equal(t, false, cto.Primary)
	require.Equal(t, contactId1, cto.Contact.ID)

	require.Equal(t, role2, ceo.ID)
	require.Equal(t, "CEO", *ceo.JobTitle)
	require.Equal(t, true, ceo.Primary)
	require.Equal(t, contactId2, ceo.Contact.ID)
}

func TestQueryResolver_Organization_WithContacts_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	organizationId := neo4jt.CreateOrganization(driver, tenantName, "organization1")
	organizationId2 := neo4jt.CreateOrganization(driver, tenantName, "organization2")
	contactId1 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId3 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId4 := neo4jt.CreateDefaultContact(driver, tenantName)
	neo4jt.LinkContactWithOrganization(driver, contactId1, organizationId)
	neo4jt.LinkContactWithOrganization(driver, contactId2, organizationId)
	neo4jt.LinkContactWithOrganization(driver, contactId3, organizationId)
	neo4jt.LinkContactWithOrganization(driver, contactId4, organizationId2)

	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 4, neo4jt.GetCountOfRelationships(driver, "CONTACT_OF"))

	rawResponse, err := c.RawPost(getQuery("get_organization_with_contacts_by_id"),
		client.Var("organizationId", organizationId),
		client.Var("limit", 1),
		client.Var("page", 1),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)
	require.Equal(t, 3, searchedOrganization.Organization.Contacts.TotalPages)
	require.Equal(t, int64(3), searchedOrganization.Organization.Contacts.TotalElements)

	contacts := searchedOrganization.Organization.Contacts.Content
	require.Equal(t, 1, len(contacts))
	require.Equal(t, contactId1, contacts[0].ID)
}