package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferentCharacter(t *testing.T) {

	str, str2, shouldBe := "ab", "aab", "a"
	result := DifferentCharacter(str, str2)
	assert.Equal(t, shouldBe, result, "#1 -> Result should be: "+shouldBe)

	str, str2, shouldBe = "abcd", "abcde", "e"
	result = DifferentCharacter(str, str2)
	assert.Equal(t, shouldBe, result, "#2 -> Result should be: "+shouldBe)

	str, str2, shouldBe = "aaabbcdd", "abdbacade", "e"
	result = DifferentCharacter(str, str2)
	assert.Equal(t, shouldBe, result, "#3 -> Result should be: "+shouldBe)

	str, str2, shouldBe = "aab", "ab", "a"
	result = DifferentCharacter(str, str2)
	assert.Equal(t, shouldBe, result, "#4 -> Result should be: "+shouldBe)
}
