# design-patterns-Golang

## Table of Contents

* [design-patterns-Golang](#design-patterns-golang)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Usage of Interfaces](#usage-of-interfaces)
  * [Factory Method](#factory-method)
  * [Builder](#builder)
  * [Decorator](#decorator)
  * [Singleton](#singleton)
  <!-- * [Observer Pattern](#observer-pattern) -->

## Summary

This repository is to capture not only design patterns in action but also any complex patterns in Golang. I'm sharing this as a starting point for others to learn Go patterns and experiment with them, and also to gather feedback on what we can improve. If you have a suggestion, feel free to file an [issue](https://github.com/ishan16696/design-patterns-Golang/issues).

## Usage of Interfaces

   - [To enforce at compile time all methods inside interface should be implemented.](https://github.com/ishan16696/design-patterns-Golang/tree/main/Interface2)
  
        ```go
        type ShapeCalculator interface {
            Area() int
            Perimeter() int
        }

        type Rectangle struct {
            Width  int
            Height int
        }

        // following line guarantee that struct `Rectangle` must have to implement all methods of interface `ShapeCalculator`.
        // purpose of this pattern is to enforce at compile time that `*Rectangle` implements the interface.
        var _ ShapeCalculator = (*Rectangle)(nil)

        // another way of using this
        var _ ShapeCalculator = &Rectangle{}
        ```

   - [Another pattern of using interface](https://github.com/ishan16696/design-patterns-Golang/tree/main/Interface)

        ```go
        type Action interface {
            // Do performs an action.
            Do(ctx context.Context)
        }

        // ActionFunc is a function that implements Action.
        type ActionFunc func(ctx context.Context)

        // Do performs an action.
        func (f ActionFunc) Do(ctx context.Context) {
            f(ctx)
        }
        ```

## Factory Method

 - Define an interface for creating an object, but let subclasses/Methods decide which class to instantiate. Defer the instantiation to subclasses/Methods. check example [here](https://github.com/ishan16696/design-patterns-Golang/tree/main/factoryMethod).

    1. Define `Factory` interface.

    2. Defer the instantiation of object to subclasses/Methods.

        ```go
        type Factory interface {
            NewPerson() *person
            NewAgent() *secretAgent
        }

        // create object of factory
        fact := factory.NewFactory(config)

        // with factory method, we can construct the instances of both person/agent.
        person := fact.NewPerson()
        agent := fact.NewAgent()
        ```

## Builder

  - It's a creational design pattern that separate the construction of a complex object or segregating the one builder into multiple builders. It is used to construct a complex object step by step and the final step will return the object. Check example [here](https://github.com/ishan16696/design-patterns-Golang/tree/main/Builder).

    ```go
    // instead of passing all parameters at once via constructor to create object, we can use builder pattern to create object step by step.
    emp1 := dev.SetName("Ishan").SetTechStack([]string{"C++", "Docker", "Go"}).SetEmpID(2).BuildDev()
    ```

## Singleton

  - It's creational desing pattern that ensures that a class has only one instance and provides a global point of access to it. This is useful when exactly one object is needed to coordinate actions across a system. Check example [here](https://github.com/ishan16696/design-patterns-Golang/tree/main/Singleton).

    ```go
    type Singleton struct {
        // Your singleton fields
    }

    var (
        instance *Singleton
        once     sync.Once
    )

    func GetInstance() *Singleton {
        once.Do(func() {
            instance = &Singleton{
                // ...
            }
        })
        return instance
    }
    ```

    > Note: Always use `sync.Once` in Go for thread safe.

## Decorator

  - It's a structural design pattern that allows you to dynamically add new behavior to existing objects by placing them inside special wrapper objects. check example [here](https://github.com/ishan16696/design-patterns-Golang/tree/main/Decorator).

    ```go
    // Base object
    type Book struct {
        ...
    }

    // Decorator which wraps the original base object
    type BorrowableBook struct {
        Book       // Embed the original object
        IsBorrowed bool
    }

    // Adds new behavior using decorator object
    func (b *BorrowableBook) Borrow() {
        if b.IsBorrowed {
            ...
        }
        ...
    }
    ```
