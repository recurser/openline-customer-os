package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"testing"
)

func TestMutationResolver_EmailMergeToContact(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact
	contactId := neo4jt.CreateDefaultContact(ctx, driver, tenantName)

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/merge_email_to_contact"),
		client.Var("contactId", contactId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var email struct {
		EmailMergeToContact model.Email
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &email)
	require.Nil(t, err, "Error unmarshalling response data")

	e := email.EmailMergeToContact

	// Check that the fields of the email struct have the expected values
	require.NotNil(t, e.ID, "Email ID is nil")
	require.NotNil(t, e.CreatedAt, "Missing createdAt field")
	require.NotNil(t, e.UpdatedAt, "Missing updatedAt field")
	require.Equal(t, true, e.Primary, "Email Primary field is not true")
	require.Equal(t, "test@gmail.com", *e.Email)
	require.Equal(t, "test@gmail.com", *e.RawEmail)
	require.False(t, *e.Validated)
	if e.Label == nil {
		t.Errorf("Email Label field is nil")
	} else {
		require.Equal(t, model.EmailLabelWork, *e.Label, "Email Label field is not expected value")
	}
	require.Equal(t, model.DataSourceOpenline, e.Source, "Email Source field is not expected value")
	require.Equal(t, model.DataSourceOpenline, e.SourceOfTruth, "Email Source of truth field is not expected value")
	require.Equal(t, "test", e.AppSource, "Email App source field is not expected value")

	// Check the number of nodes and relationships in the Neo4j database
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact"), "Incorrect number of Contact nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"), "Incorrect number of Email nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName), "Incorrect number of Email_%s nodes in Neo4j", tenantName)
	require.Equal(t, 3, neo4jt.GetTotalCountOfNodes(ctx, driver), "Incorrect total number of nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"), "Incorrect number of HAS relationships in Neo4j")

	// Check the labels on the nodes in the Neo4j database
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "Email", "Email_" + tenantName})
}

func TestMutationResolver_EmailUpdateInContact(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	contactId := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	emailId := neo4jt.AddEmailTo(ctx, driver, entity.CONTACT, tenantName, contactId, "original@email.com", true, "")

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/update_email_for_contact"),
		client.Var("contactId", contactId),
		client.Var("emailId", emailId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var email struct {
		EmailUpdateInContact model.Email
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &email)
	require.Nil(t, err, "Error unmarshalling response data")

	e := email.EmailUpdateInContact

	// Check that the fields of the email struct have the expected values
	require.Equal(t, emailId, e.ID, "Email ID is nil")
	require.Equal(t, true, e.Primary, "Email Primary field is not true")
	require.Equal(t, "original@email.com", *e.RawEmail, "Email address expected not to be changed")
	require.Equal(t, "original@email.com", *e.Email, "Email address expected not to be changed")
	require.NotNil(t, e.UpdatedAt, "Missing updatedAt field")
	if e.Label == nil {
		t.Errorf("Email Label field is nil")
	} else {
		require.Equal(t, model.EmailLabelHome, *e.Label, "Email Label field is not expected value")
	}

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"), "Incorrect number of Email nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"), "Incorrect number of HAS relationships in Neo4j")
}

func TestMutationResolver_EmailUpdateInContact_ReplaceEmail(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	contactId := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	emailId := neo4jt.AddEmailTo(ctx, driver, entity.CONTACT, tenantName, contactId, "original@email.com", true, "")

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/replace_email_for_contact"),
		client.Var("contactId", contactId),
		client.Var("emailId", emailId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the emailStruct struct
	var emailStruct struct {
		EmailUpdateInContact model.Email
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &emailStruct)
	require.Nil(t, err, "Error unmarshalling response data")

	email := emailStruct.EmailUpdateInContact

	// Check that the fields of the emailStruct struct have the expected values
	require.NotEqual(t, emailId, email.ID, "Expected new email id to be generated")
	require.Equal(t, true, email.Primary, "Email Primary field is not true")
	require.Equal(t, "new@email.com", *email.RawEmail)
	require.Equal(t, "new@email.com", *email.Email)
	require.False(t, *email.Validated, "New email is not validated yet")
	require.NotNil(t, email.CreatedAt, "Missing createdAt field")
	require.NotNil(t, email.UpdatedAt, "Missing updatedAt field")
	if email.Label == nil {
		t.Errorf("Email Label field is nil")
	} else {
		require.Equal(t, model.EmailLabelHome, *email.Label, "Email Label field is not expected value")
	}

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Email"), "Expected 2 email nodes, original one and new")
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"), "Incorrect number of HAS relationships in Neo4j")

	// Check the labels on the nodes in the Neo4j database
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "Email", "Email_" + tenantName})
}

func TestMutationResolver_EmailMergeToUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact
	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/merge_email_to_user"),
		client.Var("userId", userId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var email struct {
		EmailMergeToUser model.Email
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &email)
	require.Nil(t, err, "Error unmarshalling response data")

	e := email.EmailMergeToUser

	// Check that the fields of the email struct have the expected values
	require.NotNil(t, e.ID, "Email ID is nil")
	require.NotNil(t, e.CreatedAt, "Missing createdAt field")
	require.NotNil(t, e.UpdatedAt, "Missing updatedAt field")
	require.Equal(t, true, e.Primary, "Email Primary field is not true")
	require.Equal(t, "test@gmail.com", *e.Email)
	require.Equal(t, "test@gmail.com", *e.RawEmail)
	require.False(t, *e.Validated)
	if e.Label == nil {
		t.Errorf("Email Label field is nil")
	} else {
		require.Equal(t, model.EmailLabelWork, *e.Label, "Email Label field is not expected value")
	}
	require.Equal(t, model.DataSourceOpenline, e.Source, "Email Source field is not expected value")
	require.Equal(t, model.DataSourceOpenline, e.SourceOfTruth, "Email Source of truth field is not expected value")
	require.Equal(t, "test", e.AppSource, "Email App source field is not expected value")

	// Check the number of nodes and relationships in the Neo4j database
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"), "Incorrect number of User nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"), "Incorrect number of Email nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName), "Incorrect number of Email_%s nodes in Neo4j", tenantName)
	require.Equal(t, 3, neo4jt.GetTotalCountOfNodes(ctx, driver), "Incorrect total number of nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"), "Incorrect number of HAS relationships in Neo4j")

	// Check the labels on the nodes in the Neo4j database
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "User", "Email", "Email_" + tenantName})
}

func TestMutationResolver_EmailUpdateInUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	emailId := neo4jt.AddEmailTo(ctx, driver, entity.USER, tenantName, userId, "original@email.com", true, "")

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/update_email_for_user"),
		client.Var("userId", userId),
		client.Var("emailId", emailId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var email struct {
		EmailUpdateInUser model.Email
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &email)
	require.Nil(t, err, "Error unmarshalling response data")

	e := email.EmailUpdateInUser

	// Check that the fields of the email struct have the expected values
	require.Equal(t, emailId, e.ID, "Email ID is nil")
	require.Equal(t, true, e.Primary, "Email Primary field is not true")
	require.Equal(t, "original@email.com", *e.Email, "Email address expected not to be changed")
	require.Equal(t, "original@email.com", *e.RawEmail, "Email address expected not to be changed")
	require.NotNil(t, e.UpdatedAt, "Missing updatedAt field")
	if e.Label == nil {
		t.Errorf("Email Label field is nil")
	} else {
		require.Equal(t, model.EmailLabelHome, *e.Label, "Email Label field is not expected value")
	}

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"), "Incorrect number of Email nodes in Neo4j")
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"), "Incorrect number of HAS relationships in Neo4j")
}

func TestMutationResolver_EmailDelete(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	emailId := neo4jt.AddEmailTo(ctx, driver, entity.USER, tenantName, userId, "original@email.com", true, "")
	neo4jt.AddEmailTo(ctx, driver, entity.CONTACT, tenantName, contactId, "original@email.com", true, "")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/delete_email"),
		client.Var("emailId", emailId))
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var emailStruct struct {
		EmailDelete model.Result
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &emailStruct)
	require.Nil(t, err, "Error unmarshalling response data")

	require.Equal(t, true, emailStruct.EmailDelete.Result)

	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "User"})
}

func TestMutationResolver_EmailRemoveFromUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	neo4jt.AddEmailTo(ctx, driver, entity.USER, tenantName, userId, "original@email.com", true, "")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/remove_email_from_user"),
		client.Var("userId", userId),
		client.Var("email", "original@email.com"),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var emailStruct struct {
		EmailRemoveFromUser model.Result
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &emailStruct)
	require.Nil(t, err, "Error unmarshalling response data")

	require.Equal(t, true, emailStruct.EmailRemoveFromUser.Result)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Email", "Email_" + tenantName, "User"})
}

func TestMutationResolver_EmailRemoveFromUserById(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)

	// Create a tenant in the Neo4j database
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Create a default contact and email
	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	emailId := neo4jt.AddEmailTo(ctx, driver, entity.USER, tenantName, userId, "original@email.com", true, "")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	// Make the RawPost request and check for errors
	rawResponse, err := c.RawPost(getQuery("email/remove_email_from_user_by_id"),
		client.Var("userId", userId),
		client.Var("emailId", emailId),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	// Unmarshal the response data into the email struct
	var emailStruct struct {
		EmailRemoveFromUserById model.Result
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &emailStruct)
	require.Nil(t, err, "Error unmarshalling response data")

	require.Equal(t, true, emailStruct.EmailRemoveFromUserById.Result)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Email", "Email_" + tenantName, "User"})
}
