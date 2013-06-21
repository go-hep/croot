#include "croot/croot.h"

// ROOT includes
#include "DllImport.h"
#include "Rtypes.h"

// STL includes
#include <map>
#include <string>

// C includes
#include <stdlib.h>

// --------------------------------------------------
#include "Reflex/Builder/ReflexBuilder.h"
#include "Reflex/Builder/ClassBuilder.h"
#include "Reflex/Type.h"
#include "Cintex/Cintex.h"
/*
namespace golang {
  struct gostring {
    char *Data;
    int   Len;
  };

  template <typename T>
  struct goslice {
    T   *Data;
    int  Len;
    int  Cap;
  };

  template struct goslice<double>;
}

namespace {
  struct init_reflex_gostructs {
    init_reflex_gostructs() {
      ROOT::Cintex::Cintex::Enable();
      std::string full_name = "golang::goslice<double>";
      Reflex::ClassBuilder *bldr = new Reflex::ClassBuilder
        (full_name,
         sizeof(golang::goslice<double>),
         Reflex::PUBLIC|Reflex::ARTIFICIAL,
         Reflex::STRUCT);
      size_t offset = sizeof(double*);
      bldr->AddDataMember(Reflex::Type::ByName("int"),
                          "Len",
                          offset,
                          Reflex::PUBLIC);
      offset += sizeof(int);
      bldr->AddDataMember(Reflex::Type::ByName("int"),
                          "Cap",
                          offset,
                          Reflex::PUBLIC);

      offset = 0;
      Reflex::Type ty_double = Reflex::Type::ByName("double");
      Reflex::Type ty_double_p = Reflex::PointerBuilder(ty_double);
      bldr->AddDataMember(ty_double_p
                          "Data",
                          offset,
                          Reflex::PUBLIC);
      bldr->AddProperty("comment", "[Len]");

      Reflex::Type ty_void = Reflex::Type::ByName("void");
      

    }
  };
  init_reflex_gostructs s_init_reflex_gostructs;
}
// --------------------------------------------------
*/

struct GoObject {
  void       *Ptr;
  const char *Type;
  int32_t     Size; // Size == 0: Scalar
                    // >0 : size of elements of array in bytes
};

CRoot_GoObject
CRoot_GoObject_New(void *ptr, const char *type)
{
  GoObject *self = new GoObject;
  self->Ptr = ptr;
  self->Type = type;
  self->Size = -1;
  return ((CRoot_GoObject)self);
}

void*
CRoot_GoObject_Ptr(CRoot_GoObject self)
{
  return ((GoObject*)self)->Ptr;
}

const char*
CRoot_GoObject_Type(CRoot_GoObject self)
{
  return ((GoObject*)self)->Type;
}

int32_t
CRoot_GoObject_Size(CRoot_GoObject self)
{
  return ((GoObject*)self)->Size;
}

void
CRoot_GoObject_SetPtr(CRoot_GoObject self, void *ptr)
{
  ((GoObject*)self)->Ptr = ptr;
}

void
CRoot_GoObject_SetSize(CRoot_GoObject self, int32_t size)
{
  ((GoObject*)self)->Size = size;
}

struct GoObjConverter
{
  virtual ~GoObjConverter();

  virtual int32_t CnvFromC(CRoot_GoObject, void*) = 0;
  virtual int32_t CnvToC(CRoot_GoObject, void*) = 0;
};

GoObjConverter::~GoObjConverter()
{}

typedef CRoot_GoObject_Converter (*GoConverterFactory_t)();
typedef std::map<std::string, GoConverterFactory_t> GoCnvFactories_t;
GoCnvFactories_t g_gocnv;

int32_t
CRoot_GoObject_CnvFromC(CRoot_GoObject_Converter self, CRoot_GoObject obj, void* address)
{
  return ((GoObjConverter*)self)->CnvFromC(obj, address);
}

int32_t
CRoot_GoObject_CnvToC(CRoot_GoObject_Converter self, CRoot_GoObject obj, void* address)
{
  return ((GoObjConverter*)self)->CnvToC(obj, address);
}

CRoot_GoObject_Converter
CRoot_GoObject_Converter_Create(const char *fullname)
{
  GoCnvFactories_t::iterator itr = g_gocnv.find(fullname);
  if (itr != g_gocnv.end()) {
    return (itr->second)();
  }
  
  return NULL;
}

CRoot_GoObject_Converter
CRoot_GoObject_Converter_Get(const char *fullname)
{
  GoCnvFactories_t::iterator itr = g_gocnv.find(fullname);
  if (itr != g_gocnv.end()) {
    return (itr->second)();
  }
  return NULL;
}

#define GOCROOT_DECLARE_BASIC_CONVERTER( name )                         \
  class CRoot_##name##Converter : public GoObjConverter {               \
  public:                                                               \
  virtual int32_t CnvFromC(CRoot_GoObject, void*);                      \
  virtual int32_t CnvToC(CRoot_GoObject, void*);                        \
  }

#define GOCROOT_DECLARE_BASIC_CONVERTER2( name, base )                  \
  class CRoot_##name##Converter : public CRoot_##base##Converter {      \
  public:                                                               \
  virtual int32_t CnvFromC(CRoot_GoObject, void*);                      \
  virtual int32_t CnvToC(CRoot_GoObject, void*);                        \
  }

#define GOCROOT_DECLARE_ARRAY_CONVERTER( name )                         \
  class CRoot_##name##Converter : public GoObjConverter {               \
  public:                                                               \
  CRoot_##name##Converter( int32_t size = -1 ) : m_size( size ) {}      \
  virtual int32_t CnvFromC(CRoot_GoObject, void*);                      \
  virtual int32_t CnvToC(CRoot_GoObject, void*);                        \
  private:                                                              \
   int32_t m_size;                                                     \
  }

// converters for built-ins
GOCROOT_DECLARE_BASIC_CONVERTER( Long );
//GOCROOT_DECLARE_BASIC_CONVERTER( LongRef );
GOCROOT_DECLARE_BASIC_CONVERTER( Bool );
GOCROOT_DECLARE_BASIC_CONVERTER( Char );
GOCROOT_DECLARE_BASIC_CONVERTER( UChar );
GOCROOT_DECLARE_BASIC_CONVERTER2( Short, Long );
GOCROOT_DECLARE_BASIC_CONVERTER2( UShort, Long );
GOCROOT_DECLARE_BASIC_CONVERTER2( Int, Long );
//GOCROOT_DECLARE_BASIC_CONVERTER( IntRef );
GOCROOT_DECLARE_BASIC_CONVERTER( ULong );
GOCROOT_DECLARE_BASIC_CONVERTER2( UInt, ULong );
GOCROOT_DECLARE_BASIC_CONVERTER( LongLong );
GOCROOT_DECLARE_BASIC_CONVERTER( ULongLong );
GOCROOT_DECLARE_BASIC_CONVERTER( Float );
GOCROOT_DECLARE_BASIC_CONVERTER( Double );
//GOCROOT_DECLARE_BASIC_CONVERTER2( Float, Double );
//GOCROOT_DECLARE_BASIC_CONVERTER( DoubleRef );

// arrays
GOCROOT_DECLARE_ARRAY_CONVERTER( BoolArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( ShortArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( UShortArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( IntArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( UIntArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( LongArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( ULongArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( FloatArray );
GOCROOT_DECLARE_ARRAY_CONVERTER( DoubleArray );

// --- impls ---
#define GOCROOT_IMPLEMENT_BASIC_CONVERTER( name, type )         \
  int32_t CRoot_##name##Converter::CnvFromC(CRoot_GoObject obj, void* address ) \
  {                                                                     \
    ((GoObject*)obj)->Ptr = (void*)((type*)address);                    \
    return 1;                                                           \
  }                                                                     \
                                                                        \
  int32_t CRoot_##name##Converter::CnvToC(CRoot_GoObject obj, void* address ) \
  {                                                                     \
    address = (void*)(((GoObject*)obj)->Ptr);                           \
    return 1;                                                           \
  }

#define GOCROOT_IMPLEMENT_ARRAY_CONVERTER( name, type ) \
  int32_t                                                               \
  CRoot_##name##ArrayConverter::CnvFromC(CRoot_GoObject obj, void* address ) \
  {                                                                     \
    memcpy(((GoObject*)obj)->Ptr, *(type**)address, ((GoObject*)obj)->Size); \
    return 1;                                                           \
  }                                                                     \
                                                                        \
  int32_t                                                               \
  CRoot_##name##ArrayConverter::CnvToC(CRoot_GoObject obj, void* address ) \
  {                                                                     \
    memcpy(*(type**)address, ((GoObject*)obj)->Ptr, ((GoObject*)obj)->Size); \
    return 1;                                                           \
  }

// built-ins
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Long,   Long_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Bool,   Bool_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Char,   Char_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( UChar,  UChar_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Short,  Short_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( UShort, UShort_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Int,    Int_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( ULong,  ULong_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( UInt,   UInt_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Float,  Float_t )
GOCROOT_IMPLEMENT_BASIC_CONVERTER( Double, Double_t )

// arrays
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Bool,   Bool_t)   // signed char
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Short,  Short_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( UShort, UShort_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Int,    Int_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( UInt,   UInt_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Long,   Long_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( ULong,  ULong_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Float,  Float_t)
GOCROOT_IMPLEMENT_ARRAY_CONVERTER( Double, Double_t)

#define GOCROOT_BASIC_CONVERTER_FACTORY( name )                         \
  CRoot_GoObject_Converter Create##name##Converter()                    \
  {                                                                     \
    return (CRoot_GoObject_Converter)(new CRoot_##name##Converter());   \
  }

#define GOCROOT_ARRAY_CONVERTER_FACTORY( name )                         \
  CRoot_GoObject_Converter Create##name##Converter()                    \
  {                                                                     \
    return (CRoot_GoObject_Converter)(new CRoot_##name##Converter());   \
  }

GOCROOT_BASIC_CONVERTER_FACTORY( Bool )
GOCROOT_BASIC_CONVERTER_FACTORY( Char )
GOCROOT_BASIC_CONVERTER_FACTORY( UChar )
GOCROOT_BASIC_CONVERTER_FACTORY( Short )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstShortRef )
GOCROOT_BASIC_CONVERTER_FACTORY( UShort )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstUShortRef )
GOCROOT_BASIC_CONVERTER_FACTORY( Int )
//GOCROOT_BASIC_CONVERTER_FACTORY( IntRef )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstIntRef )
GOCROOT_BASIC_CONVERTER_FACTORY( UInt )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstUIntRef )
GOCROOT_BASIC_CONVERTER_FACTORY( Long )
//GOCROOT_BASIC_CONVERTER_FACTORY( LongRef )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstLongRef )
GOCROOT_BASIC_CONVERTER_FACTORY( ULong )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstULongRef )
GOCROOT_BASIC_CONVERTER_FACTORY( Float )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstFloatRef )
GOCROOT_BASIC_CONVERTER_FACTORY( Double )
//GOCROOT_BASIC_CONVERTER_FACTORY( DoubleRef )
//GOCROOT_BASIC_CONVERTER_FACTORY( ConstDoubleRef )
// GOCROOT_BASIC_CONVERTER_FACTORY( Void )
// GOCROOT_BASIC_CONVERTER_FACTORY( Macro )
// GOCROOT_BASIC_CONVERTER_FACTORY( LongLong )
// GOCROOT_BASIC_CONVERTER_FACTORY( ConstLongLongRef )
// GOCROOT_BASIC_CONVERTER_FACTORY( ULongLong )
// GOCROOT_BASIC_CONVERTER_FACTORY( ConstULongLongRef )
// GOCROOT_ARRAY_CONVERTER_FACTORY( CString )
// GOCROOT_ARRAY_CONVERTER_FACTORY( NonConstCString )
// GOCROOT_ARRAY_CONVERTER_FACTORY( NonConstUCString )
GOCROOT_ARRAY_CONVERTER_FACTORY( BoolArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( ShortArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( UShortArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( IntArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( UIntArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( LongArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( ULongArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( FloatArray )
GOCROOT_ARRAY_CONVERTER_FACTORY( DoubleArray )

#define ADD_CNV(name, type) \
  g_gocnv[name] = &type

namespace {

  struct init_go_cnv_factories {
  public:
    init_go_cnv_factories() {
      ADD_CNV("bool", CreateBoolConverter);
      ADD_CNV("char", CreateCharConverter);
      ADD_CNV("unsigned char",             CreateUCharConverter              );
      ADD_CNV("short",                     CreateShortConverter              );
      //ADD_CNV("const short &",             CreateConstShortRefConverter      );
      ADD_CNV("unsigned short",            CreateUShortConverter             );
      //ADD_CNV("const unsigned short&",     CreateConstUShortRefConverter     );
      ADD_CNV("int",                       CreateIntConverter                );
      //ADD_CNV("int&",                      CreateIntRefConverter             );
      //ADD_CNV("const int&",                CreateConstIntRefConverter        );
      ADD_CNV("unsigned int",              CreateUIntConverter               );
      //ADD_CNV("const unsigned int&",       CreateConstUIntRefConverter       );
      ADD_CNV("UInt_t", /* enum */         CreateUIntConverter               );
      ADD_CNV("long",                      CreateLongConverter               );
      //ADD_CNV("long&",                     CreateLongRefConverter            );
      //ADD_CNV("const long&",               CreateConstLongRefConverter       );
      ADD_CNV("unsigned long",             CreateULongConverter              );
      //ADD_CNV("const unsigned long&",      CreateConstULongRefConverter      );
      //ADD_CNV("long long",                 CreateLongLongConverter           );
      //ADD_CNV("const long long&",          CreateConstLongLongRefConverter   );
      //ADD_CNV("Long64_t",                  CreateLongLongConverter           );
      //ADD_CNV("const Long64_t&",           CreateConstLongLongRefConverter   );
      //ADD_CNV("unsigned long long",        CreateULongLongConverter          );
      //ADD_CNV("const unsigned long long&", CreateConstULongLongRefConverter  );
      //ADD_CNV("ULong64_t",                 CreateULongLongConverter          );
      //ADD_CNV("const ULong64_t&",          CreateConstULongLongRefConverter  );
      ADD_CNV("float",                     CreateFloatConverter              );
      //ADD_CNV("const float&",              CreateConstFloatRefConverter      );
                                           
      ADD_CNV("double",                    CreateDoubleConverter             );
      //ADD_CNV("double&",                   CreateDoubleRefConverter          );
      //ADD_CNV("const double&",             CreateConstDoubleRefConverter     );
      //ADD_CNV("void",                      CreateVoidConverter               );
      //ADD_CNV("#define",                   CreateMacroConverter              );

      
      // pointer/array factories
      ADD_CNV( "bool*",                     CreateBoolArrayConverter          );
      //ADD_CNV( "const unsigned char*",      CreateCStringConverter            );
      //ADD_CNV( "unsigned char*",            CreateNonConstUCStringConverter   );
      ADD_CNV( "short*",                    CreateShortArrayConverter         );
      ADD_CNV( "unsigned short*",           CreateUShortArrayConverter        );
      ADD_CNV( "int*",                      CreateIntArrayConverter           );
      ADD_CNV( "unsigned int*",             CreateUIntArrayConverter          );
      ADD_CNV( "long*",                     CreateLongArrayConverter          );
      ADD_CNV( "unsigned long*",            CreateULongArrayConverter         );
      ADD_CNV( "float*",                    CreateFloatArrayConverter         );
      ADD_CNV( "double*",                   CreateDoubleArrayConverter        );
      //ADD_CNV( "long long*",                CreateLongLongArrayConverter      );
      //ADD_CNV( "unsigned long long*",       CreateLongLongArrayConverter      );  // TODO: ULongLong
      //ADD_CNV( "void*",                     CreateVoidArrayConverter          );

    }
  };
  init_go_cnv_factories s_go_init_cnv_factories;
}

// EOF

