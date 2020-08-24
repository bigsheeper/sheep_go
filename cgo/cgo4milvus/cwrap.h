#ifdef __cplusplus
extern "C" {
#endif

typedef void* CSegmentBase;

CSegmentBase SegmentBaseInit();

bool Insert(CSegmentBase cb, const std::vector<int64_t> &primary_keys,
                           const std::vector<milvus::engine::Timestamp> &timestamps,
                           milvus::engine::DogDataChunkPtr values);

#ifdef __cplusplus
}
#endif