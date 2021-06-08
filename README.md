# Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on making slight adjustments to the randomized map, slowly making it better, this provides a good map (local optimization). The third step consists of generating a bunch of good maps and keeping the best one (global optimization).

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

# TODO

- Improve performance
- Add more maps here
- Add option to select which beacons to include
- Cleanup code, it's getting messy
- Improve beam
- Automatically render image
- Add instructions on how to modify base maps to your currently unlocked tiles in the map

# Best Combined Speed & Production Maps

## Tutorial Island

```
        ..          
        vdd    w    
       v vv   ..*O< 
       vvvv> O*.*O< 
        dd  >O*.*O< 
   >r>O ..   O . O< 
   h>>..... lhh.*.<<
   >>>O...O l  w    
   rr>O*..O<l  w    
   hr>....O<<  w    
      uu ^u<>O..*O< 
             O*.  < 
             O*.*O< 
  .    ^^w   O*.*O< 
  .h   .*.          
       .*.          
```
Score: 314.13

## Flesh World

```
        v     dvvd..
h>r> ..OOl     dvv
.rrrO*.  <<  vdddvv
h  >.....<<hOvv vO
r    *.*.hh>>....O
>    OOOOh>r>.....ll
>  ..*^^hh>>>.....<l
>>O....<hhhrr... Oll
>>O*.*O<    O^u^^O<
 >O..OO      w^^u^u
>>O . O      w^^^u
>>O . O      w^^u
>>O . O      ..O<lh
>>O . O<    *.*O<<
>>O . O<l>>O*..O
r>O*.*O<l>>O .*     
 >O....<hh-O..
```
Score: 599.64

## Planet Tronne

```
O .*.  vvvvv..vvvvvd
  .*. OOvvO .*O  vv
O*.*>>O ..OOl<O  OOO
O*.  >>....h hh...*.
..w  >>....h hh.....
vdwvr>O . OO vvO* *.
Ovw vv   uOOOOOO OO.
  vvvvv^^u  ..*^^^^^
O..OO  h >>.....l<hh
.*.* <<  >rO*.*  <lw
h......hh h ...   h|
      *.l r^^^^    w
u^ OO  .   uu     v.
^u ^^ ^^v v^^uu vvv
       OOO.   ..OO .
.* .l >.**  & *.* *.
.. h-   ... hhh. ...
```
Score: 573.11

## Candy Land

```
r...      .dw rO .*.
dw*.*           *.*O
ddv.  .*.   vdd.*.*O
vv     .  O  dvO*.
vv  <   >   .vOO<w .
..  <<hh>> ...O<<
...  <h rr ....<< h.
...l  hh>> ....<l<h.
. OOv  hr> O..   <hw
^  vdd   > u^^ O   |
u^  ddvO         d v
^ur  .*  lduw^ >OO .
r>>. ..Ol  vw      .
hr>...  < h.. hh....
 >>O*.  l<*** k  ***
  >O..hhh.... hh ...
   ^^^^O<***.    ^^u
```
Score: 477.91

## Mansions & Managers

```
..OOOO   ..OO   OO..
h..... --h.  ...  ..
vv ***   v  &***&  w
OOOOOO     OOO..OO  
                    
....hh h..... hh  ..
***ww> r*****   ****
.dd vv r....hhh ....
vvvdv   ^u^^^vv   **
      O   ^   vdO OO
r.*..   l   O       
>..... <hhh   ....<<
>.*. .    >r ...* <<
. OOOOOO< rr O..OO<l
^u^^^^    >> O .*Ol<
^uuu^. ..--h .....<<
```
Score: 495.32

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
   >>>....<<<  v    
      .. ..--.....< 
             .*.  < 
             .*.*.< 
  .    ^^^   .....< 
  ..   .*.          
       .*.          
                    
```
Score: 202.88

## Flesh World

```
        v     vvvvv.
...< vvvvv     vvv  
.*.<vvv  ..  vvvvvv 
.  >vvv..*.<<vv ..  
.    ......--.....  
|    .....<-.......<
|  >.....<--.......<
|>>.....<---.... ..<
.>.*...^    ^^^^^^^ 
 *.*.^^      ^^^^^^ 
... . ^      ^v^^^  
.*| ^ |      v...   
.*. | v      ..*.<< 
.*. v ..    *.*..<  
.*. . .*-->.....    
.*.*.*...<<* *.     
 >.....---....      
```
Score: 362.55

## Planet Tronne

```
> ...  -.....---....
  **. >>... <<<  v* 
....->> ....<<<  ...
***  >>>.... <vv*.*.
...  --..... -v.....
.*vvv<. ^ ^^ v..* *.
... v<   ^^>.... .<|
  ...<<-^^  ....<<<|
>....  - >>.....<<<.
>... .<  >.*...  <*.
>.....<-- . ..^   ..
      ..v |k^^^    .
v^ ^^  *   k^     *.
.^ ^^ ... .--.. ... 
       *.*.   **k* |
.. -> ....  - ... .<
.. .-   ... ---. ...
```
Score: 372.64

## Candy Land

```
....      ..v vv .*.
.**.<           *.*.
...<  vvv   vvv.*.*.
..     .  -  .....  
.*  <   .   v.*.*< |
..  >....- .....<   
|.*  .* >> ....<< <.
|^..  ..-> ....<<.*.
| .*v  ^>> ...   .*.
.  ..v   < ^.^ k   .
.*  *...         . |
...  .*  kvv|^ <.| |
.**. ...<  ..      |
......  - ... .<-...
 ****.  >..*. <  .*.
  ....---.... <- .*.
   ^^^>>.*..^    .*.
```
Score: 318.64

## Mansions & Managers

```
.....-   ...-   -...
**** k .***.  v  .*.
....-- ....  -..  ..
.* *.<   v  v.*.*  |
.....<     >.....<  
                    
>....- -..... .-  ..
|^^*.> >.*.*.   >.*.
vvv .> >.....<- >...
vvvv|   .*.*.<v   *.
      .   .   vvv ..
...v|   ^   ^       
.....< .--^   ...<<<
.... <    >> .... <<
> ....<<- >> ....<<<
>^.*.*    >> . ..<<<
^^^... ..--- .....<<
```
Score: 329.48
