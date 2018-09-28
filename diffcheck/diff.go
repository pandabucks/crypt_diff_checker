package diffcheck

// Operation defines the operation of a diff item.
type Operation int8

const (
	DiffDelete Operation = -1
	DiffInsert Operation = 1
	DiffEqual  Operation = 0
)

// Diff represents one diff operation
type Diff struct {
	Type Operation
	Hash string
}
