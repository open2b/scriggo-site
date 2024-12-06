{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: builtins{% end %}
{% Article %}

{% raw doc %}

# Builtins

Scriggo includes several builtins, in addition to the [Go builtins](https://pkg.go.dev/builtin), that can be used in
templates. Some of these builtins are types with methods, so you can call methods on the values of these types. The
[scriggo serve](/scriggo-command#serve-a-template) command supports all of them.

Part of the documentation on this page is copyright of
<a href="https://raw.githubusercontent.com/golang/go/master/LICENSE" target="_blank">The Go Authors</a>.

### crypto

* [hmacSHA1](#hmacsha1)
* [hmacSHA256](#hmacsha256)
* [sha1](#sha1)
* [sha256](#sha256)
  
### encoding

* [base64](#base64)
* [hex](#hex)
* [marshalJSON](#marshaljson)
* [marshalJSONIndent](#marshaljsonindent)
* [marshalYAML](#marshalyaml)
* [md5](#md5)
* [unmarshalJSON](#unmarshaljson)
* [unmarshalYAML](#unmarshalyaml)

### html

* [htmlEscape](#htmlescape)

### math

* [abs](#abs)
* [max](#max)
* [min](#min)
* [pow](#pow)

### net

* [File](#file)
* [FormData](#formdata)
* [form](#form)
* [queryEscape](#queryescape)

### regexp

* [Regexp](#regexp-2)
* [regexp](#regexp-3)

<h3>sort</h3>

* [reverse](#reverse)
* [sort](#sort)

### strconv

* [formatFloat](#formatfloat)
* [formatInt](#formatint)
* [parseFloat](#parsefloat)
* [parseInt](#parseint)

### strings

* [abbreviate](#abbreviate)
* [capitalize](#capitalize)
* [capitalizeAll](#capitalizeall)
* [hasPrefix](#hasprefix)
* [hasSuffix](#hassuffix)
* [index](#index)
* [indexAny](#indexany)
* [join](#join)
* [lastIndex](#lastindex)
* [replace](#replace)
* [replaceAll](#replaceall)
* [runeCount](#runecount)
* [split](#split)
* [splitAfter](#splitafter)
* [splitAfterN](#splitaftern)
* [splitN](#splitn)
* [sprint](#sprint)
* [sprintf](#sprintf)
* [toKebab](#tokebab)
* [toLower](#tolower)
* [toUpper](#toupper)
* [trim](#trim)
* [trimLeft](#trimleft)
* [trimPrefix](#trimprefix)
* [trimRight](#trimright)
* [trimSuffix](#trimsuffix)
* [unsafeconv](#unsafeconv)
  * [ToCSS](#tocss)
  * [ToHTML](#tohtml)
  * [ToJS](#tojs)
  * [ToJSON](#tojson)
  * [ToMarkdown](#tomarkdown)

<h3>time</h3>

* [Duration](#duration)
* [Hour](#hour)
* [Microsecond](#microsecond)
* [Millisecond](#millisecond)
* [Minute](#minute)
* [Nanosecond](#nanosecond)
* [Second](#second)
* [Time](#time)
* [date](#date)
* [now](#now)
* [parseDuration](#parseduration)
* [parseTime](#parsetime)
* [unixTime](#unixtime)

## Crypto

### hmacSHA1

```go
func hmacSHA1(message, key string) string
```

Returns the HMAC-SHA1 tag for the given message and key, as a base64 encoded string.

### hmacSHA256

```go
func hmacSHA256(message, key string) string
```

Returns the HMAC-SHA256 tag for the given message and key, as a base64 encoded string.

### sha1

```go
func sha1(s string) string
```

Returns the SHA1 checksum of s as a hexadecimal encoded string.

### sha256

```go
func sha256(s string) string
```

Returns the SHA256 checksum of s as a hexadecimal encoded string.

## encoding

### base64

```go
func base64(s string) string
```

Returns the base64 encoding of s.

### hex

```go
func hex(s string) string
```

Returns the hexadecimal encoding of s.

### marshalJSON

```go
func marshalJSON(v any) (json, error)
```

Returns the JSON encoding of v. For details, see the [json.Marshal](https://pkg.go.dev/encoding/json#Marshal) function
of the Go standard library.

### marshalJSONIndent

```go
func marshalJSONIndent(v any, prefix, indent string) (json, error)
```

Is like [marshalJSON](#marshaljson) but indents the output. Each JSON element in the output will begin on a new line
beginning with prefix followed by one or more copies of indent according to the indentation nesting. prefix and
indent can only contain whitespace: ' ', '\t', '\n' and '\r'.

### marshalYAML

```go
func marshalYAML(v any) (string, error)
```

Returns the YAML encoding of v. For details, see the [yaml.Marshal](https://pkg.go.dev/gopkg.in/yaml.v3#Marshal)
function of the yaml package.

### md5

```go
func md5(s string) string
```

Returns the MD5 hash of the string s.

### unmarshalJSON

```go
func unmarshalJSON(data string, v any) error
```

Parses the JSON-encoded data and stores the result in a new value pointed to by v. If v is nil or not a pointer,
unmarshalJSON returns an error.

Unlike [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) of the Go standard library, unmarshalJSON does not
change the value pointed to by v but instantiates a new value and then replaces the value pointed to by v, if no
errors occur.

For details, see the [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function of the Go standard library.

### unmarshalYAML

```go
func unmarshalYAML(data string, v any) error
```

Parses the YAML-encoded data and stores the result in a new value pointed to by v. If v is nil or not a pointer,
unmarshalYAML returns an error.

Unlike [yaml.Unmarshal](https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal) of the yaml package, unmarshalYAML does not
accept a map value, use a pointer to a map value instead. unmarshalYAML does not change the value pointed to by v but
instantiates a new value and then replaces the value pointed to by v, if no errors occur.

For example, given a file named "receipt.yaml", in the current template directory, with content:

```
---
receipt:     Oz-Ware Purchase Invoice
date:        2012-08-06
customer:
  first_name:   Doroth
  family_name:  Gale
```

the following code reads and decodes the "example.yaml" file:

```scriggo
{% var v map[string]any %}
{{ unmarshalYAML(render "receipt.yaml", &v) }}
{{ v.customer.first_name }} {{ v.customer.family_name }}
```

For details, see the [yaml.Unmarshal](https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal) function of the yaml package.

## html

### htmlEscape

```go
func htmlEscape(s string) html
```

Escapes s, replacing the characters <, >, &amp;, " and ' and returns the escaped string as html type.

## math

### abs

```go
func abs(x int) int
```

Returns the absolute value of x.

### max

```go
func max(x, y int) int
```

Returns the larger of x or y.

### min

```go
func min(x, y int) int
```

Returns the smaller of x or y.

### pow

```go
func pow(x, y float64) float64
```

Returns x**y. For details, see the [math.Pow](https://pkg.go.dev/math#Pow) function of the Go standard library.

## net

### File

```go
type File interface {
    Name() string // name.
    Type() string // type, as mime type.
    Size() int    // size, in bytes.
}
```

Represents a file.

### FormData

```go
type FormData struct {
    // Contains unexported fields.
}

// ParseMultipart parses a request body as multipart/form-data. If the body
// is not multipart/form-data, it does nothing. It should be called before
// calling the File and Files methods and can be called multiple times.
//
// It panics if an error occurs.
func (form FormData) ParseMultipart()

// Value returns the first form data value associated with the given field. If
// there are no files associated with the field, it returns an empty string.
//
// It panics if an error occurs.
func (form FormData) Value(field string) string

// Values returns the parsed form data, including both the URL field's query
// parameters and the POST form data.
//
// It panics if an error occurs.
func (form FormData) Values() map[string][]string

// File returns the first file associated with the given field. It returns nil
// if there are no values associated with the field.
//
// Call File only after ParseMultipart is called.
func (form FormData) File(field string) File

// Files returns the parsed files of a multipart form. It returns a non nil
// map, if ParseMultipart has been called.
//
// Call Files only after ParseMultipart is called.
func (form FormData) Files() map[string][]File
```

Represents the form data from the query string and body of an HTTP request.

### form

```go
var form FormData
```

Contains the form data from the query string and body of the HTTP request. Values and files passed as
multipart/form-data are only available after the ParseMultipart method is called.

### queryEscape

```go
func queryEscape(s string) string
```

Escapes the string, so it can be safely placed inside a URL query.

## regexp

### Regexp

```go
type Regexp struct {
    // Contains unexported fields.
}

// Match reports whether the string s contains any match of the regular
// expression.
func (re Regexp) Match(s string) bool

// Find returns a string holding the text of the leftmost match in s of the
// regular expression. If there is no match, the return value is an empty
// string, but it will also be empty if the regular expression successfully
// matches an empty string. Use FindSubmatch if it is necessary to distinguish
// these cases.
func (re Regexp) Find(s string) string

// FindAll is the 'All' version of Find; it returns a slice of all successive
// matches of the expression, as defined by the 'All' description in the Go
// regexp package comment. A return value of nil indicates no match.
func (re Regexp) FindAll(s string, n int) []string

// FindAllSubmatch is the 'All' version of FindSubmatch; it returns a slice of
// all successive matches of the expression, as defined by the 'All'
// description in the Go regexp package comment. A return value of nil
// indicates no match.
func (re Regexp) FindAllSubmatch(s string, n int) [][]string

// FindSubmatch returns a slice of strings holding the text of the leftmost
// match of the regular expression in s and the matches, if any, of its
// subexpressions, as defined by the 'Submatch' description in the Go regexp
// package comment. A return value of nil indicates no match.
func (re Regexp) FindSubmatch(s string) []string

// ReplaceAll returns a copy of s, replacing matches of the Regexp with the
// replacement string repl. Inside repl, $ signs are interpreted as in Expand
// method of the Go regexp package, so for instance $1 represents the text of
// the first submatch.
func (re Regexp) ReplaceAll(s, repl string) string

// ReplaceAllFunc returns a copy of s in which all matches of the Regexp
// have been replaced by the return value of function repl applied to the
// matched substring. The replacement returned by repl is substituted
// directly, without expanding.
func (re Regexp) ReplaceAllFunc(s string, repl func(string) string) string

// Split slices s into substrings separated by the expression and returns a
// slice of the substrings between those expression matches.
//
// The slice returned by this method consists of all the substrings of s not
// contained in the slice returned by FindAllString. When called on an
// expression that contains no metacharacters, it is equivalent to SplitN.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
func (re Regexp) Split(s string, n int) []string
```

Represents a regular expression.

### regexp

```go
func regexp(expr string) Regexp
```

Parses a regular expression and returns a [Regexp](#regexp-2) value that can be used to match against text. It panics if
the expression cannot be parsed.

For the syntax of regular expressions see [https://pkg.go.dev/regexp/syntax](https://pkg.go.dev/regexp/syntax).

<h2>sort</h2>

### reverse

```go
func reverse(slice any)
```

Reverses the order of the elements of slice. If slice is not a slice, it panics.

### sort

```go
func sort(slice any, less func(i, j int) bool)
```

Sorts the elements of slice given the provided less function. If slice is not a slice, it panics.

The less function reports whether `slice[i]` should be ordered before `slice[j]`. If less is nil, sort sorts the
elements in a natural order based on their type.

The natural order can differ between different versions of Scriggo.

#### Examples

```scriggo
{% var s = []int{3, 5, 2, 9, 7} %}
{{ sort(s, nil) }}
```
<pre class="result">2 3 5 7 9</pre>

```scriggo
{% var s = []int{3, 5, 2, 9, 7} %}
{% var before = func(i, j int) int { return s[i] > s[j] }) %}
{{ sort(s, before) }}
```
<pre class="result">9 7 5 3 2</pre>

## strconv

### formatFloat

```go
func formatFloat(f float64, format string, precision int) string
```

Converts the floating-point number f to a string, according to the given format and precision. It can round the
result.

The format is one of "e", "f" or "g"

    "e": -d.dddde±dd, a decimal exponent
    "f": -ddd.dddd, no exponent
    "g": "e" for large exponents, "f" otherwise

The precision, for -1 <= precision <= 1000, controls the number of digits (excluding the exponent). The special
precision -1 uses the smallest number of digits necessary such that _parseFloat_ will return f exactly. For "e" and "f"
it is the number of digits after the decimal point. For "g" it is the maximum number of significant digits (trailing
zeros are removed).

If the format or the precision is not valid, _formatFloat_ panics.

### formatInt

```go
func formatInt(i int, base int) string
```

Returns the string representation of i in the given base, for 2 <= base <= 36. The result uses the lower-case
letters 'a' to 'z' for digit values >= 10.

It panics if base is not in the range.

### parseFloat

```go
func parseFloat(s string) (float64, error)
```

Converts the string s to a float64 value.

If s is well-formed and near a valid floating-point number, _parseFloat_ returns the nearest floating-point number
rounded using IEEE754 unbiased rounding.

### parseInt

```go
func parseInt(s string, base int) (int, error)
```

Interprets a string s in the given base, for 2 <= base <= 36, and returns the corresponding value. It returns 0 and
an error if s is empty, contains invalid digits or the value corresponding to s cannot be represented by an int
value.

## strings

### abbreviate

```go
func abbreviate(s string, n int) string
```

Abbreviates s to almost n runes. If s is longer than n runes, the abbreviated string terminates with "...".

### capitalize

```go
func capitalize(s string) string
```

Returns a copy of the string s with the first non-separator in upper case.

### capitalizeAll

```go
func capitalizeAll(s string) string
```

Returns a copy of the string s with the first letter of each word in upper case.

### hasPrefix

```go
func hasPrefix(s, prefix string) bool
```

Tests whether the string s begins with prefix.

### hasSuffix

```go
func hasSuffix(s, suffix string) bool
```

Tests whether the string s ends with suffix.

### index

```go
func index(s, substr string) int
```

Returns the index of the first instance of substr in s, or -1 if substr is not present in s.

The returned index refers to the bytes and not the runes of s.

### indexAny

```go
func indexAny(s, chars string) int
```

Returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point
from chars is present in s.

The returned index refers to the bytes and not the runes of s.

### join

```go
func join(a []string, sep string) string
```

Concatenates the elements of its first argument to create a single string. The separator string sep is placed between
elements in the resulting string.

### lastIndex

```go
func lastIndex(s, sep string) int
```

Returns the index of the last instance of substr in s, or -1 if substr is not present in s.

The returned index refers to the bytes and not the runes of s.

### replace

```go
func replace(s, old, new string, n int) string
```

Returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If n < 0,
there is no limit on the number of replacements.

### replaceAll

```go
func replaceAll(s, old, new string) string
```

Returns a copy of the string s with all non-overlapping instances of old replaced by new.

### runeCount

```go
func runeCount(s string) int
```

Returns the number of runes in s. Erroneous and short encodings are treated as single runes of width 1 byte.

### split

```go
func split(s, sep string) []string 
```

Slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

If s does not contain sep and sep is not empty, split returns a slice of length 1 whose only element is s. If
sep is empty, split splits after each UTF-8 sequence. If both s and sep are empty, split returns an empty slice.

It is equivalent to [SplitN](#splitn) with a count of -1.

### splitAfter

```go
func splitAfter(s, sep string) []string 
```

Slices s into all substrings after each instance of sep and returns a slice of those substrings.

If s does not contain sep and sep is not empty, splitAfter returns a slice of length 1 whose only element is s.
If sep is empty, splitAfter splits after each UTF-8 sequence. If both s and sep are empty, splitAfter returns
an empty slice.

It is equivalent to [splitAfterN](#splitaftern) with a count of -1.

For example:

```scriggo
{% for s in splitAfter("<section><div>a</div><div>b</div></section>", ">") %}
{{ s }}
{% end %}
```
<div class="result"><pre>
<div>&lt;section&gt;<br>a&lt;/div&gt;<br>&lt;div&gt;<br>b&lt;/div&gt;<br>&lt;/section&gt;
</pre></div>

### splitAfterN

```go
func splitAfterN(s, sep string, n int) []string 
```

Slices s into substrings after each instance of sep and returns a slice of those substrings.

The count determines the number of substrings to return:

* n > 0: at most n substrings; the last substring will be the unsplit remainder.
* n == 0: the result is nil (zero substrings)
* n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for splitAfter.

### splitN

```go
func splitN(s, sep string, n int) []string
```

Slices s into substrings separated by sep and returns a slice of the substrings between those separators.

The count determines the number of substrings to return:

* n > 0: at most n substrings; the last substring will be the unsplit remainder.
* n == 0: the result is nil (zero substrings)
* n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for split.

### sprint

```go
func sprint(a ...any) string
```

Formats using the default formats for its operands and returns the resulting string. Spaces are added between operands
when neither is a string.

For example:

```scriggo
{% const name, age = "Sherlock Holmes", 60 %}
{{ sprint(name, " is ", age, " years old") }}
```
<div class="result"><pre>Sherlock Holmes is 60 years old</pre></div>

sprint takes the same arguments and behaves like the [fmt.Sprint](https://pkg.go.dev/fmt#Sprint) function of the
Go standard library.

### sprintf

```go
func sprintf(format string, a ...any) string
```

Formats according to a format specifier and returns the resulting string.

```scriggo
{% const name, age = "Sherlock Holmes", 60 %}
{{ sprintf("%s is %d years old", name, age) }}
```
<div class="result"><pre>Sherlock Holmes is 60 years old</pre></div>

sprintf takes the same arguments and behaves like the [fmt.Sprintf](https://pkg.go.dev/fmt#Sprintf) function of
the Go standard library.

### toKebab

```go
func toKebab(s string) string
```

Returns a copy of s in kebab case. For example:

```scriggo
{{ toKebab("borderTopColor") }}
```
<pre class="result">border-top-color</pre>

### toLower

```go
func toLower(s string) string
```

Returns s with all Unicode letters mapped to their lower case.

### toUpper

```go
func toUpper(s string) string
```

Returns s with all Unicode letters mapped to their upper case.

### trim

```go
func trim(s string, cutset string) string
```

Returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.

### trimLeft

```go
func trimLeft(s string, cutset string) string
```

Returns a slice of the string s with all leading Unicode code points contained in cutset removed.

To remove a prefix, use [trimPrefix](#trimprefix) instead.

### trimPrefix

```go
func trimPrefix(s, prefix string)
```

Returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

### trimRight

```go
func trimRight(s string, cutset string) string
```

Returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.

To remove a suffix, use [trimSuffix](#trimsuffix) instead.

### trimSuffix

```go
func trimSuffix(s, suffix string) string
```

Returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.


### unsafeconv

A package that implements functions to make unsafe conversions between string values and native types.

#### ToCSS

```go
func ToCSS(s string) css
```

Returns s but with the css type. For example:

```go
unsafeconv.ToCSS("data:image/gif;base64," + base64(image))
```

#### ToHTML

```go
func ToHTML(s string) html
```

Returns s but with the html type. For example:

```go
unsafeconv.ToHTML(`<a href="` + htmlEscape(url) + `">see this page</a>`)
```

#### ToJS

```go
func ToJS(s string) js
```

Returns s but with the js type. For example:

```go
unsafeconv.ToJS("var a = " + v + ";")
```

#### ToJSON

```go
func ToJSON(s string) json
```

Returns s but with the json type. For example:

```go
unsafeconv.ToJSON(`{ "config": ` + config + ` }`)
```

#### ToMarkdown

```go
func ToMarkdown(s string) markdown
```

Returns s but with the markdown type. For example:

```go
unsafeconv.ToMarkdown("# " + title)
```

<h2>time</h2>

### Duration

```go
type Duration int64

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64

// Microseconds returns the duration as an integer microsecond count.
func (d Duration) Microseconds() int64

// Milliseconds returns the duration as an integer millisecond count.
func (d Duration) Milliseconds() int64

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Duration) Round(m Duration) Duration

// Seconds returns the duration as a floating point number of seconds.
func (d Duration) Seconds() float64

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0s.
func (d Duration) String() string

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m Duration) Duration
```

A _Duration_ represents the elapsed time between two instants as an int64 nanosecond count. The representation limits
the largest representable duration to approximately 290 years.

### Hour

```go
const Hour Duration = 60 * Minute
```

Represents an hour.

### Microsecond

```go
const Microsecond Duration = 1000 * Nanosecond
```

Represents a microsecond.

### Millisecond

```go
const Millisecond Duration = 1000 * Microsecond
```

Represents a millisecond.

### Minute

```go
const Minute Duration = 60 * Second
```

Represents a minute.

### Nanosecond

```go
const Nanosecond Duration = 1
```

Represents a nanosecond.

### Second

```go
const Second Duration = 1000 * Millisecond
```

Represents a second.

### Time

```go
type Time struct {
    // Contains unexported fields.
}

// Add returns the time t+d.
func (t Time) Add(d Duration) Time

// AddDate returns the time corresponding to adding the given number of years,
// months and days to t. For example, AddDate(-1, 2, 3) applied to January 1,
// 2011 returns March 4, 2010.
//
// AddDate normalizes its result in the same way that the Date function does.
func (t Time) AddDate(years, months, days int) Time

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool

// Clock returns the hour, minute and second within the day specified by t.
// hour is in the range [0, 23], minute and second are in the range [0, 59].
func (t Time) Clock() (hour, minute, second int)

// Date returns the year, month and day in which t occurs. month is in the
// range [1, 12] and day is in the range [1, 31].
func (t Time) Date() (year, month, day int)

// Day returns the day of the month specified by t, in the range [1, 31].
func (t Time) Day() int

// Equal reports whether t and u represent the same time instant.
func (t Time) Equal(u Time) bool

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time
//
//	Mon Jan 2 15:04:05 -0700 MST 2006
//
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// A fractional second is represented by adding a period and zeros
// to the end of the seconds section of layout string, as in "15:04:05.000"
// to format a time stamp with millisecond precision.
func (t Time) Format(layout string) string

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool

// JS returns the time as a JavaScript date. The result is undefined if the
// year of t is not in the range [-999999, 999999].
func (t Time) JS() js

// JSON returns a time in a format suitable for use in JSON.
func (t Time) JSON() json

// Month returns the month of the year specified by t, in the range [1, 12]
// where 1 is January and 12 is December.
func (t Time) Month() int

// Minute returns the minute offset within the hour specified by t, in the
// range [0, 59].
func (t Time) Minute() int

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t Time) Nanosecond() int

// Round returns the result of rounding t to the nearest multiple of d (since
// the zero time). The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but
// otherwise unchanged.
//
// Round operates on the time as an absolute duration since the zero time; it
// does not operate on the presentation form of the time. Thus, Round(Hour)
// may return a time with a non-zero minute, depending on the time's Location.
func (t Time) Round(d Duration) Time

// Second returns the second offset within the minute specified by t, in the
// range [0, 59].
func (t Time) Second() int

// String returns the time formatted.
func (t Time) String() string

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t Time) Sub(u Time) Duration

// Truncate returns the result of rounding t down to a multiple of d (since
// the zero time). If d <= 0, Truncate returns t stripped of any monotonic
// clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the zero time;
// it does not operate on the presentation form of the time. Thus,
// Truncate(Hour) may return a time with a non-zero minute, depending on the
// time's Location.
func (t Time) Truncate(d Duration) Time

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time

// Unix returns t as a Unix time, the number of seconds elapsed since January
// 1, 1970 UTC. The result does not depend on the location associated with t.
// Unix-like operating systems often record time as a 32-bit count of seconds,
// but since the method here returns a 64-bit value it is valid for billions
// of years into the past or future.
func (t Time) Unix() int64

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// nanoseconds cannot be represented by an int64 (a date before the year 1678
// or after 2262). Note that this means the result of calling UnixNano on the
// zero Time is undefined. The result does not depend on the location
// associated with t.
func (t Time) UnixNano() int64

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() int

// Year returns the year in which t occurs.
func (t Time) Year() int

// YearDay returns the day of the year specified by t, in the range [1, 365]
// for non-leap years, and [1, 366] in leap years.
func (t Time) YearDay() int
```

A _Time_ represents an instant in time.

### date

```go
func date(year, month, day, hour, min, sec, nsec int, location string) (Time, error)
```

Returns the time corresponding to the given date with time zone determined by location. If location does not exist,
it returns an error.

For example, the following call returns March 27, 2021 11:21:14.964553705 CET.

    date(2021, 3, 27, 11, 21, 14, 964553705, "Europe/Rome")

For UTC use "" or "UTC" as location. For the system's local time zone use "Local" as location.

The month, day, hour, min, sec and nsec values may be outside their usual ranges and will be normalized
during the conversion. For example, October 32 converts to November 1.

### now

```go
func now() Time
```

Returns the current local time.

### parseDuration

```go
func parseDuration(s string) (Duration, error)
```

Parses a duration string.

A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such
as "300ms", "-1.5h" or "2h45m".

Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

### parseTime

```go
func parseTime(layout, value string) (Time, error)
```

Parses a formatted string and returns the time value it represents. The layout defines the format by showing how the
reference time

    Mon Jan 2 15:04:05 -0700 MST 2006

would be interpreted if it were the value; it serves as an example of the input format. The same interpretation will
then be made to the input string.

For more details, See the [time.Parse](https://pkg.go.dev/time#Parse) function of the Go standard library. 

As a special case, if layout is an empty string, _parseTime_ parses a time representation using a predefined list of
layouts.

It returns an error if value cannot be parsed.

### unixTime

```go
func unixTime(sec int64, nsec int64) Time
```

Returns the local Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970
UTC. It is valid to pass nsec outside the range [0, 999999999]. Not all sec values have a corresponding time value.
One such value is 1<<63-1 (the largest int64 value).

{% end raw doc %}
