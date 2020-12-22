//
// Created by perfest on 12/22/20.
//

#ifndef SOURCES_UTILS_HPP_
#define SOURCES_UTILS_HPP_

#include <nlohmann/json.hpp>
#include <suggest.pb.h>

namespace suggest {

void from_json(const nlohmann::json& j, Suggestion& suggestion);

} // namespace suggest

#endif // SOURCES_UTILS_HPP_