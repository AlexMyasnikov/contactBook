package fakeDB

import (
	pb "github.com/ChuvashPeople/contactBook/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contact struct {
	id   int
	name string
}

type db struct {
	contacts []contact
}

func (db *db) addContact(r *pb.AddRequest) int {
	var id int
	if len(db.contacts) > 0 {
		id = db.contacts[0].id
		for _, element := range db.contacts {
			if element.id > id {
				id = element.id
			}
		}
		id += 1
	} else {
		id = 1
	}
	db.contacts = append(db.contacts, contact{id, r.Name})
	return id
}

func (db *db) getContact(id int) (contact, error) {
	var cont contact
	for _, c := range db.contacts {
		if id == c.id {
			cont = c
		}
	}
	if cont.id == 0 {
		return cont, status.Errorf(codes.NotFound, "cant find this contact")
	} else {
		return cont, nil
	}
}
