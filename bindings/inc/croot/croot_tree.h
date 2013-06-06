#ifndef CROOT_CROOT_TREE_H
#define CROOT_CROOT_TREE_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TBranch */

CROOT_API
char*
CRoot_Branch_GetAddress(CRoot_Branch self);

/* CROOT_API */
/* char* */
/* CRoot_Branch_GetObject(CRoot_Branch self); */

CROOT_API
const char*
CRoot_Branch_GetClassName(CRoot_Branch self);

CROOT_API
CRoot_Leaf
CRoot_Branch_GetLeaf(CRoot_Branch self, const char* name);

CROOT_API
CRoot_ObjArray
CRoot_Branch_GetListOfLeaves(CRoot_Branch self);

/* TBranchElement */

CROOT_API
char*
CRoot_BranchElement_GetAddress(CRoot_BranchElement self);

CROOT_API
const char*
CRoot_BranchElement_GetClassName(CRoot_BranchElement self);

/* TLeaf */

CROOT_API
CRoot_Branch
CRoot_Leaf_GetBranch(CRoot_Leaf self);

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

/* TLeafI */

CROOT_API
double
CRoot_LeafI_GetValue(CRoot_LeafI self, int idx);

/* TLeafF */

CROOT_API
double
CRoot_LeafF_GetValue(CRoot_LeafF self, int idx);

/* TLeafD */

CROOT_API
double
CRoot_LeafD_GetValue(CRoot_LeafD self, int idx);

/* TTree */

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

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_TREE_H */
