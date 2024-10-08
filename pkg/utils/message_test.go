package utils

import "testing"

// func BenchmarkReplaceAtTags(b *testing.B) {
// 	text := `Hello.<at oc_id="12bd9168-958f-459a-8f37-41fcfd81e8f5"></at>.please review this document.
// <at oc_id="12bd9168-958f-459a-8f37-41fcfd81e8f5"></at>, your feedback is needed.`
// 	idToName := map[string]string{
// 		"12bd9168-958f-459a-8f37-41fcfd81e8f5": "DJJ",
// 	}
// 	for i := 0; i < b.N; i++ {
// 		ReplaceAtTags(text, idToName)
// 	}
// }

func BenchmarkReplaceMentionMessage(b *testing.B) {
	text := `Hello.<at oc_id="12bd9168-958f-459a-8f37-41fcfd81e8f5"></at>.please review this document.
<at oc_id="12bd9168-958f-459a-8f37-41fcfd81e8f5"></at>, your feedback is needed.`
	idToName := map[string]string{
		"12bd9168-958f-459a-8f37-41fcfd81e8f5": "DJJ",
	}
	for i := 0; i < b.N; i++ {
		ReplaceMentionMessage(text, idToName)
	}
}

func TestSimplifyContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Short content without mention messages",
			input:    "Go is fun.",
			expected: "Go is fun.",
		},
		{
			name:     "Short content with single mention message",
			input:    "Hello, <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
			expected: "Hello, <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
		},
		{
			name:     "Long content without mention messages",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx...",
		},
		{
			name:     "Long content with single mention message",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
			expected: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx...",
		},
		{
			name:     "Long content with multiple mention messages",
			input:    "<at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at> abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrst",
			expected: "<at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at> abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqr...",
		},
		{
			name:     "Long content with multiple mention messages in mid",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmn<at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
			expected: "abcdefghijklmnopqrstuvwxyzabcdefghijklmn<at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
		},
		{
			name:     "Long content with multiple mention messages with special unicode",
			input:    "你好世界,にほんご한국어EspañolLingua LatīnaLatīnaLatīna    <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>LatīnaLatīnaLatīna你好世界",
			expected: "你好世界,にほんご한국어EspañolLingua LatīnaLatīnaLatīna    <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>",
		},
		{
			name:     "Long content with multiple mention messages in mid with special unicode",
			input:    "你好世界,にほんご한국어EspañolLingua LatīnaLatīnaLatīna <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>LatīnaLatīnaLatīna你好世界",
			expected: "你好世界,にほんご한국어EspañolLingua LatīnaLatīnaLatīna <at oc_id=\"{123e4567-e89b-12d3-a456-426655440000}\"></at>...",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				result := SimplifyContent(tt.input)
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			},
		)
	}
}
