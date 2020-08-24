#pragma once

#include <vector>
#include <cstdint>
#include <memory>

namespace milvus {
namespace engine {

struct DogDataChunk {
    std::vector<double> vector;
    int id;
};

using DogDataChunkPtr = std::shared_ptr<DogDataChunk>;

using Timestamp = int64_t;

using Status = bool;

using QueryResult = std::vector<std::vector<int64_t >>;

class SegmentBase {
public:
    enum class SegmentState {
        Invalid = 0,
        Open,
        Closed
    };

public:
    ~SegmentBase() = default;

    Status
    Insert(const std::vector <int64_t> &primary_keys, const std::vector <Timestamp> &timestamps, DogDataChunkPtr values);

    Status
    Delete(const std::vector <int64_t> &primary_keys, const std::vector <Timestamp> &timestamps);

    Status
    Query(const std::string &query, Timestamp timestamp, QueryResult &results);

    Status
    Close();

    Status
    BuildIndex(std::shared_ptr <std::string> index_params);

    Status
    DropIndex(const std::string &field_name);

    Status
    DropRawData(const std::string &field_name);

    Status
    LoadRawData(const std::string &field_name, const char *blob, int64_t blob_size);

public:
    uint32_t
    get_row_count() const;

    SegmentState
    get_state() const;

    Timestamp
    get_max_timestamp();

    Timestamp
    get_min_timestamp();

    uint32_t
    get_deleted_count() const;
};

std::shared_ptr <SegmentBase> CreateSegment(/*args*/);

}  // namespace engine
}  // namespace milvus
