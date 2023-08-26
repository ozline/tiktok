package service

import "testing"

func TestSensitiveWordsMatch(t *testing.T) {
	_, err := interactionService.MatchSensitiveWords(commentText)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
