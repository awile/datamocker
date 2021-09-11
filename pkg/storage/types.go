package storage

import (
	"github.com/awile/datamkr/pkg/dataset"
)

type WriterOptions struct {
	Id                string
	FieldKeys         []string
	DatasetDefinition dataset.DatasetDefinition
}

func CreateWriterOptions() WriterOptions {
	return WriterOptions{}
}