# design-patterns-Golang

## Table of Contents

* [design-patterns-Golang](#design-patterns-golang)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Usage of Interfaces](#usage-of-interfaces)

## Summary

This repository is to capture not only design patterns in action but also any complex patterns in Golang. I'm sharing this as a starting point for others to learn Go patterns and experiment with them, and also to gather feedback on what we can improve. If you have a suggestion, feel free to file an [issue](https://github.com/ishan16696/design-patterns-Golang/issues).

## Usage of Interfaces

   - [To make sure all methods inside interface should be implemented](https://github.com/ishan16696/design-patterns-Golang/tree/main/Interface2)

        ```go
        type ShapeCalculator interface {
            Area() int
            Perimeter() int
        }

        type Rectangle struct {
            Width  int
            Height int
        }

        // following line guarantee that struct `Rectangle` must have to implement all methods of interface `ShapeCalculator`
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

