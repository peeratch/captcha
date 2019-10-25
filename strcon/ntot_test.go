package strcon

import "testing"

func TestNtot_ShouldConvertNumber_To_EnglishNumber_CorrectResult(t *testing.T) {
	actual := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "1024", "514"}
	expected := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"one thousand twenty-four",
		"five hundred fourteen",
	}

	for index := 0; index < len(actual); index++ {
		t.Run("expected should equal actual", func(t *testing.T) {
			if expected[index] != Ntot(actual[index]) {
				t.Errorf("expected: %s should  equal actual but got %s", expected[index], actual[index])
			}
		})
	}
}
