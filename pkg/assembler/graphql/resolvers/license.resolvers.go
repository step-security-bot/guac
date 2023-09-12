package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/backends/helper"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestLicense is the resolver for the ingestLicense field.
func (r *mutationResolver) IngestLicense(ctx context.Context, license *model.LicenseInputSpec) (string, error) {
	if err := helper.ValidateLicenseInput(license); err != nil {
		return "", err
	}
	il, err := r.Backend.IngestLicense(ctx, license)
	if err != nil {
		return "", err
	}
	return il.ID, nil
}

// IngestLicenses is the resolver for the ingestLicenses field.
func (r *mutationResolver) IngestLicenses(ctx context.Context, licenses []*model.LicenseInputSpec) ([]string, error) {
	for _, l := range licenses {
		if err := helper.ValidateLicenseInput(l); err != nil {
			return nil, err
		}
	}
	ils, err := r.Backend.IngestLicenses(ctx, licenses)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, il := range ils {
		ids = append(ids, il.ID)
	}
	return ids, nil
}

// Licenses is the resolver for the licenses field.
func (r *queryResolver) Licenses(ctx context.Context, licenseSpec model.LicenseSpec) ([]*model.License, error) {
	return r.Backend.Licenses(ctx, &licenseSpec)
}
