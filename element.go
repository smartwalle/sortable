package sortable

type Element interface {
	GetUniqueID() uint64

	GetSortIndex() uint32

	UpdateSortIndex(sortIndex uint32)
}
