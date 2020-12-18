// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

// ApiKey defines model for ApiKey.
type ApiKey struct {
	Actions     []string `json:"actions"`
	Collections []string `json:"collections"`
	Description string   `json:"description"`
	Id          *string  `json:"id,omitempty"`
	Value       *string  `json:"value,omitempty"`
	ValuePrefix *string  `json:"value_prefix,omitempty"`
}

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Message *string `json:"message,omitempty"`
}

// Collection defines model for Collection.
type Collection struct {
	// Embedded struct due to allOf(#/components/schemas/CollectionSchema)
	CollectionSchema
	// Embedded fields due to inline allOf schema

	// Number of documents in the collection
	NumDocuments int64 `json:"num_documents"`
}

// CollectionAlias defines model for CollectionAlias.
type CollectionAlias struct {

	// Name of the collection you wish to map the alias to
	CollectionName string  `json:"collection_name"`
	Name           *string `json:"name,omitempty"`
}

// CollectionSchema defines model for CollectionSchema.
type CollectionSchema struct {

	// The name of an int32 / float field that determines the order in which the search results are ranked when a sort_by clause is not provided during searching. This field must indicate some kind of popularity.
	DefaultSortingField string `json:"default_sorting_field"`

	// A list of fields for querying, filtering and faceting
	Fields []Field `json:"fields"`

	// Name of the collection
	Name string `json:"name"`
}

// CollectionsResponse defines model for CollectionsResponse.
type CollectionsResponse struct {
	Collections []*Collection `json:"collections"`
}

// Field defines model for Field.
type Field struct {
	Facet    bool   `json:"facet"`
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
	Type     string `json:"type"`
}

// SearchOverride defines model for SearchOverride.
type SearchOverride struct {

	// List of document `id`s that should be excluded from the search results.
	Excludes *[]struct {

		// document id that should be excluded from the search results.
		Id string `json:"id"`
	} `json:"excludes,omitempty"`

	// List of document `id`s that should be included in the search results with their corresponding `position`s.
	Includes *[]struct {

		// document id that should be included
		Id string `json:"id"`

		// position number where document should be included in the search results
		Position int `json:"position"`
	} `json:"includes,omitempty"`
	Rule struct {
		Id *string `json:"id,omitempty"`

		// Indicates whether the match on the query term should be `exact` or `contains`. If we want to match all queries that contained the word `apple`, we will use the `contains` match instead.
		Match string `json:"match"`

		// Indicates what search queries should be overridden
		Query string `json:"query"`
	} `json:"rule"`
}

// SearchResultHit defines model for SearchResultHit.
type SearchResultHit struct {

	// Can be any key-value pair
	Document map[string]interface{} `json:"document"`

	// Contains highlighted portions of the search fields
	Highlights *[]struct {
		Field *string `json:"field,omitempty"`

		// The indices property will be present only for string[] fields and will contain the corresponding indices of the snippets in the search field
		Indices *[]int `json:"indices,omitempty"`

		// Present only for (non-array) string fields
		Snippet *string `json:"snippet,omitempty"`

		// Present only for (array) string[] fields
		Snippets *[]string `json:"snippets,omitempty"`
	} `json:"highlights,omitempty"`
}

// SearchResultHits defines model for SearchResultHits.
type SearchResultHits []SearchResultHit

// UpsertAliasJSONBody defines parameters for UpsertAlias.
type UpsertAliasJSONBody CollectionAlias

// CreateCollectionJSONBody defines parameters for CreateCollection.
type CreateCollectionJSONBody CollectionSchema

// DeleteDocumentsParams defines parameters for DeleteDocuments.
type DeleteDocumentsParams struct {

	// Filter conditions for refining your documents to delete. Separate multiple conditions with &&.
	FilterBy string `json:"filter_by"`

	// Batch size parameter controls the number of documents that should be deleted at a time. A larger value will speed up deletions, but will impact performance of other operations running on the server.
	BatchSize int `json:"batch_size"`
}

// IndexDocumentJSONBody defines parameters for IndexDocument.
type IndexDocumentJSONBody interface{}

// IndexDocumentParams defines parameters for IndexDocument.
type IndexDocumentParams struct {

	// Additional action to perform
	Action *string `json:"action,omitempty"`
}

// ImportCollectionParams defines parameters for ImportCollection.
type ImportCollectionParams struct {

	// Action mode. Allowed action modes are `create`, `upsert` and `update`. `create` mode creates a new document. Fails if a document with the same id already exists. `upsert` mode creates a new document or updates an existing document if a document with the same id already exists. `update` mode updates an existing document. Fails if a document with the given id does not exist.
	Action *string `json:"action,omitempty"`

	// Batch size used for import. By default, Typesense ingests 40 documents at a time into Typesense. To increase this value, use the `batch_size` parameter. Larger batch sizes will consume larger transient memory during import.
	BatchSize *int `json:"batch_size,omitempty"`
}

// SearchCollectionParams defines parameters for SearchCollection.
type SearchCollectionParams struct {

	// The query text to search for in the collection. Use * as the search string to return all documents. This is typically useful when used in conjunction with filter_by.
	Q string `json:"q"`

	// A list of `string` or `string[]` fields that should be queried against. Multiple fields are separated with a comma.
	QueryBy []string `json:"query_by"`

	// Maximum number of hits returned. Increasing this value might increase search latency. Default: 500. Use `all` to return all hits found.
	MaxHits *interface{} `json:"max_hits,omitempty"`

	// Boolean field to indicate that the last word in the query should be treated as a prefix, and not as a whole word. This is used for building autocomplete and instant search interfaces. Defaults to true.
	Prefix *bool `json:"prefix,omitempty"`

	// Filter conditions for refining your search results. Separate multiple conditions with &&.
	FilterBy *string `json:"filter_by,omitempty"`

	// A list of numerical fields and their corresponding sort orders that will be used for ordering your results. Up to 3 sort fields can be specified. The text similarity score is exposed as a special `_text_match` field that you can use in the list of sorting fields. If no `sort_by` parameter is specified, results are sorted by `_text_match:desc,default_sorting_field:desc`
	SortBy *[]string `json:"sort_by,omitempty"`

	// A list of fields that will be used for faceting your results on. Separate multiple fields with a comma.
	FacetBy *[]string `json:"facet_by,omitempty"`

	// Maximum number of facet values to be returned.
	MaxFacetValues *int `json:"max_facet_values,omitempty"`

	// Facet values that are returned can now be filtered via this parameter. The matching facet text is also highlighted. For example, when faceting by `category`, you can set `facet_query=category:shoe` to return only facet values that contain the prefix "shoe".
	FacetQuery *string `json:"facet_query,omitempty"`

	// The number of typographical errors (1 or 2) that would be tolerated. Default: 2
	NumTypos *int `json:"num_typos,omitempty"`

	// Results from this specific page number would be fetched.
	Page *int `json:"page,omitempty"`

	// Number of results to fetch per page. Default: 10
	PerPage *int `json:"per_page,omitempty"`

	// You can aggregate search results into groups or buckets by specify one or more `group_by` fields. Separate multiple fields with a comma. To group on a particular field, it must be a faceted field.
	GroupBy *[]string `json:"group_by,omitempty"`

	// Maximum number of hits to be returned for every group. If the `group_limit` is set as `K` then only the top K hits in each group are returned in the response. Default: 3
	GroupLimit *int `json:"group_limit,omitempty"`

	// List of fields from the document to include in the search result
	IncludeFields *[]string `json:"include_fields,omitempty"`

	// List of fields from the document to exclude in the search result
	ExcludeFields *[]string `json:"exclude_fields,omitempty"`

	// List of fields which should be highlighted fully without snippeting
	HighlightFullFields *[]string `json:"highlight_full_fields,omitempty"`

	// The number of tokens that should surround the highlighted text on each side. Default: 4
	HighlightAffixNumTokens *int `json:"highlight_affix_num_tokens,omitempty"`

	// The start tag used for the highlighted snippets. Default: `<mark>`
	HighlightStartTag *string `json:"highlight_start_tag,omitempty"`

	// The end tag used for the highlighted snippets. Default: `</mark>`
	HighlightEndTag *string `json:"highlight_end_tag,omitempty"`

	// Field values under this length will be fully highlighted, instead of showing a snippet of relevant portion. Default: 30
	SnippetThreshold *int `json:"snippet_threshold,omitempty"`

	// If the number of results found for a specific query is less than this number, Typesense will attempt to drop the tokens in the query until enough results are found. Tokens that have the least individual hits are dropped first. Set to 0 to disable. Default: 10
	DropTokensThreshold *int `json:"drop_tokens_threshold,omitempty"`

	// If the number of results found for a specific query is less than this number, Typesense will attempt to look for tokens with more typos until enough results are found. Default: 100
	TypoTokensThreshold *int `json:"typo_tokens_threshold,omitempty"`

	// A list of records to unconditionally include in the search results at specific positions. An example use case would be to feature or promote certain items on the top of search results. A list of `record_id:hit_position`. Eg: to include a record with ID 123 at Position 1 and another record with ID 456 at Position 5, you'd specify `123:1,456:5`.
	// You could also use the Overrides feature to override search results based on rules. Overrides are applied first, followed by `pinned_hits` and  finally `hidden_hits`.
	PinnedHits *[]string `json:"pinned_hits,omitempty"`

	// A list of records to unconditionally hide from search results. A list of `record_id`s to hide. Eg: to hide records with IDs 123 and 456, you'd specify `123,456`.
	// You could also use the Overrides feature to override search results based on rules. Overrides are applied first, followed by `pinned_hits` and finally `hidden_hits`.
	HiddenHits *[]string `json:"hidden_hits,omitempty"`
}

// UpdateDocumentJSONBody defines parameters for UpdateDocument.
type UpdateDocumentJSONBody interface{}

// UpsertSearchOverrideJSONBody defines parameters for UpsertSearchOverride.
type UpsertSearchOverrideJSONBody SearchOverride

// CreateKeyJSONBody defines parameters for CreateKey.
type CreateKeyJSONBody ApiKey

// UpsertAliasRequestBody defines body for UpsertAlias for application/json ContentType.
type UpsertAliasJSONRequestBody UpsertAliasJSONBody

// CreateCollectionRequestBody defines body for CreateCollection for application/json ContentType.
type CreateCollectionJSONRequestBody CreateCollectionJSONBody

// IndexDocumentRequestBody defines body for IndexDocument for application/json ContentType.
type IndexDocumentJSONRequestBody IndexDocumentJSONBody

// UpdateDocumentRequestBody defines body for UpdateDocument for application/json ContentType.
type UpdateDocumentJSONRequestBody UpdateDocumentJSONBody

// UpsertSearchOverrideRequestBody defines body for UpsertSearchOverride for application/json ContentType.
type UpsertSearchOverrideJSONRequestBody UpsertSearchOverrideJSONBody

// CreateKeyRequestBody defines body for CreateKey for application/json ContentType.
type CreateKeyJSONRequestBody CreateKeyJSONBody