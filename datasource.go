package sortable

import "context"

type DataSource interface {
	GetSortableElements(ctx context.Context, minSortIndex, maxSortIndex uint32) ([]Element, error)

	UpateSortableElements(ctx context.Context, elements []Element) error
}
