package dependencytrack

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestGetViolations tests listing policy violations
func TestGetViolations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	now := time.Now().Truncate(time.Second)

	mux.HandleFunc("/api/v1/violation", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Method; got != http.MethodGet {
			t.Errorf("Got request method %v, want %v", got, http.MethodGet)
		}
		if got := r.URL.Query().Get("suppressed"); !cmp.Equal(got, "true") {
			t.Errorf("Got query parameters: %v, want %v", got, "true")
		}
		if r.URL.Query().Get("pageNumber") == "1" && r.URL.Query().Get("pageSize") == "100" {
			fmt.Fprintf(w,
				`
			[
 			  {
			    "analysis": {
			      "analysisState": "APPROVED",
			      "isSuppressed": true
			    },
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "foo",
			      "version": "bar",
			      "active": true,
			      "lastBomImport": %d,
			      "metrics": {
			        "critical": 0,
			        "high": 1,
			        "low": 2,
			        "medium": 3,
			        "unassigned": 4,
			        "inheritedRiskScore": 1240
			      },
			      "uuid": "fd1b10b9-678d-4af9-ad8e-877d1f357b03"
			    },
			    "type": "SECURITY"
			  },
			  {
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "bar",
			      "version": "foo",
			      "active": false,
			      "metrics": {
			        "critical": 50,
			        "high": 25,
			        "low": 12,
			        "medium": 6,
			        "unassigned": 3,
			        "inheritedRiskScore": 2560.26
			      },
			      "uuid": "9b9a702a-a8b4-49fb-bb99-c05c1a8c8d49"
			    },
			    "type": "LICENSE"
			  }
			]
			`,
				now.Unix(),
			)
		}
		if r.URL.Query().Get("pageNumber") == "2" && r.URL.Query().Get("pageSize") == "100" {
			fmt.Fprintf(w,
				`
			[
 			  {
			    "analysis": {
			      "analysisState": "APPROVED",
			      "isSuppressed": true
			    },
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "foo",
			      "version": "bar",
			      "active": true,
			      "lastBomImport": %d,
			      "metrics": {
			        "critical": 0,
			        "high": 1,
			        "low": 2,
			        "medium": 3,
			        "unassigned": 4,
			        "inheritedRiskScore": 1240
			      },
			      "uuid": "fd1b10b8-678d-4af9-ad8e-877d1f357b03"
			    },
			    "type": "SECURITY"
			  },
			  {
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "bar",
			      "version": "foo",
			      "active": false,
			      "metrics": {
			        "critical": 50,
			        "high": 25,
			        "low": 12,
			        "medium": 6,
			        "unassigned": 3,
			        "inheritedRiskScore": 2560.26
			      },
			      "uuid": "9b9a703a-a8b4-49fb-bb99-c05c1a8c8d49"
			    },
			    "type": "LICENSE"
			  }
			]
			`,
				now.Unix(),
			)
		}
		if r.Header.Get("pageNumber") == "3" && r.Header.Get("pageSize") == "100" {
			fmt.Fprintf(w, `[]`)
		}
	})

	got, err := client.GetAllViolations(true)
	if err != nil {
		t.Errorf("GetViolations returned error: %v", err)
	}

	want := []*PolicyViolation{
		{
			Analysis: &ViolationAnalysis{
				AnalysisState: "APPROVED",
				IsSuppressed:  true,
			},
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "foo",
				Version:       "bar",
				Active:        true,
				LastBomImport: Time{now},
				Metrics: ProjectMetrics{
					Critical:           0,
					High:               1,
					Low:                2,
					Medium:             3,
					Unassigned:         4,
					InheritedRiskScore: 1240,
				},
				UUID: "fd1b10b9-678d-4af9-ad8e-877d1f357b03",
			},
			Type: "SECURITY",
		},
		{
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "bar",
				Version:       "foo",
				Active:        false,
				LastBomImport: Time{},
				Metrics: ProjectMetrics{
					Critical:           50,
					High:               25,
					Low:                12,
					Medium:             6,
					Unassigned:         3,
					InheritedRiskScore: 2560.26,
				},
				UUID: "9b9a702a-a8b4-49fb-bb99-c05c1a8c8d49",
			},
			Type: "LICENSE",
		},
		{
			Analysis: &ViolationAnalysis{
				AnalysisState: "APPROVED",
				IsSuppressed:  true,
			},
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "foo",
				Version:       "bar",
				Active:        true,
				LastBomImport: Time{now},
				Metrics: ProjectMetrics{
					Critical:           0,
					High:               1,
					Low:                2,
					Medium:             3,
					Unassigned:         4,
					InheritedRiskScore: 1240,
				},
				UUID: "fd1b10b8-678d-4af9-ad8e-877d1f357b03",
			},
			Type: "SECURITY",
		},
		{
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "bar",
				Version:       "foo",
				Active:        false,
				LastBomImport: Time{},
				Metrics: ProjectMetrics{
					Critical:           50,
					High:               25,
					Low:                12,
					Medium:             6,
					Unassigned:         3,
					InheritedRiskScore: 2560.26,
				},
				UUID: "9b9a703a-a8b4-49fb-bb99-c05c1a8c8d49",
			},
			Type: "LICENSE",
		},
	}

	if !cmp.Equal(got, want) {
		t.Errorf("Got violations %v, want %v", got, want)
	}
}

// TestGetViolationsByPage tests listing policy violations by page
func TestGetViolationsByPage(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	now := time.Now().Truncate(time.Second)

	mux.HandleFunc("/api/v1/violation", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Method; got != http.MethodGet {
			t.Errorf("Got request method %v, want %v", got, http.MethodGet)
		}
		if got := r.URL.Query().Get("suppressed"); !cmp.Equal(got, "true") {
			t.Errorf("Got query parameters: %v, want %v", got, "true")
		}
		if r.URL.Query().Get("pageNumber") == "1" && r.URL.Query().Get("pageSize") == "1" {
			fmt.Fprintf(w,
				`
			[
			  {
			    "analysis": {
			      "analysisState": "APPROVED",
			      "isSuppressed": true
			    },
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "foo",
			      "version": "bar",
			      "active": true,
			      "lastBomImport": %d,
			      "metrics": {
			        "critical": 0,
			        "high": 1,
			        "low": 2,
			        "medium": 3,
			        "unassigned": 4,
			        "inheritedRiskScore": 1240
			      },
			      "uuid": "fd1b10b9-678d-4af9-ad8e-877d1f357b03"
			    },
			    "type": "SECURITY"
			  }
			]
			`,
				now.Unix(),
			)
		}
		if r.URL.Query().Get("pageNumber") == "2" && r.URL.Query().Get("pageSize") == "1" {
			fmt.Fprintf(w,
				`
			[
			  {
			    "policyCondition": {
			      "policy": {
			        "violationState": "WARN"
			      }
			    },
			    "project": {
			      "name": "bar",
			      "version": "foo",
			      "active": false,
				  "lastBomImport": %d,
			      "metrics": {
			        "critical": 50,
			        "high": 25,
			        "low": 12,
			        "medium": 6,
			        "unassigned": 3,
			        "inheritedRiskScore": 2560.26
			      },
			      "uuid": "9b9a702a-a8b4-49fb-bb99-c05c1a8c8d49"
			    },
			    "type": "LICENSE"
			  }
			]
			`,
				now.Unix(),
			)
		}
	})

	gotPageOne, err := client.GetViolations(true, 1, 1)
	if err != nil {
		t.Errorf("GetViolations returned error: %v", err)
	}

	wantPageOne := []*PolicyViolation{
		{
			Analysis: &ViolationAnalysis{
				AnalysisState: "APPROVED",
				IsSuppressed:  true,
			},
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "foo",
				Version:       "bar",
				Active:        true,
				LastBomImport: Time{now},
				Metrics: ProjectMetrics{
					Critical:           0,
					High:               1,
					Low:                2,
					Medium:             3,
					Unassigned:         4,
					InheritedRiskScore: 1240,
				},
				UUID: "fd1b10b9-678d-4af9-ad8e-877d1f357b03",
			},
			Type: "SECURITY",
		},
	}

	if !cmp.Equal(gotPageOne, wantPageOne) {
		t.Errorf("Got violations %v, want %v", gotPageOne, wantPageOne)
	}

	gotPageTwo, err := client.GetViolations(true, 2, 1)
	if err != nil {
		t.Errorf("GetViolations returned error: %v", err)
	}

	wantPageTwo := []*PolicyViolation{
		{
			PolicyCondition: PolicyCondition{
				Policy: Policy{
					ViolationState: "WARN",
				},
			},
			Project: Project{
				Name:          "bar",
				Version:       "foo",
				Active:        false,
				LastBomImport: Time{now},
				Metrics: ProjectMetrics{
					Critical:           50,
					High:               25,
					Low:                12,
					Medium:             6,
					Unassigned:         3,
					InheritedRiskScore: 2560.26,
				},
				UUID: "9b9a702a-a8b4-49fb-bb99-c05c1a8c8d49",
			},
			Type: "LICENSE",
		},
	}

	if !cmp.Equal(gotPageTwo, wantPageTwo) {
		t.Errorf("Got violations %v, want %v", gotPageTwo, wantPageTwo)
	}
}
