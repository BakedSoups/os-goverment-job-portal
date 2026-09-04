package taxonomy

import "testing"

func TestTechnologyAliasesResolveToCanonicalConcepts(t *testing.T) {
	tests := map[string]string{
		"django":      "python",
		"uvicorn":     "python",
		"postgresql":  "sql",
		"mongodb":     "nosql",
		"spring boot": "java",
		"asp.net":     "dotnet",
		"react":       "javascript",
		"restful api": "api_development",
		"docker":      "devops",
		"law":         "legal",
	}

	aliases := AliasMap()
	for alias, want := range tests {
		concept, ok := aliases[alias]
		if !ok {
			t.Errorf("AliasMap() does not contain %q", alias)
			continue
		}
		if concept.Name != want {
			t.Errorf("AliasMap()[%q].Name = %q, want %q", alias, concept.Name, want)
		}
	}
}
