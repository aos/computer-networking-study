#include <stdio.h>
#include <stdint.h>

int main(void)
{
  uint16_t val = 0x0400;
  uint8_t *ptr = (uint8_t*)&val;

  if (ptr[0] == 0x04) {
    printf("big endian\n");
  } 
  else if (ptr[1] == 0x04) {
    printf("little endian\n");
  }
  else {
    printf("Unknown endianness\n");
  }
}
