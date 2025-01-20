Link parser. A package that accepts html (io.reader or string)

In this exercise your goal is create a package that makes it easy to parse an HTML file and extract all of the links (<a href="">...</a> tags). For each extracted link you should return a data structure that includes both the href, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

- use `net/html` package
- Ignore nested linked
