package config

import (
	"time"
)

type Config struct {
	MinRacesCount   int
	InitialRating   float64
	FirstSeasonDate time.Time
	SeasonOffset    int
	CompetitionDays int
	RestDays        int
	Elo             ConfigElo
}

type ConfigElo struct {
	D float64
	K float64
}

func (c *Config) GetSeason() int {
	return c.GetSeasonAt(time.Now())
}

func (c *Config) IsCompetitionActive() bool {
	return c.IsCompetitionActiveAt(time.Now())
}

func (c *Config) GetCompetitionEndDate() time.Time {
	return c.GetCompetitionEndDateAt(time.Now())
}

func (c *Config) GetNextSeasonStartDate() time.Time {
	return c.GetNextSeasonStartDateAt(time.Now())
}

func (c *Config) GetSeasonAt(date time.Time) int {
	return c.SeasonOffset + c.GetSecondsSinceFirstAt(date)/c.GetSeasonDurationSeconds()
}

func (c *Config) IsCompetitionActiveAt(date time.Time) bool {
	secondsSinceSeasonStart := c.GetSecondsSinceFirstAt(date) % c.GetSeasonDurationSeconds()
	return secondsSinceSeasonStart <= c.CompetitionDays*24*3600
}

func (c *Config) GetCompetitionEndDateAt(date time.Time) time.Time {
	secondsSinceSeasonStart := c.GetSecondsSinceFirstAt(date) % c.GetSeasonDurationSeconds()
	secondsUntilCompetitionEnd := c.CompetitionDays*24*3600 - secondsSinceSeasonStart
	return date.Add(time.Second * time.Duration(secondsUntilCompetitionEnd))
}

func (c *Config) GetNextSeasonStartDateAt(date time.Time) time.Time {
	nextSeasonSecondsSinceFirst := c.GetSeasonDurationSeconds() * (c.GetSeasonAt(date) + 1 - c.SeasonOffset)
	return c.FirstSeasonDate.Add(time.Second * time.Duration(nextSeasonSecondsSinceFirst))
}

func (c *Config) GetSeasonDurationSeconds() int {
	return (c.CompetitionDays + c.RestDays) * 24 * 3600
}

func (c *Config) GetSecondsSinceFirstAt(date time.Time) int {
	return int(date.Unix() - c.FirstSeasonDate.Unix())
}
