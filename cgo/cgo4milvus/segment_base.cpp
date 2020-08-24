#include <iostream>
#include "segment_base.h"

namespace milvus {
namespace engine {

Status SegmentBase::Insert(const std::vector<int64_t> &primary_keys,
                           const std::vector<milvus::engine::Timestamp> &timestamps,
                           milvus::engine::DogDataChunkPtr values) {
  std::cout << "mock insert" << std::endl;

  for (auto key: primary_keys) {
    std::cout << key << std::endl;
  }

  for (auto timestamp: timestamps) {
    std::cout << timestamp <<std::endl;
  }

  auto vector = values->vector;
  auto id = values->id;

  for (auto v: vector) {
    std::cout << v << std::endl;
  }
  std::cout << "id = " << id << std::endl;

  return true;
}

Status SegmentBase::Delete(const std::vector<int64_t> &primary_keys,
                          const std::vector<milvus::engine::Timestamp> &timestamps) {
  return true;
}

Status SegmentBase::Query(const std::string &query, milvus::engine::Timestamp timestamp,
                          milvus::engine::QueryResult &results) {
  return true;
}

Status SegmentBase::Close() {
  return true;
}

Status SegmentBase::BuildIndex(std::shared_ptr<std::string> index_params) {
  return true;
}

Status SegmentBase::DropIndex(const std::string &field_name) {
  return true;
}

Status SegmentBase::DropRawData(const std::string &field_name) {
  return true;
}

Status SegmentBase::LoadRawData(const std::string &field_name, const char *blob, int64_t blob_size) {
  return true;
}

uint32_t SegmentBase::get_row_count() const {
  return true;
}

SegmentBase::SegmentState SegmentBase::get_state() const {
  return SegmentBase::SegmentState::Open;
}

Timestamp SegmentBase::get_max_timestamp() {
  return 9999;
}

Timestamp SegmentBase::get_min_timestamp() {
  return -9999;
}

uint32_t SegmentBase::get_deleted_count() const {
  return 1234;
}

}
}
