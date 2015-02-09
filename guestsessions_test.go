package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetGuestSessionRatedMovies(c *C) {
	result, err := s.tmdb.GetGuestSessionRatedMovies(guestSession, nil)
	s.baseTest(&result, err, c)
}
