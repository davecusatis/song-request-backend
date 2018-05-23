package api

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/davecusatis/song-request-backend/song-request/aggregator"
	"github.com/davecusatis/song-request-backend/song-request/datasource"
)

// API struct
type API struct {
	Aggregator *aggregator.Aggregator
	S3Uploader *s3manager.Uploader
	Datasource *datasource.Datasource
}

// NewAPI creates a new instance of an API
func NewAPI(s3Uploader *s3manager.Uploader) (*API, error) {
	a := aggregator.NewAggregator()
	a.Start()
	return &API{
		S3Uploader: s3Uploader,
		Aggregator: a,
		Datasource: datasource.NewDatasource(),
	}, nil
}
