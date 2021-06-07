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

# Beacons

Beacons of a given type have their gains added together. Beacons of different type are not added together.

- `SpeedGain = BaseSpeed * (100% + SumOfAllSpeedBeacons)`
- `ProductionGain = BaseProduction * (100% + SumOfAllProductionBeacons)`

# Tips

- Labs are only impacted by speed gain, you will get lower gains from using combined speed and production maps.

# TODO

- Improve performance
- Add more maps here
- Add donut

# Best Combined Speed & Production Maps

## Tutorial Island

## Flesh World

## Planet Tronne

## Candy Land

```
....xxxxxx...xh.x...
.**..xxxxxxxxxxx*.*.
....xx...xxxh.......
.*xxxxx.xx.xx.*.*.xx
..xx.xxx.xxx......xw
..xx.h....x..*.*&xxx
.*.xx..x*.xh.....x..
....xxh...x..h^...*.
wx.*.xx...x..hxxx...
wxx...xxx.x.w.x.xxx.
..xx.*.&xxxxxxxxx.x.
.*.xx..xx..h..x...xw
....xh...xx..xxxxxx.
..*..<xx.x.*.x.<..*.
x.....xx.....xhxx...
xx..*.h..*.*.x..x.*.
xxx...h......xxxx...
```
Score: 52706.34

## Mansions & Managers

```
.....hxxx....xxx....
.*.*x.x..*..xx.xx.*.
......xh...xx...xx..
.*x.*.xxx.xx..*..xx.
......xxxxx.>.....xx
xxxxxxxxxxxxxxxxxxxx
.....hx.....hx..xx..
.*.*..x.*.*..xxx..*.
.*.x.hx......hhx....
.....xxxh.*....xxx..
xxxxxx.xxx.xxx...x.w
w.w.wxxx.xxx&xxxxxxx
.....hx....xxx......
.*.*x.xxxx..x..*.x*.
.x....h..x..xh......
.*.*..xxxx*.x.x.*.*.
.....hxh....x.h.....
```
Score: 56396.04

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
