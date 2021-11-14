#include "MediaInfoDLL.h"

static inline void* g_MediaInfo_New() {
  void *mi = MediaInfo_New();
  if (mi)
    MediaInfo_Option(mi, "CharSet", "UTF-8");
  return mi;
}

static inline void g_MediaInfo_Delete(void *mi) {
  return MediaInfo_Delete(mi);
}

static inline size_t g_MediaInfo_Open(void *mi, const char *name) {
  return MediaInfo_Open(mi, name);
}

static inline const char* g_MediaInfo_Option(void *mi, const char *option, const char *value) {
  return MediaInfo_Option(mi, option, value);
}

static inline const char* g_MediaInfo_Inform(void *mi) {
  return MediaInfo_Inform(mi, 0);
}

static inline const char* g_MediaInfo_Get(void *mi, MediaInfo_stream_C streamKind, size_t streamNumber,
  const char *parameter, MediaInfo_info_C kindOfInfo, MediaInfo_info_C kindOfSearch) {
  return MediaInfo_Get(mi, streamKind, streamNumber, parameter, kindOfInfo, kindOfSearch);
}

static inline const char* g_MediaInfo_GetI(void *mi, MediaInfo_stream_C streamKind, size_t streamNumber,
  size_t parameter, MediaInfo_info_C kindOfInfo) {
  return MediaInfo_GetI(mi, streamKind, streamNumber, parameter, kindOfInfo);
}

static inline size_t g_MediaInfo_Count_Get(void *mi, MediaInfo_stream_C streamKind) {
  return MediaInfo_Count_Get(mi, streamKind, -1);
}

static inline void g_MediaInfo_Close(void *mi) {
  MediaInfo_Close(mi);
}
