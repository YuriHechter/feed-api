package mdb

/*
This is a modified version of the github.com/Bnei-Baruch/mdb/api/registry.go
 We take, manually, only what we need from there.
*/

import (
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/pkg/errors"

	mdbmodels "github.com/Bnei-Baruch/feed-api/databases/mdb/models"
)

var (
	CONTENT_TYPE_REGISTRY      = &ContentTypeRegistry{}
	CONTENT_ROLE_TYPE_REGISTRY = &ContentRoleTypeRegistry{}
	PERSONS_REGISTRY           = &PersonsRegistry{}
	AUTHOR_REGISTRY            = &AuthorRegistry{}
	SOURCE_TYPE_REGISTRY       = &SourceTypeRegistry{}
	TWITTER_USERS_REGISTRY     = &TwitterUsersRegistry{}
	BLOGS_REGISTRY             = &BlogsRegistry{}
)

func InitTypeRegistries(exec boil.ContextExecutor) error {
	registries := []TypeRegistry{CONTENT_TYPE_REGISTRY,
		CONTENT_ROLE_TYPE_REGISTRY,
		PERSONS_REGISTRY,
		AUTHOR_REGISTRY,
		SOURCE_TYPE_REGISTRY,
		TWITTER_USERS_REGISTRY,
		BLOGS_REGISTRY}

	for _, x := range registries {
		if err := x.Init(exec); err != nil {
			return err
		}
	}

	return nil
}

type TypeRegistry interface {
	Init(exec boil.ContextExecutor) error
}

type ContentTypeRegistry struct {
	ByName map[string]*mdbmodels.ContentType
	ByID   map[int64]*mdbmodels.ContentType
}

func (r *ContentTypeRegistry) Init(exec boil.ContextExecutor) error {
	types, err := mdbmodels.ContentTypes().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load content_types from DB")
	}

	r.ByName = make(map[string]*mdbmodels.ContentType)
	r.ByID = make(map[int64]*mdbmodels.ContentType)
	for _, t := range types {
		r.ByName[t.Name] = t
		r.ByID[t.ID] = t
	}

	return nil
}

type ContentRoleTypeRegistry struct {
	ByName map[string]*mdbmodels.ContentRoleType
}

func (r *ContentRoleTypeRegistry) Init(exec boil.ContextExecutor) error {
	types, err := mdbmodels.ContentRoleTypes().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load content_role_types from DB")
	}

	r.ByName = make(map[string]*mdbmodels.ContentRoleType)
	for _, t := range types {
		r.ByName[t.Name] = t
	}

	return nil
}

type PersonsRegistry struct {
	ByPattern map[string]*mdbmodels.Person
}

func (r *PersonsRegistry) Init(exec boil.ContextExecutor) error {
	types, err := mdbmodels.Persons(qm.Where("pattern is not null")).All(exec)
	if err != nil {
		return errors.Wrap(err, "Load persons from DB")
	}

	r.ByPattern = make(map[string]*mdbmodels.Person)
	for _, t := range types {
		r.ByPattern[t.Pattern.String] = t
	}

	return nil
}

type AuthorRegistry struct {
	ByCode map[string]*mdbmodels.Author
}

func (r *AuthorRegistry) Init(exec boil.ContextExecutor) error {
	authors, err := mdbmodels.Authors().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load authors from DB")
	}

	r.ByCode = make(map[string]*mdbmodels.Author)
	for _, a := range authors {
		r.ByCode[a.Code] = a
	}

	return nil
}

type SourceTypeRegistry struct {
	ByName map[string]*mdbmodels.SourceType
	ByID   map[int64]*mdbmodels.SourceType
}

func (r *SourceTypeRegistry) Init(exec boil.ContextExecutor) error {
	types, err := mdbmodels.SourceTypes().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load source_types from DB")
	}

	r.ByName = make(map[string]*mdbmodels.SourceType)
	r.ByID = make(map[int64]*mdbmodels.SourceType)
	for _, t := range types {
		r.ByName[t.Name] = t
		r.ByID[t.ID] = t
	}

	return nil
}

type TwitterUsersRegistry struct {
	ByUsername map[string]*mdbmodels.TwitterUser
	ByID       map[int64]*mdbmodels.TwitterUser
}

func (r *TwitterUsersRegistry) Init(exec boil.ContextExecutor) error {
	users, err := mdbmodels.TwitterUsers().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load twitter users from DB")
	}

	r.ByUsername = make(map[string]*mdbmodels.TwitterUser)
	r.ByID = make(map[int64]*mdbmodels.TwitterUser)
	for _, t := range users {
		r.ByUsername[t.Username] = t
		r.ByID[t.ID] = t
	}

	return nil
}

type BlogsRegistry struct {
	ByName map[string]*mdbmodels.Blog
	ByID   map[int64]*mdbmodels.Blog
}

func (r *BlogsRegistry) Init(exec boil.ContextExecutor) error {
	blogs, err := mdbmodels.Blogs().All(exec)
	if err != nil {
		return errors.Wrap(err, "Load blogs from DB")
	}

	r.ByName = make(map[string]*mdbmodels.Blog)
	r.ByID = make(map[int64]*mdbmodels.Blog)
	for _, b := range blogs {
		r.ByName[b.Name] = b
		r.ByID[b.ID] = b
	}

	return nil
}
