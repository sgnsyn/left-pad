package leftPad

import "strings"

func Pad(s string, length int, chars ...rune ) string {
    padAmount := len(s) - length //getting the required padding amount 
    if length <= 0 || padAmount <= 0{ 
        return s // return original string if no padding is required
    }

    charToPad := '0' //initialize char to pad to 0
    if len(chars)  > 0{
        charToPad = chars[0] //use the first char value if third argument "char ...rune" is provided
    }

    var sb strings.Builder
    for i := 0; i < padAmount; i++{
        sb.WriteRune(charToPad )  //write char to pad required times to string builder
    }
    sb.WriteString(s) //write original string to string builder

    return sb.String()  //return the string builder as a string
}
