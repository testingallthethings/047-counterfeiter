package drivinglicence_test

import (
	"drivinglicence"
	"drivinglicence/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DrivingLicenceSuite struct {
	applicant              *mock.FakeApplicant
	logger                 *mock.FakeLogger
	randomNumbersGenerator *mock.FakeRandomNumbersGenerator

	lg drivinglicence.NumberGenerator

	suite.Suite
}

func TestDrivingLicenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenceSuite))
}

func (s *DrivingLicenceSuite) SetupTest() {
	s.applicant = &mock.FakeApplicant{}
	s.logger = &mock.FakeLogger{}
	s.randomNumbersGenerator = &mock.FakeRandomNumbersGenerator{}

	s.lg = drivinglicence.NewNumberGenerator(s.logger, s.randomNumbersGenerator)
}

func (s *DrivingLicenceSuite) TestUnderageApplicant() {
	s.applicant.HoldsLicenceReturns(false)
	s.applicant.IsOver17Returns(false)

	_, err := s.lg.Generate(s.applicant)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, s.logger.LogStuffCallCount())
	s.Contains(s.logger.LogStuffArgsForCall(0), "Underaged")
}

func (s *DrivingLicenceSuite) TestNoSecondLicence() {
	s.applicant.HoldsLicenceReturns(true)

	_, err := s.lg.Generate(s.applicant)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, s.logger.LogStuffCallCount())
	s.Contains(s.logger.LogStuffArgsForCall(0), "Duplicate")
}

func (s *DrivingLicenceSuite) TestLicenceGeneration() {
	s.applicant.HoldsLicenceReturns(false)
	s.applicant.IsOver17Returns(true)
	s.applicant.GetInitialsReturns("MDB")
	s.applicant.GetDOBReturns("23082011")

	s.randomNumbersGenerator.GetRandomNumbersReturns("00000")

	ln, err := s.lg.Generate(s.applicant)

	s.NoError(err)
	s.Equal("MDB2308201100000", ln)
}

func (s *DrivingLicenceSuite) TestLicenceGenerationShorterInitials() {
	s.applicant.HoldsLicenceReturns(false)
	s.applicant.IsOver17Returns(true)
	s.applicant.GetInitialsReturns("MB")
	s.applicant.GetDOBReturns("23082011")

	s.randomNumbersGenerator.GetRandomNumbersReturns("000000")

	ln, err := s.lg.Generate(s.applicant)

	s.NoError(err)
	s.Equal("MB23082011000000", ln)
}
