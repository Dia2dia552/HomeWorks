package structs

import (
	"fmt"
)

type Zookeeper struct {
	hasAccessToPredators bool
}

type Animal struct {
	zookeeper  Zookeeper
	isPredator bool
}

type Cage struct {
	animal Animal
	isOpen bool
}

func (cage *Cage) OpenCage(zookeeper Zookeeper) {
	if zookeeper.hasAccessToPredators {
		cage.isOpen = true
		fmt.Println("Cage is now open.")
	} else {
		fmt.Println("Zookeeper does not have access to predators. Cannot open the cage.")
	}
}

func (cage *Cage) CloseCage() {
	cage.isOpen = false
	fmt.Println("Cage is now closed.")
}

func (cage *Cage) PutAnimalInCage(animal Animal) {
	if cage.isOpen {
		cage.animal = animal
		fmt.Println("Animal has been placed in the cage.")
	} else {
		fmt.Println("Cannot put the animal in the cage because the cage is not open.")
	}
}

func start() {
	zookeeperWithAccess := Zookeeper{hasAccessToPredators: true}

	zookeeperWithoutAccess := Zookeeper{hasAccessToPredators: false}

	cageWithPredator := Cage{animal: Animal{zookeeper: zookeeperWithAccess, isPredator: true}}
	cageWithPredator.OpenCage(zookeeperWithAccess)

	anotherCageWithPredator := Cage{animal: Animal{zookeeper: zookeeperWithoutAccess, isPredator: true}}
	anotherCageWithPredator.OpenCage(zookeeperWithoutAccess)

	cageWithPredator.CloseCage()
}
