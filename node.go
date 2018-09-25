package nestedset

// NodeInterface is the interface implemented by types that can be used by nodes in nested set
type NodeInterface interface {
	GetId() int        // Returns id of node
	GetParentID() int  // Returns parent id of node
	GetDepth() int     // Returns depth of node
	GetLeft() float64  // Returns left of node
	GetRight() float64 // Returns right of node

	SetId(int)        // Sets node id
	SetParentID(int)  // Sets parent id
	SetDepth(int)     // Sets node depth
	SetLeft(float64)  // Sets node left
	SetRight(float64) // Sets node right
}
