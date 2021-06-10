# Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on generating a good map (local optimization) by making slight adjustments to the randomized map, both randomly and with beam search. The third step consists of generating a bunch of good candidate maps and keeping the best candidate (global optimization), followed by one last beam search on this best candidate.

The resulting map is presented to the console where tile types are:

```
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
# Custom Map

If you have not yet unlocked all the tiles of a given map, you can go into subfolder `maps/data` and edit the corresponding map accordingly before running the optimizer.

# TODO

- Improve performance
- Add more maps here
- Improve beam
- Automatically render image
- Add heat maps

# Best Combined Speed & Production Maps

## Tutorial Island

```
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
Score: 324.62

## Flesh World

```
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
Score: 607.87

## Planet Tronne

```
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
Score: 579.34

## Candy Land

```
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
Score: 484.03

## Mansions & Managers

```
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
Score: 498.46

# Best Speed Maps

## Tutorial Island

```
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
Score: 203.38

## Flesh World

```
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
Score: 365.96

## Planet Tronne

```
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
Score: 375.36

## Candy Land

```
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
Score: 321.90

## Mansions & Managers

```
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
Score: 330.28
