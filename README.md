
## Eat This Much Sample Generator Project

### Project Requirements
- Implement the `Generator` interface and modify `getGenerator()` in `main.go` to return your generator.  
- The number of meals and foods per meal is entirely up to you but preferably somewhat similar to what a normal person might eat.
- Don't include any foods whose name contains any keywords passed in `exclusions`
- Your generator shouldn't produce the exact same output every time - incorporate some randomness.
- Aim for less than 10% calorie target deviation without any exclusions, return an error if that isn't possible 

### Testing your generator
You can run the program with `go run sample_generator` 
- can optionally specify the target calories with an argument (defaults to 2000): `-cals 1000`