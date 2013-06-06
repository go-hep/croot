#include "croot/croot.h"

#include "TH1F.h"

// TH1F
CRoot_H1F
CRoot_H1F_new(const char *name, const char *title, int32_t nbins, double xlow, double xup)
{
  TH1F *self = new TH1F(name, title, nbins, xlow, xup);
  return (CRoot_H1F)self;
}

CRoot_H1F
CRoot_H1F_new2(const char *name, 
			   const char *title,
			   int32_t nbinsx,
			   double *xbins)
{
  TH1F *self = new TH1F(name, title, nbinsx, xbins);
  return (CRoot_H1F)self;
}

void
CRoot_H1F_AddBinContent(CRoot_H1F self, int32_t bin, double weight)
{
  ((TH1F*)self)->AddBinContent(bin, weight);
}

double
CRoot_H1F_GetBinContent(CRoot_H1F self, int32_t bin)
{
  return ((TH1F*)self)->GetBinContent(bin);
}

void
CRoot_H1F_SetBinContent(CRoot_H1F self, int32_t bin, double content)
{
  ((TH1F*)self)->SetBinContent(bin, content);
}

int32_t
CRoot_H1F_Fill(CRoot_H1F self, double x, double w)
{
  return ((TH1F*)self)->Fill(x,w);
}

void
CRoot_H1F_FillN(CRoot_H1F self, int32_t ntimes, const double *x, const double *w, int32_t stride)
{
  ((TH1F*)self)->FillN(ntimes, x, w, stride);
}

int32_t
CRoot_H1F_GetBin(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBin(binx);
}

double
CRoot_H1F_GetBinCenter(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBinCenter(binx);
}

double
CRoot_H1F_GetBinError(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBinError(binx);
}

double
CRoot_H1F_GetBinErrorLow(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBinErrorLow(binx);
}

double
CRoot_H1F_GetBinErrorUp(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBinErrorUp(binx);
}

double
CRoot_H1F_GetBinWidth(CRoot_H1F self, int32_t binx)
{
  return ((TH1F*)self)->GetBinWidth(binx);
}

double
CRoot_H1F_GetEntries(CRoot_H1F self)
{
  return ((TH1F*)self)->GetEntries();
}

double
CRoot_H1F_GetMean(CRoot_H1F self)
{
  return ((TH1F*)self)->GetMean();
}

double
CRoot_H1F_GetMeanError(CRoot_H1F self)
{
  return ((TH1F*)self)->GetMeanError();
}

double
CRoot_H1F_GetRMS(CRoot_H1F self)
{
  return ((TH1F*)self)->GetRMS();
}

double
CRoot_H1F_GetRMSError(CRoot_H1F self)
{
  return ((TH1F*)self)->GetRMSError();
}

// EOF

