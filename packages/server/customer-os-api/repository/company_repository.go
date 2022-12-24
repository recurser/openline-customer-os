package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

type CompanyRepository interface {
	GetCompanyForRole(session neo4j.Session, tenant, roleId string) (*dbtype.Node, error)
	GetPaginatedCompaniesWithNameLike(tenant, companyName string, skip, limit int) (*utils.DbNodesWithTotalCount, error)
}

type companyRepository struct {
	driver *neo4j.Driver
}

func (r *companyRepository) GetCompanyForRole(session neo4j.Session, tenant, roleId string) (*dbtype.Node, error) {
	dbRecords, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		if queryResult, err := tx.Run(`
			MATCH (r:Role {id:$roleId})-[:WORKS]->(co:Company)-[:COMPANY_BELONGS_TO_TENANT]->(t:Tenant {name:$tenant})
			RETURN co`,
			map[string]any{
				"tenant": tenant,
				"roleId": roleId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect()
		}
	})
	if err != nil {
		return nil, err
	}
	if len(dbRecords.([]*neo4j.Record)) == 0 {
		return nil, nil
	}
	return utils.NodePtr(dbRecords.([]*neo4j.Record)[0].Values[0].(dbtype.Node)), nil

}

func NewCompanyRepository(driver *neo4j.Driver) CompanyRepository {
	return &companyRepository{
		driver: driver,
	}
}

func (r *companyRepository) GetPaginatedCompaniesWithNameLike(tenant, companyName string, skip, limit int) (*utils.DbNodesWithTotalCount, error) {
	session := utils.NewNeo4jReadSession(*r.driver)
	defer session.Close()

	dbNodesWithTotalCount := new(utils.DbNodesWithTotalCount)

	dbRecords, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		queryResult, err := tx.Run(`
				MATCH (:Tenant {name:$tenant})<-[:COMPANY_BELONGS_TO_TENANT]-(co:Company) 
				WHERE toLower(co.name) CONTAINS toLower($companyName)
				RETURN count(co) as count`,
			map[string]interface{}{
				"tenant":      tenant,
				"companyName": companyName,
			})
		if err != nil {
			return nil, err
		}
		count, _ := queryResult.Single()
		dbNodesWithTotalCount.Count = count.Values[0].(int64)

		queryResult, err = tx.Run(`
				MATCH (:Tenant {name:$tenant})<-[:COMPANY_BELONGS_TO_TENANT]-(co:Company) 
                WHERE toLower(co.name) CONTAINS toLower($companyName)
				RETURN co ORDER BY co.name SKIP $skip LIMIT $limit`,
			map[string]interface{}{
				"tenant":      tenant,
				"companyName": companyName,
				"skip":        skip,
				"limit":       limit,
			})
		return queryResult.Collect()
	})
	if err != nil {
		return nil, err
	}

	for _, v := range dbRecords.([]*neo4j.Record) {
		dbNodesWithTotalCount.Nodes = append(dbNodesWithTotalCount.Nodes, utils.NodePtr(v.Values[0].(neo4j.Node)))
	}

	return dbNodesWithTotalCount, nil
}
