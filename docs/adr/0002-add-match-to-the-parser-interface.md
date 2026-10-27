# 2. Add Match to the Parser interface

Date: 2026-10-27

## Status

Accepted

## Context

`StringFrom`, `StringUntil`, and `Until` call parsers and discard their values. When the parser is a slice-producing parser such as `OneOrMore`, `Times`, or `All`, it appends every match to a slice, and the caller then discards the slice. `Whitespace` is `StringFrom(OneOrMore(RuneInRanges(unicode.White_Space)))`, so each call allocated a slice that grew with the amount of whitespace.

A memory profile of templ's parser benchmarks showed that `timesParser.Parse` and `allParser.Parse` were the two largest sources of allocations, at 10.6% and 9.3% of all allocated objects, and that 79% of the `timesParser` allocations came from `StringFrom`.

The API also had type-level problems that follow from the same cause. `StringFrom` required all of its parsers to have the same value type, so mixing `Times` and `Rune` needed nested `StringFrom` calls. `StringUntil` and `Until` had a type parameter for the delimiter's value type, which they never used.

## Decision

Add a `Matcher` interface, and make it part of `Parser`:

```go
type Matcher interface {
	Match(in *Input) (ok bool, err error)
}

type Parser[T any] interface {
	Matcher
	Parse(in *Input) (item T, ok bool, err error)
}
```

`Match` must return the same `ok` and `err` as `Parse`, and leave the input at the same position, without building the value.

Combinators that discard values take `Matcher`:

```go
func StringFrom(matchers ...Matcher) Parser[string]
func StringUntil(delimiter Matcher) Parser[string]
func StringUntilEOF(delimiter Matcher) Parser[string]
func Until[T any](parser Parser[T], delimiter Matcher) Parser[[]T]
func UntilEOF[T any](parser Parser[T], delimiter Matcher) Parser[[]T]
```

Every parser in the package implements `Match`. Parsers that build slices (`Times` and its variants, `All`, and `Until`) implement `Match` as a separate method with the same control flow as `Parse`, rather than sharing one method with a flag that switches between building and discarding values. `TestMatchAgreesWithParse` checks that the two methods agree. Parsers that wrap other parsers (`Any`, `Or`, `Optional`, `SequenceOf2` to `SequenceOf9`, and `StringFrom`) call `Match` on their children, so the saving reaches nested parsers. `Func` and `Convert` call the function and discard the value, because the function, or the conversion, can return an error that `Match` must also return.

The `Match` function parses and discards the value. It implements the `Match` method for parsers that have no cheaper way to match than to parse:

```go
func (p elementParser) Match(in *parse.Input) (ok bool, err error) {
	return parse.Match(p, in)
}
```

It takes the parser's concrete type as a type parameter, so that calling it does not convert the parser to an interface, which would allocate.

## Alternatives considered

- An unexported `skipper` interface that the combinators detect with a type assertion. This keeps the public API, but it puts a hidden fast path into the combinators, a parser defined outside the package cannot take part in it, and the type assertion costs time on each call. Using it for the `StringUntil` delimiter, which runs once per byte, made `StringUntil` 33% slower.
- An exported optional interface, as `io.Copy` does with `io.WriterTo`. Parsers outside the package can take part, but the combinators still detect it with a type assertion at run time.
- Returning concrete exported types from the constructors, so that they can carry `Match` without changing `Parser`. Go does not infer a type argument from a type that implements an interface, so `parse.Any(parse.String("a"), parse.Rune('b'))` would no longer compile.
- Repetition combinators that return a `string`, such as a string-returning `OneOrMore`. Each repetition combinator would need a second version, and an `All` nested inside one would still build a slice.
- Pooling result slices with `sync.Pool`. The caller owns the returned slice, so the library cannot know when to return it to the pool.

## Consequences

This is a breaking change. Every type outside the package that implements `Parser` must add a `Match` method. In templ, that is 22 types, and each one is three lines that call the `Match` function. Code that instantiates `StringFrom`, `StringUntil`, or `Until` with explicit type arguments must drop them. The `Match[T]` struct returned by `Optional` and `Or` is renamed to `Result[T]`, because the `Match` function needs the name.

`StringFrom` accepts parsers of different value types, so `StringFrom(Times(4, ZeroToNine), Rune('-'), Times(2, ZeroToNine))` compiles.

A parser can implement a `Match` that is cheaper than its `Parse`, and callers that only need the extent of a match benefit without changes.

Measured with templ's parser benchmarks (`benchstat`, n=10, Apple M1 Max), together with the other allocation changes made at the same time (pre-sized slices in `All` and `Times`, an allocation-free `Rune` for ASCII, and a single allocation for newline offsets in `NewInput`):

| Benchmark                               | allocs/op         | B/op              | sec/op |
|-----------------------------------------|-------------------|-------------------|--------|
| `Parse`                                 | 477 to 323        | 24.4Ki to 18.3Ki  | -15.6% |
| `ParseStringManyComponents` element, 40 | 68.95k to 49.46k  | 3.34Mi to 2.85Mi  | -15.1% |
| `ParseStringManyComponents` if, 40      | 255.2k to 186.3k  | 10.19Mi to 8.92Mi | -11.6% |

The `skipper` alternative reached 361, 55.86k, and 192.7k allocations, and 6.1%, 9.3%, and 6.3% less time.
