package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	name          string
	performAction func(*Cache, string)
	verifyResult  func(*testing.T, *Cache, string)
}

func TestCache_Set(t *testing.T) {
	testData := map[string][]interface{}{
		"case1": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
	}

	expectedData := map[string][]interface{}{
		"case1": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
	}

	tests := []Test{
		{
			name: "case1",
			performAction: func(p *Cache, testName string) {
				for i, val := range testData[testName] {
					p.Set(fmt.Sprintf("%d", i), val)
				}
			},
			verifyResult: func(t *testing.T, p *Cache, testName string) {
				for i, val := range expectedData[testName] {
					eq, ok := p.Items[fmt.Sprintf("%d", i)]
					if !ok {
						t.Errorf("item key: %v and value %v not found", i, val)
					}
					assert.Equal(t, val, eq.value)
				}
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.performAction(p, test.name)
			test.verifyResult(t, p, test.name)
		})
	}
}

func TestCache_Get(t *testing.T) {
	testData := map[string][]interface{}{
		"case1": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
	}

	expectedData := map[string][]interface{}{
		"case1": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
	}

	tests := []Test{
		{
			name: "case1",
			performAction: func(p *Cache, testName string) {
				for i, val := range testData[testName] {
					p.Set(fmt.Sprintf("%d", i), val)
				}
			},
			verifyResult: func(t *testing.T, p *Cache, testName string) {
				for i, val := range expectedData[testName] {
					eq, ok := p.Get(fmt.Sprintf("%d", i))
					if !ok {
						t.Errorf("item key: %v and value %v not found", i, val)
					}
					assert.Equal(t, val, eq)
				}
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.performAction(p, test.name)
			test.verifyResult(t, p, test.name)
		})
	}
}

func TestCache_Remove(t *testing.T) {
	testData := map[string][]interface{}{
		"case1": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
		"case2": {
			11, 844.2, "test", []int{1, 2, 3, 4, 5},
		},
	}

	deleteData := map[string][]int{
		"case1": {1},
		"case2": {0, 1, 2, 3},
	}

	expectedData := map[string][]interface{}{
		"case1": {
			0, 2, 3,
		},
		"case2": {},
	}

	tests := []Test{
		{
			name: "case1",
			performAction: func(p *Cache, testName string) {
				for i, val := range testData[testName] {
					p.Set(fmt.Sprintf("%d", i), val)
				}
				for _, val := range deleteData[testName] {
					p.Remove(fmt.Sprintf("%d", val))
				}
			},
			verifyResult: func(t *testing.T, p *Cache, testName string) {
				for _, val := range deleteData[testName] {
					item, ok := p.Items[fmt.Sprintf("%d", val)]
					if ok {
						t.Errorf("item %v index %v after removal is present in items", item, val)
					}
				}
				for _, val := range expectedData[testName] {
					_, ok := p.Items[fmt.Sprintf("%d", val)]
					if !ok {
						t.Errorf("item key: %v not found", val)
					}
				}
				assert.Equal(t, len(expectedData[testName]), len(p.Items))
			},
		},
		{
			name: "case2",
			performAction: func(p *Cache, testName string) {
				for i, val := range testData[testName] {
					p.Set(fmt.Sprintf("%d", i), val)
				}

				for i := range testData[testName] {
					p.Remove(fmt.Sprintf("%d", i))
				}
			},
			verifyResult: func(t *testing.T, p *Cache, testName string) {
				assert.Equal(t, len(expectedData[testName]), len(p.Items))
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.performAction(p, test.name)
			test.verifyResult(t, p, test.name)
		})
	}
}
