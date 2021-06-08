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
        ..          
        vdv    w    
       d dv   ..*O< 
       vvvd> O*.*O< 
        dd  >O*.*O< 
   >>>O ..   O . O< 
   h>>..... lh-.*O<.
   >>>O...O <  w    
   rr>O...O<l  w    
   h>>O...O<<  w    
      ^u u^<>O..*O< 
             O*.  < 
             O*.*O< 
  .    w^w   O*.*O< 
  .h   .*.          
       .*.          
                    
```
Score: 313.81

## Flesh World

```
        v     vvvv..
.h.h ..OOO     dvd  
.*>>&**  *<  vvvvvv 
.  hh......h-vv dO  
^    vv***>>rO...O  
|    OOOOOrr>.....ll
>  *.**^^hrr>.....<l
h>-.....hhhh>... O<<
>rO*.**<    Ouu^uO< 
 >OO.OO      u^^u^u 
.>O w ^      ^w^^u  
.>O . ^      vwwu   
>>O . O      O..O<< 
>>O . O<    O*.*Ol  
>>O . O<hh>>....    
r>O*.*O<<r>> *.     
 >O....<hhh>..      
```
Score: 595.19

## Planet Tronne

```
d wdr  >....hhhh....
  dvv >...* <lv  ** 
vdvvvr> *..Ovvd  ..O
OO.  O>O&.^O vdvO^^*
O*.  l<l^^^O vdvOO..
.....<h h >O ...< <h
... .l   >>r.... <<h
  OOOOOhhr  ....lllh
^^^^^  h rrr....<<<h
^^^^ ..  >>>...  <<h
wuu^.*vvd > ^u^   l|
      OO. O^uu^    .
.* >>  *   l^     ..
.. hh ... .<huw .O* 
       *.*<   &.*O .
.. hr O..O  h *.. l<
^^ ..   ^^h >... O<<
```
Score: 567.64

## Candy Land

```
dd.*      .vv d& ..<
vvd..           .*Ol
dvdd  ...   dv.*.*O<
vv     *  O  dO*.*  
..  l   r   ..OO.< <
..  l<hh>> ...Ol<   
...  hh >> ...Ol< h.
.*.l  hhr> ...O<<lh.
^ ^vv  h>r ...   hh.
^  dvv   > ^^^ <   |
u^  vvvO         v w
>>r  .*  vv^u^ .vv v
rr>. ..Ol  w^      .
>>r...  l *.< r>&O*.
 >>O..  ....h h  ...
  OOu.OO*.*O< >. .*.
   ^u^^O..OO<    OOO
```
Score: 470.29

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
