# Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on making slight adjustments to the randomized map, slowly making it better, this provides a good map (local optimization). The third step consists of generating a bunch of good maps and keeping the best one (global optimization).

The resulting map is presented to the console where tile types are:


```
. = resource
* = box
k = knight
^v<> = up, down, left, right arrows
-| = horizontal, vertical wall
o = donut
```

# Speed Beacons

Speed beacons effect are added together, even if speed beacon types differ. This means that ff base speed is 100%, then impact of one box speed beacon (40%) results in a speed of 140% while the impact of two box beacons results in a speed of 180%. Then formula is:

`Speed = BaseSpeed * (100% + SumOfAllSpeedBeacons)`

# Production Beacons

TBD

# Efficiency Beacons

TBD

# Notes

- One beacon type does not impact a beacon of a different type.
- Labs are only impacted by speed beacons.

# TODO

- Trace to see bottlenecks and improve performance
- Merge best in here
- Production only
- Production and speed combined
- Add donut
