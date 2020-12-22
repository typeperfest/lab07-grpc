//
// Created by perfest on 12/22/20.
//
#include "utils.h"
#include <nlohmann/json.hpp>
#include <cstdint>
#include <string>

namespace suggest {

void from_json(const nlohmann::json& j, Suggestion& s) {
  s.set_text(j.at("name").get<std::string>());
  s.set_position(j.at("cost").get<uint32_t>());
}

} // namespace suggest