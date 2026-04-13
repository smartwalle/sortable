package sortable

func Sort(dataSource DataSource, source Element, target Element) error {
	if dataSource == nil || source == nil || target == nil {
		return nil
	}
	if source.GetUniqueID() < 1 || target.GetUniqueID() < 1 || source.GetUniqueID() == target.GetUniqueID() {
		return nil
	}

	var sourceSortIndex = source.GetSortIndex()
	var targetSortIndex = target.GetSortIndex()

	var minSortIndex int
	var maxSortIndex int
	if sourceSortIndex > targetSortIndex {
		minSortIndex = targetSortIndex
		maxSortIndex = sourceSortIndex
	} else {
		minSortIndex = sourceSortIndex
		maxSortIndex = targetSortIndex
	}

	// 取出 source、target 及两者之间的所有数据
	elements, err := dataSource.GetSortableList(minSortIndex, maxSortIndex)
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

	if err = dataSource.UpateSortableList(elements); err != nil {
		return err
	}

	return nil
}
