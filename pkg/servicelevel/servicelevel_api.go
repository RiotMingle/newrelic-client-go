// Code generated by tutone: DO NOT EDIT
package servicelevel

import (
	"context"

	"github.com/RiotMingle/newrelic-client-go/pkg/common"
	"github.com/RiotMingle/newrelic-client-gogo/pkg/errors"
)

// Creates a new SLI.
func (a *Servicelevel) ServiceLevelCreate(
	entityGUID common.EntityGUID,
	indicator ServiceLevelIndicatorCreateInput,
) (*ServiceLevelIndicator, error) {
	return a.ServiceLevelCreateWithContext(context.Background(),
		entityGUID,
		indicator,
	)
}

// Creates a new SLI.
func (a *Servicelevel) ServiceLevelCreateWithContext(
	ctx context.Context,
	entityGUID common.EntityGUID,
	indicator ServiceLevelIndicatorCreateInput,
) (*ServiceLevelIndicator, error) {

	resp := ServiceLevelCreateQueryResponse{}
	vars := map[string]interface{}{
		"entityGuid": entityGUID,
		"indicator":  indicator,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, ServiceLevelCreateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.ServiceLevelIndicator, nil
}

type ServiceLevelCreateQueryResponse struct {
	ServiceLevelIndicator ServiceLevelIndicator `json:"ServiceLevelCreate"`
}

const ServiceLevelCreateMutation = `mutation(
	$entityGuid: EntityGuid!,
	$indicator: ServiceLevelIndicatorCreateInput!,
) { serviceLevelCreate(
	entityGuid: $entityGuid,
	indicator: $indicator,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entityGuid
	events {
		account {
			id
			name
		}
		badEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		goodEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		validEvents {
			from
			select {
				attribute
				function
			}
			where
		}
	}
	guid
	id
	name
	objectives {
		description
		name
		target
		timeWindow {
			rolling {
				count
				unit
			}
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Deletes an existing SLI by the GUID.
// iD optional field manually removed in favor of gUID to favor it's deprecation
func (a *Servicelevel) ServiceLevelDelete(
	gUID common.EntityGUID,
) (*ServiceLevelIndicator, error) {
	return a.ServiceLevelDeleteWithContext(context.Background(),
		gUID,
	)
}

// Deletes an existing SLI by the GUID.
// iD optional field manually removed in favor of gUID to favor it's deprecation
func (a *Servicelevel) ServiceLevelDeleteWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
) (*ServiceLevelIndicator, error) {

	resp := ServiceLevelDeleteQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, ServiceLevelDeleteMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.ServiceLevelIndicator, nil
}

type ServiceLevelDeleteQueryResponse struct {
	ServiceLevelIndicator ServiceLevelIndicator `json:"ServiceLevelDelete"`
}

// id optional field manually removed in favor of guid to favor it's deprecation
const ServiceLevelDeleteMutation = `mutation(
	$guid: EntityGuid,
) { serviceLevelDelete(
	guid: $guid,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entityGuid
	events {
		account {
			id
			name
		}
		badEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		goodEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		validEvents {
			from
			select {
				attribute
				function
			}
			where
		}
	}
	guid
	id
	name
	objectives {
		description
		name
		target
		timeWindow {
			rolling {
				count
				unit
			}
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Updates an existing SLI by the GUID.
// iD optional field manually removed in favor of gUID to favor it's deprecation
func (a *Servicelevel) ServiceLevelUpdate(
	gUID common.EntityGUID,
	indicator ServiceLevelIndicatorUpdateInput,
) (*ServiceLevelIndicator, error) {
	return a.ServiceLevelUpdateWithContext(context.Background(),
		gUID,
		indicator,
	)
}

// Updates an existing SLI by the GUID.
// iD optional field manually removed in favor of gUID to favor it's deprecation
func (a *Servicelevel) ServiceLevelUpdateWithContext(
	ctx context.Context,
	gUID common.EntityGUID,
	indicator ServiceLevelIndicatorUpdateInput,
) (*ServiceLevelIndicator, error) {

	resp := ServiceLevelUpdateQueryResponse{}
	vars := map[string]interface{}{
		"guid":      gUID,
		"indicator": indicator,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, ServiceLevelUpdateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.ServiceLevelIndicator, nil
}

type ServiceLevelUpdateQueryResponse struct {
	ServiceLevelIndicator ServiceLevelIndicator `json:"ServiceLevelUpdate"`
}

// id optional field manually removed in favor of guid to favor it's deprecation
const ServiceLevelUpdateMutation = `mutation(
	$guid: EntityGuid,
	$indicator: ServiceLevelIndicatorUpdateInput!,
) { serviceLevelUpdate(
	guid: $guid,
	indicator: $indicator,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entityGuid
	events {
		account {
			id
			name
		}
		badEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		goodEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		validEvents {
			from
			select {
				attribute
				function
			}
			where
		}
	}
	guid
	id
	name
	objectives {
		description
		name
		target
		timeWindow {
			rolling {
				count
				unit
			}
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// The SLIs attached to the entity.
func (a *Servicelevel) GetIndicators(
	entityGUID common.EntityGUID,
) (*[]ServiceLevelIndicator, error) {
	return a.GetIndicatorsWithContext(context.Background(),
		entityGUID,
	)
}

// The SLIs attached to the entity.
func (a *Servicelevel) GetIndicatorsWithContext(
	ctx context.Context,
	entityGUID common.EntityGUID,
) (*[]ServiceLevelIndicator, error) {

	resp := indicatorsResponse{}
	vars := map[string]interface{}{
		"entityGUID": entityGUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getIndicatorsQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Entity.ServiceLevel.Indicators) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Entity.ServiceLevel.Indicators, nil
}

const getIndicatorsQuery = `query(
	$entityGUID: EntityGuid!,
) { actor { entity(guid: $entityGUID) { serviceLevel { indicators {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entityGuid
	events {
		account {
			id
			name
		}
		badEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		goodEvents {
			from
			select {
				attribute
				function
			}
			where
		}
		validEvents {
			from
			select {
				attribute
				function
			}
			where
		}
	}
	guid
	id
	name
	objectives {
		description
		name
		target
		timeWindow {
			rolling {
				count
				unit
			}
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} } } } }`
