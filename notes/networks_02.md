## Link Model

- Rate (bandwidth, speed, capacity) in bits/second
- Delay in seconds

- Latency: delay to send a message over a link (`L = (M / R) + D`)
    1. transmission delay: time to put M-bit message on the wire (`M / R`)
    2. propagation delay: time for bits to propagate across the wire (`length / (2/3)c = D`)

### Examples
- Dialup
    ```
    D = 5 ms, R = 56 kbps, M = 1250 bytes
    L = 5 ms + (1250 * 8) / (56 * 10^3)
      = 184 ms
    ```

- Broadband (cross-country link)
    ```
    D = 50 ms, R = 10 Mbps, M = 1250 bytes
    L = 50 ms + (1250 bytes * 8) / (10 * 10^6)
      = 50 ms
    ```

## Signal Propagation

Bits are transmitted over a wire as analog signals. Over time it is represented
by its frequency components (called a Fourier analysis). The more bandwidth
we have (EE term, the more frequencies), the higher the fidelity of the signal,
and we are better able to see the voltage high/low.

Wireless signals attenuate greatly over distance (`1 / d^2`).

## Modulation

- Clock recovery 4B/5B - map every 4 data bits into 5 code bits without long runs
of zeroes.

- Passband modulation - the carrier frequency is changed via either amplitude,
frequency, or phase shift.

## Limits

> (how rapidly can we send information over a link?)

Channel propertiers:
1. bandwidth (B)
2. signal strength (S)
3. noise strength (N)

- Nyquist limit: maximum symbol rate is 2B
- Shannon capacity: number of signal levels we can distinguish is based on
    Signal-to-Noise ratio
