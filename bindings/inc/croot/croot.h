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

#include "croot/croot_types.h"
#include "croot/croot_globals.h"

#include "croot/croot_goobject.h"

#include "croot/croot_chain.h"
#include "croot/croot_cintex.h"
#include "croot/croot_cint.h"
#include "croot/croot_class.h"
#include "croot/croot_file.h"
#include "croot/croot_hist.h"
#include "croot/croot_interpreter.h"
#include "croot/croot_leaf.h"
#include "croot/croot_math.h"
#include "croot/croot_object.h"
#include "croot/croot_objarray.h"
#include "croot/croot_random.h"
#include "croot/croot_reflex.h"
#include "croot/croot_root.h"
#include "croot/croot_tree.h"

#endif /* !CROOT_CROOT_H */
