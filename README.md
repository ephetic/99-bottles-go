# 99-bottles-go
Doing something like Sandi Metz's `99 Bottles of OOP`

`99 Bottles` used Ruby and acheived elegant reuse through inheritance and polymorphism.  This repo is an attempt to do the same with Go.

By implementing the interface for a "base" struct which is then embedded in "child" structs, the base implementations are available and overridable, **BUT** the dispatch doesn't behave as one would like for inheritance.  Namely, methods that call methods will call the versions in their implementation.  The workaround here is to copy and paste the `String` method into each of the child implementations as it does most of the cross-implementation calls.  Similar smaller methods were simply removed (see git history).  The net result is that I strayed from the goal of this exercize and turned it back into a toy that doesn't serve its goals.

Further research is needed to see if other language features/patterns can avoid this dispatch issue.
