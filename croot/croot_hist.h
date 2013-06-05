#ifndef CROOT_CROOT_HIST_H
#define CROOT_CROOT_HIST_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TH1F */

CROOT_API
CRoot_H1F
CRoot_H1F_new(const char *name, 
			  const char *title,
			  int32_t nbinsx,
			  double xlow,
			  double xup);

CROOT_API
CRoot_H1F
CRoot_H1F_new2(const char *name, 
			   const char *title,
			   int32_t nbinsx,
			   double *xbins);

CROOT_API
void
CRoot_H1F_AddBinContent(CRoot_H1F self, int32_t bin, double weight);

CROOT_API
double
CRoot_H1F_GetBinContent(CRoot_H1F self, int32_t bin);

CROOT_API
void
CRoot_H1F_SetBinContent(CRoot_H1F self, int32_t bin, double content);

CROOT_API
int32_t
CRoot_H1F_Fill(CRoot_H1F self, double x, double w);

CROOT_API
void
CRoot_H1F_FillN(CRoot_H1F self, int32_t ntimes, const double *x, const double *w, int32_t stride);

CROOT_API
int32_t
CRoot_H1F_GetBin(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetBinCenter(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetBinError(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetBinErrorLow(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetBinErrorUp(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetBinWidth(CRoot_H1F self, int32_t binx);

CROOT_API
double
CRoot_H1F_GetEntries(CRoot_H1F self);

CROOT_API
double
CRoot_H1F_GetMean(CRoot_H1F self);

CROOT_API
double
CRoot_H1F_GetMeanError(CRoot_H1F self);

CROOT_API
double
CRoot_H1F_GetRMS(CRoot_H1F self);

CROOT_API
double
CRoot_H1F_GetRMSError(CRoot_H1F self);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_HIST_H */
