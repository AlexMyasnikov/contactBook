package fakeDB

import (
	pb "github.com/ChuvashPeople/contactBook/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Contact struct {
	Id   int
	Name string
}

type Db struct {
	contacts []Contact
}

func (db *Db) AddContact(r *pb.AddRequest) int {
	var id int
	if len(db.contacts) > 0 {
		id = db.contacts[0].Id
		for _, element := range db.contacts {
			if element.Id > id {
				id = element.Id
			}
		}
		id += 1
	} else {
		id = 1
	}
	db.contacts = append(db.contacts, Contact{id, r.Name})
	return id
}

func (db *Db) GetContact(r *pb.GetRequest) (Contact, error) {
	var cont Contact
	for _, c := range db.contacts {
		if int(r.Id) == c.Id {
			cont = c
		}
	}
	if cont.Id == 0 {
		return cont, status.Errorf(codes.NotFound, "cant find this contact")
	} else {
		return cont, nil
	}
}
