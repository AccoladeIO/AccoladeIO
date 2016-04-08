// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package main

import (
	"github.com/ob27/AccoladeIO/AccoladeIO/Neo4j_Utils"
)

//we would pass in a stream of spouses and their wedding date
func main() {
	db := Neo4j_Utils.Connect("http://neo4j:foofight1!@localhost:7474/db/data/")
	//defer Neo4j_Utils.Cleanup(db)
	EntityA := Neo4j_Utils.CreateEntity(db, "WorkCorp")
	EntityB := Neo4j_Utils.CreateEntity(db, "Peter")
	EntityC := Neo4j_Utils.CreateEntity(db, "Cafe Foo")
	EntityD := Neo4j_Utils.CreateEntity(db, "Pete")
	EntityE := Neo4j_Utils.CreateEntity(db, "Dick")
	EntityF := Neo4j_Utils.CreateEntity(db, "Piet")

	AccoladeA := Neo4j_Utils.CreateAccolade(db, "MVP")
	AccoladeB := Neo4j_Utils.CreateAccolade(db, "Self Starter")
	AccoladeC := Neo4j_Utils.CreateAccolade(db, "Record Holder")
	AccoladeD := Neo4j_Utils.CreateAccolade(db, "Loyal Patron")
	AccoladeE := Neo4j_Utils.CreateAccolade(db, "Team Player")
	AccoladeF := Neo4j_Utils.CreateAccolade(db, "Loyal Patron")
	AccoladeG := Neo4j_Utils.CreateAccolade(db, "Personal Best")

	EmailA := Neo4j_Utils.CreateEmailAddress(db, "Peter@hotmail.com")
	EmailB := Neo4j_Utils.CreateEmailAddress(db, "Peter.Smith@gmail.com")
	EmailC := Neo4j_Utils.CreateEmailAddress(db, "Dick.Smith@gmail.com")
	EmailD := Neo4j_Utils.CreateEmailAddress(db, "Piet.Smith@WorkCorp.com")
	EmailE := Neo4j_Utils.CreateEmailAddress(db, "Piet42@WorkCorp.com")


	//Creation of Claims
	Neo4j_Utils.Claims(db, EntityB, EmailA, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityB, EmailB, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityD, EmailC, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityE, EmailD, "07/04/2016")
	Neo4j_Utils.Claims(db, EntityF, EmailE, "07/04/2016")


	//Creation of Accolade
	Neo4j_Utils.Issued(db, EntityA, AccoladeB, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeB, EmailB, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityA, AccoladeA, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeA, EmailA, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityA, AccoladeC, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeC, EmailB, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityC, AccoladeD, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeD, EmailB, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityA, AccoladeE, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeE, EmailC, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityA, AccoladeF, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeF, EmailD, "07/04/2016")

	Neo4j_Utils.Issued(db, EntityA, AccoladeG, "07/04/2016")
	Neo4j_Utils.Received(db, AccoladeG, EmailC, "07/04/2016")

	//Creation of Rulings
	Neo4j_Utils.Rules(db, EntityB, EntityD, "07/04/2016")
	Neo4j_Utils.Rules(db, EntityB, EntityE, "07/04/2016")
	Neo4j_Utils.Rules(db, EntityB, EntityF, "07/04/2016")

	print("All good! go check localhost:7474")
}


