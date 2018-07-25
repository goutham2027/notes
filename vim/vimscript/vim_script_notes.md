http://learnvimscriptthehardway.stevelosh.com/

## Feb 18, 2018 - Sunday
Book divided into 3 sections
1) Vim commands to use in `~/.vimrc` file to customize Vim.
2) Vimscript
3) Writing plugins


### Vim terminology
buffer
window
normal mode
insert mode
text object

### Vimrc file
- vim config
- contains Vimscript code.
- runs this everytime we open vim
- to know the vimrc path: `echo $MYVIMRC` in vim

`echo` and `echom` commands
- documenation `:help echo`
- persistent echoing using `echom` command.
- to see list of messages, `:messages`
- `echom` is used in debugging when writing complex Vim scripts.

- Comments using `"` eg: " this is a comment

### Setting Options
There are two main kinds of options: boolean and options that take a
value.

Boolean options either have on or off. All boolean operations work the
following way:
- `:set <name>` will turns the option
- `:set no<name>` will turn off the option.
- `:set <name>!` to toggle between the boolean options.
- `:set <name>?` to check what option is currently set.

Options with values
`set numberwidth=10`

Setting multiple options at once
`:set number numberwidth=10`

### Basic Mapping
Ability to map keys. This feature allow us to customize Vim, however we want.
"When I press this key, I want you to do this stuff isntead of whatever
you would normally do"
`:map - x`
`:map - dd`

Special Characters
we can use <keyname> to tell Vim about special keys.
`:map <space> viw`
`:map <c-d> dd` # ctrl+d

to delete the current line and paste it in the next line.
`:map - ddp`

to unmap `:unmap -`

### Modal Mapping

We can be more specific when we want mapping to apply by using `namp`,
`vmap` and `imap`.

`:imap <c-d> dd`
if we press `CTRL+d` in insert mode it will insert two `dd`s in the
current line. This is because in insert mode `dd` doesn't have a special
meaning.
To make this working as intended use `<esc>`.
`:imap <c-d> <esc>dd`
`<esc>` is telling Vim to press the Escape Key.
`:imap <c-d> <esc>ddi` will delete the current line and goes back to the
insert mode.

### Strict Mapping
Recursion is dangerous.
`:nmap dd O<esc>jddk` this will never exit, because there is a
recursion.

Downsides of `*map` commands:
- recursion
- their behavior can change if we install a plugin that maps keys they
  depend on.

Non recursive mapping
```
:nmap x dd
:nnoremap \ x
```
when we press `\`, Vim ignores the `x` mapping and it does what `x` will
originally do, i.e. it deletes a character instead of deleting a line.

Each of the `*map` command has a `*noremap` counterpart.

### Leaders
Whenever we map a key we are overriding it's default behavior.
http://learnvimscriptthehardway.stevelosh.com/chapters/06.html
