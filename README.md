# Beta Code and UTF-8 Transcoder
[![Go Reference](https://pkg.go.dev/badge/github.com/jllovet/betacode-utf8-transcoder.svg)](https://pkg.go.dev/github.com/jllovet/betacode-utf8-transcoder)

This transcoder allows you to convert texts that are in Beta Code to UTF-8 and back. Its primary purpose is to support working with Ancient Greek datasets that used Beta Code as their encoding.

This project owes a great deal to [github.com/matgrioni/betacode](https://github.com/matgrioni/betacode), which is an analogous transcoder written in python. Portions of this transcoder are translations of that project into Go, and the goals of our projects are very nearly the same. Thanks, [Matias](https://github.com/matgrioni)!

## Installation

```shell
go get github.com/jllovet/betacode-utf8-transcoder
```

## Usage

The most frequent uses you'll have for this package are using `BetaToUni` and `UniToBeta`.

### Example Usage of BetaToUni

```golang
// BetaToUni(beta string) (uni string, err error)
package main

import (
    "fmt"
    "log"

    "github.com/jllovet/betacode-utf8-transcoder"
)

func main() {
    b := `a)/lfa`
    u, err := transcoder.BetaToUni(b)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(b, "becomes", u)
}
```

```shell
> go run main.go
> a)/lfa becomes ἄλφα
```

### Example Usage of UniToBeta

```golang
// UniToBeta(uni string) (beta string, err error)
package main

import (
    "fmt"
    "log"

    "github.com/jllovet/betacode-utf8-transcoder"
)

func main() {
    u := `ἄλφα`
    b, err := transcoder.UniToBeta(u)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(u, "becomes", b)
}
```

```shell
> go run main.go
> ἄλφα becomes a)/lfa
```



## Resources

### Encodings and Transcording Background and Specs

- [Joel on Software - The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [What is Transcoding?](https://en.wikipedia.org/wiki/Transcoding)
- [Wikipedia - Beta_Code](https://en.wikipedia.org/wiki/Beta_Code)
- [TLGU - BCM.pdf](http://www.tlg.uci.edu/encoding/BCM.pdf)
- [TLGU - quickbeta.pdf](http://www.tlg.uci.edu/encoding/quickbeta.pdf)

### Unicode Resources Online
- [Fileformat.info](https://www.fileformat.info/info/unicode/)
- [Unicode Converter](https://www.branah.com/unicode-converter)
### Unicode Normalization
- [YouTube - Unicode Normalization for NLP in Python - James Briggs](https://www.youtube.com/watch?v=9Od9-DV9kd8)
- [Medium - What on Earth is Unicode Normalization? - James Briggs](https://towardsdatascience.com/what-on-earth-is-unicode-normalization-56c005c55ad0)
- [YouTube - Practical Serialization In Go: Unicode Normalization - Ardan Labs](https://www.youtube.com/watch?v=kOFIToH9KSY)

### Unicode Normalization and Using Unicode in Go
- [The Go Blog - Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
    - Go source code is always UTF-8.
    - A string holds arbitrary bytes.
    - A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
    - Those sequences represent Unicode code points, called runes.
    - No guarantee is made in Go that characters in strings are normalized.

- [The Go Blog - Text normalization in Go](https://blog.golang.org/normalization)
- [Unicode Programming - Unicode Normalization in Go](https://unicode-programming.readthedocs.io/en/latest/normalization/go/)

#### Trie, Trie, Trie Again

##### Videos

- [YouTube - Data Structures in Golang - The trie data structure - Junmin Lee](https://www.youtube.com/watch?v=nL7BHR5vJDc)
- [YouTube - HackerRank - Data Structures: Tries](https://www.youtube.com/watch?v=zIjfhVPRZCg)
- [YouTube - Jacob Sorber - The Trie Data Structure (Prefix Tree)](https://www.youtube.com/watch?v=3CbFFVHQrk4)
- [YouTube - Implement Trie | Leetcode #208 - Techn Dose](https://www.youtube.com/watch?v=xqsaAhQC6c8)
- [YouTube - Tech Dose - Trie Playlist](https://www.youtube.com/watch?v=6PX6wqDQE20&list=PLEJXowNB4kPyi859E6qGUs7jlpQehJndl)

##### Articles
- [Wikipedia - Trie](https://en.wikipedia.org/wiki/Trie)
- [Medium - Vaidehi Joshi - Trying to Understand Tries](https://medium.com/basecs/trying-to-understand-tries-3ec6bede0014)
- [Stack Overflow - Ukkonen's suffix tree algorithm in plain English](https://stackoverflow.com/questions/9452701/ukkonens-suffix-tree-algorithm-in-plain-english)
- [Suffix Trees and Their Applications - Bálint Márk Vásárhelyi](http://berkri.web.elte.hu//Theses/Vasarhelyi_2.pdf)
- [Wikipedia - Longest Common Prefix Array](https://en.wikipedia.org/wiki/LCP_array)
- [Suffix Arrays - A Programming Contest Approach - Adrian Vladu and Cosmin Negruşeri](https://web.stanford.edu/class/cs97si/suffix-array.pdf)
- [Simple Linear Work Suffix Array Construction - Juha Kärkkäinen and Peter Sanders](https://www.cs.helsinki.fi/u/tpkarkka/publications/icalp03.pdf)

### Other Beta Code to UTF-8 Transcoders
- [github.com/matgrioni/betacode](https://github.com/matgrioni/betacode)
- [github.com/cltk/grc_software_tlgu](https://github.com/cltk/grc_software_tlgu)
- [TypeGreek Alphabet Key](http://www.typegreek.com/alphabet.key/)
- [Translatum Online Converter](https://www.translatum.gr/converter/beta-code.htm)
- [unibetacode](http://unifoundry.com/unibetacode/)
