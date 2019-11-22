package handlers

import (
	"github.com/golang/glog"
	"github.com/hashicorp/go-memdb"
)

//Create database schema
func createDatabaseSchema() *memdb.DBSchema {
	glog.V(5).Info("Creating database schema")
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			//MeetupGroup Schema
			"meetupGroup": {
				Name: "meetupGroup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"city": {
						Name:         "city",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "City"},
					},
					"country": {
						Name:         "country",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Country"},
					},
					"latitude": {
						Name:         "latitude",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Latitude"},
					},
					"longitude": {
						Name:         "longitude",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Longitude"},
					},
				},
			},
			//Organizer Schema
			"organizer": {
				Name: "organizer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"title": {
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"email": {
						Name:         "email",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Email"},
					},
					"github": {
						Name:         "github",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Github"},
					},
					"twitter": {
						Name:         "twitter",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Twitter"},
					},
					"speakersBureau": {
						Name:         "speakersBureau",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakersBureau"},
					},
					"meetupGroupId": {
						Name:         "meetupGroupId",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
				},
			},
			//Company Schema
			"company": {
				Name: "company",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": {
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": {
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
					"organizerID": {
						Name:         "organizerID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "OrganizerID"},
					},
					"speakerID": {
						Name:         "speakerID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
				},
			},
			//Meetup Schema
			"meetup": {
				Name: "meetup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"date": {
						Name:         "date",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Date"},
					},
					"duration": {
						Name:         "duration",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Duration"},
					},
					"attendees": {
						Name:         "attendees",
						Unique:       false,
						AllowMissing: false,
						Indexer:      &memdb.IntFieldIndex{Field: "Attendees"},
					},
					"address": {
						Name:         "address",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Address"},
					},
					"meetupGroupID": {
						Name:         "meetupGroupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
				},
			},
			//MeetupSponsor Schema
			"meetupSponsor": {
				Name: "meetupSponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": {
						Name:         "meetupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//Member Schema
			"member": {
				Name: "member",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": {
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": {
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
				},
			},
			//Sponsor Schema
			"sponsor": {
				Name: "sponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": {
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": {
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
				},
			},
			//Venue Schema
			"venue": {
				Name: "venue",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": {
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": {
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
				},
			},
			//Other Schema
			"other": {
				Name: "other",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": {
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": {
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
				},
			},
			//Presentation Schema
			"presentation": {
				Name: "presentation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"duration": {
						Name:         "duration",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Duration"},
					},
					"title": {
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"slides": {
						Name:         "slides",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Slides"},
					},
					"meetupID": {
						Name:         "meetupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//Speaker Schema
			"speaker": {
				Name: "speaker",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"title": {
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"email": {
						Name:         "email",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Email"},
					},
					"github": {
						Name:         "github",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Github"},
					},
					"speakersBureau": {
						Name:         "speakersBureau",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakersBureau"},
					},
					"presentationID": {
						Name:         "presentationID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "PresentationID"},
					},
				},
			},
			//Country Schema
			"country": {
				Name: "country",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			//CountryToEntity Schema
			"entityToCountry": {
				Name: "entityToCountry",
				Indexes: map[string]*memdb.IndexSchema{
					"countryId": {
						Name:    "countryId",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CountryID"},
					},
					"id": {
						Name:   "id",
						Unique: true,

						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"entityType": {
						Name:    "entityType",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "EntityType"},
					},
				},
			},
		},
	}
	return schema
}
