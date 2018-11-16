# bfdump

bfdump is a test utility for debugging the 
[blackfriday.v2](https://github.com/russross/blackfriday/tree/v2) markdown
parser tree output format. This is just a tool I made to help me create a
markdown to markdown formatting renderer based off of *blackfriday.v2* 
library.

See also, the project I wrote this for: <https://github.com/bobertlo/vmd>

And, finally, for a graphical demo of what this tool is, see the output
of this file, `README.md`:

```
Document
  Heading Level 1
    Text "bfdump"
  Paragraph
    Text "bfdump is a test utility for debugging the"
    Text " (blackfriday.v2)[https://github.com/russross/blackfriday/tree/v2] markdown parser tree output format. This is just a tool I made to help me create a markdown to markdown formatting renderer based off of "
    Emph
      Text "blackfriday.v2"
    Text ""
    Text " library."
  Paragraph
    Text "See also, the project I wrote this for: "
    Link dst:https://github.com/bobertlo/vmd
      Text "https://github.com/bobertlo/vmd"
  Paragraph
    Text "And, finally, for a graphical demo of what this tool is, see the output of this file, "
    Code: `README.md`
    Text ":"
```
