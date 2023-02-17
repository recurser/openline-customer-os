# CHECK 1 - Cross tenant data link - Result should be 0

CALL {
 MATCH (t1:Tenant)--(c:Contact)-[rel]-(:Email)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(c:Contact)-[rel]-(:Location)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(c:Contact)-[rel]-(:Tag)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Contact)-[rel]-(:Organization)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Contact)-[rel]-(:ExternalSystem)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Organization)-[rel]-(:Location)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Organization)-[rel]-(:Email)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Organization)-[rel]-(:Tag)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Organization)-[rel]-(:ExternalSystem)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(:Note)-[rel]-(:ExternalSystem)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(c1:Contact)-[rel]-(:Conversation)--(c2:Contact)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(u1:User)-[rel]-(:Conversation)--(c2:Contact)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
 UNION
 MATCH (t1:Tenant)--(u1:User)-[rel]-(:Conversation)--(u2:User)--(t2:Tenant)
 WHERE t1.name <> t2.name
 return count(rel) as x
} return sum(x) as Problematic_relationships;

# CHECK 2 - Check amount of labels per node
CALL {
 MATCH (node:Tenant) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 1 return count(nodeCount) as x
 UNION
 MATCH (node:EmailDomain) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 1 return count(nodeCount) as x
 UNION
 MATCH (node:AlternateOrganization) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 1 return count(nodeCount) as x
 UNION
 MATCH (node:AlternateContact) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 1 return count(nodeCount) as x
 UNION
 MATCH (node:Tag) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:ExternalSystem) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:OrganizationType) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:User) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Contact) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Organization) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:JobRole) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Email) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Conversation) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Location) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:Note) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:PhoneNumber) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 2 return count(nodeCount) as x
 UNION
 MATCH (node:CustomField) with node, labels(node) as labs unwind labs as labsList with node, count(node) as nodeCount where nodeCount <> 3 return count(nodeCount) as x
} return sum(x) as Problematic_nodes;

# CHECK 3

CALL {
 MATCH (t:Tenant)--(n:Contact) WHERE NOT 'Contact_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:Organization) WHERE NOT 'Organization_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:Organization) WHERE NOT 'Organization_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:ExternalSystem) WHERE NOT 'ExternalSystem_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:OrganizationType) WHERE NOT 'OrganizationType_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:User) WHERE NOT 'User_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:Tag) WHERE NOT 'Tag_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:Email) WHERE NOT 'Email_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:Email) WHERE NOT 'Email_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Organization)--(n:Email) WHERE NOT 'Email_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:User)--(n:Email) WHERE NOT 'Email_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:User)--(n:Conversation) WHERE NOT 'Conversation_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:Conversation) WHERE NOT 'Conversation_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:JobRole) WHERE NOT 'JobRole_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(n:Location) WHERE NOT 'Location_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:Location) WHERE NOT 'Location_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Organization)--(n:Location) WHERE NOT 'Location_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:CustomField) WHERE NOT 'CustomField_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:PhoneNumber) WHERE NOT 'PhoneNumber_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:Note) WHERE NOT 'Note_'+t.name  in labels(n) return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)--(n:Action) WHERE NOT 'Action_'+t.name  in labels(n) return count(n) as x
} return sum(x) as Problematic_nodes;

#=========== CHECK 4 - Missing tenant links

CALL {
 MATCH (t:Tenant)--(:Contact)-[rel]-(n:Email) WHERE NOT (n)--(t)
 return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)-[rel]-(n:Tag) WHERE NOT (n)--(t)
 return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Contact)-[rel]-(n:Location) WHERE NOT (n)--(t)
 return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Organization)-[rel]-(n:Email) WHERE NOT (n)--(t)
 return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Organization)-[rel]-(n:Tag) WHERE NOT (n)--(t)
 return count(n) as x
 UNION
 MATCH (t:Tenant)--(:Organization)-[rel]-(n:Location) WHERE NOT (n)--(t)
 return count(n) as x
} return sum(x) as Problematic_nodes;

#=========== CHECK 5 - Query to verify labels mix
MATCH (n) RETURN count(labels(n)), labels(n);