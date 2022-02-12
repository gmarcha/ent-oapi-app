#! /bin/bash

#go get entgo.io/ent/cmd/ent
#go get github.com/masseelch/elk #(dependencies for code generation: go get github.com/mailru/easyjson github.com/go-chi/chi/v5 go.uber.org/zap)
#go get entgo.io/contrib/entoas

#Create ent schemas for database tables:
ent init User Event

#Generate ent code, after defining fields and relations in schemas:
#go generate ./ent