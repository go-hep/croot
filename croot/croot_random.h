#ifndef CROOT_CROOT_RANDOM_H
#define CROOT_CROOT_RANDOM_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TRandom */

CROOT_API
int32_t
CRoot_Random_Binomial(CRoot_Random self, int32_t ntot, double prob);

CROOT_API
double
CRoot_Random_Gaus(CRoot_Random self,
                  double mean, double sigma);

CROOT_API
void
CRoot_Random_Rannorf(CRoot_Random self,
                     float *a, float *b);

CROOT_API
void
CRoot_Random_Rannord(CRoot_Random self,
                     double *a, double *b);

CROOT_API
double
CRoot_Random_Rndm(CRoot_Random self,
                  int32_t i);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_RANDOM_H */
