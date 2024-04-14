package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/watariRyo/baby-map-server/domain/repository"
)

type UsersRepositoryTestSuite struct {
	suite.Suite
	ur   UsersRepository
	conn repository.DBConnection
	ctx  context.Context
}

func (s *UsersRepositoryTestSuite) SetupSuite() {
	s.conn = testConnection(s.T())
	s.ctx = context.Background()
	s.ur = UsersRepository{}
}

func (s *UsersRepositoryTestSuite) TearDownTest() {
	truncateTables(s.T())
}

func TestSuiteUsersRepository(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}

// func (s *UsersRepositoryTestSuite) Test_UserRepositoryTestSuite_Insert() {
// 	t := s.T()

// 	displayName := "hogeDisp"
// 	got, err := s.ur.Insert(s.ctx, s.conn, &model.User{
// 		UUID:        "hoge",
// 		DisplayName: &displayName,
// 	})
// 	assert.Nil(t, err)

// 	want := &model.User{
// 		UUID:        "hoge",
// 		DisplayName: &displayName,
// 	}

// 	if d := cmp.Diff(got, want, cmpopts.IgnoreFields(model.User{}, "ID")); len(d) != 0 {
// 		t.Errorf("differs: (-got +want)\n%s", d)
// 	}
// }
