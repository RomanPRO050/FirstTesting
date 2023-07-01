package counter

import (
	"FirstTesting/internal/app/services/post"
	mock_post "FirstTesting/test/gomock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetList(t *testing.T) {
	req := require.New(t)
	//any := gomock.Any()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	newClient := mock_post.NewMockClient(mockCtrl)
	newClient.EXPECT().GetList().Return([]post.Post{}, nil).Times(1)
	_, err := PostCount(newClient)
	req.NoError(err)
}
