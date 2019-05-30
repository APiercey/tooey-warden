package state

import "testing"

func TestAppendFilterString(t *testing.T) {
	var transformation Transformation
	transformation = &AppendFilterString{Value: "a"}
	var transformation2 Transformation
	transformation2 = &AppendFilterString{Value: "b"}

	app := createApplicationState()
	result := transformation.Run(app.State)
	result2 := transformation2.Run(result)

	if result2.FilterString != "ab" {
		t.Errorf("did not work")
	}
}

func TestChopFilterString(t *testing.T) {
	var transformation Transformation
	transformation = &AppendFilterString{Value: "abcdef"}
	var transformation2 Transformation
	transformation2 = &ChopFilterString{}

	app := createApplicationState()
	result := transformation.Run(app.State)
	result2 := transformation2.Run(result)

	if result2.FilterString != "abcde" {
		t.Errorf("did not work")
	}
}

func TestAddingItemsToState(t *testing.T) {
	login := &BwLogin{Username: "test", Password: "testc"}
	bwItem := &BwItem{ID: "1", Name: "Hello", Login: login}
	collection := []BwItem{}
	collection = append(collection, bwItem)
	var transformation Transformation
	transformation = &SetBwItems{Items: collection}
	app := createApplicationState()

	result := transformation.Run(app.State)

	if result2.FilterString != "abcde" {
		t.Errorf("did not work")
	}
}
