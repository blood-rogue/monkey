# THE MONKEY PROGRAMMING LANGUAGE

The [monkey programming language](https://monkeylang.org/) interpreter/compiler from the books [Writing an Interpreter in Go](https://interpreterbook.com/) and [Writing a Compiler in Go](https://compilerbook.com/) by Thorsten Ball.

**CHANGES**:
- added `exit()` function.
- added `keys()` and `values()` for hash values.
- added `globals()` and `locals()` functions to check environment (*`interpreter` only*).
- added `toInt()` and `toBool()` for type conversion.

**TODO**:
- implement `globals()` and `locals()` in compiler/vm.