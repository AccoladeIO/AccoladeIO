// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Credit to: Jason McVetta


package Neo4j_Utils

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
	"strings"
)


//strict "Person" entity model
func CreatePerson(db *neoism.Database, Name string, Gender string, DOB string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"Name": Name, "Gender": Gender, "DOB": DOB})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Person")

	return p
}

//strict "Entity" entity model
func CreateEntity(db *neoism.Database, Name string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"Name": Name})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Entity")

	return p
}

//strict "Accolade" entity model
func CreateAccolade(db *neoism.Database, Name string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"Name": Name})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Accolade")

	return p
}

//strict "Property" entity model
func CreateProperty(db *neoism.Database, TitleRef string, Address string) {
	p, err := db.CreateNode(neoism.Props{"TitleRef": TitleRef, "Address":Address })
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Property")

}

//strict "Mortgage" entity model
func CreateMortgage(db *neoism.Database, MortgageRef string) {
	p, err := db.CreateNode(neoism.Props{"MortgageRef": "M123A4SD560"})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Mortgage")
}

//strict "EmailAddress" entity model
func EmailAddress(db *neoism.Database, Email string) *neoism.Node {

	p, err := db.CreateNode(neoism.Props{"DomainKey": Email[0:strings.Index(Email, "@")], "Domain": Email[strings.Index(Email, "@"):len(Email)]})
	if err != nil {
		log.Fatal(err)
	}
	p.AddLabel("Email")

	return p
}




//Sample code not simplified yet
func create_example(db *neoism.Database) {
	kirk, err := db.CreateNode(neoism.Props{"name": "Kirk", "shirt": "yellow"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kirk.Properties()) // map[shirt:yellow name:Kirk] <nil>
	// Ignoring subsequent errors for brevity
	spock, _ := db.CreateNode(neoism.Props{"name": "Spock", "shirt": "blue"})
	mccoy, _ := db.CreateNode(neoism.Props{"name": "McCoy", "shirt": "blue"})
	tom, _ := db.CreateNode(neoism.Props{"name": "Tom", "shirt": "white"})
	r, _ := kirk.Relate("outranks", spock.Id(), nil) // No properties on this relationship
	start, _ := r.Start()
	fmt.Println(start.Properties()) // map[name:Kirk shirt:yellow] <nil>
	kirk.Relate("outranks", mccoy.Id(), nil)
	spock.Relate("outranks", mccoy.Id(), nil)
	mccoy.Relate("outranks", tom.Id(), nil)
	kirk.AddLabel("Person")
	mccoy.AddLabel("Person")
	spock.AddLabel("Person")
	tom.AddLabel("Person")
}

func cypher(db *neoism.Database) {
	cq := neoism.CypherQuery{
		Statement: `
			START n=node(*)
			MATCH (n)-[r:outranks]->(m)
			WHERE n.shirt = {color}
			RETURN n.name, type(r), m.name
			`,
		Parameters: neoism.Props{"color": "blue"},
		Result: &[]struct {
			N   string `json:"n.name"`
			Rel string `json:"type(r)"`
			M   string `json:"m.name"`
		}{},
	}
	// db.Session.Log = true
	db.Cypher(&cq)
	fmt.Println(cq.Result)
	// &[{Spock outranks McCoy} {Spock outranks Scottie} {McCoy outranks Scottie}]
}

//Clears all nodes from the database
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