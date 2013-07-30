#ifndef CROOT_CROOT_LEAF_H
#define CROOT_CROOT_LEAF_H 1

#ifdef __cplusplus
extern "C" {
#endif

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

CROOT_API
void
CRoot_Leaf_SetAddress(CRoot_Leaf self, void* addr);

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

/* TLeafO */

CROOT_API
double
CRoot_LeafO_GetValue(CRoot_LeafO self, int idx);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_LEAF_H */
