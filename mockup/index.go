package mockup

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct{}

func (m *MockCollection) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions,
) (int64, error) {
	return 0, nil
}
