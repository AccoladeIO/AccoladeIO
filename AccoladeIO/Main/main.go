// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package main

import (
	"github.com/ob27/PropertyGraph/PropertyGraph/Neo4j_Utils"
)

//we would pass in a stream of spouses and their wedding date
func main() {
	db := Neo4j_Utils.Connect("http://neo4j:foofight1!@localhost:7474/db/data/")
	//defer Neo4j_Utils.Cleanup(db)
	EntityA := Neo4j_Utils.CreateEntity(db, "WorkCorp")
	EntityB := Neo4j_Utils.CreateEntity(db, "Jo")
	AccoladeA := Neo4j_Utils.CreateAccolade(db, "MVP")
	AccoladeB := Neo4j_Utils.CreateAccolade(db, "Early Starter")
	AccoladeC := Neo4j_Utils.CreateAccolade(db, "Record Holder")
	EmailA := Neo4j_Utils.EmailAddress(db, "ob27@hotmail.com")
	EmailB := Neo4j_Utils.EmailAddress(db, "thomas.m.obeirne@gmail.com")
	Neo4j_Utils.Issued(db, EntityA, AccoladeA, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeA, EmailA, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityB, EmailA, "07/04/2016")
	Neo4j_Utils.Issued(db, EntityA, AccoladeB, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeB, EmailB, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityB, EmailB, "07/04/2016")
	Neo4j_Utils.Issued(db, EntityA, AccoladeC, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeC, EmailB, "07/04/2016")

	print("All good")
}


