package config

import (
	"testing"
	"time"
)

var testConfig = Config{
	FirstSeasonDate: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
	SeasonOffset:    1,
	CompetitionDays: 14,
	RestDays:        14,
}

type testCase struct {
	now                        time.Time
	expectedSeason             int
	expectedCompetitionActive  bool
	expectedCompetitionEndDate time.Time
	expectedSeasonEndDate      time.Time
}

var testCases = []testCase{
	{
		now:                        time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
		expectedSeason:             1,
		expectedCompetitionActive:  true,
		expectedCompetitionEndDate: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.Local),
		expectedSeasonEndDate:      time.Date(2000, time.January, 29, 0, 0, 0, 0, time.Local),
	},
	{
		now:                        time.Date(2000, time.January, 16, 0, 0, 0, 0, time.Local),
		expectedSeason:             1,
		expectedCompetitionActive:  false,
		expectedCompetitionEndDate: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.Local),
		expectedSeasonEndDate:      time.Date(2000, time.January, 29, 0, 0, 0, 0, time.Local),
	},
	{
		now:                        time.Date(2000, time.February, 1, 0, 0, 0, 0, time.Local),
		expectedSeason:             2,
		expectedCompetitionActive:  true,
		expectedCompetitionEndDate: time.Date(2000, time.February, 12, 0, 0, 0, 0, time.Local),
		expectedSeasonEndDate:      time.Date(2000, time.February, 26, 0, 0, 0, 0, time.Local),
	},
}

func TestSeasonConfig(t *testing.T) {
	for i, tc := range testCases {
		actualSeason := testConfig.GetSeasonAt(tc.now)
		if actualSeason != tc.expectedSeason {
			t.Errorf("testcase #%v: incorrect season. expected %v, got %v", i, tc.expectedSeason, actualSeason)
		}

		actualCompetitionActive := testConfig.IsCompetitionActiveAt(tc.now)
		if actualCompetitionActive != tc.expectedCompetitionActive {
			t.Errorf("testcase #%v: incorrect competition active: expected %v, got %v", i, tc.expectedCompetitionActive, actualCompetitionActive)
		}

		actualCompetitionEndDate := testConfig.GetCompetitionEndDateAt(tc.now)
		if actualCompetitionEndDate.Unix() != tc.expectedCompetitionEndDate.Unix() {
			t.Errorf("testcase #%v: incorrect competition end date: expected %v, got %v", i, tc.expectedCompetitionEndDate, actualCompetitionEndDate)
		}

		actualSeasonEndDate := testConfig.GetNextSeasonStartDateAt(tc.now)
		if actualSeasonEndDate.Unix() != tc.expectedSeasonEndDate.Unix() {
			t.Errorf("testcase #%v: incorrect season end date: expected %v, got %v", i, tc.expectedSeasonEndDate, actualSeasonEndDate)
		}
	}

}
