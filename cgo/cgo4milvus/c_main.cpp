#include "cwrap.h"
#include "segment_base.h"

//void test_raw_segment() {
//  milvus::engine::SegmentBase segmentBase;
//
//  std::vector<int64_t> primary_keys{1, 2, 3};
//  std::vector<int64_t> timestamps{100, 200, 300};
//
//  std::vector<double> vector{0.1111, 0.222, 0.333};
//
//  milvus::engine::DogDataChunk dogDataChunk;
//  dogDataChunk.id = 1;
//  dogDataChunk.vector = vector;
//
//  auto dataChunkPtr = std::make_shared<milvus::engine::DogDataChunk>(dogDataChunk);
//
//  segmentBase.Insert(primary_keys, timestamps, dataChunkPtr);
//}

void test_cwrap_segment() {
  std::vector<int64_t> primary_keys{1, 2, 3};
  std::vector<int64_t> timestamps{100, 200, 300};

  std::vector<double> vector{0.1111, 0.222, 0.333};

  milvus::engine::DogDataChunk dogDataChunk;
  dogDataChunk.id = 1;
  dogDataChunk.vector = vector;

  auto dataChunkPtr = std::make_shared<milvus::engine::DogDataChunk>(dogDataChunk);

  auto segment = SegmentBaseInit();
  Insert(segment, primary_keys, timestamps, dataChunkPtr);
}

int main() {
  test_cwrap_segment();
  return 0;
}


