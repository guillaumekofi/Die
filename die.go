package main

// imports
import (
	"math/rand"
	"time"
)

// create a die of type struct
type die struct {
	numberOfFaces int
	faceValue int
}

// declare default variables
const (
	defaultNumFaces = 6
	defaultFaceValue = 0
)

func defaultDie() die {
	// create a new die with default values
	return die{numberOfFaces: defaultNumFaces, faceValue: defaultFaceValue}
}

func copyDie(d *die) die {
	// imitates the java copy constructor
	// creates a new die using an existing die
	return die{d.numberOfFaces, d.faceValue}
}

func newDie(faceValue int, numFaces int) die {
	// just to conform to the Java constructor style
	// can use this method if willing

	// creates a new die with user provided values
	return die{numberOfFaces: numFaces, faceValue: faceValue}
}

func otherDie(faceValue int) die {
	// just to conform to the Java constructor style
	// can use this method if willing

	// creates a new die with user provided face value and number of faces to default
	return die{numberOfFaces: defaultNumFaces, faceValue: faceValue}
}

func (d die) getFaceValue() int {
	// return the face value
	return d.faceValue
}

func (d die) getNumFaces() int {
	// return the number of faces
	return d.numberOfFaces
}

func (d *die) setFaceValue(value int) {
	// update the face value with user provided value
	d.faceValue = value
}

func (d *die) roll() {
	// randomly generates a face value and update faceValue
	source := rand.NewSource(time.Now().UnixNano())	// generate source
	r := rand.New(source)	// create new random with source
	d.faceValue =  r.Intn(d.getNumFaces()) + 1	// generate value and assign
}
