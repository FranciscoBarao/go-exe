package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCompress(t *testing.T) {

	str, shouldBe := "AABBCC", "AABBCC"
	result := StringCompress(str)
	assert.Equal(t, shouldBe, result, "#1 -> Result should be: "+shouldBe)

	str, shouldBe = "AAABCCDDDD", "A3BC2D4"
	result = StringCompress(str)
	assert.Equal(t, shouldBe, result, "#2 -> Result should be: "+shouldBe)

	str, shouldBe = "BAAACCDDDD", "BA3C2D4"
	result = StringCompress(str)
	assert.Equal(t, shouldBe, result, "#2 -> Result should be: "+shouldBe)

	str, shouldBe = "AAABAACCDDDD", "A3BA2C2D4"
	result = StringCompress(str)
	assert.Equal(t, shouldBe, result, "#2 -> Result should be: "+shouldBe)
}
