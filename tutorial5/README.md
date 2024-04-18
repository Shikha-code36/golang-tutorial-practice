# Strings, Runes, and UTF-8 Encoding in Go

This will guide you through some basic concepts in Go, including strings, runes, and UTF-8 encoding. The code we'll be discussing demonstrates these concepts in a simple and understandable way.

## Strings and UTF-8 Encoding

In Go, a string is a sequence of bytes. This means when you create a string in Go, it's encoded in UTF-8 by default. UTF-8 is a variable-width character encoding that can represent every character in the Unicode standard, yet the initial encoding of byte sequences is ASCII.

Let's take a look at this piece of code:

```go
var myString = "résumé"
var indexed = myString[1]
fmt.Println(indexed)
fmt.Printf("%v %T \n", indexed, indexed)
```

Here, `myString[1]` doesn't give you the second character 'é', but instead it gives you a number which is the Unicode of the second byte of 'é'. This is because 'é' is a non-ASCII character represented by two bytes in UTF-8.

## Runes

A rune in Go is a type that represents a Unicode CodePoint. It does not matter if the character is ASCII or not, it can be represented as a rune. A rune literal is just a number. We usually use it to represent a character's unicode code point.

Let's modify the code to use runes:

```go
var myString2 = []rune("résumé")
var indexed2 = myString2[1]
fmt.Println(indexed2)
fmt.Printf("%v %T \n", indexed2, indexed2)
```

Now `myString2[1]` gives you the Unicode of 'é', because we're treating the string as a sequence of runes instead of a sequence of bytes.

## String Concatenation

In Go, you can concatenate strings using the `+` operator:

```go
var strSlice = []string{"s", "h", "i", "k", "h", "a"}
var catStr = ""
for i := range strSlice{
	catStr += strSlice[i]
}
fmt.Printf("\nString building: %v",catStr)
```

However, since strings are immutable in Go, string concatenation using `+` creates a new string, which can be inefficient when concatenating a large number of strings. A more efficient way is to use the `strings.Builder`:

```go
var strBuilder strings.Builder
for i := range strSlice{
	strBuilder.WriteString(strSlice[i])
}
var catStr1 = strBuilder.String()
fmt.Printf("\nString building through built in package: %v",catStr1)
```

### Checkout the code

 - [main.go](main.go)