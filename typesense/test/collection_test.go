//go:build integration
// +build integration

package test

import (
	"testing"

	"github.com/KSerik/typesensego/typesense/api"
	"github.com/KSerik/typesensego/typesense/api/pointer"
	"github.com/stretchr/testify/require"
)

func TestCollectionRetrieve(t *testing.T) {
	collectionName := createNewCollection(t, "companies")
	expectedResult := expectedNewCollection(collectionName)

	result, err := typesenseClient.Collection(collectionName).Retrieve()
	result.CreatedAt = pointer.Int64(0)

	require.NoError(t, err)
	require.Equal(t, expectedResult, result)
}

func TestCollectionDelete(t *testing.T) {
	collectionName := createNewCollection(t, "companies")
	expectedResult := expectedNewCollection(collectionName)

	result, err := typesenseClient.Collection(collectionName).Delete()
	result.CreatedAt = pointer.Int64(0)
	require.NoError(t, err)
	require.Equal(t, expectedResult, result)

	_, err = typesenseClient.Collection(collectionName).Retrieve()
	require.Error(t, err)
}

func TestCollectionUpdate(t *testing.T) {
	collectionName := createNewCollection(t, "companies")

	updateSchema := &api.CollectionUpdateSchema{
		Fields: []api.Field{
			{
				Name: "country",
				Drop: pointer.True(),
			},
		},
	}

	result, err := typesenseClient.Collection(collectionName).Update(updateSchema)
	require.NoError(t, err)
	require.Equal(t, 1, len(result.Fields))
	require.Equal(t, "country", result.Fields[0].Name)
	require.Equal(t, pointer.True(), result.Fields[0].Drop)
}
