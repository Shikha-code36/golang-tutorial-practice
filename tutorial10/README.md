# Generics in GO

## Key Concepts

### Generics

Generics is a concept in programming where a function or a data type can be written in a way that it can handle different data types. In Go, generics are defined using square brackets `[]` followed by a type parameter list.

### Generic Functions

In this program, we have three generic functions: `sumSlice`, `isEmpty`, and `LoadJSON`.

```go
func sumSlice[T int | float32 | float64](slice []T) T{
    var sum T
    for _, v := range slice{
        sum += v
    }
    return sum
}

func isEmpty[T any](slice []T) bool{
    return len(slice) == 0
}

func LoadJSON[T contactInfo | purchaseInfo] (filePath string) []T{
    data, _ := ioutil.ReadFile(filePath)    
    var loaded = []T{}  
    json.Unmarshal(data, &loaded)

    return  loaded
}
```

`sumSlice` takes a slice of type `T` where `T` can be `int`, `float32`, or `float64`. It returns the sum of all elements in the slice.

`isEmpty` takes a slice of any type `T` and returns a boolean indicating whether the slice is empty.

`LoadJSON` takes a file path as input and returns a slice of type `T` where `T` can be `contactInfo` or `purchaseInfo`. It reads the JSON file from the given file path and unmarshals the data into a slice of the appropriate type.

### Generic Structs

In this program, we have two struct types: `contactInfo` and `purchaseInfo`.

```go
type contactInfo struct{
    Name string 
    Email string    
}
    
type purchaseInfo struct{   
    Name string
    Price float32
    Amount int
}
```

These structs are used as type parameters in the `LoadJSON` function to load JSON data of different types.

## Usage

In the `main` function, we demonstrate the usage of these generic functions.

```go
func main(){
    var intSlice = []int{1,2,3}
    fmt.Println(sumSlice[int](intSlice))

    var float32Slice = []float32{1,2,3}
    fmt.Println(sumSlice[float32](float32Slice))

    var sliceint = []int{}
    fmt.Println(isEmpty[int](sliceint))
    fmt.Println(isEmpty[int](intSlice))

    var contacts []contactInfo = LoadJSON[contactInfo]("./files/contactInfo.json")
    fmt.Printf("\n%+v", contacts)

    var purchases []purchaseInfo = LoadJSON [purchaseInfo]("./files/PurchaseInfo.json")
    fmt.Printf("\n%+v", purchases)
}
```

We create slices of `int` and `float32` and use `sumSlice` to calculate their sums. We also check if a slice is empty using `isEmpty`. Finally, we use `LoadJSON` to load JSON data into slices of `contactInfo` and `purchaseInfo`.

### Checkout the code

- [main.go](main.go)
