package lwwset

import (
	"errors"
	"time"
)

// package lwwset implements the LWWSet (Last Writer Wins Set) CRDT data type along with the functionality
// to append, list & lookup values in a LWWSet. It also provides the functionality to merge multiple
// LWWSets together and a utility function to clear a LWWSet used in tests

// LWWSet is the LWWSet CRDT data type
// It is implemented by combining two LWWNodes,
// One to store the values added & another
// to store the values removed
type LWWSet struct {
	// Add is a LWWNodeSlice to store the values added
	Add LWWNodeSlice `json:"add"`
	// Remove is a LWWNodeSlice to store the values removed
	Remove LWWNodeSlice `json:"remove"`
}

// LWWNode TODO: ...
type LWWNode struct {
	Value     string
	Timestamp time.Time
}

// LWWNodeSlice ...
type LWWNodeSlice []LWWNode

// Initialize returns a new empty LWWSet
func Initialize() LWWSet {
	return LWWSet{
		Add:    LWWNodeSlice{},
		Remove: LWWNodeSlice{},
	}
}

// Addition adds a new unique value to the Add LWWSet
func (lwwset LWWSet) Addition(value string) (LWWSet, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return lwwset, errors.New("empty value provided")
	}

	// Set = Set U value
	lwwset.Add = append(lwwset.Add, LWWNode{Value: value, Timestamp: time.Now()})

	// Return the new LWWSet followed by nil error
	return lwwset, nil
}

// Removal adds a new unique value to the Remove LWWSet
func (lwwset LWWSet) Removal(value string) (LWWSet, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return lwwset, errors.New("empty value provided")
	}

	// Set = Set U value
	lwwset.Remove = append(lwwset.Remove, LWWNode{Value: value, Timestamp: time.Now()})

	// Return the new LWWSet followed by nil error
	return lwwset, nil
}

// GetValues extracts all the values present in the LWWNode slice
func (lwwnodeslice LWWNodeSlice) GetValues() []string {
	if len(lwwnodeslice) == 0 {
		return []string{}
	}

	values := make([]string, 0)

	for _, lwwnode := range lwwnodeslice {
		values = append(values, lwwnode.Value)
	}

	return values
}

// List returns all the elements present in the LWWSet
func (lwwset LWWSet) List() []string {
	if len(lwwset.Remove) == 0 {
		return lwwset.Add.GetValues()
	}

	// An element is a member of the LWW-Element-Set if it is in the add set, and either not in the remove
	// set, or in the remove set but with an earlier timestamp than the latest timestamp in the add set.

	result := make([]string, 0)

	for _, lwwNode := range lwwset.Add {
		if !isPresent(lwwNode.Value, lwwset.Remove) || latestValue(lwwNode, lwwset.Remove).Timestamp.UnixNano() < lwwNode.Timestamp.UnixNano() {
			result = append(result, lwwNode.Value)
		}
	}

	return result
}

// isPresent ...
func isPresent(value string, list LWWNodeSlice) bool {
	for _, element := range list {
		if element.Value == value {
			return true
		}
	}
	return false
}

func latestValue(node LWWNode, list LWWNodeSlice) LWWNode {
	maxNode := node
	for _, element := range list {
		if element.Value == maxNode.Value && element.Timestamp.UnixNano() > maxNode.Timestamp.UnixNano() {
			maxNode = node
		}
	}
	return maxNode
}

// Lookup returns either boolean true/false indicating
// if a given value is present in the LWWSet or not
func (lwwset LWWSet) Lookup(value string) (bool, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return false, errors.New("empty value provided")
	}

	list := lwwset.List()

	// Iterative over the LWWSet and check if the
	// value is the one we're searching
	// return true if the value exists
	for _, element := range list {
		if element == value {
			return true, nil
		}
	}

	// If the value isn't found after iterating
	// over the entire LWWSet we return false
	return false, nil
}

// Merge conbines multiple LWWSets together using Union
// and returns a single merged LWWSet
func Merge(LWWSets ...LWWSet) LWWSet {
	var LWWSetMerged LWWSet

	// GSetMerged = GSetMerged U GSetToMergeWith
	for _, lwwset := range LWWSets {
		for _, lwwnode := range lwwset.Add {
			if lwwnode.Value == "" {
				continue
			}
			LWWSetMerged, _ = LWWSetMerged.Addition(lwwnode.Value)
		}
		for _, lwwnode := range lwwset.Remove {
			if lwwnode.Value == "" {
				continue
			}
			LWWSetMerged, _ = LWWSetMerged.Removal(lwwnode.Value)
		}
	}

	// Return the merged LWWSet followed by nil error
	return LWWSetMerged
}

// Clear is utility function used only for tests
// to empty the contents of a given LWWSet
func (lwwset LWWSet) Clear() LWWSet {
	lwwset.Add = LWWNodeSlice{}
	lwwset.Remove = LWWNodeSlice{}
	return lwwset
}
