package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T){
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.Equal(t, "", at.AccessToken, "new access token should not have defined access token id" )

	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TextAccessTokenIsExpired(t *testing.T)  {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}