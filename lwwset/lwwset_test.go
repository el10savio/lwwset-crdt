package lwwset

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lwwset LWWSet
)

func init() {
	lwwset = Initialize()
}

// TestList checks the basic functionality of LWWSet List()
// List() should return all unique values added to the LWWSet
func TestList(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")

	expectedValue := []string{"xx"}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestList_UpdatedValue checks the functionality of LWWSet List() when
// multiple values are added to LWWSet it should return
// all the unique values added to the LWWSet
func TestList_UpdatedValue(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")
	lwwset, _ = lwwset.Addition("yy")
	lwwset, _ = lwwset.Addition("zz")

	expectedValue := []string{"xx", "yy", "zz"}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestList_ReAddValue checks the functionality of LWWSet List() when
// a value is added to LWWSet after it got removed it should return
// all the unique values added to the LWWSet
func TestList_ReAddValue(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")
	lwwset, _ = lwwset.Removal("xx")
	lwwset, _ = lwwset.Addition("xx")

	expectedValue := []string{"xx"}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestList_RemoveValue checks the functionality of LWWSet List() when
// multiple values are added & removed to LWWSet it should return
// all the unique values finally present to the LWWSet
func TestList_RemoveValue(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")
	lwwset, _ = lwwset.Removal("xx")
	lwwset, _ = lwwset.Removal("zz")

	expectedValue := []string{}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestList_RemoveEmpty checks the functionality of LWWSet List() when
// multiple values are removed to LWWSet it should return
// all the unique values finally present to the LWWSet
func TestList_RemoveEmpty(t *testing.T) {
	lwwset, _ = lwwset.Removal("zz")

	expectedValue := []string{}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestList_NoValue checks the functionality of LWWSet List() when
// no values are added to LWWSet, it should return
// an empty string slice when the LWWSet is empty
func TestList_NoValue(t *testing.T) {
	expectedValue := []string{}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestClear checks the basic functionality of LWWSet Clear()
// utility function it clears all the values in a LWWSet
func TestClear(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx1")
	lwwset, _ = lwwset.Addition("xx2")
	lwwset = Clear()

	expectedValue := []string{}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestClear_EmptyStore checks the functionality of LWWSet Clear() utility function
// when no values are in it, it clears all the values in a LWWSet set
func TestClear_EmptyStore(t *testing.T) {
	lwwset = Clear()

	expectedValue := []string{}
	_, actualValue := lwwset.List()

	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestLookup checks the basic functionality of LWWSet Lookup() function
// it returns a boolean if a value passed is present in the LWWSet set or not
func TestLookup(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")

	expectedValue := true
	actualValue, actualError := lwwset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestLookup_NotPresent checks the functionality of LWWSet Lookup() function
// it returns false if a value passed is not present in the LWWSet
func TestLookup_NotPresent(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")

	expectedValue := false
	actualValue, actualError := lwwset.Lookup("yy")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestLookup_NotPresent checks the functionality of LWWSet Lookup() function
// it returns false if a value passed is not present in the LWWSet
func TestLookup_Removed(t *testing.T) {
	lwwset, _ = lwwset.Addition("xx")
	lwwset, _ = lwwset.Removal("xx")

	expectedValue := false
	actualValue, actualError := lwwset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestLookup_EmptySet checks the functionality of LWWSet Lookup() function
// it returns false if the LWWSet is empty irrespective of the value passed
func TestLookup_EmptySet(t *testing.T) {
	expectedValue := false
	actualValue, actualError := lwwset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}

// TestLookup_EmptyLookup checks the functionality of LWWSet Lookup() function
// it returns an error if the value passed is nil irrespective of the LWWSet
func TestLookup_EmptyLookup(t *testing.T) {
	expectedValue := false
	expectedError := errors.New("empty value provided")

	actualValue, actualError := lwwset.Lookup("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	lwwset = Clear()
}
