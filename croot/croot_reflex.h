#ifndef CROOT_CROOT_REFLEX_H
#define CROOT_CROOT_REFLEX_H 1

#ifdef __cplusplus
extern "C" {
#endif

/* --- Reflex API --- */

/**
 * typedef for function member type (necessary for return value of
 * getter function)
 */
typedef void (*CRoot_Reflex_StubFunction)(void*, void*, void*, void*);

/** typedef for function for Offset calculation */
typedef size_t (*CRoot_Reflex_OffsetFunction)(void*);

// these defines are used for the modifiers they are used in the following
// classes
// BA = BASE
// CL = CLASS
// FU = FUNCTION
// DM = DATAMEMBER
// FM = FUNCTIONMEMBER
// TY = TYPE
// ME = MEMBER
//                                             BA  CL  DM  FM  TY  ME
typedef enum {
   CRoot_Reflex_PUBLIC = (1 << 0),            //  X       X   X       X
   CRoot_Reflex_PROTECTED = (1 << 1),         //  X       X   X       X
   CRoot_Reflex_PRIVATE = (1 << 2),           //  X       X   X       X
   CRoot_Reflex_REGISTER = (1 << 3),          //          X   X       X
   CRoot_Reflex_STATIC = (1 << 4),            //          X   X       X
   CRoot_Reflex_CONSTRUCTOR = (1 << 5),       //              X       X
   CRoot_Reflex_DESTRUCTOR = (1 << 6),        //              X       X
   CRoot_Reflex_EXPLICIT = (1 << 7),          //              X       X
   CRoot_Reflex_EXTERN = (1 << 8),            //          X   X       X
   CRoot_Reflex_COPYCONSTRUCTOR = (1 << 9),   //              X       X
   CRoot_Reflex_OPERATOR = (1 << 10),         //              X       X
   CRoot_Reflex_INLINE = (1 << 11),           //              X       X
   CRoot_Reflex_CONVERTER = (1 << 12),        //              X       X
   CRoot_Reflex_AUTO = (1 << 13),             //          X           X
   CRoot_Reflex_MUTABLE = (1 << 14),          //          X           X
   CRoot_Reflex_CONST = (1 << 15),            //          X       X   X
   CRoot_Reflex_VOLATILE = (1 << 16),         //          X       X   X
   CRoot_Reflex_REFERENCE = (1 << 17),        //          X           X
   CRoot_Reflex_ABSTRACT = (1 << 18),         //      X       X   X
   CRoot_Reflex_VIRTUAL = (1 << 19),          //  X   X           X
   CRoot_Reflex_TRANSIENT = (1 << 20),        //          X           X
   CRoot_Reflex_ARTIFICIAL = (1 << 21),       //  X   X   X   X   X   X
   // the bits 31 - 28 are reserved for template default arguments
   CRoot_Reflex_TEMPLATEDEFAULTS1 = (0 << 31) & (0 << 30) & (0 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS2 = (0 << 31) & (0 << 30) & (1 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS3 = (0 << 31) & (0 << 30) & (1 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS4 = (0 << 31) & (1 << 30) & (0 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS5 = (0 << 31) & (1 << 30) & (0 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS6 = (0 << 31) & (1 << 30) & (1 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS7 = (0 << 31) & (1 << 30) & (1 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS8 = (1 << 31) & (0 << 30) & (0 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS9 = (1 << 31) & (0 << 30) & (0 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS10 = (1 << 31) & (0 << 30) & (1 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS11 = (1 << 31) & (0 << 30) & (1 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS12 = (1 << 31) & (1 << 30) & (0 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS13 = (1 << 31) & (1 << 30) & (0 << 29) & (1 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS14 = (1 << 31) & (1 << 30) & (1 << 29) & (0 << 28),
   CRoot_Reflex_TEMPLATEDEFAULTS15 = (1 << 31) & (1 << 30) & (1 << 29) & (1 << 28)
} CRoot_Reflex_ENTITY_DESCRIPTION;


/** enum for printing names */
typedef enum {
   CRoot_Reflex_FINAL = (1 << 0),
   CRoot_Reflex_QUALIFIED = (1 << 1),
   CRoot_Reflex_SCOPED = (1 << 2),
   CRoot_Reflex_F = (1 << 4),
   CRoot_Reflex_Q = (1 << 5),
   CRoot_Reflex_S = (1 << 6)
} CRoot_Reflex_ENTITY_HANDLING;

/** enum containing all possible types and scopes */
typedef enum {
   CRoot_Reflex_CLASS = 0,
   CRoot_Reflex_STRUCT,
   CRoot_Reflex_ENUM,
   CRoot_Reflex_FUNCTION,
   CRoot_Reflex_ARRAY,
   CRoot_Reflex_FUNDAMENTAL,
   CRoot_Reflex_POINTER,
   CRoot_Reflex_POINTERTOMEMBER,
   CRoot_Reflex_TYPEDEF,
   CRoot_Reflex_UNION,
   CRoot_Reflex_TYPETEMPLATEINSTANCE,
   CRoot_Reflex_MEMBERTEMPLATEINSTANCE,
   CRoot_Reflex_NAMESPACE,
   CRoot_Reflex_DATAMEMBER,
   CRoot_Reflex_FUNCTIONMEMBER,
   CRoot_Reflex_UNRESOLVED
} CRoot_Reflex_TYPE;

/** enum containing all possible 'representation' types */
typedef enum {
   CRoot_Reflex_REPRES_POINTER = 'a' - 'A',                 // To be added to the other value to refer to a pointer to
   CRoot_Reflex_REPRES_CHAR = 'c',
   CRoot_Reflex_REPRES_SIGNED_CHAR = 'c',
   CRoot_Reflex_REPRES_SHORT_INT = 's',
   CRoot_Reflex_REPRES_INT = 'i',
   CRoot_Reflex_REPRES_LONG_INT = 'l',
   CRoot_Reflex_REPRES_UNSIGNED_CHAR = 'b',
   CRoot_Reflex_REPRES_UNSIGNED_SHORT_INT = 'r',
   CRoot_Reflex_REPRES_UNSIGNED_INT = 'h',
   CRoot_Reflex_REPRES_UNSIGNED_LONG_INT = 'k',
   CRoot_Reflex_REPRES_BOOL = 'g',
   CRoot_Reflex_REPRES_FLOAT = 'f',
   CRoot_Reflex_REPRES_DOUBLE = 'd',
   CRoot_Reflex_REPRES_LONG_DOUBLE = 'q',
   CRoot_Reflex_REPRES_VOID = 'y',
   CRoot_Reflex_REPRES_LONGLONG = 'n',
   CRoot_Reflex_REPRES_ULONGLONG = 'm',
   CRoot_Reflex_REPRES_STRUCT = 'u',
   CRoot_Reflex_REPRES_CLASS = 'u',
   CRoot_Reflex_REPRES_ENUM = 'i',                   // Intentionally equal to REPRES_INT
   CRoot_Reflex_REPRES_NOTYPE = '\0'
                   // '1' is also a value used (for legacy implementation of function pointer)
                   // 'E' is also a value used (for legacy implementation of FILE*)
                   // 'a', 'j', 'T', 'o', 'O', 'p', 'P', 'z', 'Z', '\011', '\001', 'w' are also a value used (for support of various interpreter types)
} CRoot_Reflex_REPRESTYPE;

typedef enum {
   CRoot_Reflex_INHERITEDMEMBERS_DEFAULT,    // NO by default, set to ALSO by UpdateMembers()
   CRoot_Reflex_INHERITEDMEMBERS_NO,
   CRoot_Reflex_INHERITEDMEMBERS_ALSO
} CRoot_Reflex_EMEMBERQUERY;

typedef enum {
   CRoot_Reflex_DELAYEDLOAD_OFF,
   CRoot_Reflex_DELAYEDLOAD_ON
} CRoot_Reflex_EDELAYEDLOADSETTING;

CROOT_API
void
CRoot_Reflex_FireClassCallback(CRoot_Reflex_Type self);

CROOT_API
void
CRoot_Reflex_FireFunctionCallback(CRoot_Reflex_Member self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Type_new(const char* name, unsigned int modifiers);

CROOT_API
void
CRoot_Reflex_Type_delete(CRoot_Reflex_Type self);

CROOT_API
void*
CRoot_Reflex_Type_Id(CRoot_Reflex_Type self);

CROOT_API
size_t
CRoot_Reflex_Type_ArrayLength(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Type_ByName(const char *name);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_Type_FunctionMemberAt(CRoot_Reflex_Type self,
                                   size_t nth,
                                   CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
size_t
CRoot_Reflex_Type_FunctionMemberSize(CRoot_Reflex_Type self,
                                     CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_Type_DataMemberAt(CRoot_Reflex_Type self,
                               size_t nth,
                               CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
size_t
CRoot_Reflex_Type_DataMemberSize(CRoot_Reflex_Type self,
                                 CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
bool
CRoot_Reflex_Type_IsAbstract(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsArray(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsClass(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsComplete(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsEquivalentTo(CRoot_Reflex_Type self,
                                 CRoot_Reflex_Type other,
                                 unsigned int modifiers_mask);

CROOT_API
bool
CRoot_Reflex_Type_IsFunction(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsFundamental(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsPrivate(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsProtected(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsPublic(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsPointer(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsPointerToMember(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsReference(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsStruct(CRoot_Reflex_Type self);

CROOT_API
bool
CRoot_Reflex_Type_IsVirtual(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_Type_MemberAt(CRoot_Reflex_Type self,
                           size_t nth,
                           CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
size_t
CRoot_Reflex_Type_MemberSize(CRoot_Reflex_Type self,
                             CRoot_Reflex_EMEMBERQUERY inh);

CROOT_API
const char*
CRoot_Reflex_Type_Name(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_PropertyList
CRoot_Reflex_Type_Properties(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Type_RawType(CRoot_Reflex_Type self);

CROOT_API
size_t
CRoot_Reflex_Type_SizeOf(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Type_ToType(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Type_TypeAt(size_t nth);

CROOT_API
size_t
CRoot_Reflex_Type_TypeSize();

CROOT_API
CRoot_Reflex_TYPE
CRoot_Reflex_Type_TypeType(CRoot_Reflex_Type self);

CROOT_API
void
CRoot_Reflex_Type_Unload(CRoot_Reflex_Type self);

CROOT_API
void
CRoot_Reflex_Type_UpdateMembers(CRoot_Reflex_Type self);

CROOT_API
void
CRoot_Reflex_Type_AddDataMember(CRoot_Reflex_Type self,
                                CRoot_Reflex_Member dm);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_Type_AddDataMember2(CRoot_Reflex_Type self,
                                 const char* name,
                                 CRoot_Reflex_Type type,
                                 size_t offset,
                                 unsigned int modifiers,
                                 char *interpreterOffset);

CROOT_API
void
CRoot_Reflex_Type_RemoveDataMember(CRoot_Reflex_Type self,
                                   CRoot_Reflex_Member dm);

CROOT_API
void
CRoot_Reflex_Type_SetSize(CRoot_Reflex_Type self,
                          size_t s);

CROOT_API
CRoot_Reflex_REPRESTYPE
CRoot_Reflex_Type_RepresType(CRoot_Reflex_Type self);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_Member_new();

CROOT_API
void
CRoot_Reflex_Member_delete(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsDataMember(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsPrivate(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsProtected(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsPublic(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsTransient(CRoot_Reflex_Member self);

CROOT_API
bool
CRoot_Reflex_Member_IsVirtual(CRoot_Reflex_Member self);

CROOT_API
CRoot_Reflex_TYPE
CRoot_Reflex_Member_MemberType(CRoot_Reflex_Member self);

CROOT_API
const char*
CRoot_Reflex_Member_Name(CRoot_Reflex_Member self);

CROOT_API
size_t
CRoot_Reflex_Member_Offset(CRoot_Reflex_Member self);

CROOT_API
void
CRoot_Reflex_Member_InterpreterOffset(CRoot_Reflex_Member self,
                                      char *offset);

CROOT_API
CRoot_Reflex_PropertyList
CRoot_Reflex_Member_Properties(CRoot_Reflex_Member self);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_Member_TypeOf(CRoot_Reflex_Member self);

CROOT_API
void*
CRoot_Reflex_Member_Stubcontext(CRoot_Reflex_Member self);

CROOT_API
CRoot_Reflex_StubFunction
CRoot_Reflex_Member_Stubfunction(CRoot_Reflex_Member self);

/* propertylist api */
CROOT_API
const char*
CRoot_Reflex_PropertyList_PropertyAsString(CRoot_Reflex_PropertyList self,
                                           size_t idx);

CROOT_API
size_t
CRoot_Reflex_PropertyList_PropertyCount(CRoot_Reflex_PropertyList self);

CROOT_API
const char*
CRoot_Reflex_PropertyList_PropertyKeys(CRoot_Reflex_PropertyList self);

/* type builder API */

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_PointerBuilder_new(CRoot_Reflex_Type t);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_ArrayBuilder_new(CRoot_Reflex_Type t,
                              size_t n);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new(CRoot_Reflex_Type r);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new1(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new2(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0,
                                      CRoot_Reflex_Type t1);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_FunctionTypeBuilder_new3(CRoot_Reflex_Type r,
                                      CRoot_Reflex_Type t0,
                                      CRoot_Reflex_Type t1,
                                      CRoot_Reflex_Type t2);

CROOT_API
CRoot_Reflex_ClassBuilder
CRoot_Reflex_ClassBuilder_new(const char *name,
                              void* typeinfo,
                              size_t size,
                              unsigned int modifiers,
                              CRoot_Reflex_TYPE type);

CROOT_API
void
CRoot_Reflex_ClassBuilder_delete(CRoot_Reflex_ClassBuilder self);

CROOT_API
void
CRoot_Reflex_ClassBuilder_AddDataMember(CRoot_Reflex_ClassBuilder self,
                                        CRoot_Reflex_Type type,
                                        const char* name,
                                        size_t offset,
                                        unsigned int modifiers);

CROOT_API
void
CRoot_Reflex_ClassBuilder_AddFunctionMember(CRoot_Reflex_ClassBuilder self,
                                            CRoot_Reflex_Type type,
                                            const char *name,
                                            CRoot_Reflex_StubFunction stubFP,
                                            void *stubCtx,
                                            const char *params,
                                            unsigned int modifiers);

CROOT_API
void
CRoot_Reflex_ClassBuilder_AddProperty(CRoot_Reflex_ClassBuilder self,
                                      const char *key,
                                      const char *value);

CROOT_API
CRoot_Reflex_Type
CRoot_Reflex_ClassBuilder_ToType(CRoot_Reflex_ClassBuilder self);


CROOT_API
CRoot_Reflex_FunctionBuilder
CRoot_Reflex_FunctionBuilder_new(CRoot_Reflex_Type type,
                                 const char* name,
                                 CRoot_Reflex_StubFunction stubFP,
                                 void *stubCtx,
                                 const char *params,
                                 unsigned char modifiers);

CROOT_API
void
CRoot_Reflex_FunctionBuilder_delete(CRoot_Reflex_FunctionBuilder self);

CROOT_API
CRoot_Reflex_Member
CRoot_Reflex_FunctionBuilder_ToMember(CRoot_Reflex_FunctionBuilder self);


CROOT_API
size_t
CRoot_Reflex_PropertyList_AddProperty(CRoot_Reflex_PropertyList self,
                                      const char *key,
                                      const char *value);

#ifdef __cplusplus
}
#endif

#endif /* !CROOT_CROOT_REFLEX_H */
