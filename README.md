# usgs-iv-details

You can download the most recent version from the [releases page](https://github.com/andrew-fcx/usgs-iv-details/releases).

To use the downloaded script, run the executable and pass the site code you want to inspect the instantaneous value data for.

For example, if you want to check on the Delaware River at Trenton NJ site (site code: 01463500), you'd run the following:

```cmd
usgs-iv-details.exe -site=01463500
```

This will give a response similar to this:

```
Getting timeseries data order from endpoint: https://waterservices.usgs.gov/nwis/iv/?format=json&site=01463500

Delaware River at Trenton NJ (01463500)
========================================
0 | 00010 | Temperature, water, &#176;C
1 | 00060 | Streamflow, ft&#179;/s
2 | 00065 | Gage height, ft
3 | 00095 | Specific conductance, water, unfiltered, microsiemens per centimeter at 25&#176;C
4 | 00300 | Dissolved oxygen, water, unfiltered, mg/L
5 | 00301 | Dissolved oxygen, water, unfiltered, % saturation
6 | 00400 | pH, water, unfiltered, field, standard units
7 | 32318 | Chlorophylls, water, in situ, fluorometric method, excitation at 470 &#177;15 nm, emission at 685 &#177;20 nm, micrograms per liter
8 | 63680 | Turbidity, water, unfiltered, monochrome near infra-red LED light, 780-900 nm, detection angle 90 &#177;2.5&#176;, formazin nephelometric units (FNU)
9 | 99133 | Nitrate plus nitrite, water, in situ, mg/L as N
```

The first column gives the index in the array from the response, the 2nd column shows the parameter codes, and the last column shows the name of the variable.
