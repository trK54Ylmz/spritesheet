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
spritesheet --input doc --output output.png [-trim]
```

The input,

<p align="center">
<kbd><img src="/doc/1.png"></kbd>
<kbd><img src="/doc/2.png"></kbd>
<kbd><img src="/doc/3.png"></kbd>
</p>

The output,

<p align="center">
<kbd><img src="/doc/output.png"></kbd>
</p>

The output with trim enabled,

<p align="center">
<kbd><img src="/doc/output-trim.png"></kbd>
</p>
