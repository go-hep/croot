/* @brief a C-API wrapper for (some of) the C++ classes of ROOT
 */

#ifndef CROOT_CROOT_H
#define CROOT_CROOT_H 1

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>

#if __GNUC__ >= 4
#  define CROOT_HASCLASSVISIBILITY
#endif

#if defined(CROOT_HASCLASSVISIBILITY)
#  define CROOT_IMPORT __attribute__((visibility("default")))
#  define CROOT_EXPORT __attribute__((visibility("default")))
#  define CROOT_LOCAL  __attribute__((visibility("hidden")))
#else
#  define CROOT_IMPORT
#  define CROOT_EXPORT
#  define CROOT_LOCAL
#endif

#define CROOT_API CROOT_EXPORT

#ifdef __cplusplus
extern "C" {
#endif

typedef int CRoot_Bool;

/* Option_t */
typedef const char CRoot_Option;

/* TObject */
typedef void* CRoot_Object;

CROOT_API
const char*
CRoot_Object_ClassName(CRoot_Object self);

CROOT_API
CRoot_Object
CRoot_Object_Clone(CRoot_Object self,
                   const char *newname);

CROOT_API
CRoot_Object
CRoot_Object_FindObject(CRoot_Object self, 
                        const char *name);

CROOT_API
const char*
CRoot_Object_GetName(CRoot_Object self);

CROOT_API
const char*
CRoot_Object_GetTitle(CRoot_Object self);

CROOT_API
CRoot_Bool
CRoot_Object_InheritsFrom(CRoot_Object self, 
                          const char *classname);

CROOT_API
void
CRoot_Object_Print(CRoot_Object self,
                   CRoot_Option *option);

/* TObjArray */
typedef void *CRoot_ObjArray;

CROOT_API
int64_t
CRoot_ObjArray_GetSize(CRoot_ObjArray self);

CROOT_API
CRoot_Object
CRoot_ObjArray_At(CRoot_ObjArray self, int64_t idx);

CROOT_API
const char*
CRoot_ObjArray_GetName(CRoot_ObjArray self);

/* TBranch */
typedef void *CRoot_Branch;

CROOT_API
char*
CRoot_Branch_GetAddress(CRoot_Branch self);

CROOT_API
const char*
CRoot_Branch_GetClassName(CRoot_Branch self);

/* TBranchElement */
typedef void *CRoot_BranchElement;

CROOT_API
char*
CRoot_BranchElement_GetAddress(CRoot_BranchElement self);

CROOT_API
const char*
CRoot_BranchElement_GetClassName(CRoot_BranchElement self);

/* TLeaf */
typedef void *CRoot_Leaf;

CROOT_API
int
CRoot_Leaf_GetLenStatic(CRoot_Leaf self);

CROOT_API
CRoot_Leaf
CRoot_Leaf_GetLeafCount(CRoot_Leaf self);

CROOT_API
const char*
CRoot_Leaf_GetTypeName(CRoot_Leaf self);

CROOT_API
void*
CRoot_Leaf_GetValuePointer(CRoot_Leaf self);

/* TTree */
typedef void* CRoot_Tree;

CROOT_API
CRoot_Tree
CRoot_Tree_new(const char *name, const char *title, int32_t splitlevel);

CROOT_API
void
CRoot_Tree_delete(CRoot_Tree self);

CROOT_API
CRoot_Branch
CRoot_Tree_Branch(CRoot_Tree self,
                  const char *name, const char *classname,
                  void *addobj, int32_t bufsize, int32_t splitlevel);

CROOT_API
CRoot_Branch
CRoot_Tree_Branch2(CRoot_Tree self,
                   const char *name, void *address, const char *leaflist,
                   int32_t bufsize);

CROOT_API
int
CRoot_Tree_Fill(CRoot_Tree self);

CROOT_API
CRoot_Branch
CRoot_Tree_GetBranch(CRoot_Tree self,
                     const char *name);

CROOT_API
int64_t
CRoot_Tree_GetEntries(CRoot_Tree self);

CROOT_API
int32_t
CRoot_Tree_GetEntry(CRoot_Tree self,
                    int64_t entry, int32_t getall);

CROOT_API
CRoot_Leaf
CRoot_Tree_GetLeaf(CRoot_Tree self,
                   const char *name);

CROOT_API
CRoot_ObjArray
CRoot_Tree_GetListOfBranches(CRoot_Tree self);

CROOT_API
CRoot_ObjArray
CRoot_Tree_GetListOfLeaves(CRoot_Tree self);

CROOT_API
int64_t
CRoot_Tree_GetSelectedRows(CRoot_Tree self);

CROOT_API
double*
CRoot_Tree_GetVal(CRoot_Tree self,
                  int32_t i);

CROOT_API
double*
CRoot_Tree_GetV1(CRoot_Tree self);

CROOT_API
double*
CRoot_Tree_GetV2(CRoot_Tree self);

CROOT_API
double*
CRoot_Tree_GetV3(CRoot_Tree self);

CROOT_API
double*
CRoot_Tree_GetV4(CRoot_Tree self);

CROOT_API
double*
CRoot_Tree_GetW(CRoot_Tree self);

CROOT_API
int64_t
CRoot_Tree_LoadTree(CRoot_Tree self,
                    int64_t entry);

CROOT_API
int32_t
CRoot_Tree_MakeClass(CRoot_Tree self,
                     const char *classname, CRoot_Option *option);

CROOT_API
CRoot_Bool
CRoot_Tree_Notify(CRoot_Tree self);

CROOT_API
void
CRoot_Tree_Print(CRoot_Tree self,
                 CRoot_Option *option);

CROOT_API
int64_t
CRoot_Tree_Process(CRoot_Tree self,
                   const char *filename, CRoot_Option *option,
                   int64_t nentries, int64_t firstentry);

CROOT_API
int64_t
CRoot_Tree_Project(CRoot_Tree self,
                   const char *hname, const char *varexp,
                   const char *selection, CRoot_Option *option,
                   int64_t nentries, int64_t firstentry);

CROOT_API
int32_t
CRoot_Tree_SetBranchAddress(CRoot_Tree self,
                            const char *bname, void *addr, CRoot_Branch *ptr);

CROOT_API
void
CRoot_Tree_SetBranchStatus(CRoot_Tree self,
                           const char *bname, CRoot_Bool status, 
                           uint32_t *found);

CROOT_API
int32_t
CRoot_Tree_Write(CRoot_Tree self,
                 const char *name, int32_t option, int32_t bufsize);

/* TChain */
typedef void* CRoot_Chain;

CROOT_API
CRoot_Chain
CRoot_Chain_new(const char *name, const char *title);

CROOT_API
void
CRoot_Chain_delete(CRoot_Chain self);

CROOT_API
int32_t
CRoot_Chain_Add(CRoot_Chain self,
                const char *name, int64_t nentries);

CROOT_API
int32_t
CRoot_Chain_AddFile(CRoot_Chain self,
                    const char *name, int64_t nentries, const char *tname);

CROOT_API
int64_t
CRoot_Chain_GetEntries(CRoot_Chain self);

CROOT_API
int32_t
CRoot_Chain_GetEntry(CRoot_Chain self,
                     int64_t entry, int32_t getall);

/* TFile */
typedef void* CRoot_File;

typedef enum {
    kDefault =0,
    kLocal   =1,
    kNet     =2,
    kWeb     =3,
    kFile    =4
} CRoot_FileType;

CROOT_API
CRoot_File
CRoot_File_Open(const char *name, 
                CRoot_Option *option,
                const char *ftitle,
                int32_t compress,
                int32_t netopt);

CROOT_API
CRoot_Bool
CRoot_File_cd(CRoot_File self, const char *path);

CROOT_API
void
CRoot_File_Close(CRoot_File self, CRoot_Option *option);

CROOT_API
int
CRoot_File_GetFd(CRoot_File self);

CROOT_API
CRoot_Object
CRoot_File_Get(CRoot_File self, const char *namecycle);

CROOT_API
CRoot_Bool
CRoot_File_IsOpen(CRoot_File self);

CROOT_API
CRoot_Bool
CRoot_File_ReadBuffer(CRoot_File self,
                      char *buf, int64_t pos, int32_t len);

CROOT_API
CRoot_Bool
CRoot_File_ReadBuffers(CRoot_File self,
                       char *buf, int64_t *pos, int32_t *len, int32_t nbuf);

CROOT_API
int32_t
CRoot_File_WriteBuffer(CRoot_File self,
                       const char *buf, int32_t len);

CROOT_API
int32_t
CRoot_File_Write(CRoot_File self, 
                 const char *name, int32_t opt, int32_t bufsiz);

/* TROOT */
typedef void *CRoot_ROOT;

/* The global ROOT object */
extern CRoot_ROOT CRoot_GRoot;

CROOT_API
CRoot_File
CRoot_ROOT_GetFile(CRoot_ROOT self,
                   const char *name);

/* TRandom */
typedef void *CRoot_Random;

extern CRoot_Random CRoot_gRandom;

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

#include "croot/croot_hist.h"
#include "croot/croot_reflex.h"
#include "croot/croot_cintex.h"
#include "croot/croot_cint.h"

#endif /* !CROOT_CROOT_H */
