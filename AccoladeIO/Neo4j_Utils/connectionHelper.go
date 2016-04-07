// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Credit to: Jason McVetta


package Neo4j_Utils

import (
	"github.com/jmcvetta/neoism"
	"log"
)

//Connect () creates a connection to a Neo4j database passing in
//three params (UserName, Password, RestAPI URL)
func Connect(conn string) *neoism.Database {
	db, err := neoism.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}