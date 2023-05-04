package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestCertifyGood is the resolver for the ingestCertifyGood field.
func (r *mutationResolver) IngestCertifyGood(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyGood model.CertifyGoodInputSpec) (*model.CertifyGood, error) {
	return r.Backend.IngestCertifyGood(ctx, subject, pkgMatchType, certifyGood)
}

// CertifyGood is the resolver for the CertifyGood field.
func (r *queryResolver) CertifyGood(ctx context.Context, certifyGoodSpec *model.CertifyGoodSpec) ([]*model.CertifyGood, error) {
	return r.Backend.CertifyGood(ctx, certifyGoodSpec)
}
