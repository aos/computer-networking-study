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

## Framing

1. Byte counts (difficult to synchronize after an error)
2. Byte stuffing - special flag byte that means start/end of frame (must be
   escaped)

## Error Correction

1. Error detection codes - check bits
2. Error correction codes

- Using codewords (systematic block codes): consist of `D` data plus `R` check
  bits. `R = fn(D)`


### Hamming Distance

Distance: number of bit flips needed to change D+R_1 to D+R_2.

Example:
```
1 -> 111, 0 -> 000

Distance = 3
```

- **Hamming distance** of a code: the minimum distance between any pair of
  codewords.

- Error detection: for a code of D+1, up to `d` errors will always be detected.
- Error correction: for a code of distance 2D+1, up to `d` errors can always be
  corrected by mapping to the closest codeword

## Error Detection

1. Parity bit: take D data bits, add 1 check bit that is the sum of the D bits,
   modulo 2. (XOR)
    - Example: `1001100 -> 3 mod 2 = 1 -> 10011001`
2. Internet checksum: The 16bit one's complement of the one's complement sum of
   all 16 bit words.
    - Sum up all words, 16 bits a time, then negate it
    - Finds all burst errors up to 16
3. Cyclic Redundancy Check: generate `k` check bits such that the `n+k` bits
   are evenly divisible by a generator `C`

## Error Correction

aaa
