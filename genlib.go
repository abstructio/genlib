package genlib

import (
	"math/rand"
	"sort"
)

type Population []*Genom
type Genom interface {
	//IMPORTANT: the struct must contain Attribute []interface{}!!!!

	//Get Fitness
	Fit() float64
	//Get Attribute
	Attribute(index int) interface{}
	//Set Attribute
	SetAttribute(index int, value interface{})
	//Mutate a Attribute
	Mutate()
}

func (p Population) Reproduction() Population {
	nextGen := make(Population, 0)
	for (len(p) > 0) && (len(nextGen) < len(p)) {
		//find parents and do a cross over
		for {
			first := rand.Intn(len(p))
			second := rand.Intn(len(p))
			if first != second {
				p[first].crossOver(p[second])
				p[first].Mutate()
				p[second].Mutate()
				nextGen = append(nextGen, p[first])
				nextGen = append(nextGen, p[second])
				p.deleteGenom(first)
				p.deleteGenom(second)
			}
		}
	}
	return nextGen
}
func (p Population) deleteGenom(index int) {
	copy(p[index:], p[index+1:])
	p[len(p)-1] = nil // or the zero value of T
	p = p[:len(p)-1]
}
func (p Population) Selection() Population {
	return stochasticUniversalSampling(p, 0)
}

func (g *Genom) crossOver(o *Genom) {
	q := rand.Intn(len(g.Attribute))
	p := rand.Intn(len(g.Attribute))
	if p > q {
		p, q = q, p
	}
	for i := p; i < q; i++ {
		g.Attribute[i], o.Attribute[i] = o.Attribute[i], g.Attribute[i]
	}

}

//Selection functions
//better selection function
func stochasticUniversalSampling(p Population, selectionSize int) Population {
	var (
		aggregateFitness, startOffset, cumulativeExpectation float64
		index                                                int
	)
	if selectionSize == 0 {
		selectionSize = len(p)
	}
	for _, i := range Population {
		aggregateFitness += i.Fit()
	}
	selection := make(Population, 0)
	startOffset = rand.Float64()
	cumulativeExpectation = 0
	for _, i := range p {
		cumulativeExpectation += ((i.Fit() / aggregateFitness) * selectionSize)
		for cumulativeExpectation > (startOffset + index) {
			selection = append(selection, i)
		}
	}
	return selection
}

//simplest selection function
func (p Population) getBestN(n int) Population {
	selection = make(Population, 0)
	sort.Sort(p)
	for i := 0; i < n; i++ {
		selection = append(selection, p[i])
	}
	return selection
}

func (p Population) Len() int {
	return len(p)
}
func (p Population) Less(i, j int) bool {
	return p[i].Fit() < p[j].Fit()
}
func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
