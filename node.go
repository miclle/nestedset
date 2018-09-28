package nestedset

// NodeInterface is the interface implemented by types that can be used by nodes in nested set
type NodeInterface interface {
	GetId() int          // Returns id of node
	GetParentID() int    // Returns parent id of node
	GetDepth() int       // Returns depth of node
	GetLeft() int        // Returns left of node
	GetRight() int       // Returns right of node
	GetSortPath() string // Returns sort path of node
	GetNodeNumber() int  // Returns node number of node
	GetNodeCount() int   // Returns node count of node

	SetId(int)          // Sets node id
	SetParentID(int)    // Sets parent id
	SetDepth(int)       // Sets node depth
	SetLeft(int)        // Sets node left
	SetRight(int)       // Sets node right
	SetSortPath(string) // set node sort path
	SetNodeNumber(int)  // set node number
	SetNodeCount(int)   // set node count
}

// Node type
type Node struct {
	ID         int
	ParentID   int
	Depth      int
	Left       int
	Right      int
	SortPath   string
	NodeNumber int
	NodeCount  int
}
