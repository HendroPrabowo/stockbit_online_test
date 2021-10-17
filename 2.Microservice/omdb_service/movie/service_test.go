package movie

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"omdb_service/infrastructure/omdb"
)

type ServiceTestSuite struct {
	suite.Suite
	omdb *omdb.MockService
	sut  *ServiceImpl
}

func (svc *ServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(svc.T())
	svc.omdb = omdb.NewMockService(ctrl)
	svc.sut = &ServiceImpl{svc.omdb}
}

func (svc ServiceTestSuite) TestNewService() {
	actual := NewService(svc.omdb)
	svc.Assert().Equal(svc.sut, actual)
}

func (svc ServiceTestSuite) TestProceedGetMovie() {
	filter := omdb.Filter{"id11", "batman"}
	svc.omdb.EXPECT().GetMovie(filter).Return(omdb.MovieInformation{}, nil)
	actual, err := svc.sut.ProceedGetMovie(filter.ID, filter.Title)
	svc.Assert().NoError(err)
	svc.Assert().Equal(omdb.MovieInformation{}, actual)
}

func (svc ServiceTestSuite) TestProceedGetMovies() {
	svc.omdb.EXPECT().GetMovies(gomock.Any(), gomock.Any()).Return(omdb.Result{}, nil)
	actual, err := svc.sut.ProceedGetMovies("batman", "1")
	svc.Assert().NoError(err)
	svc.Assert().Equal(omdb.Result{}, actual)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
