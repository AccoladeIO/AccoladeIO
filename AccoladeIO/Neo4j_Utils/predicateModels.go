// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Credit to: Jason McVetta


package Neo4j_Utils

import (
	"github.com/jmcvetta/neoism"
)

//Strict "Received" relation model.
//The only object that may be received by a subject is an Accolade.
func Received(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "received"
	subject.Relate(p, object.Id(), nil)
}

//Strict "Issued" relation model.
//The only object that may be issued from an entity is an Accolade.
func Issued(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "issued"
	subject.Relate(p, object.Id(), nil)
}

//Strict "Claims" relation model.
//Claims are made over a tract a domain.
//One might claim an email address as their own.
//No ones is said to own a domain tract, they may only claim it and are given certain liberties regarding the associated data based on that claim.
//That claim can be disputed, revoked or severed
func Claims(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "claims"
	subject.Relate(p, object.Id(), nil)
}

//Strict "Rules" relation model.
//Where one entity rules another, both must represent the same body and it is the ruling entity that takes primacy over its subordinates, which is to say that it is the single entity from the collection said to represent the whole.
func Rules(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "rules"
	subject.Relate(p, object.Id(), nil)
}