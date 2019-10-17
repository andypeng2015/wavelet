// Copyright (c) 2019 Perlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package wavelet

import (
	"sync"

	"github.com/perlin-network/wavelet/conf"
)

type Snowball struct {
	sync.RWMutex
	beta  int
	alpha int

	count  int
	counts map[BlockID]uint16

	preferred *snowballVote
	last      *snowballVote

	decided bool
}

func NewSnowball() *Snowball {
	return &Snowball{
		counts: make(map[BlockID]uint16),
	}
}

func (s *Snowball) Reset() {
	s.Lock()

	s.preferred = nil
	s.last = nil

	s.counts = make(map[BlockID]uint16)
	s.count = 0

	s.decided = false
	s.Unlock()
}

func (s *Snowball) Tick(tallies map[VoteID]float64, votes map[VoteID]*snowballVote) {
	s.Lock()
	defer s.Unlock()

	if s.decided {
		return
	}

	var majority *snowballVote
	var majorityTally float64 = 0

	for id, tally := range tallies {
		if tally > majorityTally {
			majority, majorityTally = votes[id], tally
		}
	}

	denom := float64(len(votes))

	if denom < 2 {
		denom = 2
	}

	if majority == nil || majorityTally < conf.GetSnowballAlpha()*2/denom {
		s.count = 0
	} else {
		s.counts[majority.id] += 1

		if s.counts[majority.id] > s.counts[s.preferred.id] {
			s.preferred = majority
		}

		if s.last == nil || majority.id != s.last.id {
			s.last, s.count = majority, 1
		} else {
			s.count += 1

			if s.count > conf.GetSnowballBeta() {
				s.decided = true
			}
		}
	}
}

func (s *Snowball) Prefer(b *snowballVote) {
	s.Lock()
	s.preferred = b
	s.Unlock()
}

func (s *Snowball) Preferred() *snowballVote {
	s.RLock()
	preferred := s.preferred
	s.RUnlock()

	return preferred
}

func (s *Snowball) Decided() bool {
	s.RLock()
	decided := s.decided
	s.RUnlock()

	return decided
}

func (s *Snowball) Progress() int {
	s.RLock()
	progress := s.count
	s.RUnlock()

	return progress
}
