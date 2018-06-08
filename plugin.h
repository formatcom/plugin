static uintptr_t pluginOpen(const char* path, char** err);
static void* pluginLookup(uintptr_t h, const char* name, char** err);
int pluginClose(uintptr_t h, char** err);
