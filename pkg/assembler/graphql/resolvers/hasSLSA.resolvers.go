package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestSlsa is the resolver for the ingestSLSA field.
func (r *mutationResolver) IngestSlsa(ctx context.Context, subject model.ArtifactInputSpec, builtFrom []*model.ArtifactInputSpec, builtBy model.BuilderInputSpec, slsa model.SLSAInputSpec) (string, error) {
	if len(builtFrom) < 1 {
		return "", gqlerror.Errorf("IngestSLSA :: Must have at least 1 builtFrom")
	}

	ingestedSLSA, err := r.Backend.IngestSLSA(ctx, subject, builtFrom, builtBy, slsa)
	if err != nil {
		return "", err
	}
	return ingestedSLSA.ID, err
}

// IngestSLSAs is the resolver for the ingestSLSAs field.
func (r *mutationResolver) IngestSLSAs(ctx context.Context, subjects []*model.ArtifactInputSpec, builtFromList [][]*model.ArtifactInputSpec, builtByList []*model.BuilderInputSpec, slsaList []*model.SLSAInputSpec) ([]string, error) {
	funcName := "IngestSLSAs"
	ingestedSLSAIDS := []string{}
	if len(subjects) != len(slsaList) {
		return ingestedSLSAIDS, gqlerror.Errorf("%v :: uneven subjects and slsa attestation for ingestion", funcName)
	}
	if len(subjects) != len(builtFromList) {
		return ingestedSLSAIDS, gqlerror.Errorf("%v :: uneven subjects and built from artifact list for ingestion", funcName)
	}
	if len(subjects) != len(builtByList) {
		return ingestedSLSAIDS, gqlerror.Errorf("%v :: uneven subjects and built by for ingestion", funcName)
	}

	ingestedSLSAs, err := r.Backend.IngestSLSAs(ctx, subjects, builtFromList, builtByList, slsaList)
	if err == nil {
		for _, SLSA := range ingestedSLSAs {
			ingestedSLSAIDS = append(ingestedSLSAIDS, SLSA.ID)
		}
	}
	return ingestedSLSAIDS, err
}

// HasSlsa is the resolver for the HasSLSA field.
func (r *queryResolver) HasSlsa(ctx context.Context, hasSLSASpec model.HasSLSASpec) ([]*model.HasSlsa, error) {
	return r.Backend.HasSlsa(ctx, &hasSLSASpec)
}
