package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature schema/snapshot --feature sql/upsert --template ping.tmpl --template get_client.tmpl ./schema
