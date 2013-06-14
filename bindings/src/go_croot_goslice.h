#ifndef GO_CROOT_GOSLICE_H
#define GO_CROOT_GOSLICE_H 1

#include <stdint.h>

struct go_slice {
  void *Data;
  int64_t Len;
  int64_t Cap;
};

#endif /* !GO_CROOT_GOSLICE_H */
