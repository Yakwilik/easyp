package rules_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"github.com/easyp-tech/easyp/internal/core"
	"github.com/easyp-tech/easyp/internal/rules"
)

func TestEnumNoAllowAlias_Message(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	const expMessage = "enum must not allow alias"

	rule := rules.EnumNoAllowAlias{}
	message := rule.Message()

	assert.Equal(expMessage, message)
}

func TestEnumNoAllowAlias_Validate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		fileName   string
		wantIssues *core.Issue
		wantErr    error
	}{
		"invalid": {
			fileName: invalidAuthProto,
			wantIssues: &core.Issue{
				Position: meta.Position{
					Filename: "",
					Offset:   864,
					Line:     49,
					Column:   1,
				},
				SourceName: "social_network",
				Message:    "enum must not allow alias",
				RuleName:   "ENUM_NO_ALLOW_ALIAS",
			},
			wantErr: nil,
		},
		"invalid_nested": {
			fileName: invalidAuthProto,
			wantIssues: &core.Issue{
				Position: meta.Position{
					Filename: "",
					Offset:   610,
					Line:     31,
					Column:   3,
				},
				SourceName: "social_network",
				Message:    "enum must not allow alias",
				RuleName:   "ENUM_NO_ALLOW_ALIAS",
			},
			wantErr: nil,
		},
		"valid": {
			fileName: validAuthProto,
			wantErr:  nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, protos := start(t)

			rule := rules.EnumNoAllowAlias{}
			issues, err := rule.Validate(protos[tc.fileName])
			r.ErrorIs(err, tc.wantErr)
			switch {
			case tc.wantIssues != nil:
				r.Contains(issues, *tc.wantIssues)
			case len(issues) > 0:
				r.Empty(issues)
			}
		})
	}
}
