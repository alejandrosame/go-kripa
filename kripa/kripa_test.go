package kripa_test

import (
	"testing"

	"github.com/alejandrosame/go-kripa/kripa"
)

func TestEscapeEucKr(t *testing.T) {
	got := kripa.EscapeEucKr("확인하기")
	expected := "%C8%AE%C0%CE%C7%CF%B1%E2"
	if got != expected {
		t.Errorf("kripa.EscapeEucKr(\"확인하기\") = \"%s\"; want \"%s\"",
			got, expected)
	}
}

func TestGetTranscriptIPA(t *testing.T) {
	got, err := kripa.GetTranscriptIPA("그는 괜찮은 척하려고 애쓰는 것 같았다")
	if err != nil {
		t.Errorf("Error Found: %v", err)
	}

	expected := "kɯnɯn kwɛntsʰanɯn tsʰʌkʰaryʌgo ɛs̕ɯnɯn gʌt̚ katʰat̚t̕a"
	if got != expected {
		t.Errorf("kripa.GetTranscriptIPA(\"그는 괜찮은 척하려고 애쓰는 것 같았다\") = \"%s\"; want \"%s\"",
			got, expected)
	}
}
