// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Credit to: Jason McVetta


package Neo4j_Utils

import (
	"github.com/jmcvetta/neoism"
	"log"
	"strings"
)


//Strict "Entity" entity model
func CreateEntity(db *neoism.Database, Name string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"Name": Name})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Entity")

	return p
}


//Strict "Accolade" entity model
func CreateAccolade(db *neoism.Database, Name string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"Name": Name})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Accolade")

	return p
}


//Strict "EmailAddress" entity model.
//This is a form of Domain Tract.
//It it stakes a claim over an identity that exists within a domain.
func CreateEmailAddress(db *neoism.Database, Email string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"DomainKey": Email[0:strings.Index(Email, "@")], "Domain": Email[strings.Index(Email, "@"):len(Email)]})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Email")

	return p
}


//Clears all nodes from the database.
//Best used for testing and removed entirely from production code.
func Cleanup(db *neoism.Database) {
	qs := []*neoism.CypherQuery{
		&neoism.CypherQuery{
			Statement: `START r=rel(*) DELETE r`,
		},
		&neoism.CypherQuery{
			Statement: `START n=node(*) DELETE n`,
		},
	}
	err := db.CypherBatch(qs)
	if err != nil {
		log.Fatal(err)
	}
}