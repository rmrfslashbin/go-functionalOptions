# go-functionalOptions
PoC Go code to do functional options. 

## Purpose
This concept of functional options allows greater flexibility in passing config items into a struct. This basic example will get things going, but the concept is extendable to do all sorts of nifty things. A few of the significant benefits:
- No required order to pass in config data.
- Extendable API. Adding setters/getters won't break functionality.
- Ability to keep a config stuct private.
- Removes the need to init and pass a config struct to functions.
- Add ability to manipulate config data per setting (ex: reformat a text string before saving it to the config struct).
- Easily define presets and defaults.

## How It Works
### main.go
[main.go](https://github.com/rmrfslashbin/go-functionalOptions/blob/main/main.go) is the entry point. The function returns an initialized `thing.Config{}` struct called `myThing{}`. The `thing.New()` function accepts a variable amount of args (variadic function). Each optional function sets a value within the returned struct.

### thing/defs.go
[thing/defs.go](https://github.com/rmrfslashbin/go-functionalOptions/blob/main/thing/defs.go) defines two essential items:
- An `Option` func which is used as a return from the "set" functions.
- A `Config` struct that stores the various config settings. Private fields are defined within the struct to prevent manipulation/lookup outside the package.

### thing/functions.go
[thing/functions.go](https://github.com/rmrfslashbin/go-functionalOptions/blob/main/thing/functions.go) is where the magic happens.
- The `New()` function accepts a variable amount of functional args (`opts`). The function defines a new instance of the `Config{}` struct, sets the `color` field to `blue` (example of a default setting), processes the `opts` via a `for` loop, and executes each setter function. The function returns the newly created/configed `Config` struct.
- Next, getters are defined to return individual fields from the `Config` struct.
- Finally, setters are defined to set `Config` fields via variadic functions passed to `New()`.

## Summary
This PoC is a basic example to help bootstrap the concept. See the below links for more guidance and recipes. As always appreciate feedback and PRs to improve the PoC!

## Credit
Several writings inspire this PoC, but credit where due. I found these two posts to be most helpful:
- https://www.sohamkamani.com/golang/options-pattern/
- https://sagikazarmark.hu/blog/functional-options-on-steroids/

