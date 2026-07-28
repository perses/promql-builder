// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promqlbuilder

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/prometheus/prometheus/promql/parser"
)

// Well-known Perses dashboard variables and their syntactically valid PromQL
// placeholder values. The placeholders are intentionally unique durations or
// strings unlikely to collide with real metric/label content.
var defaultVarReplacements = map[string]string{
	"$__rate_interval": "2d20h8m7s",
	"$__interval":      "2d20h8m8s",
	"$__interval_ms":   "7d19h59m27s",
	"$__dashboard":     "PERSESVAR_dashboard",
	"$__project":       "PERSESVAR_project",
	"$__from":          "1715222400000.000",
	"$__to":            "1715222400000.000",
	"$__range":         "2d20h8m9s",
	"$__range_s":       "1h2m17s",
	"$__range_ms":      "3737373",
}

// userVarPattern matches user-defined Perses dashboard variables like $job,
// $namespace, $cluster — but not $__ prefixed ones (handled separately).
var userVarPattern = regexp.MustCompile(`\$([a-zA-Z_][a-zA-Z0-9_]*)`)

// Validate checks whether the given expression produces syntactically valid
// PromQL. It handles Perses dashboard variables ($__rate_interval, $job, etc.)
// by substituting them with dummy values before parsing.
func Validate(expr parser.Expr) error {
	return ValidateWithVars(expr, nil)
}

// ValidateWithVars checks whether the given expression produces syntactically
// valid PromQL. extraVars is an optional list of additional variable names
// (without the $ prefix) to substitute before parsing. Built-in Perses
// variables and any $varname patterns are handled automatically.
func ValidateWithVars(expr parser.Expr, extraVars []string) (retErr error) {
	defer func() {
		if r := recover(); r != nil {
			retErr = fmt.Errorf("invalid PromQL expression: rendering failed: %v", r)
		}
	}()

	rendered := expr.Pretty(0)
	sanitized := substituteVars(rendered, extraVars)

	p := parser.NewParser(parser.Options{})
	_, err := p.ParseExpr(sanitized)
	if err != nil {
		return fmt.Errorf("invalid PromQL expression: %w\nrendered: %s", err, strings.TrimSpace(rendered))
	}
	return nil
}

func substituteVars(query string, extraVars []string) string {
	// Build sorted keys (longest first to avoid partial replacements).
	type replacement struct {
		varName     string
		placeholder string
	}

	var reps []replacement
	for k, v := range defaultVarReplacements {
		reps = append(reps, replacement{k, v})
	}
	for _, name := range extraVars {
		varName := "$" + name
		if _, exists := defaultVarReplacements[varName]; !exists {
			reps = append(reps, replacement{varName, "PERSESVAR_" + name})
		}
	}

	sort.Slice(reps, func(i, j int) bool {
		return len(reps[i].varName) > len(reps[j].varName)
	})

	for _, r := range reps {
		query = strings.ReplaceAll(query, r.varName, r.placeholder)
	}

	// Catch remaining user-defined variables not in defaultVarReplacements or extraVars.
	query = userVarPattern.ReplaceAllStringFunc(query, func(match string) string {
		name := match[1:]
		// Skip if it looks like it was already replaced.
		if strings.HasPrefix(name, "_") {
			return match
		}
		return "PERSESVAR_" + name
	})

	return query
}
