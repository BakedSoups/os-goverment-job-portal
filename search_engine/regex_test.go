package search_engine

import "testing"

func TestMatcherRequiresITContextForInfrastructure(t *testing.T) {
	matcher := NewMatcher()

	physical := matcher.Match("Lead waterfront infrastructure development, construction coordination, and climate resilience projects.")
	if physical["cloud_infrastructure"] != 0 {
		t.Fatalf("physical infrastructure matched cloud_infrastructure: %#v", physical)
	}

	technical := matcher.Match("Manage AWS cloud infrastructure with Terraform, Kubernetes, and disaster recovery controls.")
	if technical["cloud_infrastructure"] == 0 {
		t.Fatalf("technical infrastructure did not match cloud_infrastructure: %#v", technical)
	}
}

func TestMatcherRequiresEnterpriseContextForArchitecture(t *testing.T) {
	matcher := NewMatcher()

	building := matcher.Match("Review historic architecture and building preservation requirements.")
	if building["enterprise_architecture"] != 0 {
		t.Fatalf("building architecture matched enterprise_architecture: %#v", building)
	}

	technical := matcher.Match("Define enterprise architecture and systems architecture standards.")
	if technical["enterprise_architecture"] == 0 {
		t.Fatalf("technical architecture did not match enterprise_architecture: %#v", technical)
	}
}

func TestMatcherMapsFrameworksAndToolsToCanonicalSkills(t *testing.T) {
	matcher := NewMatcher()
	hits := matcher.Match("Built Django APIs served by Uvicorn, stored data in PostgreSQL, and deployed containers with Docker.")

	for _, concept := range []string{"python", "sql", "devops"} {
		if hits[concept] == 0 {
			t.Errorf("expected %q canonical concept match, got %#v", concept, hits)
		}
	}
}

func TestMatcherDoesNotPromoteGenericWordsToDomainSkills(t *testing.T) {
	matcher := NewMatcher()
	hits := matcher.Match("Maintain a software application and protect the property of its users.")

	for _, concept := range []string{"facilities_maintenance", "real_estate_property"} {
		if hits[concept] != 0 {
			t.Errorf("generic text unexpectedly matched %q: %#v", concept, hits)
		}
	}
}
