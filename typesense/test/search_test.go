//go:build integration
// +build integration

package test

import (
	"testing"

	"github.com/KSerik/typesensego/typesense/api"
	"github.com/KSerik/typesensego/typesense/api/pointer"
	"github.com/stretchr/testify/require"
)

func TestCollectionSearch(t *testing.T) {
	collectionName := createNewCollection(t, "companies")
	documents := []interface{}{
		newDocument("123", withCompanyName("Company 1"), withNumEmployees(50)),
		newDocument("125", withCompanyName("Company 2"), withNumEmployees(150)),
		newDocument("127", withCompanyName("Company 3"), withNumEmployees(250)),
		newDocument("129", withCompanyName("Stark Industries 4"), withNumEmployees(500)),
		newDocument("131", withCompanyName("Stark Industries 5"), withNumEmployees(1000)),
	}

	params := &api.ImportDocumentsParams{Action: pointer.String("create")}
	_, err := typesenseClient.Collection(collectionName).Documents().Import(documents, params)
	require.NoError(t, err)

	searchParams := &api.SearchCollectionParams{
		Q:              "Company",
		QueryBy:        "company_name, company_name",
		QueryByWeights: pointer.String("2, 1"),
		FilterBy:       pointer.String("num_employees:>=100"),
		SortBy:         pointer.String("num_employees:desc"),
		NumTypos:       pointer.String("2"),
		Page:           pointer.Int(1),
		PerPage:        pointer.Int(10),
	}

	expectedDocs := []map[string]interface{}{
		newDocumentResponse("127", withResponseCompanyName("Company 3"),
			withResponseNumEmployees(250)),
		newDocumentResponse("125", withResponseCompanyName("Company 2"),
			withResponseNumEmployees(150)),
	}

	result, err := typesenseClient.Collection(collectionName).Documents().Search(searchParams)

	require.NoError(t, err)
	require.Equal(t, 2, *result.Found, "found documents number is invalid")
	require.Equal(t, 2, len(*result.Hits), "number of hits is invalid")

	docs := make([]map[string]interface{}, len(*result.Hits))
	for i, hit := range *result.Hits {
		docs[i] = *hit.Document
	}

	require.Equal(t, expectedDocs, docs)
}

func TestCollectionSearchRange(t *testing.T) {
	collectionName := createNewCollection(t, "companies")
	documents := []interface{}{
		newDocument("123", withCompanyName("Company 1"), withNumEmployees(50)),
		newDocument("125", withCompanyName("Company 2"), withNumEmployees(150)),
		newDocument("127", withCompanyName("Company 3"), withNumEmployees(250)),
		newDocument("129", withCompanyName("Stark Industries 4"), withNumEmployees(500)),
	}

	params := &api.ImportDocumentsParams{Action: pointer.String("create")}
	_, err := typesenseClient.Collection(collectionName).Documents().Import(documents, params)
	require.NoError(t, err)

	searchParams := &api.SearchCollectionParams{
		Q:        "*",
		FilterBy: pointer.String("num_employees:>=100&&num_employees:<=300"),
		SortBy:   pointer.String("num_employees:asc"),
		Page:     pointer.Int(1),
		PerPage:  pointer.Int(10),
		QueryBy:  "company_name, country",
	}

	expectedDocs := []map[string]interface{}{
		newDocumentResponse("125", withResponseCompanyName("Company 2"),
			withResponseNumEmployees(150)),
		newDocumentResponse("127", withResponseCompanyName("Company 3"),
			withResponseNumEmployees(250)),
	}

	result, err := typesenseClient.Collection(collectionName).Documents().Search(searchParams)

	require.NoError(t, err)
	require.Equal(t, 2, *result.Found, "found documents number is invalid")
	require.Equal(t, 2, len(*result.Hits), "number of hits is invalid")

	docs := make([]map[string]interface{}, len(*result.Hits))
	for i, hit := range *result.Hits {
		docs[i] = *hit.Document
	}

	require.Equal(t, expectedDocs, docs)
}
