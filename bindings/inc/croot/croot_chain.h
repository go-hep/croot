#ifndef CROOT_CROOT_CHAIN_H
#define CROOT_CROOT_CHAIN_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* TChain */

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

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_CHAIN_H */
