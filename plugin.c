#include <dlfcn.h>
#include "plugin.h"

static uintptr_t pluginOpen(const char* path, char** err){
	void *h = dlopen(path, RTLD_NOW|RTLD_GLOBAL);
	if (h == NULL){
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}

static void* pluginLookup(uintptr_t h, const char* name, char** err){
	void *sym = dlsym((void*)h, name);
	if (sym == NULL){
		*err = (char*)dlerror();
	}
	return sym;
}

int pluginClose(uintptr_t h, char** err){
	if (dlclose((void*)h) != 0){
		*err = (char*)dlerror();
		return 1;
	}
	return 0;
}
