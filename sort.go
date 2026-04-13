package sortable

import "context"

func Sort(ctx context.Context, dataSource DataSource, source, target Element) error {
	if dataSource == nil || source == nil || target == nil {
		return nil
	}
	if source.GetUniqueID() < 1 || target.GetUniqueID() < 1 || source.GetUniqueID() == target.GetUniqueID() {
		return nil
	}

	var sourceSortIndex = source.GetSortIndex()
	var targetSortIndex = target.GetSortIndex()

	var minSortIndex uint32
	var maxSortIndex uint32
	if sourceSortIndex > targetSortIndex {
		minSortIndex = targetSortIndex
		maxSortIndex = sourceSortIndex
	} else {
		minSortIndex = sourceSortIndex
		maxSortIndex = targetSortIndex
	}

	// 取出 source、target 及两者之间的所有数据
	elements, err := dataSource.GetSortableElements(ctx, minSortIndex, maxSortIndex)
	if err != nil {
		return err
	}

	for _, ele := range elements {
		if ele.GetUniqueID() == source.GetUniqueID() {
			ele.UpdateSortIndex(targetSortIndex)
			continue
		}

		if sourceSortIndex > targetSortIndex {
			// 往前移动，其它元素的排序 +1
			ele.UpdateSortIndex(ele.GetSortIndex() + 1)
		} else {
			// 往后移动，其它元素的排序 -1
			ele.UpdateSortIndex(ele.GetSortIndex() - 1)
		}
	}

	if err = dataSource.UpateSortableElements(ctx, elements); err != nil {
		return err
	}

	return nil
}
