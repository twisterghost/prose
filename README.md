# Prose

A CLI for managing prosefiles.

## Introduction

`prose` is a command line interface for updating and managing a "prosefile" - a
file which stores written word entries with a timestamp, ID and optional title
and author.

A prosefile looks like:

```json
{
  "filetype": "prosefile",
  "version": "0.0.1",
  "entries": [
    {
      "title": "My entry",
      "body": "Hello, prose! This is an entry in my prosefile.\n\nI will later publish this somewhere!",
      "id": "f0abfc1e-5fb8-40b7-b78b-0b756ba0d509",
      "author": "Michael",
      "time": "2020-05-12T15:45:19.621931532-04:00",
      "metadata": {}
    }
  ]
}
```

By default, `prose` assumes your prosefile exists at `~/prosefile.json`, but you
can configure it to work with a different location, or with multiple files by
using the `--file` flag, or configuring `file` in your configuration.

The `prose` configuration file lives at `~/.prose.yaml` and can also be
configured using the `--config` flag when using `prose`.

### Example Config

```yaml
author: Michael
file: /home/michael/snippets.json
pretty: true
```

## Sending

The main purpose of a prosefile is to store your written entries for other
programs to consume, modify and render. Consider a program called `prose-to-md`
which reads a prosefile from stdin and outputs each entry rendered as markdown,
with the entry title as the first level header:

```shell
$ prose send | prose-to-md
# My Entry

Hello, prose! This is an entry in my prosefile.

I will later publish this somewhere!
$ ...
```

`prose-to-md` is referred to as a "prose renderer", as it takes a prosefile
shaped input and creates a rendered output.

Now, consider a program called `prose-add-date` which reads a prosefile from
stdin, modifies each entry's `body` by prepending a human readable date, then
outputs back to stdout. We could then pipe that into `prose-to-md` for greater
effect:

```shell
$ prose send | prose-add-date | prose-to-md
# My Entry

May 12th, 2020

Hello, prose! This is an entry in my prosefile.

I will later publish this somewhere!
```

`prose-add-date` is referred to as a "prose modifier", as it takes a prosefile
shaped input, modifies it, then outputs a prosefile.

As you can imagine, you can continue to chain modifiers to greater benefit
before ultimately piping to a renderer. Perhaps you want to inject author
bylines before rendering to static html, or you want to modify all pronouns to
be ungendered before automating a publish to a web service:

```shell
$ prose send | prose-inject-byline | prose-to-html > my-website.html
$ prose send | prose-degender | prose-writeas-submit
```
