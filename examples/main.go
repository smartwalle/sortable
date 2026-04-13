package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartwalle/sortable"
)

func main() {
	var ds = &DataSource{}
	for i := 1; i <= 10; i++ {
		ds.users = append(ds.users, &User{
			ID:        uint64(i),
			SortIndex: uint32(i),
		})
	}

	var ctx = context.Background()

	sortable.Sort(ctx, ds, ds.users[2], ds.users[1])
	sortable.Sort(ctx, ds, ds.users[2], ds.users[3])
	sortable.Sort(ctx, ds, ds.users[2], ds.users[4])
}

type User struct {
	ID        uint64
	SortIndex uint32
}

func (u *User) String() string {
	return fmt.Sprintf("[%d-%d]", u.ID, u.SortIndex)
}

func (u *User) GetUniqueID() uint64 {
	return u.ID
}

func (u *User) GetSortIndex() uint32 {
	return u.SortIndex
}

func (u *User) UpdateSortIndex(sortIndex uint32) {
	u.SortIndex = sortIndex
}

type DataSource struct {
	users []*User
}

func (ds *DataSource) GetSortableElements(ctx context.Context, minSortIndex, maxSortIndex uint32) ([]sortable.Element, error) {
	var elements = make([]sortable.Element, 0, len(ds.users))
	for _, u := range ds.users {
		if u.SortIndex <= maxSortIndex && u.SortIndex >= minSortIndex {
			elements = append(elements, u)
		}
	}
	return elements, nil
}

func (ds *DataSource) UpateSortableElements(ctx context.Context, elements []sortable.Element) error {
	sort.SliceStable(ds.users, func(i, j int) bool {
		if ds.users[i].SortIndex < ds.users[j].SortIndex {
			return true
		}
		return false
	})
	fmt.Println("更新后的顺序：", ds.users)
	return nil
}
