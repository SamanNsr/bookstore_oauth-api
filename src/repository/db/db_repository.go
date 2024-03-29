package db

import (
	"github.com/SamanNsr/bookstore_oauth-api/src/clients/cassandra"
	"github.com/SamanNsr/bookstore_oauth-api/src/domain/access_token"
	"github.com/SamanNsr/bookstore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const(
	queryGetAccessToken="SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) (*errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) (*errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err !=nil {
		
		if err==gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return nil, errors.NewInternalServerError("database connection not implemented")
}

func (r *dbRepository) Create(at access_token.AccessToken) (*errors.RestErr){
	session := cassandra.GetSession()
	if err := session.Query(
		queryCreateAccessToken, 
		at.AccessToken, 
		at.UserId, 
		at.ClientId, 
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) (*errors.RestErr){
	session := cassandra.GetSession()
	if err := session.Query(
		queryUpdateExpires, 
		at.Expires,
		at.AccessToken, 
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}