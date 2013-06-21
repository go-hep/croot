#ifndef CROOT_CROOT_TYPES_H
#define CROOT_CROOT_TYPES_H 1

#ifdef __cplusplus
extern "C" {
#endif

  typedef const char CRoot_Option; /* Option_t */
  typedef int CRoot_Bool;

  typedef void *CRoot_Branch; /* TBranch */
  typedef void *CRoot_BranchElement; /* TBranchElement */
  typedef void *CRoot_Chain; /* TChain */
  typedef void *CRoot_Leaf; /* TLeaf */
  typedef void *CRoot_LeafD; /* TLeafD */
  typedef void *CRoot_LeafF; /* TLeafF */
  typedef void *CRoot_LeafI; /* TLeafI */
  typedef void *CRoot_LeafO; /* TLeafO */
  typedef void *CRoot_ObjArray; /* TObjArray */
  typedef void *CRoot_Object; /* TObject */
  typedef void *CRoot_ROOT;  /*TROOT*/
  typedef void *CRoot_Random; /* TRandom */
  typedef void *CRoot_Cint_TagInfo;
  typedef void *CRoot_Class; /* TClass */
  typedef void *CRoot_File; /* TFile */
  typedef void *CRoot_H1F; /* TH1F */
  typedef void *CRoot_Tree; /* TTree */

  typedef void* CRoot_Reflex_Type;
  typedef void* CRoot_Reflex_Member;
  typedef void* CRoot_Reflex_PropertyList;
  typedef void* CRoot_Reflex_Scope;
  typedef void* CRoot_Reflex_ClassBuilder;
  typedef void* CRoot_Reflex_FunctionBuilder;

  typedef void *CRoot_GoObject;
  typedef void *CRoot_GoObject_Converter;

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_TYPES_H */
