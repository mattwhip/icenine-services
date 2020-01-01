package sampling

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

// A Population is a collection of weighted items that can be randomly sampled
type Population struct {
	Items       []*Item
	TotalWeight int64
	searchSpace []*searchRange
}

// An Item is a weighted item
type Item struct {
	Value  interface{}
	Weight int64
	Index  int
}

type searchRange struct {
	item *Item
	min  int64
	max  int64
}

// NewPopulation creates a new population of weighted items
func NewPopulation(items []interface{}, weights []int64) (*Population, error) {
	if len(items) != len(weights) {
		panic(fmt.Errorf("length of items %v not equal to length of weights %v",
			len(items), len(weights)))
	}
	size := len(items)
	popItems := make([]*Item, 0, size)
	searchSpace := make([]*searchRange, 0, size)
	totalWeight := int64(0)
	rangeStart := int64(0)
	popIndex := 0
	for i, v := range items {
		// Skip items with negative or zero weight
		if weights[i] <= 0 {
			popIndex++
			continue
		}
		item := &Item{
			Value:  v,
			Weight: weights[i],
			Index:  popIndex,
		}
		searchRange := &searchRange{
			item: item,
			min:  rangeStart,
			max:  rangeStart + weights[i],
		}
		popItems = append(popItems, item)
		searchSpace = append(searchSpace, searchRange)
		totalWeight += weights[i]
		rangeStart += weights[i]
		popIndex++
	}
	if len(popItems) == 0 {
		return nil, errors.New("no items in population with positive weight")
	}
	return &Population{
		Items:       popItems,
		TotalWeight: totalWeight,
		searchSpace: searchSpace,
	}, nil
}

// Sample a single random item from the population
func (p *Population) Sample() (*Item, error) {
	val := rand.Int63n(p.TotalWeight)
	item, err := p.SampleAtValue(val)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sample item from population")
	}
	return item, nil
}

// SampleAtValue samples a single item from the population at the specified value (deterministic)
func (p *Population) SampleAtValue(value int64) (*Item, error) {
	if value < 0 || value >= p.TotalWeight {
		return nil, fmt.Errorf("value %v outside valid range of population with total weight %v", value, p.TotalWeight)
	}
	result, _ := find(p.searchSpace, value, 0, len(p.Items))
	return result.item, nil
}

// SampleWithReplacement samples count random items from the population with replacment
func (p *Population) SampleWithReplacement(count int) ([]*Item, error) {
	if count < 0 {
		return nil, fmt.Errorf("cannot sample negative number of items: %v", count)
	}
	results := make([]*Item, 0, count)
	for i := 0; i < count; i++ {
		item, err := p.Sample()
		if err != nil {
			return nil, errors.Wrap(err, "failed to sample item from population")
		}
		results = append(results, item)
	}
	return results, nil
}

// SampleWithoutReplacement samples count random items from the population without replacement
func (p *Population) SampleWithoutReplacement(count int) ([]*Item, error) {
	// Ensure there are enough items to sample
	if count < len(p.Items) {
		return nil, fmt.Errorf("cannot sample %v items from a population of size %v", count, len(p.Items))
	}
	// Create copy of search space to be modified throughout sampling (without replacement)
	searchSpace := make([]*searchRange, 0, len(p.searchSpace))
	for _, sr := range p.searchSpace {
		searchSpace = append(searchSpace, &searchRange{
			item: sr.item,
			min:  sr.min,
			max:  sr.max,
		})
	}
	// Sample count times, removing the selected searchRange from the slice each time
	results := make([]*Item, 0, count)
	totalWeight := p.TotalWeight
	for i := 0; i < count; i++ {
		// Sample a value from the remaining search space
		sampleVal := rand.Int63n(totalWeight)
		sample, index := find(searchSpace, sampleVal, 0, len(searchSpace))
		results = append(results, sample.item)
		// Remove searchRange from the search space
		searchSpace = append(searchSpace[:index], searchSpace[index+1:]...)
		// Update the searchRange values to the right of the removed index in the remaining search space by
		// decrementing the range bounds of each searchRange by the removed weight
		for si := index; si < len(searchSpace); si++ {
			searchSpace[si].min -= sample.item.Weight
			searchSpace[si].max -= sample.item.Weight
		}
	}
	return results, nil
}

// find executes a recursive binary search to find the searchRange at the specified value
func find(space []*searchRange, value int64, minIndex int, maxIndex int) (*searchRange, int) {
	midIndex := (minIndex + maxIndex) / 2
	midItem := space[midIndex]
	if value >= midItem.min && value < midItem.max {
		return midItem, midIndex
	} else if value < midItem.min {
		return find(space, value, minIndex, midIndex)
	}
	return find(space, value, midIndex+1, maxIndex)
}
