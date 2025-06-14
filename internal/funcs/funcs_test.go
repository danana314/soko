package funcs

import (
	"1008001/splitwiser/internal/models"
	"html/template"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "standard date",
			input:    time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			expected: "2023-12-25",
		},
		{
			name:     "leap year date",
			input:    time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: "2024-02-29",
		},
		{
			name:     "new year",
			input:    time.Date(2025, 1, 1, 23, 59, 59, 0, time.UTC),
			expected: "2025-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatTime(tt.input)
			if result != tt.expected {
				t.Errorf("formatTime(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDateRange(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			name:     "same day",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 1, 23, 59, 59, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "one day apart",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "week apart",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: 7,
		},
		{
			name:     "reversed order",
			start:    time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dateRange(tt.start, tt.end)
			if len(result) != tt.expected {
				t.Errorf("dateRange() returned %d dates, want %d", len(result), tt.expected)
			}
		})
	}
}

func TestApproxDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "less than second",
			duration: 500 * time.Millisecond,
			expected: "less than 1 second",
		},
		{
			name:     "one second",
			duration: time.Second,
			expected: "1 second",
		},
		{
			name:     "multiple seconds",
			duration: 30 * time.Second,
			expected: "30 seconds",
		},
		{
			name:     "one minute",
			duration: time.Minute,
			expected: "1 minute",
		},
		{
			name:     "multiple minutes",
			duration: 45 * time.Minute,
			expected: "45 minutes",
		},
		{
			name:     "one hour",
			duration: time.Hour,
			expected: "1 hour",
		},
		{
			name:     "multiple hours",
			duration: 12 * time.Hour,
			expected: "12 hours",
		},
		{
			name:     "one day",
			duration: 24 * time.Hour,
			expected: "1 day",
		},
		{
			name:     "multiple days",
			duration: 30 * 24 * time.Hour,
			expected: "30 days",
		},
		{
			name:     "one year",
			duration: 365 * 24 * time.Hour,
			expected: "1 year",
		},
		{
			name:     "multiple years",
			duration: 2 * 365 * 24 * time.Hour,
			expected: "2 years",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := approxDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("approxDuration(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestPluralize(t *testing.T) {
	tests := []struct {
		name      string
		count     any
		singular  string
		plural    string
		expected  string
		expectErr bool
	}{
		{
			name:     "one item",
			count:    1,
			singular: "item",
			plural:   "items",
			expected: "item",
		},
		{
			name:     "zero items",
			count:    0,
			singular: "item",
			plural:   "items",
			expected: "items",
		},
		{
			name:     "multiple items",
			count:    5,
			singular: "item",
			plural:   "items",
			expected: "items",
		},
		{
			name:     "string number",
			count:    "1",
			singular: "cat",
			plural:   "cats",
			expected: "cat",
		},
		{
			name:      "invalid type",
			count:     "invalid",
			singular:  "item",
			plural:    "items",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := pluralize(tt.count, tt.singular, tt.plural)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("pluralize(%v, %q, %q) = %q, want %q", tt.count, tt.singular, tt.plural, result, tt.expected)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple text",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			name:     "with special characters",
			input:    "Hello, World!",
			expected: "hello-world",
		},
		{
			name:     "with numbers and underscores",
			input:    "Test_123",
			expected: "test_123",
		},
		{
			name:     "with dashes",
			input:    "already-slugified",
			expected: "already-slugified",
		},
		{
			name:     "with unicode",
			input:    "café naïve",
			expected: "caf-nave",
		},
		{
			name:     "multiple spaces",
			input:    "multiple   spaces",
			expected: "multiple---spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slugify(tt.input)
			if result != tt.expected {
				t.Errorf("slugify(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSafeHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected template.HTML
	}{
		{
			name:     "plain text",
			input:    "Hello World",
			expected: template.HTML("Hello World"),
		},
		{
			name:     "html tags",
			input:    "<b>Bold</b>",
			expected: template.HTML("<b>Bold</b>"),
		},
		{
			name:     "complex html",
			input:    `<div class="test">Content</div>`,
			expected: template.HTML(`<div class="test">Content</div>`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := safeHTML(tt.input)
			if result != tt.expected {
				t.Errorf("safeHTML(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIncr(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  int64
		expectErr bool
	}{
		{
			name:     "int",
			input:    5,
			expected: 6,
		},
		{
			name:     "int64",
			input:    int64(10),
			expected: 11,
		},
		{
			name:     "string number",
			input:    "42",
			expected: 43,
		},
		{
			name:      "invalid string",
			input:     "invalid",
			expectErr: true,
		},
		{
			name:      "unsupported type",
			input:     3.14,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := incr(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("incr(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecr(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  int64
		expectErr bool
	}{
		{
			name:     "int",
			input:    5,
			expected: 4,
		},
		{
			name:     "int64",
			input:    int64(10),
			expected: 9,
		},
		{
			name:     "string number",
			input:    "42",
			expected: 41,
		},
		{
			name:      "invalid string",
			input:     "invalid",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := decr(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("decr(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFormatInt(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  string
		expectErr bool
	}{
		{
			name:     "small number",
			input:    42,
			expected: "42",
		},
		{
			name:     "large number",
			input:    1000000,
			expected: "1,000,000",
		},
		{
			name:     "string number",
			input:    "1234",
			expected: "1,234",
		},
		{
			name:      "invalid string",
			input:     "invalid",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatInt(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("formatInt(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		dp       int
		expected string
	}{
		{
			name:     "two decimal places",
			input:    3.14159,
			dp:       2,
			expected: "3.14",
		},
		{
			name:     "no decimal places",
			input:    42.9,
			dp:       0,
			expected: "43",
		},
		{
			name:     "large number with formatting",
			input:    1234567.89,
			dp:       2,
			expected: "1,234,567.89",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatFloat(tt.input, tt.dp)
			if result != tt.expected {
				t.Errorf("formatFloat(%v, %d) = %q, want %q", tt.input, tt.dp, result, tt.expected)
			}
		})
	}
}

func TestYesno(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected string
	}{
		{
			name:     "true",
			input:    true,
			expected: "Yes",
		},
		{
			name:     "false",
			input:    false,
			expected: "No",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := yesno(tt.input)
			if result != tt.expected {
				t.Errorf("yesno(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUrlSetParam(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		key      string
		value    any
		expected string
	}{
		{
			name:     "add new param",
			url:      "https://example.com",
			key:      "foo",
			value:    "bar",
			expected: "foo=bar",
		},
		{
			name:     "update existing param",
			url:      "https://example.com?foo=old",
			key:      "foo",
			value:    "new",
			expected: "foo=new",
		},
		{
			name:     "add to existing params",
			url:      "https://example.com?existing=value",
			key:      "new",
			value:    123,
			expected: "existing=value&new=123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			result := urlSetParam(u, tt.key, tt.value)
			if !strings.Contains(result.RawQuery, tt.expected) {
				t.Errorf("urlSetParam() query = %q, want to contain %q", result.RawQuery, tt.expected)
			}
		})
	}
}

func TestUrlDelParam(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		key         string
		shouldExist bool
	}{
		{
			name:        "delete existing param",
			url:         "https://example.com?foo=bar&baz=qux",
			key:         "foo",
			shouldExist: false,
		},
		{
			name:        "delete non-existing param",
			url:         "https://example.com?foo=bar",
			key:         "missing",
			shouldExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			result := urlDelParam(u, tt.key)
			exists := strings.Contains(result.RawQuery, tt.key+"=")
			if exists != tt.shouldExist {
				t.Errorf("urlDelParam() param %q exists = %v, want %v", tt.key, exists, tt.shouldExist)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  int64
		expectErr bool
	}{
		{
			name:     "int",
			input:    42,
			expected: 42,
		},
		{
			name:     "int64",
			input:    int64(100),
			expected: 100,
		},
		{
			name:     "string number",
			input:    "123",
			expected: 123,
		},
		{
			name:      "invalid string",
			input:     "not-a-number",
			expectErr: true,
		},
		{
			name:      "unsupported type",
			input:     3.14,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := toInt64(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("toInt64(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConcatUserDate(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		time     time.Time
		expected string
	}{
		{
			name:     "basic concat",
			user:     models.User{Id: "user123"},
			time:     time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			expected: "user123_2023-12-25",
		},
		{
			name:     "different date",
			user:     models.User{Id: "abc"},
			time:     time.Date(2024, 1, 1, 15, 30, 45, 0, time.UTC),
			expected: "abc_2024-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := concatUserDate(tt.user, tt.time)
			if result != tt.expected {
				t.Errorf("concatUserDate(%v, %v) = %q, want %q", tt.user, tt.time, result, tt.expected)
			}
		})
	}
}