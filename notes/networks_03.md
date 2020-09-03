## Link Layer

## Retransmissions

### Automatic Repeat Request (ARQ)

Used when errors are common or must be corrected.

Rules:
- Receiver automatically acknowledges correct frames with an ACK
- Sender resends after a timeout, until an ACK is received

Sequence numbers: frames and ACKs must both carry sequence numbers for
correctness. To distinguish the current frame from the next one, a single bit
(two numbers) is sufficient. Called **Stop-and-Wait**

Limitation of Stop-and-Wait: can only have one frame outstanding at a time.

**Sliding Window**
- allow W frames to be outstanding
- can send W frames per RTT (= 2D)

## Multiplexing

Sharing a link among different users.

1. Time division: users take turns on a fixed schedule, sends high bursts a
   fraction of the time
2. Frequency division: users share channel on different frequency bands, sends
   low rate all the time

TDM/FDM used widely in telecommunications.

### Multiplexing Network Traffic

- Network traffic is bursty (on/off sources). Load varies greatly over time.
- Multiple access schemes multiplex users according to their demands!

## Randomized Multiple Access

Used in Classical Ethernet!

- ALOHA protocol
    - node just sends when it has traffic. If no ACK received (collision), wait
      a random time then resend
    - works well under low load, **not efficient** under high load

- CSMA (Carrier Sense Multiple Access)
    - improved ALOHA by listening for activity before we send
    - only good when BD (bandwidth-delay product) is small

- CSMA/CD (with Collision Detection)
    - reduce cost of collisions by detecting them and aborting "Jam" the rest
      of the frame time
    - impose a minimum frame size that lasts for 2D seconds, so nodes can't
      finish before collision
    - Ethernet minimum frame is 64 bytes
    - if there are `N` queued senders, we want each to send next with
      probability `1/N`

- Binary Exponential Backoff
    - estimates probability (doubles interval for each successive collision):
        - 1st collision: wait 0 or 1 frame times
        - 2nd collision: wait from 0 to 3 times
        - 3rd collsion: wait from 0 to 7 times

## Wireless Multiple Access

More complicated than wired case:
1. nodes may have different areas of coverage (cannot do Carrier Sense)
2. nodes can't hear while sending (cannot do Collision Detection)

- Hidden terminals: nodes `A` and `C` are **hidden terminals** when sending to `B`
    - They cannot hear each other (to coordinate), yet collide at `B`

    ```
    ---     ---     ---
    |A| --> |B| <-- |C|
    ---     ---     ---
    ```

- Exposed terminals: `B` and `C` are **exposed terminals** when sending to `A` and `D`
    - Can hear each other yet don't collide at receivers
    - Send concurrently to increase performance

    ```
    ---     ---         ---     ---
    |A| <-- |B| --> <-- |C| --> |D|
    ---     ---         ---     ---
    ```

### 802.11, WiFi

- Physical layer
    - Uses 20/40 MHz channels on ISM bands (b/g/n on 2.4 GHz, a/n on 5 GHz)
    - OFDM modulation (except b)

- Link layer
    - multiple access uses CMSA/CA; RTS/CTS optional
    - frames are ACKed and retransmitted with ARQ
    - uses 3 addresses (due to AP)

### 802.11 CSMA/CA for multiple access

- Sender avoids collisions by inserting small random gaps
