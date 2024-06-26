package resolvers

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
)

type complianceAggregationResponseWithDomainResolver struct {
	complianceAggregation_ResponseResolver
	domainMap map[*storage.ComplianceAggregation_Result]*storage.ComplianceDomain
}

func (r *complianceAggregationResponseWithDomainResolver) Results(ctx context.Context) ([]*complianceAggregationResultWithDomainResolver, error) {
	results, err := r.complianceAggregation_ResponseResolver.Results(ctx)
	if err != nil {
		return nil, err
	}

	wrappedResults := make([]*complianceAggregationResultWithDomainResolver, len(results))
	for i, result := range results {
		wrappedResults[i] = &complianceAggregationResultWithDomainResolver{
			complianceAggregation_ResultResolver: *result,
			domain:                               r.domainMap[results[i].data],
		}
	}
	return wrappedResults, nil
}

type complianceAggregationResultWithDomainResolver struct {
	complianceAggregation_ResultResolver
	domain *storage.ComplianceDomain
}
