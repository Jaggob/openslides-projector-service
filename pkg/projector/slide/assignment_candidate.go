package slide

import (
	"context"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func AssignmentCandidateSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	cQ := req.Fetch.AssignmentCandidate(*req.ContentObjectID)
	candidate, err := cQ.Preload(cQ.MeetingUser().User()).Preload(cQ.Assignment()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load assignment candidate id %w", err)
	}

	candidateName := req.Locale.Get("Unknown user")
	if candidate.MeetingUser != nil {
		if mu, ok := candidate.MeetingUser.Value(); ok && mu.User != nil {
			candidateName = viewmodels.User_ShortName(mu.User)
		}
	}

	return map[string]any{
		"Assignment":    candidate.Assignment,
		"CandidateName": candidateName,
		"Application":   template.HTML(candidate.Application),
	}, nil
}
