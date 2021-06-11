# NGU Optimizer

## How to Use

Download the latest release and use command `ngu optimize -h` to see usage instructions:

```ascii
Optimize placement of beacons on NGU Industries map.

Usage:
  ngu optimize [flags]

Flags:
  -a, --adjust int    how many adjustments to perform on each random map (default 10000)
  -b, --beacons int   optimization beacon types available: (1)Box, (2)Box & Knight, (3)Box, Knight & Arrow, (4) Box, Knight, Arrow & Wall, (5)All (default 5)
  -c, --cycle int     how many global optimization cycles to run (default 100)
  -g, --goal int      optimization goal: (1)Speed&Production, (2)Speed, (3)Production (default 1)
  -h, --help          help for optimize
  -m, --map int       map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers (default 1)
  -r, --random int    how many random map to generate per cycle (default 100)
  -s, --spread int    optimization modifies up to X tiles at once during randomised hill climbing where X is the spread (default 3)
```

Important flags are `-m` for choosing the map, `-b` for specifying beacons you have unlocked in the game and `-g` for choosing the type of optimization (speed, production or combination of speed&production). For instance, to launch a speed optimization for Candy Land with box and knight beacons, use command `ngu optimize -m 4 -b 2 -g 2`. It will run for a few minutes before offering an optimized map in console and will also generate an output png image on disk.

If you want to make modifications to the code, the only dependency is [GoLang](https://golang.org/). Once GoLang is installed, clone the repo and from its root use command `go run ngu.go`.

## Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on generating a good map (local optimization) by making slight adjustments to the randomized map, both randomly and with beam search. The third step consists of generating a bunch of good candidate maps and keeping the best candidate (global optimization), followed by one last beam search on this best candidate.

The resulting map is presented to the console where tile types are:

```ascii
  = unusable (space character)
. = resource

* = box (speed)
k = knight (speed)
^v<> = up, down, left, right arrows (speed)
-| = horizontal, vertical wall (speed)
o = donut (speed)

b = box (production)
& = knight (production)
udlr = up, down, left, right arrows (production)
hw = horizontal, vertical wall (production)
O = donut (production)
```

## Custom Map

If you have not yet unlocked all the tiles of a given map, you can go into subfolder `plugins/locations/<mapname>/data` and edit the corresponding txt file before running the optimizer. This is currently only available by running the optimizer from code with Go.

## TODO / Upcoming Features

- Improve performance
- Improve beam
- Add heat maps of both speed and production gains
- Add run/improve forever option

## Best Combined Speed & Production Maps

```ascii
Tutorial Island (324.62)

        ww
        www    .
       v dv   *.*O<
       vvvvd O*.*O<
        dd  >O*.*O<
   hr>> ..   < . O<
   hr>r.... <<hw.O<.
   hr>r.... l  w
   h>>>...OO<  w
   .>>O^uu^O<  w
      ^^ ^u^>.*.*O<
             O*.  <
             O*.*O<
  .    www   O*.*O<
  .h   .*.
       .*.

```

```ascii
Flesh World (607.87)

        *     www...
.*.h dw...     vdd
.*.vvvv  *.  vvvvd.
.  vvvvvO..>Ovv vO
w    d.OO<>>OO..OO
r    ..*hhrrr....<l<
h  .....hhh>>....<<<
>>>..*.llh>rr... l<l
>>O*.OO<    OO.uOO<
 >O*.^^      ^^^^^<
.>O . w      ^u^^u
..O w w      wwww   
.*O w w      www..l
..O w vv    >.*.*<
^>> . ..Ohhhh...
^>r>.*.*<<<l **
 >h>.....lhh-.
```

```ascii
Planet Tronne (579.34)

w wvw  dddvvvdvdvdvv
  vvv vvvvv vvv  vv
vdvvvvv OOOOOOv  OOO
OOO  OO.*.*. O.OO.*.
.*.  hh..... .......
....... h h. .... *.
.*. .*   OOOOOOO OO.
  OOOOO^^^  ^^^^u^^^
^^u^^  ^ ^^^^^^^^u^w
^^u^ ^.  w.w.^w  ^^w
wwwuwuvvv v vvw   ww
      OOO OkOOw    |
vv >.  *   *.     v.
.. hh ... ...h. h-.
       ***&   .*OO .
.. hh O..O  O .*o *.
.. .<   ^^^ ^^.. ...
```

```ascii
Candy Land (484.03)

.OO.      .ww >O .b.
w***.           *.*.
....  -h.   dvv...*.
vv     <  O  vdO*.
OO  O   r   .dOO<w |
.*  .<hh>> ...O<<
...  hh >r ....<l hw
.*.l  hh>> ....<l<hw
. Ovd  vr> O..   <vw
^  vvv   > ^^^ O   .
^^  ddvO         v .
w>r  ..  <lww^ >>O .
>>>. ...l  hw      .
rr>...  < -.. hh....
 >>O..  <<O*. O  **w
  OOuu^^>>O*. O< OO|
   ^uuu^>>...    ^^^
```

```ascii
Mansions & Managers (498.46)

>OO..O   h..O   OO..
>O*. O <<*.*  *  *.*
>.*..O <-h.  ...  .-
>O .*O   d  v***&  w
>OO..O     OOO..OO

..^wwh h..... hh  ..
*wdww> >*****   &***
vdd vd r.....hh h...
dvvdv   ^u^^^vv   **
      O   O   vwO ..
r.*..   <   ^
>..... <hhh   ...O<l
>.*. .    r> .*.* <<
r OO.OO<l hr O...O<l
^^^^^^    r> O .*O<<
^uuu^^ ..-h> O...O<<
```

## Best Speed Maps

```ascii
Tutorial Island (203.38)

        ..
        .*.    v
       v v.   ..vv<
       vvv>> ....<<
        vv  >.....<
   >>.. ..   > . *.
   >>>..... ---....<
   >>>....< <  ^
   >>>....<<<  |
   >>>....<<<  |
      .. .---.....<
             .*.  < 
             .*.*.<
  .    ^^|   .....<
  ..   .*.
       .*.

```

```ascii
Flesh World (365.96)

        v     vvv...
...< vvvv.     vvv
.*.<<vv  *.  vvvvvk
.  >vv.....--.. ..
.    .....<--.....
.    ....<>>>.....<<
|  >....<<>>>.....<<
|>>....^<<>>>... .<<
|>.*...^    >.....<
 *.*.^|      ||^^^.
.*. . |      ||^^^
.*. ^ |      |.^^
.*. o .      ..*.<<
.*. o .*    o....<
.*. o .*...oo..*
.*.*.*.*.*.* ..
 ......--.....
```

```ascii
Planet Tronne (375.36)

> ...  --.....<-vvvv
  **. >>... .<<  .v
....--> ....<<<  *.v
.**  >>>.... --.....
...  -->.... vv>....
|vvv..o ^ ^^ v.oo ..
.oo .*   ^^ov..* o..
  o.....--  .....<^.
.....  < >>.....<<|.
.*.. oo  >.....  <||
.....---. . ..^   ||
      ... ooo^|    |
|^ ^^  *   <^     ..
|^ ^> ... ...-. -..
       *.*.   ..oo .
.. -> ...*  o .*o ..
.. .-   ... --.. ...
```

```ascii
Candy Land (321.90)

....      ..- vv ...
.**.<           v**.
...<  vvv   >......<
..     v  >  .*..*
.*  <   v   >..... <
..  -..... ->.*..
|^|  >. .. <-.... ..
||.>  .... <<|^^^.*.
| vv>  ... <<|   .*.
.  vvv   . vv. ^   .
.*  vv^^         | |
...  .<  |vv.. <<| |
.*.. ..-<  ..      |
>.....  - ... .<-...
 >*.*.  >.*.. <  .*.
  ....---.... << .*.
   ^^^>>>...^    .*.
```

```ascii
Mansions & Managers (330.28)

>....-   ....   -...
>.*. o .*.**  .  .*.
>.*..o ..-.  .*.  ..
>. .*.   v  >...*  |
.oo...     vvo*..o

.....- -..... -.  ..
||^|^> >...*.   .**.
vvv |> >.....<- ....
vvv|v   .....<v   *.
      .   ^   vv| ..
....v   ^   ^
.....< <^^^   ...<<<
.... <    >> .... <<
> .....<- -> ....<<<
^^^.**    >> . ..<<<
^^^^.. ..--- .....<<
```

## Best Production Maps

```ascii
Tutorial Island (191.50)
        &d
        OO.    .
       . ..   .....
       ..... hh....
        OO  OOOOOOO
   .... OO   . . O.
   .h...... h.......
   OOO..O.. O  w
   rOO.OOO.OO  w
   .....Oh...  .
      .. ...hh.....
             O..  l
             O..OOl
  .    ...   h..u..
  ..   .b.
       O..

```

```ascii
Flesh World (360.53)

        d     .....h
...h ..O..     O&.
.b.hO..  ..  ..OOw.
.  OO..OO..OO.. O.
.    ..OO..OO..OO.
.    ..OO..OO..OO..O
w  OO..OO..OO..OO..O
w..OO..OO..OO..O ..O
w..OO..O    O..OO..
 ..OO..      .wOO..
.b. O w      w.OO.
... O w      ..OO
.b. . .      ..OO..
.OO O ..    O..OO.
.OO O .OO..OO..O
.....h.hh... ..
 ....h.hh.....
```

```ascii
Planet Tronne (376.35)

. O..  .....hh.h....
  O.. ...h. h..  ..
.OO..OO OO..OO.  Oww
.OO  OO.OO.. Ow.OO..
...  h...O.. O..OO..
.....h. . .. O..O ..
wOO OO   Ow.OO.. O..
  O.OO..OO  OO..OO..
.....  . h......hh..
.b.. O.  ......  .ww
.OO.OO.wO . OOw   ww
      ..O OOOOw    w
.. ..  .   ..     ..
wu wr O.. ..... O&.
       .bOO   ..OO .
.. h. h...  O .b. ..
.. .h   ... hh.. ...
```

```ascii
Candy Land (320.27)

....      ... .. ...
..b.l           OO..
..OO  ddd   OOOOOOO.
..     O  .  .....
..  .   .   ...... .
w.  ...... O..OOO
wwO  .O .. O..OOO Ow
wwOO  OO.. O......h.
. OO.  Ow. Ow.   ...
.  ...   . OwO O   w
..  ....         w w
..O  wO  w.... ... .
..OO .OO.  ..      .
..OO..  . OO. OO&O&w
 .OO..  ..OO. h  ...
  OO..OO..O.. h. .b.
   ...hh.....    O..
```

```ascii
Mansions & Managers (328.04)

....h.   ..h.   ....
...O . O..OO  O  .b.
OOOOO. O..O  .OO  Ow
.. O..   .  ..b..  .
......     .......  
                     
wwOOOw OwwOOO OO  O.
.....h h.....   ..b.
... .. h...h... ....
wOOOw   ..OO.wO   dw
      O   O   OOO Ow
.....   .   .       
..h... ...h   ......
.OOO O    .O .h.. ..
. OO.OOOO OO .OOOOO.
..O...    .O . O.Ob.
.....h ..... hh.....
```
