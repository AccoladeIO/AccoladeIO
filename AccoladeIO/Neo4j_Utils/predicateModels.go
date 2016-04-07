// Copyright (c) 2015 Thomas O'Beirne.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Credit to: Jason McVetta


package Neo4j_Utils

import (
	"github.com/jmcvetta/neoism"
)

//strict "Marriage" relation model
func Married(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "married to"
	subject.Relate(p, object.Id(), neoism.Props{"Date": date})
}

//strict "Received" relation model
func Received(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "received"
	subject.Relate(p, object.Id(), nil)
}

//strict "Issued" relation model
func Issued(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "issued"
	subject.Relate(p, object.Id(), nil)
}

//strict "Claims" relation model
func Claims(db *neoism.Database, subject *neoism.Node, object *neoism.Node, date string) {
	p := "claims"
	subject.Relate(p, object.Id(), nil)
}