#include "segment_base.h"
#include "cwrap.h"

CSegmentBase SegmentBaseInit() {
  CSegmentBase cb = new milvus::engine::SegmentBase();
  return (void*)cb;
}

bool Insert(CSegmentBase cb, const std::vector<int64_t> &primary_keys,
            const std::vector<milvus::engine::Timestamp> &timestamps,
            milvus::engine::DogDataChunkPtr values) {
  auto seg = (milvus::engine::SegmentBase*)cb;
  return seg->Insert(primary_keys, timestamps, values);
}