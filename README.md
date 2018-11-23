# bfdump

bfdump is a test utility for debugging the
[blackfriday.v2](https://github.com/russross/blackfriday/tree/v2) markdown
parser tree output format. This is just a tool I made to help me create a
markdown to markdown formatting renderer based off of *blackfriday.v2*.

And, here is a graphical demo of what this tool is, see the output of this file,
`README.md`:

```
Document
  Heading Level 1
    Text "bfdump"
  Paragraph
    Text "bfdump is a test utility for debugging the"
    Text "\n"
    Link dst:https://github.com/russross/blackfriday/tree/v2
      Text "blackfriday.v2"
    Text " markdown\nparser tree output format. This is just a tool I made to help me create a\nmarkdown to markdown formatting renderer based off of "
    Emph
      Text "blackfriday.v2"
    Text "."
  Paragraph
    Text "And, here is a graphical demo of what this tool is, see the output\nof this file, "
    Code: `README.md`
    Text ":"
```
