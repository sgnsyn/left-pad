# leftPad: Padding Without the Drama  

Need to pad strings to the left? Look no further—this Go package ensures your strings stay properly aligned without any unexpected surprises. Because, as we’ve all learned, padding is serious business.  

## Installation  

Just run:  

```bash  
go get github.com/sgnsyn/left-pad
```

## Usage

```go
  leftPad.Pad("foo", 5)             // Output: "  foo"
  leftPad.Pad("foobar", 6)          // Output: "foobar"  
  leftPad.Pad(1, 2, '0')            // Output: "01"  
  leftPad.Pad(17, 5, '0')           // Output: "00017"
  leftPad.Pad("hi", 5, '😊','a')    // Output: "😊😊😊hi"
```
# Notes

- If you supply more than one character for padding, we’ll quietly focus on the first one. Go is chill like that.
- Fancy characters outside the Basic Multilingual Plane (BMP) are treated as two. Because padding isn’t just for ASCII anymore.
