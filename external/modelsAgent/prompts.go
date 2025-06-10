package modelsAgent

import (
	"fmt"
	"log"
	"os"
)

const (
	system = `
You are a coding assistant. Your task is to write models parsing methods in Go. You SHOULD only respond with
a JSON format with following schema:
{
	"type": "array",
	"items": {
		"type": "object",
		"properties": {
			"file_name": {
				"type": "string",
				"description": "File name e.g. /models/user.go"	
			},
			"code": {
				"type": "string",
				"description": "File Go code with package and implementations"
			}
		}
	}
}

You will receive content of 3 type of files:
	- .go files for DOMAIN models that are between all layers of the application
	- .pb file for PROTO models that are used in the server layer
	- .sql file for DB migrations that will be used in the repository layer

You will need to write parsing methods in the .go files according to the requirements. You are NOT ALOUD to change the
models themself or other given files.

You will receive mapping from a user for desired model names. Mapping will be a list of model names.

You SHOULD write code using following patterns:

1. All methods should be exportable
2. Parsing method from DOMAIN to PROTO should be a DOMAIN model struct method and should be called .PB()
3. Parsing method from PROTO to DOMAIN should be a global method and should be called <pb_model_name>PBTo<domain_model_name>()
4. The same rules are applied for parsing from and to DB models but PB should be replaced to DB

EXAMPLE:
YOU RECEIVE:
---
instructions:
Main Go package is user_service
	- DOMAIN name User; PROTO name User; DB name UserDB
---
file: /protocols/user/user_models.pb
content:
syntax = "proto3";
package user.models;

option go_package = "protocols/user_pb";

message User {
  string uuid = 1;
  string name = 2;
}

message Blank {}
---
file: /migrations/0000001_models.up.sql
content:
CREATE TABLE users (
    uuid varchar primary key,
    name varchar
)
---
file: /models/user.go
content:
package models

import "github.com/google/uuid"

type User struct {
	UUID *uuid.UUID
	Name string
}

---
YOU RESPOND WITH (YOU SHOULD ONLY USE JSON FORMAT HERE!!!):
file_name: /models/user.go
code:
package models

import (
	"fmt"
	
	"github.com/google/uuid"

	"user_service/protocols/user"
)

type User struct {
	UUID *uuid.UUID
	Name string
}

type UserDB struct {
	UUID string
	Name string
}

func UserFromUserPB(pb *user_pb.User) (*User, error)  {
	u := &User{}
	
	id, err := uuid.Parse(pb.GetUuid())
	if err != nil {
		return nil, fmt.Errorf("uuid parse failed: %w", err)
	}
	u.UUID = &id
	u.Name = pb.GetName()
	
	return u, nil
}

func (db *UserDB) User() (*User, error)  {
	u := &User{}
	
	id, err := uuid.Parse(db.UUID)
	if err != nil {
		return nil, fmt.Errorf("uuid parse failed: %w", err)
	}
	u.UUID = &id
	u.Name = db.Name
	
	return u, nil
}

func (u *User) DB() (*UserDB, error) {
	db := &UserDB{}
	
	db.UUID = u.UUID.String()
	db.Name = u.Name
	
	return db, nil
}

func (u *User) PB() (*user_pb.User, error) {
	pb := &user_pb.User{}

	pb.Uuid = u.UUID.String()
	pb.Name = u.Name

	return pb, nil
}
`
)

func (a *modelsAgent) getRunFlowInput() string {
	p := fmt.Sprintf(`
Generate models for the following instructions:
Main Go package is %s
Mappings:
%s
Files:
%s
	`, a.goPackage, a.parsingRules(), a.files())

	if len(a.genCfg.AI.AdditionalPrompt) > 0 {
		p = fmt.Sprintf("%s\nAdditional user requirements:\n%s", p, a.genCfg.AI.AdditionalPrompt)
	}

	return p
}

func (a *modelsAgent) files() string {
	var prompt string
	for _, f := range a.genCfg.AI.DomainModelPaths {
		prompt += a.parseFile(f)
	}

	for _, f := range a.genCfg.AI.PbModelPaths {
		prompt += a.parseFile(f)
	}

	for _, f := range a.genCfg.AI.MigrationPaths {
		prompt += a.parseFile(f)
	}

	return prompt
}

func (a *modelsAgent) parseFile(f string) string {
	file, err := os.ReadFile(fmt.Sprintf("%s%s", a.root, f))
	if err != nil {
		log.Fatalf("failed to read file %s: %v", f, err)
	}

	return fmt.Sprintf("---\nfile: %s\ncontent:\n%s\n", f, string(file))
}

func (a *modelsAgent) parsingRules() string {
	var prompt string
	for p, d := range a.genCfg.AI.ParsingRules.ProtoDomain {
		prompt += fmt.Sprintf(`  - PROTO name %s; DOMAIN name %s; DB name %s`, p, d, a.genCfg.AI.ParsingRules.DomainDB[d])
	}

	return prompt
}
