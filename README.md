# Go VM

An experimental proof-of-concept that illustrates the new instruction set and
interpreter architecture in Rubinius.

A few notes:

1. Everything is starting in a single file until I learn how to properly
   modularize things in Go.
1. One of the more interesting aspects of this is the difference in memory
   management in C/C++ and Go. For example, in Rubinius, we can cast a pointer
   to a function as an `intptr_t` to store in the opcode sequence. This is
   strictly forbidden in Go.
1. One of the challenges writing a just-in-time (JIT) compiler is the
   potential mismatch between a program's semantics as interpreted versus as
   JIT compiled. It would be easy to cast this a checking whether the JIT is
   "correct", but it's not that simple. A better way to think about it is
   whether the JIT is the "same as" the interpreter because bugs can become
   features. The best way to ensure the JIT is the same is to use the same
   code that the interpreter uses.
1. The point above has informed the design of the instruction set and
   interpreter, specifically the separate of the interpreter from the
   instruction set functions. The instruction mutates program state, while the
   interpreter mutates the PC/IP (program counter / instruction pointer). The
   (naive) JIT serializes (calls one after the other) the instructions in the
   opcode sequence.
