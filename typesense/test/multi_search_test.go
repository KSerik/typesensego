//go:build integration
// +build integration

package test

import (
	"testing"

	"github.com/KSerik/typesensego/typesense/api"
	"github.com/KSerik/typesensego/typesense/api/pointer"
	"github.com/stretchr/testify/require"
)

func TestMultiSearch(t *testing.T) {
	collectionName1 := createNewCollection(t, "companies")
	collectionName2 := createNewCollection(t, "companies")
	documents := []interface{}{
		newDocument("123", withCompanyName("Company 1"), withNumEmployees(50)),
		newDocument("125", withCompanyName("Company 2"), withNumEmployees(150)),
		newDocument("127", withCompanyName("Company 3"), withNumEmployees(250)),
		newDocument("129", withCompanyName("Stark Industries 4"), withNumEmployees(500)),
		newDocument("131", withCompanyName("Stark Industries 5"), withNumEmployees(1000)),
	}

	params := &api.ImportDocumentsParams{Action: pointer.String("create")}
	_, err := typesenseClient.Collection(collectionName1).Documents().Import(documents, params)
	require.NoError(t, err)

	_, err = typesenseClient.Collection(collectionName1).Documents().Import(documents, params)
	require.NoError(t, err)

	searchParams := &api.MultiSearchParams{
		FilterBy: pointer.String("num_employees:>100"),
		Q:        pointer.String("Company"),
		QueryBy:  pointer.String("company_name"),
	}

	searches := api.MultiSearchSearchesParameter{
		Searches: []api.MultiSearchCollectionParameters{
			{
				Collection: collectionName1,
				FilterBy:   pointer.String("num_employees:>100"),
				SortBy:     pointer.String("num_employees:desc"),
			},
			{
				Collection: collectionName1,
				FilterBy:   pointer.String("num_employees:>1000"),
			},
			{
				Collection: collectionName2,
				Q:          pointer.String("Stark"),
				FilterBy:   pointer.String("num_employees:>=1000"),
			},
		},
	}

	expectedDocs1 := []map[string]interface{}{
		newDocumentResponse("127", withResponseCompanyName("Company 3"), withResponseNumEmployees(250)),
		newDocumentResponse("125", withResponseCompanyName("Company 2"), withResponseNumEmployees(150)),
	}

	expectedDocs2 := []map[string]interface{}{
		newDocumentResponse("131", withResponseCompanyName("Stark Industries 5"), withResponseNumEmployees(1000)),
	}

	result, err := typesenseClient.MultiSearch.Perform(searchParams, searches)
	require.NoError(t, err)

	require.Equal(t, 3, len(result.Results))

	// Check first result
	for i, doc := range *result.Results[0].Hits {
		require.Equal(t, *doc.Document, expectedDocs1[i])
	}

	// Check second result
	require.Equal(t, 0, len(*result.Results[1].Hits))

	// Check third result
	for i, doc := range *result.Results[2].Hits {
		require.Equal(t, *doc.Document, expectedDocs2[i])
	}
}

func TestMultiSearchVectorQuery(t *testing.T) {
	_, err := typesenseClient.Collection("embeddings").Delete()

	collSchema := api.CollectionSchema{
		Name: "embeddings",
		Fields: []api.Field{
			{
				Name: "title",
				Type: "string",
			},
			{
				Name:   "vec",
				Type:   "float[]",
				NumDim: pointer.Int(4),
			},
		},
	}

	_, err = typesenseClient.Collections().Create(&collSchema)
	require.NoError(t, err)

	type vecDocument struct {
		ID    string    `json:"id"`
		Title string    `json:"title"`
		Vec   []float32 `json:"vec"`
	}

	vecDoc := &vecDocument{
		ID:    "0",
		Title: "Stark Industries",
		Vec:   []float32{0.45, 0.222, 0.021, 0.1323},
	}

	_, err = typesenseClient.Collection("embeddings").Documents().Create(vecDoc)
	require.NoError(t, err)

	searchParams := &api.MultiSearchParams{}
	searches := api.MultiSearchSearchesParameter{
		Searches: []api.MultiSearchCollectionParameters{
			{
				Collection:  "embeddings",
				Q:           pointer.String("*"),
				VectorQuery: pointer.String("vec:([0.96826,0.94,0.39557,0.306488], k: 10)"),
			},
		},
	}

	searchResp, err := typesenseClient.MultiSearch.Perform(searchParams, searches)
	require.NoError(t, err)

	require.NotNil(t, searchResp.Results[0].Hits)
	require.Equal(t, 1, len(*searchResp.Results[0].Hits))
}
