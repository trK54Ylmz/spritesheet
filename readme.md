### Spritesheet

The spritesheet is a small utility to generate sprite sheets to animate, tile, etc. from group of PNG files for games engines such as Unity, Godot, etc.

#### Installation

The utility is plain simple and does not contain any dependencies.

```bash
go install github.com/trk54ylmz/spritesheet
```

#### Example

Here is an example usage,

```bash
spritesheet --input doc --output output.png
```

The input,

<p align="center">
<img src="/doc/1.png">
<img src="/doc/2.png">
<img src="/doc/3.png">
</p>

The output,

<p align="center">
<img src="/doc/output.png">
</p>
