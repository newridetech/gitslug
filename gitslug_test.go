package gitslug

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreatedByName(t *testing.T) {
    assert.Equal(t, "test17-pl-newridetech-63d3ef86f6b16907eb52be24a94a6f11", SlugifyGitRepositoryName("git@github.com:newridetech/test17.pl.git"))
    assert.Equal(t, "unifiedapi-io-newridetech-ab5c1cb3200f083df1299cfbe524b520", SlugifyGitRepositoryName("git@github.com:newridetech/unifiedapi.io.git"))
}
