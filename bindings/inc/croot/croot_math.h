#ifndef CROOT_CROOT_MATH_H
#define CROOT_CROOT_MATH_H 1

#ifdef __cplusplus
extern "C" {
#endif


/* TMath */

CROOT_API
double
CRoot_Math_Sin(double);

CROOT_API
double
CRoot_Math_Cos(double);

CROOT_API
double
CRoot_Math_Tan(double);

CROOT_API
double
CRoot_Math_SinH(double);

CROOT_API
double
CRoot_Math_CosH(double);

CROOT_API
double
CRoot_Math_TanH(double);

CROOT_API
double
CRoot_Math_ASin(double);

CROOT_API
double
CRoot_Math_ACos(double);

CROOT_API
double
CRoot_Math_ATan(double);

CROOT_API
double
CRoot_Math_ATan2(double, double);

CROOT_API
double
CRoot_Math_ASinH(double);

CROOT_API
double
CRoot_Math_ACosH(double);

CROOT_API
double
CRoot_Math_ATanH(double);

CROOT_API
double
CRoot_Math_Hypot(double x, double y);

CROOT_API
double
CRoot_Math_Sqrt(double);

CROOT_API
double
CRoot_Math_Ceil(double);

CROOT_API
int32_t
CRoot_Math_CeilNint(double);

CROOT_API
double
CRoot_Math_Floor(double);

CROOT_API
int32_t
CRoot_Math_FloorNint(double);

CROOT_API
double
CRoot_Math_Exp(double);

CROOT_API
double
CRoot_Math_Ldexp(double x, int32_t exp);

CROOT_API
double
CRoot_Math_Factorial(int32_t i);

CROOT_API
double
CRoot_Math_Power(double x, double y);

CROOT_API
double
CRoot_Math_Log(double);

CROOT_API
double
CRoot_Math_Log2(double);

CROOT_API
double
CRoot_Math_Log10(double);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_MATH_H */
