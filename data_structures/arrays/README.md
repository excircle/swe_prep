# 'Array' Data Structure

### What Is An 'Array?'

An 'Array' is formally defined as:

>...a data structure consisting of a collection of elements (values or variables), each identified by at least one array index or key. 

In it's most basic and non-descriptive terms, it is a list.

An array in Go is a fixed-length data type that contains a contiguous block of elements of the same type.

### Arrays In Golang

To create an Array in Go, you need to provide the number of items that will be stored and the type that each item in the array will be.  Consider the example below:

```Go
var arr [5]int
```

The above shows that variable "arr" will hold 5 items of type "int."

If the brackets '[5]' have a number in them, it is considered an array in Go.

If the brackets '[]' DO NOT have a number in them, they are considered a 'Slice' (which we will discuss in greater length later).

<b>Once an array is declared in Go, the type it holds and it's length cannot be changed.</b>

Like "list" or "dictionary" objects in other languages, the declarations of the values at certain indexs can arbitrarily declared like so:

```Go
var arr [5]int
arr[1] = 6
fmt.Println(arr)
// Returns: [0 6 0 0 0]
```

Notice how we did not have to write to the first available index.


<b>Golang Array Internals</b>

The Golang array is simple. It is indexed (starting at 0) and provides the ability to hold in each index a typed object. Below we see an array with 5 indexes holding int type objects.

![alt text](http://ascetism.com/static/array_internals.png "Golang_Array_Internals")

# Golang 'Slices' As Dynamic, Multi-Dimensional Arrays

While Golang arrays are not dynamic, they are multi-dimensional due to the fact that they can hold multiple subarrays.

An example of this is shown below:

```Go
func main() {
	a := [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6}}
	fmt.Println(a)
}
// Returns: [[1 2] [3 4] [5 6]]
```

<b>But how do we increase the capacity of this array without re-creating the data??</b>

SLICES.

### How are Slices different than arrays?

