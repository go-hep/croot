#include "croot/croot.h"

#include "TBranch.h"
#include "TBranchElement.h"
#include "TLeaf.h"
#include "TLeafI.h"
#include "TLeafF.h"
#include "TLeafD.h"
#include "TTree.h"
#include "TChain.h"

#include "TFile.h"

#include "TObject.h"
#include "TObjArray.h"

#include "TROOT.h"
#include "TMath.h"
#include "TRandom.h"

#include "Api.h"

/* TObject */
const char*
CRoot_Object_ClassName(CRoot_Object self)
{
  return ((TObject*)self)->ClassName();
}

CRoot_Object
CRoot_Object_Clone(CRoot_Object self,
                   const char *newname)
{
  return (TObject*)(((TObject*)self)->Clone(newname));
}

CRoot_Object
CRoot_Object_FindObject(CRoot_Object self, 
                        const char *name)
{
  return (TObject*)(((TObject*)self)->FindObject(name));
}

const char*
CRoot_Object_GetName(CRoot_Object self)
{
  return ((TObject*)self)->GetName();
}

const char*
CRoot_Object_GetTitle(CRoot_Object self)
{
  return ((TObject*)self)->GetTitle();
}

CRoot_Bool
CRoot_Object_InheritsFrom(CRoot_Object self, 
                          const char *classname)
{
  return (CRoot_Bool)(((TObject*)self)->InheritsFrom(classname));
}

void
CRoot_Object_Print(CRoot_Object self,
                   CRoot_Option *option)
{
  return ((TObject*)self)->Print((Option_t*)option);
}

/* TObjArray */
int64_t
CRoot_ObjArray_GetSize(CRoot_ObjArray self)
{
  return int64_t(((TObjArray*)self)->GetSize());
}

int64_t
CRoot_ObjArray_GetEntries(CRoot_ObjArray self)
{
  return int64_t(((TObjArray*)self)->GetEntries());
}

CRoot_Object
CRoot_ObjArray_At(CRoot_ObjArray self, int64_t idx)
{
  return (TObject*)(((TObjArray*)self)->At(idx));
}

const char*
CRoot_ObjArray_GetName(CRoot_ObjArray self)
{
  return ((TObjArray*)self)->GetName();
}

/* TROOT */
CRoot_ROOT CRoot_GRoot;

CRoot_File
CRoot_ROOT_GetFile(CRoot_ROOT self,
                   const char *name)
{
  TFile *f = ((TROOT*)self)->GetFile(name);
  return (CRoot_File)f;
}


/* TTree */
CRoot_Tree
CRoot_Tree_new(const char *name, const char *title, int32_t splitlevel)
{
  TTree *self = new TTree(name, title, splitlevel);
  return (CRoot_Tree)self;
}

void
CRoot_Tree_delete(CRoot_Tree self)
{
  TTree *tree = (TTree*)self;
  delete tree;
  self = 0;
}

CRoot_Branch
CRoot_Tree_Branch(CRoot_Tree self,
                  const char *name, const char *classname,
                  void *addobj, int32_t bufsize, int32_t splitlevel)
{
  return (CRoot_Branch)(((TTree*)self)->Branch(name, classname, addobj,
                                               bufsize, splitlevel));
}

CRoot_Branch
CRoot_Tree_Branch2(CRoot_Tree self,
                   const char *name, void *address, const char *leaflist,
                   int32_t bufsize)
{
  return (CRoot_Branch)(((TTree*)self)->Branch(name, address, 
                                               leaflist, bufsize));
}

int
CRoot_Tree_Fill(CRoot_Tree self)
{
  return ((TTree*)self)->Fill();
}

CRoot_Branch
CRoot_Tree_GetBranch(CRoot_Tree self,
                     const char *name)
{
  return (CRoot_Branch)(((TTree*)self)->GetBranch(name));
}

int64_t
CRoot_Tree_GetEntries(CRoot_Tree self)
{
  return ((TTree*)self)->GetEntries();
}

int32_t
CRoot_Tree_GetEntry(CRoot_Tree self,
                    int64_t entry, int32_t getall)
{
  return ((TTree*)self)->GetEntry(entry, getall);
}

CRoot_Leaf
CRoot_Tree_GetLeaf(CRoot_Tree self,
                   const char *name)
{
  return (CRoot_Leaf)(((TTree*)self)->GetLeaf(name));
}

CRoot_ObjArray
CRoot_Tree_GetListOfBranches(CRoot_Tree self)
{
  return (CRoot_ObjArray)(((TTree*)self)->GetListOfBranches());
}

CRoot_ObjArray
CRoot_Tree_GetListOfLeaves(CRoot_Tree self)
{
  return (CRoot_ObjArray)(((TTree*)self)->GetListOfLeaves());
}

int64_t
CRoot_Tree_GetSelectedRows(CRoot_Tree self)
{
  return ((TTree*)self)->GetSelectedRows();
}

double*
CRoot_Tree_GetVal(CRoot_Tree self,
                  int32_t i)
{
  return (double*)(((TTree*)self)->GetVal(i));
}

double*
CRoot_Tree_GetV1(CRoot_Tree self)
{
  return (double*)(((TTree*)self)->GetV1());
}

double*
CRoot_Tree_GetV2(CRoot_Tree self)
{
  return (double*)(((TTree*)self)->GetV2());
}

double*
CRoot_Tree_GetV3(CRoot_Tree self)
{
  return (double*)(((TTree*)self)->GetV3());
}

double*
CRoot_Tree_GetV4(CRoot_Tree self)
{
  return (double*)(((TTree*)self)->GetV4());
}

double*
CRoot_Tree_GetW(CRoot_Tree self)
{
  return (double*)(((TTree*)self)->GetW());
}

int64_t
CRoot_Tree_LoadTree(CRoot_Tree self,
                    int64_t entry)
{
  return ((TTree*)self)->LoadTree(entry);
}

int32_t
CRoot_Tree_MakeClass(CRoot_Tree self,
                     const char *classname, CRoot_Option *option)
{
  return ((TTree*)self)->MakeClass(classname, (Option_t*)option);
}

CRoot_Bool
CRoot_Tree_Notify(CRoot_Tree self)
{
  return (CRoot_Bool)(((TTree*)self)->Notify());
}

void
CRoot_Tree_Print(CRoot_Tree self,
                 CRoot_Option *option)
{
  return ((TTree*)self)->Print((Option_t*)option);
}

int64_t
CRoot_Tree_Process(CRoot_Tree self,
                   const char *filename, CRoot_Option *option,
                   int64_t nentries, int64_t firstentry)
{
  return ((TTree*)self)->Process(filename, (Option_t*)option,
                                 nentries, firstentry);
}

int64_t
CRoot_Tree_Project(CRoot_Tree self,
                   const char *hname, const char *varexp,
                   const char *selection, CRoot_Option *option,
                   int64_t nentries, int64_t firstentry)
{
  return ((TTree*)self)->Project(hname, varexp, selection,
                                 (Option_t*)option, nentries, firstentry);
}

int32_t
CRoot_Tree_SetBranchAddress(CRoot_Tree self,
                            const char *bname, void *addr, CRoot_Branch *ptr)
{
  return ((TTree*)self)->SetBranchAddress(bname, addr, (TBranch**)ptr);
}


void
CRoot_Tree_SetBranchStatus(CRoot_Tree self,
                           const char *bname, CRoot_Bool status, 
                           uint32_t* found)
{
  return ((TTree*)self)->SetBranchStatus(bname, (Bool_t)status, found);
}

int32_t
CRoot_Tree_Write(CRoot_Tree self,
                 const char *name, int32_t option, int32_t bufsize)
{
  return ((TTree*)self)->Write(name, option, bufsize);
}

/* TChain */

CRoot_Chain
CRoot_Chain_new(const char *name, const char *title)
{
  TChain *self = new TChain(name, title);
  return (CRoot_Chain)self;
}

void
CRoot_Chain_delete(CRoot_Chain self)
{
  TChain *tree = (TChain*)self;
  delete tree;
  self = 0;
}

int32_t
CRoot_Chain_Add(CRoot_Chain self,
                const char *name, int64_t nentries)
{
  return ((TChain*)self)->Add(name, nentries);
}

int32_t
CRoot_Chain_AddFile(CRoot_Chain self,
                    const char *name, int64_t nentries, const char *tname)
{
  return ((TChain*)self)->AddFile(name, nentries, tname);
}

int64_t
CRoot_Chain_GetEntries(CRoot_Chain self)
{
  return ((TChain*)self)->GetEntries();
}

int32_t
CRoot_Chain_GetEntry(CRoot_Chain self,
                     int64_t entry, int32_t getall)
{
  return ((TChain*)self)->GetEntry(entry, getall);
}

/* TBranch */
char*
CRoot_Branch_GetAddress(CRoot_Branch self)
{
  return ((TBranch*)self)->GetAddress();
}

char*
CRoot_Branch_GetObject(CRoot_Branch self)
{
  return ((TBranch*)self)->GetObject();
}

const char*
CRoot_Branch_GetClassName(CRoot_Branch self)
{
  return ((TBranch*)self)->GetClassName();
}

/* TLeaf */
int
CRoot_Leaf_GetLenStatic(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetLenStatic();
}

CRoot_Leaf
CRoot_Leaf_GetLeafCount(CRoot_Leaf self)
{
  return (CRoot_Leaf)(((TLeaf*)self)->GetLeafCount());
}

const char*
CRoot_Leaf_GetTypeName(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetTypeName();
}

void*
CRoot_Leaf_GetValuePointer(CRoot_Leaf self)
{
  return ((TLeaf*)self)->GetValuePointer();
}

/* TLeafI */
double
CRoot_LeafI_GetValue(CRoot_LeafI self, int idx)
{
  return ((TLeafI*)self)->GetValue(idx);
}

/* TLeafF */
double
CRoot_LeafF_GetValue(CRoot_LeafF self, int idx)
{
  return ((TLeafF*)self)->GetValue(idx);  
}

/* TLeafD */
double
CRoot_LeafD_GetValue(CRoot_LeafD self, int idx)
{
  return ((TLeafD*)self)->GetValue(idx);    
}


/* TBranchElement */
char*
CRoot_BranchElement_GetAddress(CRoot_BranchElement self)
{
  return ((TBranchElement*)self)->GetAddress();
}

const char*
CRoot_BranchElement_GetClassName(CRoot_BranchElement self)
{
  return ((TBranchElement*)self)->GetClassName();
}

/* TFile */
CRoot_File
CRoot_File_Open(const char *name, 
                CRoot_Option *option,
                const char *ftitle,
                int32_t compress,
                int32_t netopt)
{
  return (CRoot_File)(TFile::Open(name, (Option_t*)option, 
                                  ftitle, 
                                  compress, netopt));
}

CRoot_Bool
CRoot_File_cd(CRoot_File self, const char *path)
{
  return (CRoot_Bool)(((TFile*)self)->cd(path));
}

void
CRoot_File_Close(CRoot_File self, CRoot_Option *option)
{
  return ((TFile*)self)->Close((Option_t*)option);
}

int
CRoot_File_GetFd(CRoot_File self)
{
  return ((TFile*)self)->GetFd();
}

CRoot_Object
CRoot_File_Get(CRoot_File self, const char *namecycle)
{
  return (TObject*)((TFile*)self)->Get(namecycle);
}

CRoot_Bool
CRoot_File_IsOpen(CRoot_File self)
{
  return (CRoot_Bool)(((TFile*)self)->IsOpen());
}

CRoot_Bool
CRoot_File_ReadBuffer(CRoot_File self,
                      char *buf, int64_t pos, int32_t len)
{
  return (CRoot_Bool)(((TFile*)self)->ReadBuffer(buf, pos, len));
}

CRoot_Bool
CRoot_File_ReadBuffers(CRoot_File self,
                       char *buf, int64_t *pos, int32_t *len, int32_t nbuf)
{
  return (CRoot_Bool)(((TFile*)self)->ReadBuffers(buf, 
                                                  (Long64_t*)pos, 
                                                  (Int_t*)len, 
                                                  nbuf));
}

int32_t
CRoot_File_WriteBuffer(CRoot_File self,
                       const char *buf, int32_t len)
{
  return ((TFile*)self)->WriteBuffer(buf, len);
}

int32_t
CRoot_File_Write(CRoot_File self, 
                 const char *name, int32_t opt, int32_t bufsiz)
{
  return ((TFile*)self)->Write(name, opt, bufsiz);
}

/* TRandom */
CRoot_Random CRoot_gRandom = (CRoot_Random)gRandom;

int32_t
CRoot_Random_Binomial(CRoot_Random self, int32_t ntot, double prob)
{
  return ((TRandom*)self)->Binomial(ntot, prob);
}

double
CRoot_Random_Gaus(CRoot_Random self,
                  double mean, double sigma)
{
  return ((TRandom*)self)->Gaus(mean, sigma);
}

void
CRoot_Random_Rannorf(CRoot_Random self,
                     float *a, float *b)
{
  return ((TRandom*)self)->Rannor(*a, *b);
}

void
CRoot_Random_Rannord(CRoot_Random self,
                     double *a, double *b)
{
  return ((TRandom*)self)->Rannor(*a, *b);
}

double
CRoot_Random_Rndm(CRoot_Random self,
                  int32_t i)
{
  return ((TRandom*)self)->Rndm(i);
}

/* TMath */


double
CRoot_Math_Sin(double x)
{
  return TMath::Sin(x);
}


double
CRoot_Math_Cos(double x)
{
  return TMath::Cos(x);
}

double
CRoot_Math_Tan(double x)
{
  return TMath::Tan(x);
}

double
CRoot_Math_SinH(double x)
{
  return TMath::SinH(x);
}

double
CRoot_Math_CosH(double x)
{
  return TMath::CosH(x);
}

double
CRoot_Math_TanH(double x)
{
  return TMath::TanH(x);
}

double
CRoot_Math_ASin(double x)
{
  return TMath::ASin(x);
}

double
CRoot_Math_ACos(double x)
{
  return TMath::ACos(x);
}

double
CRoot_Math_ATan(double x)
{
  return TMath::ATan(x);
}

double
CRoot_Math_ATan2(double x, double y)
{
  return TMath::ATan2(x,y);
}

double
CRoot_Math_ASinH(double x)
{
  return TMath::ASinH(x);
}

double
CRoot_Math_ACosH(double x)
{
  return TMath::ACosH(x);
}

double
CRoot_Math_ATanH(double x)
{
  return TMath::ATanH(x);
}

double
CRoot_Math_Hypot(double x, double y)
{
  return TMath::Hypot(x, y);
}

double
CRoot_Math_Sqrt(double x)
{
  return TMath::Sqrt(x);
}

double
CRoot_Math_Ceil(double x)
{
  return TMath::Ceil(x);
}

int32_t
CRoot_Math_CeilNint(double x)
{
  return TMath::CeilNint(x);
}

double
CRoot_Math_Floor(double x)
{
  return TMath::Floor(x);
}

int32_t
CRoot_Math_FloorNint(double x)
{
  return TMath::FloorNint(x);
}

double
CRoot_Math_Exp(double x)
{
  return TMath::Exp(x);
}

double
CRoot_Math_Ldexp(double x, int32_t exp)
{
  return TMath::Ldexp(x, exp);
}

double
CRoot_Math_Factorial(int32_t i)
{
  return TMath::Factorial(i);
}


double
CRoot_Math_Power(double x, double y)
{
  return TMath::Power(x, y);
}

double
CRoot_Math_Log(double x)
{
  return TMath::Log(x);
}

double
CRoot_Math_Log2(double x)
{
  return TMath::Log2(x);
}
 
double
CRoot_Math_Log10(double x)
{
  return TMath::Log10(x);
}


/* -- CINT-API -- */
#if 1
CRoot_Cint_TagInfo
CRoot_Cint_TagInfo_new()
{
  G__linked_taginfo *self = new G__linked_taginfo;
  return (CRoot_Cint_TagInfo)self;
}

void
CRoot_Cint_TagInfo_delete(CRoot_Cint_TagInfo self)
{
  G__linked_taginfo *ti = (G__linked_taginfo*)self;
  delete ti;
  self = 0;
}

void
CRoot_Cint_TagInfo_SetTagName(CRoot_Cint_TagInfo self, const char* tagname)
{
  ((G__linked_taginfo*)self)->tagname = tagname;
}

void
CRoot_Cint_TagInfo_SetTagType(CRoot_Cint_TagInfo self, char tagtype)
{
  ((G__linked_taginfo*)self)->tagtype = tagtype;  
}

void
CRoot_Cint_TagInfo_SetTagNum(CRoot_Cint_TagInfo self, short tagnum)
{
  ((G__linked_taginfo*)self)->tagnum = tagnum;
}

const char*
CRoot_Cint_TagInfo_GetTagName(CRoot_Cint_TagInfo self)
{
  return ((G__linked_taginfo*)self)->tagname;
}

char
CRoot_Cint_TagInfo_GetTagType(CRoot_Cint_TagInfo self)
{
  return ((G__linked_taginfo*)self)->tagtype;
}

short
CRoot_Cint_TagInfo_GetTagNum(CRoot_Cint_TagInfo self)
{
  return ((G__linked_taginfo*)self)->tagnum;
}

int
CRoot_Cint_TagInfo_GetLinkedTagNum(CRoot_Cint_TagInfo self)
{
  return G__get_linked_tagnum((G__linked_taginfo*)self);
}

int
CRoot_Cint_Defined_TagName(const char* tagname, int noerror)
{
  return G__defined_tagname(tagname, noerror);
}

int
CRoot_Cint_TagTable_Setup(int tagnum, int size, int cpplink, int isabstract,
                          const char* comment, 
                          CRoot_Cint_incsetup setup_memvar,
                          CRoot_Cint_incsetup setup_memfunc)
{
  return G__tagtable_setup
    (tagnum,
     size,
     cpplink,
     isabstract,
     comment,
     (G__incsetup)setup_memvar,
     (G__incsetup)setup_memfunc);
}

int
CRoot_Cint_Tag_MemVar_Setup(int tagnum)
{
  return G__tag_memvar_setup(tagnum);
}

int
CRoot_Cint_MemVar_Setup(void *p, int type, int reftype,
                        int constvar, 
                        int tagnum,
                        int typenum,
                        int statictype,
                        int var_access,
                        const char* expr,
                        int definemacro,
                        const char* comment)
{
  return G__memvar_setup(p, type, reftype, constvar, tagnum, 
                         typenum,
                         statictype,
                         var_access,
                         expr,
                         definemacro,
                         comment);
}
 
int
CRoot_Cint_Tag_MemVar_Reset()
{
  return G__tag_memvar_reset();
}
#endif

// reflex API
#include "Reflex/Reflex.h"

void
CRoot_Reflex_FireClassCallback(CRoot_Reflex_Type self)
{
  Reflex::FireClassCallback(*((Reflex::Type*)self));
}

void
CRoot_Reflex_FireFunctionCallback(CRoot_Reflex_Member self)
{
  Reflex::FireFunctionCallback(*((Reflex::Member*)self));
}

CRoot_Reflex_Type
CRoot_Reflex_Type_new(const char* name, unsigned int modifiers)
{
  Reflex::Type t = Reflex::TypeBuilder(name, modifiers);
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}


void
CRoot_Reflex_Type_delete(CRoot_Reflex_Type self)
{
  Reflex::Type *t = (Reflex::Type*)self;
  delete t;
  self = 0;
}

void*
CRoot_Reflex_Type_Id(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->Id();
}

size_t
CRoot_Reflex_Type_ArrayLength(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->ArrayLength();
}


CRoot_Reflex_Type
CRoot_Reflex_Type_ByName(const char *name)
{
  Reflex::Type t = Reflex::Type::ByName(std::string(name));
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}


CRoot_Reflex_Member
CRoot_Reflex_Type_FunctionMemberAt(CRoot_Reflex_Type self,
                                   size_t nth,
                                   CRoot_Reflex_EMEMBERQUERY inh)
{
  Reflex::Member mbr = ((Reflex::Type*)self)->FunctionMemberAt
    (nth, (Reflex::EMEMBERQUERY)inh);
  // FIXME: leak
  return (CRoot_Reflex_Member)new Reflex::Member(mbr);
}

size_t
CRoot_Reflex_Type_FunctionMemberSize(CRoot_Reflex_Type self,
                                     CRoot_Reflex_EMEMBERQUERY inh)
{
  return ((Reflex::Type*)self)->FunctionMemberSize((Reflex::EMEMBERQUERY)inh);
}

CRoot_Reflex_Member
CRoot_Reflex_Type_DataMemberAt(CRoot_Reflex_Type self,
                               size_t nth,
                               CRoot_Reflex_EMEMBERQUERY inh)
{
  Reflex::Member mbr = ((Reflex::Type*)self)->DataMemberAt
    (nth, (Reflex::EMEMBERQUERY)inh);
  // FIXME: leak
  return (CRoot_Reflex_Member)new Reflex::Member(mbr);
}


size_t
CRoot_Reflex_Type_DataMemberSize(CRoot_Reflex_Type self,
                                 CRoot_Reflex_EMEMBERQUERY inh)
{
  return ((Reflex::Type*)self)->DataMemberSize((Reflex::EMEMBERQUERY)inh);
}


bool
CRoot_Reflex_Type_IsAbstract(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsAbstract();
}


bool
CRoot_Reflex_Type_IsArray(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsArray();
}

bool
CRoot_Reflex_Type_IsClass(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsClass();
}

bool
CRoot_Reflex_Type_IsComplete(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsComplete();
}

bool
CRoot_Reflex_Type_IsEquivalentTo(CRoot_Reflex_Type self,
                                 CRoot_Reflex_Type other,
                                 unsigned int modifiers_mask)
{
  return ((Reflex::Type*)self)->IsEquivalentTo(*((Reflex::Type*)other),
                                               modifiers_mask);
}

bool
CRoot_Reflex_Type_IsFunction(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsFunction();
}

bool
CRoot_Reflex_Type_IsFundamental(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsFundamental();
}

bool
CRoot_Reflex_Type_IsPrivate(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsPrivate();
}

bool
CRoot_Reflex_Type_IsProtected(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsProtected();
}

bool
CRoot_Reflex_Type_IsPublic(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsPublic();
}

bool
CRoot_Reflex_Type_IsPointer(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsPointer();
}

bool
CRoot_Reflex_Type_IsPointerToMember(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsPointerToMember();
}

bool
CRoot_Reflex_Type_IsReference(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsReference();
}

bool
CRoot_Reflex_Type_IsStruct(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsStruct();
}

bool
CRoot_Reflex_Type_IsVirtual(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->IsVirtual();
}

CRoot_Reflex_Member
CRoot_Reflex_Type_MemberAt(CRoot_Reflex_Type self,
                           size_t nth,
                           CRoot_Reflex_EMEMBERQUERY inh)
{
  Reflex::Member mbr = ((Reflex::Type*)self)->MemberAt(nth, (Reflex::EMEMBERQUERY)inh);
  // FIXME: leak
  return (CRoot_Reflex_Member)new Reflex::Member(mbr);
}


size_t
CRoot_Reflex_Type_MemberSize(CRoot_Reflex_Type self,
                             CRoot_Reflex_EMEMBERQUERY inh)
{
  return ((Reflex::Type*)self)->MemberSize((Reflex::EMEMBERQUERY)inh);
}


const char*
CRoot_Reflex_Type_Name(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->Name_c_str();
}

CRoot_Reflex_PropertyList
CRoot_Reflex_Type_Properties(CRoot_Reflex_Type self)
{
  Reflex::PropertyList plist = ((Reflex::Type*)self)->Properties();
  // FIXME: leak
  return (CRoot_Reflex_PropertyList)new Reflex::PropertyList(plist);
}

CRoot_Reflex_Type
CRoot_Reflex_Type_RawType(CRoot_Reflex_Type self)
{
  Reflex::Type tt = ((Reflex::Type*)self)->RawType();
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(tt);
}


size_t
CRoot_Reflex_Type_SizeOf(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->SizeOf();
}

CRoot_Reflex_Type
CRoot_Reflex_Type_ToType(CRoot_Reflex_Type self)
{
  Reflex::Type t = ((Reflex::Type*)self)->ToType();
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}


CRoot_Reflex_Type
CRoot_Reflex_Type_TypeAt(size_t nth)
{
  Reflex::Type t = Reflex::Type::TypeAt(nth);
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}

size_t
CRoot_Reflex_Type_TypeSize()
{
  return Reflex::Type::TypeSize();
}


CRoot_Reflex_TYPE
CRoot_Reflex_Type_TypeType(CRoot_Reflex_Type self)
{
  return (CRoot_Reflex_TYPE)(((Reflex::Type*)self)->TypeType());
}

void
CRoot_Reflex_Type_Unload(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->Unload();
}

void
CRoot_Reflex_Type_UpdateMembers(CRoot_Reflex_Type self)
{
  return ((Reflex::Type*)self)->UpdateMembers();
}


void
CRoot_Reflex_Type_AddDataMember(CRoot_Reflex_Type self,
                                CRoot_Reflex_Member dm)
{
  return ((Reflex::Type*)self)->AddDataMember(*((Reflex::Member*)dm));
}


CRoot_Reflex_Member
CRoot_Reflex_Type_AddDataMember2(CRoot_Reflex_Type self,
                                 const char* name,
                                 CRoot_Reflex_Type type,
                                 size_t offset,
                                 unsigned int modifiers,
                                 char *interpreterOffset)
{
  Reflex::Member mbr = ((Reflex::Type*)self)->AddDataMember
    (name,
     *((Reflex::Type*)type),
     offset,
     modifiers,
     interpreterOffset);
  // FIXME: leak
  return (CRoot_Reflex_Member)new Reflex::Member(mbr);
}


void
CRoot_Reflex_Type_RemoveDataMember(CRoot_Reflex_Type self,
                                   CRoot_Reflex_Member dm)
{
  return ((Reflex::Type*)self)->RemoveDataMember(*(Reflex::Member*)dm);
}

void
CRoot_Reflex_Type_SetSize(CRoot_Reflex_Type self,
                          size_t s)
{
  return ((Reflex::Type*)self)->SetSize(s);
}

CRoot_Reflex_REPRESTYPE
CRoot_Reflex_Type_RepresType(CRoot_Reflex_Type self)
{
  return (CRoot_Reflex_REPRESTYPE)(((Reflex::Type*)self)->RepresType());
}

CRoot_Reflex_Member
CRoot_Reflex_Member_new()
{
  Reflex::Member *mbr = new Reflex::Member;
  // FIXME: leak
  return (CRoot_Reflex_Member)mbr;
}

void
CRoot_Reflex_Member_delete(CRoot_Reflex_Member self)
{
  Reflex::Member *mbr = (Reflex::Member*)self;
  delete mbr;
  self = 0;
}

bool
CRoot_Reflex_Member_IsDataMember(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsDataMember();
}

bool
CRoot_Reflex_Member_IsPrivate(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsPrivate();
}

bool
CRoot_Reflex_Member_IsProtected(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsProtected();
}

bool
CRoot_Reflex_Member_IsPublic(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsPublic();
}

bool
CRoot_Reflex_Member_IsTransient(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsTransient();
}

bool
CRoot_Reflex_Member_IsVirtual(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->IsVirtual();
}

CRoot_Reflex_TYPE
CRoot_Reflex_Member_MemberType(CRoot_Reflex_Member self)
{
  return (CRoot_Reflex_TYPE)((Reflex::Member*)self)->MemberType();
}

const char*
CRoot_Reflex_Member_Name(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->Name_c_str();
}

size_t
CRoot_Reflex_Member_Offset(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->Offset();
}

void
CRoot_Reflex_Member_InterpreterOffset(CRoot_Reflex_Member self,
                                      char *offset)
{
  return ((Reflex::Member*)self)->InterpreterOffset(offset);
}

CRoot_Reflex_PropertyList
CRoot_Reflex_Member_Properties(CRoot_Reflex_Member self)
{
  Reflex::PropertyList plist = ((Reflex::Member*)self)->Properties();
  //FIXME: leak
  return (CRoot_Reflex_PropertyList)new Reflex::PropertyList(plist);
}

CRoot_Reflex_Type
CRoot_Reflex_Member_TypeOf(CRoot_Reflex_Member self)
{
  Reflex::Type t = ((Reflex::Member*)self)->TypeOf();
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}

void*
CRoot_Reflex_Member_Stubcontext(CRoot_Reflex_Member self)
{
  return ((Reflex::Member*)self)->Stubcontext();
}

CRoot_Reflex_StubFunction
CRoot_Reflex_Member_Stubfunction(CRoot_Reflex_Member self)
{
  return (CRoot_Reflex_StubFunction)((Reflex::Member*)self)->Stubfunction();
}

const char*
CRoot_Reflex_PropertyList_PropertyAsString(CRoot_Reflex_PropertyList self,
                                           size_t idx)
{
  std::string str = ((Reflex::PropertyList*)self)->PropertyAsString(idx);
  return strdup(str.c_str());
}

size_t
CRoot_Reflex_PropertyList_PropertyCount(CRoot_Reflex_PropertyList self)
{
  return ((Reflex::PropertyList*)self)->PropertyCount();
}

const char*
CRoot_Reflex_PropertyList_PropertyKeys(CRoot_Reflex_PropertyList self)
{
  std::string str = ((Reflex::PropertyList*)self)->PropertyKeys();
  return strdup(str.c_str());
}

#include "Reflex/Builder/ReflexBuilder.h"

CRoot_Reflex_Type
CRoot_Reflex_PointerBuilder_new(CRoot_Reflex_Type t)
{
  Reflex::Type *ty = (Reflex::Type*)t;
  Reflex::Type ty_ptr = Reflex::PointerBuilder(*ty);
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(ty_ptr);
}

CRoot_Reflex_Type
CRoot_Reflex_ArrayBuilder_new(CRoot_Reflex_Type t,
                              size_t n)
{
  Reflex::Type *ty = (Reflex::Type*)t;
  Reflex::Type ty_arr = Reflex::ArrayBuilder(*ty, n);
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(ty_arr);
}

CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new(CRoot_Reflex_Type r)
{
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type
    (Reflex::FunctionTypeBuilder(*(Reflex::Type*)r));
}

CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new1(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0)
{
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type
    (Reflex::FunctionTypeBuilder(*(Reflex::Type*)r,
                                 *(Reflex::Type*)t0));
}

CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new2(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0,
                                      CRoot_Reflex_Type t1)
{
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type
    (Reflex::FunctionTypeBuilder(*(Reflex::Type*)r,
                                 *(Reflex::Type*)t0,
                                 *(Reflex::Type*)t1));
}


CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new3(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0,
                                      CRoot_Reflex_Type t1,
                                      CRoot_Reflex_Type t2)
{
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type
    (Reflex::FunctionTypeBuilder(*(Reflex::Type*)r,
                                 *(Reflex::Type*)t0,
                                 *(Reflex::Type*)t1,
                                 *(Reflex::Type*)t2));
}

CRoot_Reflex_ClassBuilder
CRoot_Reflex_ClassBuilder_new(const char *name,
                              void* typeinfo,
                              size_t size,
                              unsigned int modifiers,
                              CRoot_Reflex_TYPE type)
{
  const std::type_info *ti = &typeid(void);
  if (typeinfo) {
    ti = (const std::type_info*)typeinfo;
  }
  Reflex::ClassBuilder * cb = new Reflex::ClassBuilder(name, *ti, size, modifiers, (Reflex::TYPE)type);
  return (CRoot_Reflex_ClassBuilder)cb;
}

void
CRoot_Reflex_ClassBuilder_delete(CRoot_Reflex_ClassBuilder self)
{
  Reflex::ClassBuilder *cb = (Reflex::ClassBuilder*)self;
  delete cb; cb = 0;
  self = 0;
}

void
CRoot_Reflex_ClassBuilder_AddDataMember(CRoot_Reflex_ClassBuilder self,
                                        CRoot_Reflex_Type type,
                                        const char* name,
                                        size_t offset,
                                        unsigned int modifiers)
{
  Reflex::ClassBuilder *cb = (Reflex::ClassBuilder*)self;
  cb->AddDataMember(*((Reflex::Type*)type), name, offset, modifiers);
  //return (CRoot_Reflex_ClassBuilder)cb;
}

void
CRoot_Reflex_ClassBuilder_AddFunctionMember(CRoot_Reflex_ClassBuilder self,
                                            CRoot_Reflex_Type type,
                                            const char *name,
                                            CRoot_Reflex_StubFunction c_stubFP,
                                            void *stubCtx,
                                            const char *params,
                                            unsigned int modifiers)
{
  Reflex::ClassBuilder *cb = (Reflex::ClassBuilder*)self;
  Reflex::Type *typ = (Reflex::Type*)type;
  Reflex::StubFunction stubFP = (Reflex::StubFunction)c_stubFP;
  cb->AddFunctionMember(*typ,
                        name,
                        stubFP,
                        stubCtx,
                        params,
                        modifiers);
}

void
CRoot_Reflex_ClassBuilder_AddProperty(CRoot_Reflex_ClassBuilder self,
                                      const char *key,
                                      const char *val)
{
  std::string value(val); // otherwise C++ type conversions rules confuse Reflex::Any...
  ((Reflex::ClassBuilder*)self)->AddProperty(key, value);
}

CRoot_Reflex_Type
CRoot_Reflex_ClassBuilder_ToType(CRoot_Reflex_ClassBuilder self)
{
  Reflex::Type t = ((Reflex::ClassBuilder*)self)->ToType();
  // FIXME: leak
  return (CRoot_Reflex_Type)new Reflex::Type(t);
}

CRoot_Reflex_FunctionBuilder
CRoot_Reflex_FunctionBuilder_new(CRoot_Reflex_Type c_type,
                                 const char* name,
                                 CRoot_Reflex_StubFunction c_stubFP,
                                 void *stubCtx,
                                 const char *params,
                                 unsigned char modifiers)
{
  Reflex::Type *typ = (Reflex::Type*)c_type;
  Reflex::StubFunction stubFP = (Reflex::StubFunction)c_stubFP;
  // FIXME: leak
  Reflex::FunctionBuilder *fb = new Reflex::FunctionBuilder(*typ,
                                                            name,
                                                            stubFP,
                                                            stubCtx,
                                                            params,
                                                            modifiers);
  return (CRoot_Reflex_FunctionBuilder)fb;
}

void
CRoot_Reflex_FunctionBuilder_delete(CRoot_Reflex_FunctionBuilder self)
{
  Reflex::FunctionBuilder *fb = (Reflex::FunctionBuilder*)self;
  delete fb; fb = 0;
  self = 0;
}

CRoot_Reflex_Member
CRoot_Reflex_FunctionBuilder_ToMember(CRoot_Reflex_FunctionBuilder self)
{
  Reflex::Member mbr = ((Reflex::FunctionBuilder*)self)->ToMember();
  // FIXME: leak
  return (CRoot_Reflex_Member)new Reflex::Member(mbr);
}

size_t
CRoot_Reflex_PropertyList_AddProperty(CRoot_Reflex_PropertyList self,
                                      const char *key,
                                      const char *value)
{
  return ((Reflex::PropertyList*)self)->AddProperty(key, value);
}

void croot_reflex_init();

#include "Cintex/Cintex.h"
void
CRoot_Cintex_Enable()
{
  croot_reflex_init();
  ROOT::Cintex::Cintex::Enable();
  //ROOT::Cintex::Cintex::SetDebug(100000);
}

void
CRoot_Cintex_SetDebug(int level)
{
  ROOT::Cintex::Cintex::SetDebug(level);
}

//#include "Reflex/Builder/ReflexBuilder.h"
namespace {
  template <typename T>
  void
  croot_reflex_add_typedef(const char *name) {
    Reflex::TypedefBuilder<T> tmp(name);
  }
}

void croot_reflex_init() {

  // initialize <stdint.h> types
  ::croot_reflex_add_typedef<int8_t>("int8_t");
  ::croot_reflex_add_typedef<int16_t>("int16_t");
  ::croot_reflex_add_typedef<int32_t>("int32_t");
  ::croot_reflex_add_typedef<int64_t>("int64_t");

  ::croot_reflex_add_typedef<uint8_t>("uint8_t");
  ::croot_reflex_add_typedef<uint16_t>("uint16_t");
  ::croot_reflex_add_typedef<uint32_t>("uint32_t");
  ::croot_reflex_add_typedef<uint64_t>("uint64_t");
}

void __attribute__ ((constructor)) croot_init();

void croot_init()
{
  CRoot_GRoot = (CRoot_ROOT)gROOT;
}

