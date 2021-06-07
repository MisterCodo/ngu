# Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on making slight adjustments to the randomized map, slowly making it better, this provides a good map (local optimization). The third step consists of generating a bunch of good maps and keeping the best one (global optimization).

The resulting map is presented to the console where tile types are:

```
x = unusable space
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

# Tips

- Labs are only impacted by speed gain, you will get max output by using speed maps.

# TODO

- Improve performance
- Add more maps here
- Add option to select which beacons to include
- Add extra optimization step (beam?)

# Best Combined Speed & Production Maps

## Tutorial Island

```
xxxxxxxx..xxxxxxxxxx
xxxxxxxxvdvxxxxwxxxx
xxxxxxxdxdvxxx..*O<x
xxxxxxxvvvd>xO*.*O<x
xxxxxxxxddxx>O*.*O<x
xxx>>>Ox..xxxOx.xO<x
xxxh>>.....xlh-.*O<.
xxx>>>O...Ox<xxwxxxx
xxxrr>O...O<lxxwxxxx
xxxh>>O...O<<xxwxxxx
xxxxxx^uxu^<>O..*O<x
xxxxxxxxxxxxxO*.xx<x
xxxxxxxxxxxxxO*.*O<x
xx.xxxxw^wxxxO*.*O<x
xx.hxxx.*.xxxxxxxxxx
xxxxxxx.*.xxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
```
Score: 313.81

## Flesh World

```
xxxxxxxxvxxxxxvvvv..
.h.hx..OOOxxxxxdvdxx
.*>>&**xx*<xxvvvvvvx
.xxhh......h-vvxdOxx
^xxxxvv***>>rO...Oxx
|xxxxOOOOOrr>.....ll
>xx*.**^^hrr>.....<l
h>-.....hhhh>...xO<<
>rO*.**<xxxxOuu^uO<x
x>OO.OOxxxxxxu^^u^ux
.>Oxwx^xxxxxx^w^^uxx
.>Ox.x^xxxxxxvwwuxxx
>>Ox.xOxxxxxxO..O<<x
>>Ox.xO<xxxxO*.*Olxx
>>Ox.xO<hh>>....xxxx
r>O*.*O<<r>>x*.xxxxx
x>O....<hhh>..xxxxxx
```
Score: 595.19

## Planet Tronne

```
dxwdrxx>....hhhh....
xxdvvx>...*x<lvxx**x
vdvvvr>x*..Ovvdxx..O
OO.xxO>O&.^OxvdvO^^*
O*.xxl<l^^^OxvdvOO..
.....<hxhx>Ox...<x<h
...x.lxxx>>r....x<<h
xxOOOOOhhrxx....lllh
^^^^^xxhxrrr....<<<h
^^^^x..xx>>>...xx<<h
wuu^.*vvdx>x^u^xxxl|
xxxxxxOO.xO^uu^xxxx.
.*x>>xx*xxxl^xxxxx..
..xhhx...x.<huwx.O*x
xxxxxxx*.*<xxx&.*Ox.
..xhrxO..Oxxhx*..xl<
^^x..xxx^^hx>...xO<<
```
Score: 567.64

## Candy Land

```
dd.*xxxxxx.vvxd&x..<
vvd..xxxxxxxxxxx.*Ol
dvddxx...xxxdv.*.*O<
vvxxxxx*xxOxxdO*.*xx
..xxlxxxrxxx..OO.<x<
..xxl<hh>>x...Ol<xxx
...xxhhx>>x...Ol<xh.
.*.lxxhhr>x...O<<lh.
^x^vvxxh>rx...xxxhh.
^xxdvvxxx>x^^^x<xxx|
u^xxvvvOxxxxxxxxxvxw
>>rxx.*xxvv^u^x.vvxv
rr>.x..Olxxw^xxxxxx.
>>r...xxlx*.<xr>&O*.
x>>O..xx....hxhxx...
xxOOu.OO*.*O<x>.x.*.
xxx^u^^O..OO<xxxxOOO
```
Score: 470.29

## Mansions & Managers

```
..OOOOxxx..OOxxxOO*.
****x*x&*.*&xx*xx.*.
-.....x--.hxx...xx.-
vvx***xxxwxxv***&xxw
OOOOOOxxxxxOOO..OOxx
xxxxxxxxxxxxxxxxxxxx
....hhxhh....xhhxx..
****l<x>*****xxx&***
...xhhxh.....hhxh...
^^^dvxxx.*^^^vdxxx**
xxxxxxOxxxOxxxvdOx..
uuwdvxxx^xxx^xxxxxxx
>>....x<hhhxxx...O<<
r>**x*xxxx>rx.*.*x<<
rx....OlhxhrxO...O<<
r>**.*xxxx>>xOx.*O<l
>>O...x<lhh>xO...O<<
```
score: 493.96

# Best Speed Maps

## Tutorial Island

```
xxxxxxxx..xxxxxxxxxx
xxxxxxxx.*.xxxxvxxxx
xxxxxxxvxv.xxx..vv<x
xxxxxxxvvv>>x....<<x
xxxxxxxxvvxx>.....<x
xxx>>..x..xxx>x.x*.x
xxx>>>.....x---....<
xxx>>>....<x<xx^xxxx
xxx>>>....<<<xx|xxxx
xxx>>>....<<<xxvxxxx
xxxxxx..x..--.....<x
xxxxxxxxxxxxx.*.xx<x
xxxxxxxxxxxxx.*.*.<x
xx.xxxx^^^xxx.....<x
xx..xxx.*.xxxxxxxxxx
xxxxxxx.*.xxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
```
Score: 202.88

## Flesh World

```
xxxxxxxxvxxxxxvvvvv.
...<xvvvvvxxxxxvvvxx
.*.<vvvxx..xxvvvvvvx
.xx>vvv..*.<<vvx..xx
.xxxx......--.....xx
|xxxx.....<-.......<
|xx>.....<--.......<
|>>.....<---....x..<
.>.*...^xxxx^^^^^^^x
x*.*.^^xxxxxx^^^^^^x
...x.x^xxxxxx^v^^^xx
.*|x^x|xxxxxxv...xxx
.*.x|xvxxxxxx..*.<<x
.*.xvx..xxxx*.*..<xx
.*.x.x.*-->.....xxxx
.*.*.*...<<*x*.xxxxx
x>.....---....xxxxxx
```
Score: 362.55

## Planet Tronne

```
>x...xx-.....---....
xx**.x>>...x<<<xxv*x
....->>x....<<<xx...
***xx>>>....x<vv*.*.
...xx--.....x-v.....
.*vvv<.x^x^^xv..*x*.
...xv<xxx^^>....x.<|
xx...<<-^^xx....<<<|
>....xx-x>>.....<<<.
>...x.<xx>.*...xx<*.
>.....<--x.x..^xxx..
xxxxxx..vx|k^^^xxxx.
v^x^^xx*xxxk^xxxxx*.
.^x^^x...x.--..x...x
xxxxxxx*.*.xxx**k*x|
..x->x....xx-x...x.<
..x.-xxx...x---.x...
```
Score: 372.64

## Candy Land

```
....xxxxxx..vxvvx.*.
.**.<xxxxxxxxxxx*.*.
...<xxvvvxxxvvv.*.*.
..xxxxx.xx-xx.....xx
.*xx<xxx.xxxv.*.*<x|
..xx>....-x.....<xxx
|.*xx.*x>>x....<<x<.
|^..xx..->x....<<.*.
|x.*vxx^>>x...xxx.*.
.xx..vxxx<x^.^xkxxx.
.*xx*...xxxxxxxxx.x|
...xx.*xxkvv|^x<.|x|
.**.x...<xx..xxxxxx|
......xx-x...x.<-...
x****.xx>..*.x<xx.*.
xx....---....x<-x.*.
xxx^^^>>.*..^xxxx.*.
```
Score: 318.64

## Mansions & Managers

```
.....-xxx...-xxx-...
****xkx.***.xxvxx.*.
....--x....xx-..xx..
.*x*.<xxxvxxv.*.*xx|
.....<xxxxx>.....<xx
xxxxxxxxxxxxxxxxxxxx
>....-x-.....x.-xx..
|^^*.>x>.*.*.xxx>.*.
vvvx.>x>.....<-x>...
vvvv|xxx.*.*.<vxxx*.
xxxxxx.xxx.xxxvvvx..
...v|xxx^xxx^xxxxxxx
.....<x.--^xxx...<<<
....x<xxxx>>x....x<<
>x....<<-x>>x....<<<
>^.*.*xxxx>>x.x..<<<
^^^...x..---x.....<<
```
Score: 329.48
